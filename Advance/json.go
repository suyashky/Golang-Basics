package main

import (
	"encoding/json"
	"fmt"
)

// fields that are needed to be exported,
// starts with capital letter
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"jsonKey1"`
	Fruits []string `json:"jsonKey2"`
}

func main() {

	// working with JSON

	// encoding basic data types to JSON
	// marshal() returns json encoding of given data
	intData := 10
	intData2, _ := json.Marshal(intData)
	fmt.Println(string(intData2))

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// maps and slices can also be encoded to json
	sl1 := []int{2, 3, 4}
	sl2, _ := json.Marshal(sl1)
	fmt.Println(string(sl2))

	// endcoding custom types like struct
	var temp response1
	temp.Page = 1
	temp.Fruits = []string{"apple", "banana", "orange"}
	temp2, _ := json.Marshal(temp)
	fmt.Println(string(temp2))

	// NOTE!-> if Fruits was writeen as fruits,
	// it will not be encoded to json, line 43 will only print page

	// another way to initilize struct
	temp3 := response1{1, []string{"a", "b"}}
	temp4, _ := json.Marshal(temp3)
	fmt.Println(string(temp4))

	// decoding

	// a generic data containing an encoded data
	// similar to what we have in a text file
	byt := []byte(`{"Page":1,"Fruits":["apple","banana","orange"]}`)

	// json will put decoed data into this datype
	// map of string to arbitrary data
	var data map[string]interface{}

	// unmarshall parses the encoded data, returns value pointed by second arg
	// if returns not pointer, returns error
	// if all ok, returns nil
	if err := json.Unmarshal(byt, &data); err != nil { // ($$)
		panic(err)
	}
	fmt.Println(data)

	// ($$)
	// if we have a string instead of byte, write unmarshall([]byte(str), &data)

	// to access the decoded data, convert them to apropriate types
	nums := data["Page"].(float64)
	fmt.Println(nums)

	//for nested data, like array
	arr := data["Fruits"].([]interface{}) // first convert to inteface
	val1 := arr[0].(string)
	val2 := arr[1].(string)
	fmt.Println(val1, val2)

	// dedcoding to custom types like struct
	file := `{"jsonKey1": 45, "jsonKey2": ["apple", "peach"]}` // key can be name anything
	structVar := response2{}

	json.Unmarshal([]byte(file), &structVar)
	fmt.Println(structVar)
	fmt.Println(structVar.Page)
	fmt.Println(structVar.Fruits[1])
}
