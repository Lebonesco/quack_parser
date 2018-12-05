package main

import (
	"bytes"
	"fmt"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/codegen"
	"github.com/Lebonesco/quack_parser/environment"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/typechecker"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

const DIR = "./samples"

var results = map[string]string{
	"Pt_missing_fields.qk":                         typechecker.CREATE_CLASS_FAIL,
	"circular_dependency.qk":                       typechecker.CLASS_CYCLE,
	"duplicate_class.qk":                           typechecker.DUPLICATE_CLASS,
	"invalid_super.qk":                             typechecker.CLASS_NOT_EXIST,
	"invalid_super_type.qk":                        typechecker.CREATE_CLASS_FAIL,
	"robot.qk":                                     typechecker.CREATE_CLASS_FAIL,
	"Inheritance_Types_bad.qk":                     typechecker.CREATE_CLASS_FAIL,
	"short_test_bad.qk":                            typechecker.INCOMPATABLE_TYPES,
	"bad_init.qk":                                  typechecker.VARIABLE_NOT_INITIALIZED,
	"binop_sugar.qk":                               typechecker.VARIABLE_NOT_INITIALIZED,
	"duplicate_method.qk":                          typechecker.ALREADY_INITIALIZED,
	"init_before_use_bad.qk":                       typechecker.VARIABLE_NOT_INITIALIZED,
	"Plus_types_bad.qk":                            typechecker.INCOMPATABLE_TYPES,
	"simple_inheritingvariables_bad_wrongtype.qk":  typechecker.CREATE_CLASS_FAIL,
	"typing_test.qk":                               typechecker.CLASS_NOT_EXIST,
	"GoodWalk.qk":                                  typechecker.CLASS_NOT_EXIST,
	"LexChallenge.qk":                              typechecker.VARIABLE_NOT_INITIALIZED,
	"unknown_return_type.qk":                       typechecker.CLASS_NOT_EXIST,
	"simple_classes_tree_bad_nosuchsuper.qk":       typechecker.CLASS_NOT_EXIST,
	"simple_inheritingvariables_bad_notdefined.qk": typechecker.CREATE_CLASS_FAIL,
	"simple_classes_tree_bad_alreadydefined.qk":    typechecker.DUPLICATE_CLASS,
	"simple_naming_bad_classandmethodsamename.qk":  typechecker.ALREADY_INITIALIZED,
	"hands.qk":                                     typechecker.METHOD_NOT_EXIST,
	"Comparison_TRUE_FALSE_bad.qk":                 typechecker.INCOMPATABLE_TYPES,
	"Another_plus_types_bad.qk":                    typechecker.VARIABLE_NOT_INITIALIZED,
	"dot_priority.qk":                              typechecker.VARIABLE_NOT_INITIALIZED,
	"method_madness.qk":                            typechecker.VARIABLE_NOT_INITIALIZED,
	"method_madness_2.qk":                          typechecker.VARIABLE_NOT_INITIALIZED,
	"simple_classes_tree_bad_circular.qk":          typechecker.CLASS_CYCLE,
	"not_a_duck.qk":                                typechecker.METHOD_NOT_EXIST,
	"bad_typecase_invalid_type.qk":                 typechecker.CLASS_NOT_EXIST,
	"joseph_test_1.qk":                             typechecker.INVALID_RETURN_TYPE,
	"joseph_test_3.qk":                             typechecker.INVALID_RETURN_TYPE,
	"simple_overridingmethod_bad_numberargs.qk":    typechecker.INCORRECT_ARGUMENT_COUNT,
	"simple_method_return_bad_wrongtype.qk":        typechecker.INVALID_RETURN_TYPE,
	"subclass_method_return_mismatch.qk":           typechecker.INVALID_RETURN_TYPE,
	"TypeWalk.qk":                                  typechecker.METHOD_NOT_EXIST,
	"if_false_init.qk":                             typechecker.VARIABLE_NOT_INITIALIZED,
	"schroedinger.qk":                              typechecker.VARIABLE_NOT_INITIALIZED,
	"joseph_test_6.qk":                             typechecker.METHOD_NOT_EXIST}

var compileError = map[string]string{
	"Another_plus_types_good.qk": "‘struct class_Obj_struct’ has no member named ‘PLUS’"}

func TestFiles(t *testing.T) {
	counter := 0
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		// if file.Name() == "sort.qk" {
		// 	continue
		// }

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
			t.Log(file.Name(), " parse error", err.Error())
			counter += 1
			continue
		}

		program, _ := res.(*ast.Program)

		env := environment.CreateEnvironment() // create new environment
		_, typeErr := typechecker.TypeCheck(program, env)
		if typeErr != nil {
			if val, ok := results[file.Name()]; ok && typeErr.Type == val {
				continue
			}
			t.Errorf(file.Name() + ": " + string(typeErr.Type) + " - " + typeErr.Message.Error())
			counter += 1
		}

		if val, ok := results[file.Name()]; ok && typeErr == nil {
			t.Errorf(file.Name() + ": " + "should be " + string(val))
			counter += 1
		}

		// code generatinon
		code, err := codegen.CodeGen(program)
		check(err)

		f, err := os.Create("./build/" + file.Name() + ".c")
		check(err)

		defer f.Close()
		f.Write(code.Bytes())

		var out bytes.Buffer
		cmd1 := exec.Command("gcc", "-w", "./build/"+file.Name()+".c", "./build/Builtins.c", "./build/Builtins.h")
		cmd1.Stderr = &out
		err = cmd1.Run()
		if len(out.String()) != 0 {
			fmt.Println("error: ", out.String())
			counter += 1
			continue
		}

		cmd := exec.Command("./a.exe")
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err = cmd.Run()

		if err != nil {
			t.Fatalf(err.Error())
		}
		fmt.Println("out:", outb.String(), "error: ", errb.String())

	}
	t.Log(counter)
}
