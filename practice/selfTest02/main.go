// Level 1: Hash
// Create a function that takes a slice of strings as input and returns a hash map. The keys in the hash map should be the distinct strings in the input slice, and the corresponding values should be the counts of the occurrences of those strings in the input.

// Level 2: Hash + Sorting
// Modify the function from Level 1 to sort the keys of the hash map in descending order of their corresponding counts. If two keys have the same count, they should appear in lexicographical order. The function should return a slice of the keys in this sorted order.

// Level 3: Extend + Capacity
// Extend the function from Level 2 to only return the 'n' most common strings, where 'n' is an additional input to the function. If 'n' is greater than the number of distinct strings, the function should return all the distinct strings in the sorted order.

// Level 4: Backup and Restore
// The function should backup the full sorted list of strings (not limited by 'n') to a file in the local file system. Write a new function that reads the backup file and restores the full sorted list of strings.

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sort"
)

// Level 1 answer

// func countString(inputString []string) map[string]int {
// 	return_ob := make(map[string]int)
// 	for _, i := range inputString {
// 		return_ob[i] += 1
// 	}
// 	return return_ob
// }

// Level 2 answer

type kv struct {
	Key   string
	Value int
}

// func countString(inputString []string) []string {
// 	if len(inputString) == 0 {
// 		return nil
// 	}

// 	returnOb := make(map[string]int)
// 	for _, i := range inputString {
// 		returnOb[i]++
// 	}

// 	returnSort := make([]kv, 0, len(returnOb))
// 	for k, v := range returnOb {
// 		returnSort = append(returnSort, kv{k, v})
// 	}

// 	sort.Slice(returnSort, func(i, j int) bool {
// 		if returnSort[i].Value == returnSort[j].Value {
// 			return returnSort[i].Key < returnSort[j].Key
// 		}
// 		return returnSort[i].Value > returnSort[j].Value
// 	})

// 	returnKey := make([]string, len(returnSort))
// 	for i, kv := range returnSort {
// 		returnKey[i] = kv.Key
// 	}

// 	return returnKey
// }

// Level 3 answer

// func countString(inputString []string, n int) []string {
// 	if len(inputString) == 0 {
// 		return nil
// 	}

// 	returnOb := make(map[string]int)
// 	for _, i := range inputString {
// 		returnOb[i]++
// 	}

// 	returnSort := make([]kv, 0, len(returnOb))
// 	for k, v := range returnOb {
// 		returnSort = append(returnSort, kv{k, v})
// 	}

// 	sort.Slice(returnSort, func(i, j int) bool {
// 		if returnSort[i].Value == returnSort[j].Value {
// 			return returnSort[i].Key < returnSort[j].Key
// 		}
// 		return returnSort[i].Value > returnSort[j].Value
// 	})

// 	if n > len(returnSort) {
// 		n = len(returnSort)
// 	}
// 	returnSort = returnSort[:n]

// 	returnKey := make([]string, n)
// 	for i, kv := range returnSort {
// 		returnKey[i] = kv.Key
// 	}

// 	return returnKey
// }

// Level 4 answer

func countString(inputString []string, n int, backupFile string) ([]string, error) {
	if len(inputString) == 0 {
		return nil, nil
	}

	returnOb := make(map[string]int)
	for _, i := range inputString {
		returnOb[i]++
	}

	returnSort := make([]kv, 0, len(returnOb))
	for k, v := range returnOb {
		returnSort = append(returnSort, kv{k, v})
	}

	sort.Slice(returnSort, func(i, j int) bool {
		if returnSort[i].Value == returnSort[j].Value {
			return returnSort[i].Key < returnSort[j].Key
		}
		return returnSort[i].Value > returnSort[j].Value
	})

	// Back up the full sorted list
	file, err := os.Create(backupFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	if err := enc.Encode(returnSort); err != nil {
		return nil, err
	}

	if n > len(returnSort) {
		n = len(returnSort)
	}
	returnSort = returnSort[:n]

	returnKey := make([]string, n)
	for i, kv := range returnSort {
		returnKey[i] = kv.Key
	}

	return returnKey, nil
}

func restoreFromFile(backupFile string) ([]string, error) {
	file, err := os.Open(backupFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var returnSort []kv
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&returnSort); err != nil {
		return nil, err
	}

	returnKey := make([]string, len(returnSort))
	for i, kv := range returnSort {
		returnKey[i] = kv.Key
	}

	return returnKey, nil
}

// func main() {
// 	inputString := []string{"str1", "str2", "str3", "str1", "str1", "str3"}
// 	return_ob, _ := countString(inputString, 2, "backup.gob")

//		fmt.Println(return_ob)
//	}
var deckSize int

func main() {
	inputString := []string{"apple", "banana", "apple", "cherry", "banana", "cherry", "apple"}
	sortedKeys, err := countString(inputString, 3, "backup.gob")
	fmt.Println(sortedKeys) // This will output: ["apple", "banana", "cherry"]

	// Now we want to recover the data from the backup file
	recoveredData, err := restoreFromFile("backup.gob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recoveredData) // This will output the full sorted list, not limited by 'n'

	deckSize = 50
	fmt.Println(deckSize)
}
