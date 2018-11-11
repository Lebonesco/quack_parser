package typechecker

// import "fmt"

// tracks Objects at each layer of scope
type Environment struct {
	Vals   map[string]Variable
	Parent *Environment
	TypeTable *Objects
}

// check if dependency cycle with topological sort
func (e *Environment) CycleExist() bool {
	deps := map[ObjectType][]ObjectType{}

	// generate dependency graph
	for k := range (*e.TypeTable) {
		if _, ok := deps[k]; !ok {
			deps[k] = []ObjectType{}
		}
		if (*e.TypeTable)[k].Type != "Obj" {
			deps[k] = append(deps[k], (*e.TypeTable)[k].Parent)
		}
	}
	return false
}

// create new scope
func CreateEnvironment() *Environment {
	v := make(map[string]Variable)
	return &Environment{Vals: v, Parent: nil, TypeTable: &Objects{}}
}

// set item in current scope
func (e *Environment) Set(name string, val Variable) {
	e.Vals[name] = val
}

// get item in shortest scope
func (e *Environment) Get(name string) (Variable, bool) {
	obj, ok := e.Vals[name]
	if !ok && e.Parent != nil {
		obj, ok = e.Parent.Get(name)
	}
	return obj, ok
}

// check if type already exists
func (e *Environment) TypeExist(name ObjectType) bool {
	if _, ok := (*e.TypeTable)[name]; ok {
		return true
	}
	return false
}