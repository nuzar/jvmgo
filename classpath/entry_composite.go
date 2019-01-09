package classpath

import (
	"errors"
	"os"
	"strings"
)

type CompositeEntry struct {
	entries []Entry
}

func newCompositeEntry(pathList string) *CompositeEntry {
	cEntry := &CompositeEntry{}
	for _, path := range strings.Split(pathList, string(os.PathListSeparator)) {
		entry := newEntry(path)
		cEntry.entries = append(cEntry.entries, entry)
	}
	return cEntry
}

func (cEntry *CompositeEntry) String() string {
	strs := make([]string, len(cEntry.entries))
	for i, entry := range cEntry.entries {
		strs[i] = entry.String()
	}
	return strings.Join(strs, string(os.PathListSeparator))
}

func (cEntry *CompositeEntry) readClass(classname string) ([]byte, Entry, error) {
	for _, entry := range cEntry.entries {
		data, from, err := entry.readClass(classname)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + classname)
}

func (cEntry *CompositeEntry) addEntry(e Entry) {
	cEntry.entries = append(cEntry.entries, e)
}
