package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"reflect"
	"struct-various-sample/model"
)

func createInstance(typ reflect.Type) (val reflect.Value) {
	val = reflect.New(typ).Elem()
	return
}

// export GO111MODULE=on
//
func main() {

	/*
		Create Instance from Struct Map SAMPLE
	*/
	// make map
	var structCache = make(map[string]reflect.Type)
	structCache["User"] = reflect.TypeOf(model.User{})
	// get reflect type
	t, ok := structCache["User"]
	if !ok {
		log.Fatal("err")
	}
	// create instance
	r := createInstance(t)
	// set field val
	r.FieldByName("UserID").SetString("12345")
	r.FieldByName("Name").SetString("name")
	r.FieldByName("Age").SetInt(30)
	fmt.Println(r)

	/*
		go/parser SAMPLE
	*/
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "model/user.go", nil, 0)
	if err != nil {
		panic(err)
	}
	for _, decl := range file.Decls {
		switch td := decl.(type) {
		case *ast.GenDecl:
			switch td.Tok {
			case token.TYPE:
				fmt.Println("### type")
				for _, sp := range td.Specs {
					s := sp.(*ast.TypeSpec)
					fmt.Println(s.Name)
					fmt.Println(s.Type)

					switch t := s.Type.(type) {
					case *ast.StructType:
						fmt.Println(t.Struct)
						fmt.Println(t.Incomplete)
						for _, f := range t.Fields.List {
							fmt.Println(f)
						}
					default:
						fmt.Println(3, t)
					}
				}
			default:

			}
		default:
		}
		fmt.Println()
	}

	/*
		types.Package SAMPLE
	*/
	conf := types.Config{
		Importer: importer.Default(),
		Error: func(err error) {
			fmt.Printf("!!! %#v\n", err)
		},
	}
	if err != nil {
		log.Fatal(err)
	}
	pkg, err := conf.Check("model", fset, []*ast.File{file}, nil)
	if err != nil {
		log.Fatal(err)
	}

	S := pkg.Scope().Lookup("User")
	internal := S.Type().Underlying().(*types.Struct)

	fmt.Println(S.Type())
	fmt.Println(internal.String())
	fmt.Println(internal.Underlying())

	for i := 0; i < internal.NumFields(); i++ {
		jsonname, found := reflect.StructTag(internal.Tag(i)).Lookup("json")
		field := internal.Field(i)
		fmt.Printf("%v (exported=%t, jsonname=%s, found=%t)\n", field, field.Exported(), jsonname, found)
	}
}
