package main

import "fmt"
import "reflect"

type User struct {
	UserID string
	Name   string
	Score  int
}

func AbstractSlice(arr interface{}) []interface{} {
	dest := []interface{}{}
	switch sl := arr.(type) {
	case []interface{}:
		// こういうのは無理なんですよ
	case string:
		// rangeかけるためにcase内じゃないとダメ
		for _, b := range sl {
			dest = append(dest, b)
		}
	case []string:
		for _, str := range sl {
			dest = append(dest, str)
		}
	case []int:
		for _, i := range sl {
			dest = append(dest, i)
		}
	case []User:
		for _, user := range sl {
			dest = append(dest, user)
		}
	}
	return dest
}

func main() {

	a := "slice of byte"
	fmt.Println(
		reflect.TypeOf(a),
		reflect.TypeOf(AbstractSlice(a)),
	)

	b := []string{"Hello", "interface"}
	fmt.Println(
		reflect.TypeOf(b),
		reflect.TypeOf(AbstractSlice(b)),
	)
	fmt.Println(AbstractSlice(b))

	c := []int{2, 3, 4, 5}
	fmt.Println(
		reflect.TypeOf(c),
		reflect.TypeOf(AbstractSlice(c)),
	)

	u := []User{{UserID: "test", Name: "test", Score: 123}, {UserID: "test2", Name: "test", Score: 123}}
	fmt.Println(
		reflect.TypeOf(u),
		reflect.TypeOf(AbstractSlice(u)),
	)
	fmt.Println(AbstractSlice(u))

	var users interface{}
	users = AbstractSlice(u)
	fmt.Println(users)
	if _, ok := users.([]User); ok { // ng
		fmt.Println("cast ok")
	} else {
		fmt.Println("cast ng")
	}
	//users2 := &[]User{users} // ng
}
