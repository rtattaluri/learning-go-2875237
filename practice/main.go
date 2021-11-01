package main

import (
	"fmt"
	"sort"
)

func main() {
	maps := make(map[string]string)
	fmt.Println("Maps: ", maps)
	maps2 := 

	maps["TG"] = "Telengana"
	maps["AP"] = "Andhra Pradesh"
	maps["KA"] = "Karnataka"
	maps["OD"] = "Odisha"
	fmt.Println("Maps: ", maps)

	delete(maps, "OD")
	maps["KL"] = "Kerala"
	fmt.Println("Maps: ", maps)

	for k, v := range maps {
		fmt.Printf("%s:%s\n", k, v)
	}
	var keys []string
	for k := range maps {
		keys = append(keys, k)
	}
	fmt.Println("Keys: ", keys)
	sort.Strings(keys)
	fmt.Println("Sorted Keys: ", keys)
	fmt.Println("Sorted Maps: ")
	for k := range keys {
		fmt.Println(keys[k], ":", maps[keys[k]])
	}

}
