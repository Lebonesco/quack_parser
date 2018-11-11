package typechecker

import (
	"fmt"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
)

// start of checker
func TypeCheck(node ast.Node, env *Environment) (Object, error) {
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

func evalProgram(p *ast.Program, env *Environment) (Object, error) {
	var result Object

	evalBuiltIns(env)

	for _, class := range p.Classes {
		err := evalClass(&class, env)
		if err != nil {
			return result, fmt.Errorf("Error checking class: %s", err.Error())
		}
	}

	// check for existence 
	if env.CycleExist() {
		return result, fmt.Errorf("Error class cycle exists")
	}
	// check for loops

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

func evalClass(class *ast.Class, env *Environment) (error) {
	newObj := &Object{}
	if err := evalClassSignature(class.Signature, newObj, env); err != nil {
		return err
	}

	if err := evalClassBody(class.Body, newObj); err != nil {
		return err
	}

	(*env.TypeTable)[newObj.Type] = newObj	
	return nil
}

func evalClassSignature(sig *ast.ClassSignature, newObj *Object, env *Environment) (error) {
	constructor := []ObjectType{}

	for _, arg := range sig.Args {
		constructor = append(constructor, ObjectType(arg.Type))
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

func evalClassBody(body *ast.ClassBody, newObj *Object) (error) {
	// store internal variables and methods

	return nil
}

func evalBlockStatement(block *ast.BlockStatement, env *Environment) (Object, error) {
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

func evalPrefixExpression(expr *ast.PrefixExpression, env *Environment) (Object, error) {
	// switch operator {
	// case "-":
	// 	return TypeCheck(right)
	// }
	return Object{}, nil
}

func evalOpeExpression(left, right Object) Object {
	return Object{}
}