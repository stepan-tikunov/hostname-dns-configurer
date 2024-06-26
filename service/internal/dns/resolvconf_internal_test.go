package dns

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// region Test parseEntry()

func formatEntry(e *entry) string {
	if e == nil {
		return "nil"
	}

	return fmt.Sprintf(`&entry{option: "%s", value: "%s"}`, e.option, e.value)
}

func expectParseEntryToReturn(t *testing.T, line string, want *entry) {
	got := parseEntry(line)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`parseEntry("%s") -> %s; want %s`, line, formatEntry(got), formatEntry(want))
	}
}

func Test_parseEntry(t *testing.T) {
	type test struct {
		s string
		e *entry
	}
	tests := []test{
		{"", nil},
		{";comment", nil},
		{"nameserver 1.2.3.4", &entry{"nameserver", "1.2.3.4"}},
		{"\tnameserver \t   1.2.3.4\t   ", &entry{"nameserver", "1.2.3.4"}},
		{"nameserver", &entry{"nameserver", ""}},
		{"nameserver ;1.2.3.4", &entry{"nameserver", ""}},
	}

	for _, test := range tests {
		expectParseEntryToReturn(t, test.s, test.e)
	}
}

// endregion

func formatEntries(e entries) string {
	var builder strings.Builder
	builder.WriteString("[ ")
	for _, entry := range e {
		builder.WriteString(formatEntry(&entry))
		builder.WriteString(" ")
	}
	builder.WriteString("]")

	return builder.String()
}

func TestEntries_Checksum(t *testing.T) {
	e := entries{
		{"nameserver", "1.2.3.4"},
	}

	beforeVal := formatEntries(e)
	before := e.Checksum()

	e = append(e, entry{"nameserver", "5.6.7.8"})

	afterAppendVal := formatEntries(e)
	afterAppend := e.Checksum()

	if before == afterAppend {
		t.Errorf("entries.Checksum() -> same hashes after appending entry. before: %s, after: %s",
			beforeVal,
			afterAppendVal,
		)
	}

	e[0].value = "4.3.2.1"

	afterChangeVal := formatEntries(e)
	afterChange := e.Checksum()

	if afterAppend == afterChange {
		t.Errorf("entries.Checksum() -> same hashes after changing entry. before: %s, after: %s",
			afterAppendVal,
			afterChangeVal,
		)
	}

	sameAsBefore := entries{
		{"nameserver", "1.2.3.4"},
	}

	if before != sameAsBefore.Checksum() {
		t.Errorf("entries.Checksum() -> different hashes for same values. first: %s, second: %s",
			beforeVal,
			formatEntries(sameAsBefore),
		)
	}
}

func TestEntries_IndexNthNameserver(t *testing.T) {
	e := entries{}

	got := e.IndexNthNameserver(0)
	want := -1

	if got != want {
		t.Errorf("entries.IndexNthNameserver(0) -> %d, want %d, passed: %s", got, want, formatEntries(e))
	}

	e = append(e,
		entry{"option", "trash"},
		entry{"nameserver", "1.2.3.4"},
		entry{"nameserver", "5.6.7.8"},
		entry{"option", "trash"},
		entry{"option", "trash"},
		entry{"nameserver", "9.10.11.12"},
		entry{"option", "trash"},
		entry{"option", "trash"},
		entry{"option", "trash"},
		entry{"nameserver", "13.14.15.16"},
	)

	tests := []int{0, 1, 2, 3, 4, -1}
	wants := []int{1, 2, 5, 9, -1, -1}

	for i := range len(tests) {
		got = e.IndexNthNameserver(tests[i])
		want = wants[i]

		if got != want {
			t.Errorf("entries.IndexNthNameserver(%d) -> %d, want %d, passed %s", tests[i], got, want, formatEntries(e))
		}
	}
}

func TestResolvConf_getNameservers(t *testing.T) {
	r := ResolvConf{
		path:    "/path/to/resolv.conf",
		entries: nil,
	}

	got := r.getNameservers()
	if len(got) != 0 {
		t.Errorf("ResolvConf.getNameservers() -> %v, want empty, r.entries: %s", got, formatEntries(r.entries))
	}
	r.entries = entries{
		entry{"trash", "1.1.1.1"},
		entry{"trash", "2.2.2.2"},
		entry{"trash", "3.3.3.3"},
		entry{"trash", "4.4.4.4"},
	}

	got = r.getNameservers()
	if len(got) != 0 {
		t.Errorf("ResolvConf.getNameservers() -> %v, want empty, r.entries: %s", got, formatEntries(r.entries))
	}

	r.entries[0].option = "nameserver"
	got = r.getNameservers()
	want := []string{"1.1.1.1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ResolvConf.getNameservers() -> %v, want %v, r.entries: %s", got, want, formatEntries(r.entries))
	}

	r.entries[1].option = "nameserver"
	r.entries[2].option = "nameserver"
	r.entries[3].option = "nameserver"

	got = r.getNameservers()
	want = []string{"1.1.1.1", "2.2.2.2", "3.3.3.3", "4.4.4.4"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ResolvConf.getNameservers() -> %v, want %v, r.entries: %s", got, want, formatEntries(r.entries))
	}
}

