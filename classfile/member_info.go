package classfile

type MemberInfo struct {
	cp               ConstantPool
	accessFlags      uint16
	nameIndex        uint16
	descriptionIndex uint16
	attributes       []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
}
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:               cp,
		accessFlags:      reader.readUint16(),
		nameIndex:        reader.readUint16(),
		descriptionIndex: reader.readUint16(),
		attributes:       readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}

func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptionIndex)
}
