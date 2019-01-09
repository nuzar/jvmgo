package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (zEntry *ZipEntry) readClass(classname string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(zEntry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == classname {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, zEntry, nil
		}
	}

	return nil, nil, errors.New("class not found: " + classname)
}

func (zEntry *ZipEntry) String() string {
	return zEntry.absPath
}
