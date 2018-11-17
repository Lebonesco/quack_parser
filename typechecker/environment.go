package typechecker

//import "fmt"

// tracks Objects at each layer of scope
type Environment struct {
	Vals      map[string]ObjectType // change to ObjectType?
	Parent    *Environment
	TypeTable *Objects
	Class ObjectType
}

// returns union of variables with same types between two environments
func GetUnion(e1, e2 *Environment) map[string]ObjectType {
	result := map[string]ObjectType{}
	for k, val1 := range e1.Vals {
		if val2, ok := e2.Vals[k]; ok { // if both blocks have the variable
			result[k] = e1.GetLowestCommonType(val1, val2)
		}
	}
	return result
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
	newEnv.Class = e.Class // if in class pass down
	return newEnv;
}

// create new scope
func CreateEnvironment() *Environment {
	v := make(map[string]ObjectType)
	return &Environment{Vals: v, Parent: nil, TypeTable: &Objects{}, Class: NOTHING_CLASS}
}

func (e *Environment) SetClass(class ObjectType) {
	e.Class = class
}

func (e *Environment) GetClassObject() *Object {
	return e.GetClass(e.Class)
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
func (e *Environment) ValidSubType(sub, parent ObjectType) (bool) {
	if parent == OBJ_CLASS { // supertype for everything
		return true
	}
	next := sub
	for next != parent {
		if next == OBJ_CLASS && parent != OBJ_CLASS {
			return false
		}

		tmp := e.GetClass(next)
		if tmp == nil {
			return false
		}
		tmp = e.GetClass(tmp.Parent)
		if tmp == nil {
			return false
		}
		next = tmp.Type
	}
	return true
}

// brute force solution to get type
func (e *Environment) GetLowestCommonType(val1, val2 ObjectType) (ObjectType) {
	if val1 == OBJ_CLASS || val2 == OBJ_CLASS {
		return OBJ_CLASS
	}

	objects := []*Object{} // store type pointers
	
	for val1 != OBJ_CLASS { // traverse all the way up the lattice of types
		objects = append(objects, e.GetClass(val1))
		val1 = e.GetParentType(val1)
	}

	for val2 != OBJ_CLASS {
		for _, obj := range objects {
			if obj == e.GetClass(val2) {
				return obj.Type
			}
		}
		val2 = e.GetParentType(val2) // go one type up
	}

	return OBJ_CLASS
}

func (e *Environment) GetParentType(kind ObjectType) (ObjectType) {
	if kind == OBJ_CLASS || kind == NOTHING_CLASS {
		return kind
	}

	return e.GetClass(e.GetClass(kind).Parent).Type
}

func (e *Environment) GetClassVariable(class ObjectType, variable string) (ObjectType, bool) {
	obj := e.GetClass(class)
	if sig, ok := obj.Variables[variable]; ok {
		return sig, ok
	}
	return "", false	
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

func (e *Environment) GetClassMethod(class ObjectType, method string) (MethodSignature, bool) {
	obj := e.GetClass(class)
	for {
		if sig, ok := obj.MethodTable[method]; ok {
			return sig, ok
		}

		if obj.Type == OBJ_CLASS {
			break
		}

		obj = e.GetParent(obj.Parent)
		if obj == nil {
			break
		}

	}
	return MethodSignature{}, false	
}

// fix this to actually return parent Object
func (e *Environment) GetParent(parent ObjectType) (*Object) {
	return e.GetClass(parent)
}