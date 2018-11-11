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

type Object struct {
	MethodTable map[string]MethodSignature // each method method name maps to an array of input types and returns
	parent string
	Type ObjectType
}

func (co *Object) Parent(e *Environment) *Object { 
	res, ok := e.Get(co.parent)
	if !ok {
		return nil
	} 
	tmp := &res
	return tmp
}

func NewObject() *Object { return &Object{} }

func (co *Object) AddMethod(signature *MethodSignature) {
	co.MethodTable[signature.Name] = *signature
}

func (co *Object) ValidSubType(parent string) bool {
	return false
}