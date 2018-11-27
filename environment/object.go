package environment

// Variable Type
type ObjectType string

// built in class types
const (
	OBJ_CLASS        = "Obj"
	STRING_CLASS     = "String"
	INTEGER_CLASS    = "Int"
	BOOL_CLASS       = "Boolean"
	RETURN_VALUE_OBJ = "RETURN_VALUE_OBJ"
	NOTHING_CLASS    = "NOTHING_CLASS"
	TYPE_HOLDER      = "$TYPE_HOLDER" // represents unknown type
)

// built in methods
const (
	PLUS    = "PLUS"
	EQUALS  = "EQUALS"
	ATMOST  = "ATMOST"
	ATLEAST = "ATLEAST"
	LESS    = "LESS"
	MORE    = "MORE"
	MINUS   = "MINUS"
	DIVIDE  = "DIVIDE"
	TIMES   = "TIMES"
	AND     = "AND"
	OR      = "OR"
)

type MethodSignature struct {
	Name string
	Params []Variable
	Return ObjectType
}

// handles tracking of Type Hierarchy
type Object struct {
	MethodTable []MethodSignature // each method name maps to an array of input types and returns
	Variables   map[string]ObjectType      // variables initialized in constructor
	Constructor []Variable                 // list of constructor arguments
	Parent      ObjectType                 // parent type name
	Type        ObjectType                 // Object Type
}

// mapping of Object Types
type Objects map[ObjectType]*Object

// every variable
type Variable struct {
	Name string
	Type ObjectType
}

// init Object with default fields
func NewObject() *Object {
	return &Object{
		Variables:   map[string]ObjectType{},
		Parent:      ObjectType(OBJ_CLASS),
		MethodTable: []MethodSignature{},
		Constructor: []Variable{}}
}

func (o *Object) AddMethod(name string, signature MethodSignature) {
	signature.Name = name
	o.MethodTable = append(o.MethodTable, signature) 
}

// recursive checks for inherited method
func (o *Object) GetMethod(name string) (MethodSignature, bool) {
	for _, method := range o.MethodTable {
		if method.Name == name {
			return method, true
		}
	}

	return MethodSignature{}, false
}

func (o *Object) InConstructor(ident string) (*Variable, bool) {
	for _, val := range o.Constructor {
		if ident == val.Name {
			return &val, true
		}
	}
	return nil, false
}
