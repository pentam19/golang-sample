package main

import (
	"fmt"
	"sort"
)

type Item struct {
	id   string
	key1 int
	key2 int
}
type ItemSlice []Item

func (s ItemSlice) Len() int {
	return len(s)
}

func (s ItemSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ItemSlice) Less(i, j int) bool {
	if s[i].key1 == s[j].key1 {
		return s[i].key2 < s[j].key2
	} else {
		return s[i].key1 < s[j].key1
	}
}

func sortSearchItem(items []Item, searchItem Item) (i int, isFound bool) {
	i = sort.Search(len(items), func(i int) bool {
		if items[i].key1 == searchItem.key1 {
			return items[i].key2 >= searchItem.key2
		} else {
			return items[i].key1 >= searchItem.key1
		}
	})
	if i < len(items) &&
		(items[i].key1 == searchItem.key1 && items[i].key2 == searchItem.key2) {
		isFound = true
	} else {
		isFound = false
	}
	return
}

func printResult(isFound bool, items []Item, i int) {
	if isFound {
		fmt.Printf("found!!! [%v]\n", items[i])
	} else {
		fmt.Println("not found")
	}
}

func main() {
	items := []Item{}
	items = append(items, Item{id: "A", key1: 10, key2: 200})
	items = append(items, Item{id: "B", key1: 50, key2: 100})
	items = append(items, Item{id: "C", key1: 10, key2: 100})
	items = append(items, Item{id: "D", key1: 30, key2: 400})
	items = append(items, Item{id: "E", key1: 30, key2: 300})
	items = append(items, Item{id: "F", key1: 10, key2: 300})
	fmt.Println(items)

	var itemSlice ItemSlice
	itemSlice = items
	sort.Sort(itemSlice)
	items = itemSlice
	fmt.Println(items)

	searchItem := Item{key1: 10, key2: 200}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound := sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 11, key2: 201}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 10, key2: 150}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 10, key2: 250}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 30, key2: 400}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 30, key2: 1}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 30, key2: 301}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	searchItem = Item{key1: 1000, key2: 1000}
	fmt.Printf("searchItem --> %v\n", searchItem)
	i, isFound = sortSearchItem(items, searchItem)
	fmt.Printf("returnIndex --> %v\n", i)
	printResult(isFound, items, i)

	// string 比較
	fmt.Println(`"abc" > "abd"`)
	fmt.Println("abc" > "abd")
	fmt.Println(`"abc" > "aba"`)
	fmt.Println("abc" > "aba")
	fmt.Println(`"abc" > "abc"`)
	fmt.Println("abc" > "abc")
	fmt.Println(`"abc" >= "abc"`)
	fmt.Println("abc" >= "abc")
	fmt.Println(`"abc" > "abbzzzzzz"`)
	fmt.Println("abc" > "abbzzzzzz")
	fmt.Println(`"zaa" > "abb"`)
	fmt.Println("zaa" > "abb")
	fmt.Println(`"azzzzzzzzz" > "b"`)
	fmt.Println("azzzzzzzzz" > "b")

	/*
		sort package
		http://golang.jp/pkg/sort

		== RESULT ==
		[{A 10 200} {B 50 100} {C 10 100} {D 30 400} {E 30 300} {F 10 300}]
		[{C 10 100} {A 10 200} {F 10 300} {E 30 300} {D 30 400} {B 50 100}]
		searchItem --> { 10 200}
		returnIndex --> 1
		found!!! [{A 10 200}]
		searchItem --> { 11 201}
		returnIndex --> 3
		not found
		searchItem --> { 10 150}
		returnIndex --> 1
		not found
		searchItem --> { 10 250}
		returnIndex --> 2
		not found
		searchItem --> { 30 400}
		returnIndex --> 4
		found!!! [{D 30 400}]
		searchItem --> { 30 1}
		returnIndex --> 3
		not found
		searchItem --> { 30 301}
		returnIndex --> 4
		not found
		searchItem --> { 1000 1000}
		returnIndex --> 6
		not found

		"abc" > "abd"
		false
		"abc" > "aba"
		true
		"abc" > "abc"
		false
		"abc" >= "abc"
		true
		"abc" > "abbzzzzzz"
		true
		"zaa" > "abb"
		true
		"azzzzzzzzz" > "b"
		false
	*/
}
