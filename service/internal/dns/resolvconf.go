package dns

import (
	"bufio"
	"errors"
	"fmt"
	"hash/crc32"
	"os"
	"slices"
	"strings"
	"sync"
)

type ResolvConf struct {
	mu      sync.RWMutex
	path    string
	entries entries
}

var instance *ResolvConf

func GetResolvConfInstance() *ResolvConf {
	if instance == nil {
		instance = &ResolvConf{
			mu:      sync.RWMutex{},
			path:    "/etc/resolv.conf",
			entries: nil,
		}
	}

	return instance
}

type entry struct {
	option string
	value  string
}

const OptionTypeNameserver = "nameserver"

// Returns nil if line is blank or comment.
func parseEntry(line string) *entry {
	words := strings.Fields(line)

	var e *entry
	for _, word := range words {
		firstChar := word[0]
		if firstChar == ';' || firstChar == '#' {
			break
		}

		if e == nil {
			e = &entry{
				option: word,
			}
			continue
		}

		e.value = word
		break
	}

	return e
}

type entries []entry

func (e entries) Checksum() uint32 {
	bytes := make([]byte, 0)
	for _, entry := range e {
		entryBytes := []byte(entry.option + entry.value)
		bytes = append(bytes, entryBytes...)
	}

	result := crc32.ChecksumIEEE(bytes)

	return result
}

func (e entries) IndexNthNameserver(n int) int {
	cur := 0

	for i, entry := range e {
		if entry.option != OptionTypeNameserver {
			continue
		}

		if cur == n {
			return i
		}

		cur++
	}

	return -1
}

func (r *ResolvConf) read() error {
	f, err := os.Open(r.path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	r.entries = make(entries, 0)
	for scanner.Scan() {
		e := parseEntry(scanner.Text())
		if e == nil {
			continue
		}

		r.entries = append(r.entries, *e)
	}

	return nil
}

func (r *ResolvConf) write() error {
	f, err := os.OpenFile(r.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, e := range r.entries {
		line := fmt.Sprintf("%s %s\n", e.option, e.value)
		_, err = writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func (r *ResolvConf) getNameservers() []string {
	result := make([]string, 0, len(r.entries))
	for _, e := range r.entries {
		if e.option == OptionTypeNameserver {
			result = append(result, e.value)
		}
	}

	return result
}

func (r *ResolvConf) GetNameservers() (nameservers []string, checksum uint32, err error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	err = r.read()
	if err != nil {
		return
	}

	nameservers = r.getNameservers()
	checksum = r.entries.Checksum()

	return
}

func (r *ResolvConf) GetNameserverAt(n int) (nameserver string, checksum uint32, err error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	err = r.read()
	if err != nil {
		return
	}

	index := r.entries.IndexNthNameserver(n)
	if index < 0 {
		err = errors.New("index out of range")
		return
	}

	nameserver = r.entries[index].value
	checksum = r.entries.Checksum()

	return
}

func (r *ResolvConf) createNameserverAt(checksum uint32, index int, nameserver string) error {
	if r.entries.Checksum() != checksum {
		return errors.New("file changed on disk since last read")
	}

	nameservers := r.getNameservers()

	if index > len(nameservers) {
		return errors.New("index out of range")
	}

	if index == len(nameservers) {
		return r.createNameserverLast(nameserver)
	}

	i := r.entries.IndexNthNameserver(index)
	newEntry := entry{
		option: OptionTypeNameserver,
		value:  nameserver,
	}
	r.entries = slices.Insert(r.entries, i, newEntry)

	return nil
}

func (r *ResolvConf) CreateNameserverAt(checksum uint32, index int, nameserver string) (updChecksum uint32, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	err = r.read()
	if err != nil {
		return
	}

	err = r.createNameserverAt(checksum, index, nameserver)
	if err != nil {
		return
	}

	updChecksum = r.entries.Checksum()
	err = r.write()

	return
}

func (r *ResolvConf) createNameserverLast(nameserver string) error {
	newEntry := entry{
		option: OptionTypeNameserver,
		value:  nameserver,
	}

	r.entries = append(r.entries, newEntry)

	return nil
}

func (r *ResolvConf) CreateNameserverLast(nameserver string) (index int, checksum uint32, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	err = r.read()
	if err != nil {
		return
	}

	err = r.createNameserverLast(nameserver)
	if err != nil {
		return
	}

	index = len(r.getNameservers()) - 1
	checksum = r.entries.Checksum()

	err = r.write()

	return
}

func (r *ResolvConf) deleteNameserverAt(checksum uint32, index int) (nameserver string, err error) {
	if r.entries.Checksum() != checksum {
		err = errors.New("file changed on disk since last read")
		return
	}

	deleteAt := r.entries.IndexNthNameserver(index)
	if deleteAt < 0 {
		err = errors.New("index out of range")
		return
	}

	nameserver = r.entries[deleteAt].value
	r.entries = append(r.entries[:deleteAt], r.entries[deleteAt+1:]...)

	return
}

func (r *ResolvConf) DeleteNameserverAt(checksum uint32, index int) (nameserver string, updChecksum uint32, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	err = r.read()
	if err != nil {
		return
	}

	nameserver, err = r.deleteNameserverAt(checksum, index)
	if err != nil {
		return
	}

	updChecksum = r.entries.Checksum()
	err = r.write()

	return
}
