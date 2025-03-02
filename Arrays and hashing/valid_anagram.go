package arrays_and_hashing

import "reflect"

func isAnagram(s string, t string) bool {
	mapS := createMap(s)
	mapT := createMap(t)
	eq := reflect.DeepEqual(mapS, mapT)
	return eq

}

func createMap(s string) map[string]int {
	myMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		char := string(s[i]) // Convert byte to string
		myMap[char]++
	}
	return myMap
}
