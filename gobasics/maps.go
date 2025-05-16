package basics

import (
	"fmt"
	"maps"
)

func mapss() {
	// var m map[keyType]valueType
	// m1 = make(map[keyType]valueType)

	myMap := make(map[string]int)

	myMap["key1"] = 9
	myMap["key2"] = 60

	fmt.Println(myMap)
	fmt.Println(myMap["key2"])
	fmt.Println(myMap["code"])
	myMap["key1"] = 70
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key5"] = 10
	myMap["key3"] = 20
	myMap["key4"] = 30

	fmt.Println(myMap)

	// clear(myMap)

	// fmt.Println(myMap)

	_, ok := myMap["key3"]
	if ok{
		fmt.Println("value exists in a map")
	} else{
		fmt.Println("value does not exist in a map")
	}

	myMap1 := map[string]int{"a":1,"b":2}
	fmt.Println(myMap1)

	mymap2 :=map[string]int{"a":1,"b":2}
	if maps.Equal(myMap1,mymap2){
		fmt.Println("both maps contain same keys and values")
	}

	for k,v:=range myMap1{
		fmt.Println(k,v)
	}

	mymap3:= make(map[string]map[string]int)
	mymap3["myMap1"]=mymap2
	fmt.Println(mymap3)
}
