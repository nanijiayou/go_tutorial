package mapType

import "fmt"

func printMap(mapName map[string]string) {
	for key, value := range mapName {
		fmt.Println("key = ", key, "value = ", value)
	}
}

func TestMap() {
	fmt.Println("============= map ===============")
	var map1 map[string]string

	if map1 == nil {
		fmt.Println("mp1 is empty")
	}

	// make(map[string]string, 10)
	map1 = make(map[string]string, 10)
	map1["one"] = "one"
	map1["two"] = "two"

	fmt.Println(map1)

	map2 := make(map[string]int)
	map2["one"] = 1
	map2["two"] = 2
	fmt.Println(map2)

	map3 := map[string]string{
		"one": "c++",
		"two": "java",
	}
	printMap(map3)

	for key, value := range map3 {
		fmt.Println("key = ", key, "value = ", value)
	}

	map3["one"] = "python"
	printMap(map3)

	delete(map3, "one")
	printMap(map3)
}
