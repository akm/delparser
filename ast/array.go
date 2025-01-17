package ast

// https://docwiki.embarcadero.com/RADStudio/Alexandria/ja/%E6%A7%8B%E9%80%A0%E5%8C%96%E5%9E%8B%EF%BC%88Delphi%EF%BC%89#.E9.9D.99.E7.9A.84.E9.85.8D.E5.88.97

func NewArray(name, typ string, indexTypes []IndexType) *Type {
	return &Type{
		Name: name,
		Expr: &ArrayExpr{
			Type:       typ,
			IndexTypes: indexTypes,
		},
	}
}

func (*ArrayExpr) isVarType()  {}
func (*ArrayExpr) isTypeExpr() {}

type ArrayExpr struct {
	Dynamic    bool
	IndexTypes []IndexType
	Type       string
}

type IndexType interface {
	isIndexType()
}

func (*NumRange) isIndexType() {}

type NumRange struct {
	Packed bool
	Low    int
	High   int
}

func (*NamedIndexType) isIndexType() {}

type NamedIndexType struct {
	Packed bool
	Name   string
}
