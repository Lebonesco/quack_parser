package typechecker

// Variable Type 
type ObjectType string

const (
	OBJ_CLASS = "Obj"
	STRING_CLASS = "String"
	INTEGER_CLASS = "Int"
	BOOL_CLASS = "Boolean"
)

type MethodSignature struct {
	Name string
	Parameters []ObjectType
	Returns []ObjectType
}

// handles tracking of Type Hierarchy
type Object struct {
	MethodTable map[string]MethodSignature // each method name maps to an array of input types and returns
	Variables map[string]ObjectType // variables initialized in constructor
	Constructor []Variable // list of constructor arguments
	Parent ObjectType// parent type name
	Type ObjectType // Object Type
}

// mapping of Object Types
type Objects map[ObjectType]*Object

// every variable
type Variable struct {
	Name string
	Type ObjectType
}

func (o *Object) AddMethod(signature *MethodSignature) {
	o.MethodTable[signature.Name] = *signature
}

func (o *Object) InConstructor(ident string) (*Variable, bool) {
	for _, val := range o.Constructor {
		if ident == val.Name {
			return &val, true
		}
	}
	return nil, false 
}