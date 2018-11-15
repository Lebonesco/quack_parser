package typechecker

//import "fmt"

// tracks Objects at each layer of scope
type Environment struct {
	Vals      map[string]ObjectType // change to ObjectType?
	Parent    *Environment
	TypeTable *Objects
}

// check if dependency cycle with topological sort
func (e *Environment) CycleExist() bool {
	deps := map[ObjectType][]ObjectType{}
	// generate dependency graph
	for k := range *e.TypeTable {
		if _, ok := deps[k]; !ok {
			deps[k] = []ObjectType{}
		}
		if (*e.TypeTable)[k].Type != "Obj" {
			deps[k] = append(deps[k], (*e.TypeTable)[k].Parent)
		}
	}

	visited := map[ObjectType]bool{}
	visiting := map[ObjectType]bool{}
	for k := range deps {
		if _, ok := visited[k]; !ok { // if hasn't been visited check for cycle
			if cycle(k, deps, visited, visiting) {
				return true
			}
		}
	}
	return false
}

// used topological sort to check for cycles
func cycle(cur ObjectType, deps map[ObjectType][]ObjectType, visited, visiting map[ObjectType]bool) bool {
	visiting[cur] = true
	for _, dep := range deps[cur] {
		if _, ok := visiting[dep]; ok { // if currently visiting node, cycle found
			return true
		}

		if _, ok := visited[dep]; ok { // if already visited don't need to check further
			continue
		}

		if cycle(dep, deps, visited, visiting) {
			return true
		}

	}
	delete(visiting, cur)
	visited[cur] = true
	return false
}

// makes sure that all extended types exist
func (e *Environment) TypesExist() bool {
	// check that all parents have types
	for k := range *e.TypeTable {
		parent := (*e.TypeTable)[k].Parent;
		if ok := e.TypeExist(parent); !ok && parent != "Obj" {
			return false
		}
	}
	return true
}

// returns new environment scope
func (e *Environment) NewScope() *Environment {
	newEnv := CreateEnvironment()
	newEnv.Parent = e
	return newEnv;
}

// create new scope
func CreateEnvironment() *Environment {
	v := make(map[string]ObjectType)
	return &Environment{Vals: v, Parent: nil, TypeTable: &Objects{}}
}

// set item in current scope
func (e *Environment) Set(name string, val ObjectType) {
	e.Vals[name] = val
}

// get item in shortest scope
func (e *Environment) Get(name string) (ObjectType, bool) {
	obj, ok := e.Vals[name]
	if !ok && e.Parent != nil {
		obj, ok = e.Parent.Get(name)
	}
	return obj, ok
}

// check if type already exists
func (e *Environment) TypeExist(name ObjectType) bool {
	parent := e
	for parent != nil {
		if _, ok := (*parent.TypeTable)[name]; ok {
			return true
		}
		parent = parent.Parent // check next scope up
	}
	return false
}

// checks if sub is subtype of parent
func (e *Environment) ValidSubType(sub, parent ObjectType) bool {
	if parent == "Obj" { // supertype for everything
		return true
	}
	next := sub
	for next != parent {
	  //  fmt.Println(next, parent)
		if next == "Obj" && parent != "Obj" {
			return false
		}

		next = (*e.TypeTable)[(*e.TypeTable)[next].Parent].Type
	}
	return true
}

func (e *Environment) GetClass(class ObjectType) *Object {
	scope := e 
	for scope != nil {
		if val, ok := (*scope.TypeTable)[class]; ok {
			return val
		}
		scope = scope.Parent
	}
	return nil
}