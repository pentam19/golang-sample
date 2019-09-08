package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antonholmquist/jason"
)

// JSONの中から特定のキー名を持つ値を書き換える
func main() {
	// string
	jsonStr := `
{\"key1\":12345, \"key2\":\"abc\", \"_key3\":\"key3val\", \"key4\":\"xyz\", 
\"_key5\":\"key5val\", \"_key6\":\"key6val\", \"key7\":\"12345a\"}`

	searchStr := jsonStr
	searchFullStr := jsonStr
	forwardStr := ""
	for {
		idxKeyS := strings.Index(searchStr, "\"_")
		if idxKeyS < 0 {
			break
		}

		forwardStr = forwardStr + searchStr[:idxKeyS-1]
		fmt.Printf("forwardStr: %s\n", forwardStr)

		searchStr = searchStr[idxKeyS+1:]
		fmt.Printf("searchStr: %s\n", searchStr)
		fmt.Printf("sIdx: %d\n", idxKeyS)

		fmt.Printf("searchStr: %s\n", searchStr)
		idxKeyE := strings.Index(searchStr, "\"")

		keyName := searchStr[:idxKeyE]
		fmt.Printf("keyName: %s\n", keyName)

		fmt.Printf("eIdx: %d\n", idxKeyE)

		idxColon := strings.Index(searchStr, ":")
		searchStr = searchStr[idxColon+1:]
		fmt.Printf("searchStr: %s\n", searchStr)

		idxValS := strings.Index(searchStr, "\"")
		searchStr = searchStr[idxValS+1:]
		idxValE := strings.Index(searchStr, "\"")
		valueStr := searchStr[:idxValE]

		newValue := replaceID(keyName, valueStr)
		fmt.Printf("oldValue: %s\n", valueStr)
		fmt.Printf("newValue: %s\n", newValue)

		searchStr = searchStr[idxValE+1:]
		fmt.Printf("searchStr: %s\n", searchStr)
		forwardStr = forwardStr + "\"" + keyName + "\":\"" + newValue + "\""
		fmt.Printf("forwardStr: %s\n", forwardStr)
	}
	fmt.Printf("==============\n")
	fmt.Printf("searchFullStr: %s\n", searchFullStr)
	fmt.Printf("resultFullStr: %s\n", forwardStr+searchStr)
	fmt.Printf("==============\n")

	// jason
	fmt.Printf("==============\n")
	jsonStr2 := "{\"key1\":12345, \"key2\":\"abc\", \"_key3\":\"key3val\", \"key4\":\"xyz\", \"_key5\":\"key5val\", \"_key6\":\"key6val\", \"key7\":\"12345a\"}"
	fmt.Printf("%d\n", strings.Index(jsonStr, "_key"))

	v, err := jason.NewObjectFromBytes([]byte(jsonStr2))
	if err != nil {
		fmt.Printf("jason err: %v\n", err)
	} else {
		key3val, _ := v.GetString("_key3")
		fmt.Printf("_key3: %s\n", key3val)
		jsonMap := v.Map()
		key3data := jsonMap["_key3"]
		fmt.Printf("Map _key3 %+v\n", key3data)
		fmt.Printf("Map: %+v\n", jsonMap)

		for k, v := range jsonMap {
			if strings.HasPrefix(k, "_") {
				fmt.Printf("Target key: %+v, val: %+v\n", k, v)
			}
		}
	}
	fmt.Printf("==============\n")

	// interface
	fmt.Printf("==============\n")
	var jsonBlob = []byte("{\"key1\":12345, \"key2\":\"abc\", \"_key3\":\"key3val\", \"key4\":\"xyz\", \"_key5\":\"key5val\", \"_key6\":\"key6val\", \"key7\":{ \"key8\":\"12345a\", \"_key3\":\"key3val2\"}}")
	var jsonIFobj interface{}
	err = json.Unmarshal(jsonBlob, &jsonIFobj)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v\n", jsonIFobj)

	fmt.Printf("_key3val: %+v\n", jsonIFobj.(map[string]interface{})["_key3"])

	fmt.Println(string(jsonBlob))

	// replace
	valueReplace(jsonIFobj)
	jsonBytes, err := json.Marshal(jsonIFobj)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	fmt.Println(string(jsonBytes))
	fmt.Printf("==============\n")
}

func replaceID(keyName string, src string) (dst string) {
	switch {
	case strings.Contains(keyName, "key3"):
		dst = src + "-mod"
	case strings.Contains(keyName, "key5"):
		dst = src + "-mod"
	default:
		dst = "notFoundKeyValue"
	}
	return
}

func valueReplace(jsonIFobj interface{}) {
	for k, v := range jsonIFobj.(map[string]interface{}) {
		switch valtype := v.(type) {
		case int:
			fmt.Printf("[%s] type of value int: %T\n", k, valtype)
		case string:
			fmt.Printf("[%s] type of value string: %T\n", k, valtype)
			if strings.HasPrefix(k, "_") {
				fmt.Printf("Target key: %+v, val: %+v\n", k, v)
				modval := replaceID(k, v.(string))
				jsonIFobj.(map[string]interface{})[k] = modval
			}
		case map[string]interface{}:
			fmt.Printf("[%s] type of value map[string]interface{}: %T\n", k, valtype)
			valueReplace(v)
		default:
			fmt.Printf("[WARN] [%s] Can't handle this type of value : %T\n", k, valtype)
		}
	}
}
