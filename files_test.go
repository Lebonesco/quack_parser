package main

import (
	"fmt"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/typechecker"
	"io/ioutil"
	"testing"
)

const DIR = "./samples"


var results = map[string]string{
	"Pt_missing_fields.qk": typechecker.CREATE_CLASS_FAIL,
	"SqrDecl.qk": typechecker.CREATE_CLASS_FAIL,
	"SqrDeclEQ.qk": typechecker.CREATE_CLASS_FAIL,
	"circular_dependency.qk": typechecker.CLASS_CYCLE,
	"duplicate_class.qk": typechecker.CREATE_CLASS_FAIL,
	"invalid_super.qk": typechecker.CLASS_NOT_EXIST,
	"invalid_super_type.qk": typechecker.CREATE_CLASS_FAIL,
	"robot.qk": typechecker.CREATE_CLASS_FAIL,
	"Inheritance_Types_bad.qk": typechecker.CREATE_CLASS_FAIL,	
	"short_test_bad.qk": typechecker.INVALID_OPERATION_TYPE,
	"Another_plus_types_bad.qk": typechecker.VARIABLE_NOT_INITIALIZED,
	"Another_plus_types_good.qk": typechecker.VARIABLE_NOT_INITIALIZED,
	"bad_init.qk": typechecker.VARIABLE_NOT_INITIALIZED,
	"binop_sugar.qk": typechecker.VARIABLE_NOT_INITIALIZED,
	"duplicate_method.qk": typechecker.ALREADY_INITIALIZED,
	"init_before_use_bad.qk": typechecker.VARIABLE_NOT_INITIALIZED,
	"Plus_types_bad.qk": typechecker.INVALID_CONSTRUCTOR_TYPE,
	"simple_inheritingvariables_bad_wrongtype.qk": typechecker.CREATE_CLASS_FAIL,
	"typing_test.qk": typechecker.CREATE_CLASS_FAIL,
	"GoodWalk.qk": typechecker.CLASS_NOT_EXIST}


func TestFiles(t *testing.T) {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		fmt.Printf("Testing file %d/%d - %s\n", i+1, len(files), file.Name())
		data, err := ioutil.ReadFile(DIR + "/" + file.Name())
		if err != nil {
			t.Fatalf(err.Error())
			continue
		}

		l := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		res, err := p.Parse(l)
		if err != nil {
			//t.Log("\n------------------------------------------------------------------")
			t.Log(file.Name(), " parse error")
		}

		program, _ := res.(*ast.Program)

		env := typechecker.CreateEnvironment() // create new environment
		_ , typeErr := typechecker.TypeCheck(program, env)
		if typeErr != nil {
			if val, ok := results[file.Name()]; ok && typeErr.Type == val {
				//t.Log(typeErr.Type)
				continue
			}
			t.Errorf(file.Name() + ": " +  string(typeErr.Type) + " - " + typeErr.Message.Error())
		}

		if val, ok := results[file.Name()]; ok && typeErr == nil {
				t.Errorf(file.Name() + ": " + "should be " + string(val))
		}
	}
}
