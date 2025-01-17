package ast

type FileKind int

const (
	Program FileKind = 0
	Unit    FileKind = 1
	Library FileKind = 2
	Package FileKind = 3
)

func (k FileKind) String() string {
	if k == Program {
		return "program"
	} else if k == Unit {
		return "unit"
	} else if k == Library {
		return "library"
	} else if k == Package {
		return "package"
	}
	return "unknown FileKind"
}

type File struct {
	Kind     FileKind
	Name     string
	Sections []*FileSection
}

type FileSection struct {
	Kind   FileSectionKind
	Uses   []string
	Blocks []FileSectionBlock
}

type FileSectionKind int

const (
	InterfaceSection      FileSectionKind = 0
	ImplementationSection FileSectionKind = 1
	InitializationSection FileSectionKind = 2
	FinalizationSection   FileSectionKind = 3
)

func (k FileSectionKind) String() string {
	if k == InterfaceSection {
		return "interface"
	} else if k == ImplementationSection {
		return "implementation"
	} else if k == InitializationSection {
		return "initialization"
	} else if k == FinalizationSection {
		return "finalization"
	}
	return "unknown FileSectionKind"
}

type FileSectionBlock interface {
	isFileSectionBlock()
}
