package typechecker

// tracks Objects at each layer of scope
type Environment struct {
	Vals   map[string]Object
	Parent *Environment
}

func (e *Environment) CreateEnvironment() *Environment {
	v := make(map[string]Object)
	return &Environment{Vals: v, Parent: nil}
}

func (e *Environment) Set(name string, val Object) {
	e.Vals[name] = val
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.Vals[name]
	if !ok && e.Parent != nil {
		obj, ok = e.Parent.Get(name)
	}
	return obj, ok
}