func TestResolvConf_createNameserverAt(t *testing.T) {
	r := ResolvConf{
		path:    "/path/to/resolv.conf",
		entries: entries{},
	}

	c := r.entries.Checksum()

	err := r.createNameserverAt(c, 100, "1.1.1.1")
	if err == nil {
		t.Errorf("ResolvConf.createNameserverAt(..., 100, ...) must return error (index out of range), r.entries: %s",
			formatEntries(r.entries),
		)
	}

	err = r.createNameserverAt(c, 0, "1.1.1.1")
	if err != nil {
		t.Errorf("ResolvConf.createNameserverAt(..., 0, ...) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	want := entries{
		{"nameserver", "1.1.1.1"},
	}
	if !reflect.DeepEqual(r.entries, want) {
		t.Errorf("ResolvConf.createNameserverAt(..., 0, ...): got r.entries == %s, want %s",
			formatEntries(r.entries),
			formatEntries(want),
		)
	}

	c = r.entries.Checksum()
	err = r.createNameserverAt(c, 0, "2.2.2.2")
	if err != nil {
		t.Errorf("ResolvConf.createNameserverAt(..., 0, ...) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	want = entries{
		{"nameserver", "2.2.2.2"},
		{"nameserver", "1.1.1.1"},
	}
	if !reflect.DeepEqual(r.entries, want) {
		t.Errorf("ResolvConf.createNameserverAt(..., 0, ...): got r.entries == %s, want %s",
			formatEntries(r.entries),
			formatEntries(want),
		)
	}

	c = r.entries.Checksum()
	err = r.createNameserverAt(c, 1, "3.3.3.3")
	if err != nil {
		t.Errorf("ResolvConf.createNameserverAt(..., 1, ...) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	want = entries{
		{"nameserver", "2.2.2.2"},
		{"nameserver", "3.3.3.3"},
		{"nameserver", "1.1.1.1"},
	}
	if !reflect.DeepEqual(r.entries, want) {
		t.Errorf("ResolvConf.createNameserverAt(..., 1, ...): got r.entries == %s, want %s",
			formatEntries(r.entries),
			formatEntries(want),
		)
	}
}

func TestResolvConf_createNameserverLast(t *testing.T) {
	r := ResolvConf{
		path:    "/path/to/resolv.conf",
		entries: entries{},
	}

	err := r.createNameserverLast("1.1.1.1")
	if err != nil {
		t.Errorf("ResolvConf.createNameserverLast(...) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	want := entries{
		{"nameserver", "1.1.1.1"},
	}
	if !reflect.DeepEqual(r.entries, want) {
		t.Errorf("ResolvConf.createNameserverLast(...): got r.entries == %s, want %s",
			formatEntries(r.entries),
			formatEntries(want),
		)
	}

	err = r.createNameserverLast("2.2.2.2")
	if err != nil {
		t.Errorf("ResolvConf.createNameserverLast(...) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	want = entries{
		{"nameserver", "1.1.1.1"},
		{"nameserver", "2.2.2.2"},
	}
	if !reflect.DeepEqual(r.entries, want) {
		t.Errorf("ResolvConf.createNameserverLast(...): got r.entries == %s, want %s",
			formatEntries(r.entries),
			formatEntries(want),
		)
	}
}

func TestResolvConf_deleteNameserverAt(t *testing.T) {
	r := ResolvConf{
		path:    "/path/to/resolv.conf",
		entries: entries{},
	}

	c := r.entries.Checksum()

	_, err := r.deleteNameserverAt(c, 10)
	if err == nil {
		t.Errorf("ResolvConf.deleteNameserverAt(..., 10) must return error (index out of range), r.entries: %s", r.entries)
	}

	r.entries = append(r.entries, entry{"nameserver", "1.1.1.1"})
	r.entries = append(r.entries, entry{"nameserver", "2.2.2.2"})

	c = r.entries.Checksum()

	deleted, err := r.deleteNameserverAt(c, 0)
	if err != nil {
		t.Errorf("ResolvConf.deleteNameserverAt(..., 0) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	wantDeleted := "1.1.1.1"
	if deleted != wantDeleted {
		t.Errorf(`ResolvConf.deleteNameserverAt(..., 0) deleted "%s", want "%s"`, deleted, wantDeleted)
	}

	wantEntries := entries{
		{"nameserver", "2.2.2.2"},
	}
	if !reflect.DeepEqual(r.entries, wantEntries) {
		t.Errorf("ResolvConf.deleteNameserverAt(..., 0) got r.entries == %s, want %s",
			formatEntries(r.entries),
			formatEntries(wantEntries),
		)
	}

	c = r.entries.Checksum()

	deleted, err = r.deleteNameserverAt(c, 0)
	if err != nil {
		t.Errorf("ResolvConf.deleteNameserverAt(..., 0) must not return error, r.entries: %s, error: %s",
			formatEntries(r.entries),
			err,
		)
	}

	wantDeleted = "2.2.2.2"
	if deleted != wantDeleted {
		t.Errorf(`ResolvConf.deleteNameserverAt(..., 0) deleted "%s", want "%s"`, deleted, wantDeleted)
	}

	if len(r.entries) != 0 {
		t.Errorf("ResolvConf.deleteNameserverAt(..., 0) got r.entries == %s, want empty", formatEntries(r.entries))
	}
}
