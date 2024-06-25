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

func (e entries) Checksum() int {
	bytes := make([]byte, 0)
	for _, entry := range e {
		entryBytes := []byte(entry.option + entry.value)
		bytes = append(bytes, entryBytes...)
	}

	result := crc32.ChecksumIEEE(bytes)

	return int(result)
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

func (r *ResolvConf) GetNameservers() ([]string, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	err := r.read()
	if err != nil {
		return nil, 0, err
	}

	return r.getNameservers(), r.entries.Checksum(), nil
}

func (r *ResolvConf) GetNameserverAt(n int) (string, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	err := r.read()
	if err != nil {
		return "", 0, err
	}

	index := r.entries.IndexNthNameserver(n)
	if index < 0 {
		return "", 0, errors.New("index out of range")
	}
	nameserver := r.entries[index].value

	return nameserver, r.entries.Checksum(), nil
}

func (r *ResolvConf) createNameserverAt(checksum, index int, nameserver string) error {
	if r.entries.Checksum() != checksum {
		return errors.New("file changed on disk since last read")
	}

	nameservers := r.getNameservers()

	if index > len(nameservers) {
		return errors.New("index out of range")
	}

	if index == len(nameservers) {
		return r.createNameserverLast(checksum, nameserver)
	}

	i := r.entries.IndexNthNameserver(index)
	newEntry := entry{
		option: OptionTypeNameserver,
		value:  nameserver,
	}
	r.entries = slices.Insert(r.entries, i, newEntry)

	return nil
}

func (r *ResolvConf) CreateNameserverAt(checksum, index int, nameserver string) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.read()
	if err != nil {
		return 0, err
	}

	err = r.createNameserverAt(checksum, index, nameserver)
	if err != nil {
		return 0, err
	}

	return r.entries.Checksum(), r.write()
}

func (r *ResolvConf) createNameserverLast(checksum int, nameserver string) error {
	if r.entries.Checksum() != checksum {
		return errors.New("file changed on disk since last read")
	}

	newEntry := entry{
		option: OptionTypeNameserver,
		value:  nameserver,
	}

	r.entries = append(r.entries, newEntry)

	return nil
}

func (r *ResolvConf) CreateNameserverLast(checksum int, nameserver string) (int, int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.read()
	if err != nil {
		return 0, 0, err
	}

	err = r.createNameserverLast(checksum, nameserver)
	if err != nil {
		return 0, 0, err
	}

	return len(r.getNameservers()), r.entries.Checksum(), r.write()
}

func (r *ResolvConf) deleteNameserverAt(checksum, index int) (string, error) {
	if r.entries.Checksum() != checksum {
		return "", errors.New("file changed on disk since last read")
	}

	deleteAt := r.entries.IndexNthNameserver(index)
	if deleteAt < 0 {
		return "", errors.New("index out of range")
	}

	nameserver := r.entries[index].value
	r.entries = append(r.entries[:deleteAt], r.entries[deleteAt+1:]...)

	return nameserver, nil
}

func (r *ResolvConf) DeleteNameserverAt(checksum, index int) (string, int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.read()
	if err != nil {
		return "", 0, err
	}

	nameserver, err := r.deleteNameserverAt(checksum, index)
	if err != nil {
		return "", 0, err
	}

	return nameserver, r.entries.Checksum(), r.write()
}
