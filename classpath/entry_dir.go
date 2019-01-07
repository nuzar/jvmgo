package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (d *DirEntry) readClass(classname string) ([]byte, Entry, error) {
	filename := filepath.Join(d.absDir, classname)
	data, err := ioutil.ReadFile(filename)
	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
