package classfile

import "fmt"

/*
GO           JAVA
-----------------
int8         byte
uint8(byte)  N/A     u1
int16        short
uint16       char    u2
int32(rune)  int
uint32       N/A     u4
int64        long
uint64       N/A
float32      float
float64      double
*/

type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttriuteInfo
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstanPool(reader)
	cf.accessFlags = reader.readUint16()
}
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()

	if cf.majorVersion == 45 {
		// J2SE 1.2
		return
	} else if cf.minorVersion == 0 && cf.majorVersion >= 46 && cf.majorVersion <= 52 {
		// J2SE 1.2 ~ Java SE 8
		return
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func Parse(classData []byte) (*ClassFile, error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf := &ClassFile
	cf.read(cr)
	return cf, nil // TODO: error
}

func (cf *ClassFile) MajorVersion(uint16) {
	return cf.majorVersion
}
func (cf *ClassFile) MinorVersion(uint16) {
	return cf.minorVersion
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) InterfaceNames() []string {
	names := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		names[i] = cf.constantPool.getClassName(cpIndex)
	}
	return names
}
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}
func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return "" // java.lang.Object 没有父类
}
