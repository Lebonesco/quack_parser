package typechecker

//import "fmt"

// tracks Objects at each layer of scope
type Environment struct {
	Vals      map[string]Variable
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

func (e *Environment) TypesExist() bool {
	// check that all parents have types
	for k := range *e.TypeTable {
		parent := (*e.TypeTable)[k].Parent;
		if _, ok := (*e.TypeTable)[parent]; !ok && parent != "Obj" {
			return false
		}
	}
	return true
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
