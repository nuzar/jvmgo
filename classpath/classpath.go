package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}

func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/ext/*
	jreExtPath := filepath.Join(jreDir, "ext", "*")
	cp.extClasspath = newWildcardEntry(jreExtPath)

}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		return javaHome
	}

	panic("Cannot find jre folder")
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

func (cp *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	classname = classname + ".class"
	if data, entry, err := cp.bootClasspath.readClass(classname); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.extClasspath.readClass(classname); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(classname)
}

func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}
