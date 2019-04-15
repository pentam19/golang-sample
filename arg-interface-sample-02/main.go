package main

import (
	"fmt"
	"strings"
)

type User struct {
	UserID string
	Name   string
	Score  int
}

func printIf(src interface{}) {
	switch value := src.(type) {
	case int:
		fmt.Printf("parameter is integer. [value: %d]\n", value)
	case string:
		value = strings.ToUpper(value) // 対象がstring型なのでstringを引数に取る関数が実行できる
		fmt.Printf("parameter is string. [value: %s]\n", value)
	case []string:
		value = append(value, "<不明>") // 対象がsliceなのでAppendができる
		fmt.Printf("parameter is slice string. [value: %s]\n", value)
	case []User:
		value = append(value, User{UserID: "test", Name: "test", Score: 123})
		fmt.Println(value)
	case *User:
		value = &User{UserID: "test", Name: "test", Score: 123}
		fmt.Println(*value)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", src)
	}
}

func main() {
	var users []User
	printIf(users)
	printIf(&User{})
}
