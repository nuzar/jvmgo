package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := reader.readUint16()
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ { // i start from 1
		cp[i] = getConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	cpInfo := cp[index]
	if cpInfo != nil {
		return cpInfo
	}
	panic("invalid constant pool index")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8info.str
}
