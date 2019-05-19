package main

import (
	"fmt"
	"os"

	"gopkg.in/go-playground/validator.v8"
)

// Preparation
//  export GO111MODULE=on
//  go mod init project-name

type User struct {
	ID    int64
	Email string `validate:"required,email"`
	Name  string `validate:"required"`
}

func main() {

	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)

	user := &User{ID: 1, Email: "hoge.com"}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		// Err Type
		switch e := err.(type) {
		case *os.PathError:
			fmt.Println("case: *os.PathError")
			fmt.Printf("type: %T\n", e)
		case validator.ValidationErrors:
			fmt.Println("case: validator.ValidationErrors")
			fmt.Printf("type: %T\n", e)
		default:
			fmt.Println("case: default")
			fmt.Printf("type: %T\n", e)
		}
	}
}

/*
Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag
Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
case: validator.ValidationErrors
type: validator.ValidationErrors
*/
