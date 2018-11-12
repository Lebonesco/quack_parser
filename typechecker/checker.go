package typechecker

import (
	"fmt"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"errors"
)

type ErrorType string

type CheckError struct {
	Type string
	Message error
}

// error types
const (
	INVALID_SUBCLASS = "INVALID_SUBCLASS"
	CREATE_CLASS_FAIL = "CREATE_CLASS_FAIL"
	CLASS_NOT_EXIST = "CLASS_NOT_EXIST"
	DUPLICATE_CLASS = "DUPLICATE_CLASS"
	CLASS_CYCLE = "CLASS_CYCLE"
)

func createError(errorType, message string) *CheckError {
	return &CheckError{Type: errorType, Message: errors.New(message)}
}

// start of checker
func TypeCheck(node ast.Node, env *Environment) (Object, *CheckError) {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	}
	return Object{}, nil
}

func evalProgram(p *ast.Program, env *Environment) (Object, *CheckError) {
	var result Object

	evalBuiltIns(env)

	for _, class := range p.Classes {
		err := evalClass(&class, env) // injects class into environment
		if err != nil {
			return result, createError(CREATE_CLASS_FAIL, err.Error())
		}
	}

	// check for type loops
	if env.CycleExist() {
		return result, createError(CLASS_CYCLE, "class cycle exists")
	}
	// check for class existence
	if !env.TypesExist() {
		return result, createError(CLASS_NOT_EXIST, "class not exist")
	}

	// for _, statement := range p.Statements {
	// 	// switch result := result.(type) {
	// 	// case ReturnStatement:
	// 	// 	return result.Value
	// 	// case Error:
	// 	// 	return result
	// 	// }
	// }
	return result, nil
} 

	// populate built in classes
func evalBuiltIns(env *Environment) error {
	code := BUIILT_IN_CLASSES
	l := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	program, err := p.Parse(l)
	if err != nil {
		return err
	}

	classes := program.(*ast.Program).Classes
	for _, class := range classes {
		err := evalClass(&class, env)
		if err != nil {
			return err
		}
	}

	return nil
}

func evalClass(class *ast.Class, env *Environment) error {
	newObj := &Object{
				Variables: map[string]ObjectType{}}
	if err := evalClassSignature(class.Signature, newObj, env); err != nil {
		return err
	}

	if err := evalClassBody(class.Body, newObj, env); err != nil {
		return err
	}

	(*env.TypeTable)[newObj.Type] = newObj	// add type to table

	// if extends a parent make sure that variables align and constructor works
	if newObj.Parent != OBJ_CLASS && 
		!compareClassVars(newObj.Parent, newObj.Type, env) {
		return fmt.Errorf("not same variable/types defined")
	}

	return nil
}

func compareClassVars(block1, block2 ObjectType, env *Environment) bool {
	block1Vars := (*env.TypeTable)[block1].Variables
	block2Vars := (*env.TypeTable)[block2].Variables
	return compareBlockVars(block1Vars, block2Vars) // return count so can use for class and not
}

// compares 2 environments
func compareStmtBlockVars(env1, env2 *Environment) bool {
	vars1 := env1.Vals 
	vars2 := env2.Vals
	return compareBlockVars(vars1, vars2)
}

func compareBlockVars(vars1, vars2 map[string]ObjectType) bool {
	for k := range vars1 {
		if val2, ok := vars2[k]; ok && vars1[k] == val2 {
			continue
		} else {
			return false
		}
	}
	return true
}

func evalClassSignature(sig *ast.ClassSignature, newObj *Object, env *Environment) error {
	constructor := []Variable{}

	for _, arg := range sig.Args {
		constructor = append(constructor, Variable{arg.Arg, ObjectType(arg.Type)})
	}

	newObj.Constructor = constructor // add constuctor

	newObj.Parent = ObjectType("Obj");
	if sig.Extend != nil {
		newObj.Parent  = ObjectType(sig.Extend.Parent) // ok if not exist first time around
	}

	newObj.Type = ObjectType(sig.Name)
	// check if class with same name has already been created
	if env.TypeExist(newObj.Type) {
		return fmt.Errorf("class already created with the name %s", sig.Name)
	}

	return nil
}

func evalClassBody(body *ast.ClassBody, newObj *Object, env *Environment) error {
	// store internal variables and methods

	for _, statement := range body.Statements {
		// extract all identifiers
		switch statement.(type) {
		case *ast.LetStatement:
			letStmt := statement.(*ast.LetStatement)
			// if type explicitly declared
			if letStmt.Kind != "" {
				newObj.Variables[letStmt.Name.Value] = ObjectType(letStmt.Kind) // set type
			} else { 
				newObj.Variables[letStmt.Name.Value] = ObjectType("Obj") // default to lowest subtype
			}

			// get type of expression
			switch letStmt.Value.(type) {
			case *ast.Identifier:
				val := letStmt.Value.(*ast.Identifier)
				// check if ident from constructor
				if con, ok := newObj.InConstructor(val.Value); ok {
					// check left side is subtype 
					if !env.ValidSubType(con.Type, newObj.Variables[letStmt.Name.Value]) {
						return fmt.Errorf("whoops, invalid type assignment")
					}
				}
			// case *ast.IntegerLiteral:
			// 	val := letStmt.Value.(*ast.IntegerLiteral)
			// case *ast.StringLiteral:
			// 	val := letStmt.Value.(*ast.StringLiteral)
			// case *ast.Boolean:
			// 	val := letStmt.Value.(*ast.Boolean)
			// case *ast.MethodCall:
			// 	val := letStmt.Value.(*ast.MethodCall)
			}
		}
	}

	return nil
}

func evalBlockStatement(block *ast.BlockStatement, env *Environment) (Object, *CheckError) {
	var result Object

	//for _, statement := range block.Statements {
		// switch result := result.(type) {
		// case *ast.ReturnStatement:
		// 	return result.Value
		// case *ast.Error:
		// 	return result, nil
		// }
	//}
	return result, nil
}

func evalPrefixExpression(expr *ast.PrefixExpression, env *Environment) (Object, *CheckError) {
	// switch operator {
	// case "-":
	// 	return TypeCheck(right)
	// }
	return Object{}, nil
}

func evalOpeExpression(left, right Object) Object {
	return Object{}
}