package typechecker

type ObjectType string

const (
	STRING_OBJ = "STRING_OBJ"
	INTEGER_OBJ = "INTEGER_OBJ"
	BOOL_OBJ = "BOOL_OBJ"
	FUNCTION_OBJ = "FUNCTION_OBJ"
	CLASS_OBJ = "CLASS_OBJ"
)

type MethodSignature struct {
	Name string
	Parameters []ObjectType
	Returns []ObjectType
}

// handles tracking of Type Hierarchy
type Object struct {
	MethodTable map[string]MethodSignature // each method name maps to an array of input types and returns
	Constructor []ObjectType // list of constructor arguments
	Parent ObjectType// parent type name
	Type ObjectType // Object Type
}

// mapping of Object Types
type Objects map[ObjectType]*Object

// every variable
type Variable struct {
	Name string
	Type *Object
}

func NewObject() *Object { return &Object{} }

func (o *Object) AddMethod(signature *MethodSignature) {
	o.MethodTable[signature.Name] = *signature
}

// func (o *Object) ValidSubType(parent string) bool {
// 	search := o.Type
// 	for search != parent {
// 		search := o.
// 	}
// 	return true
// }