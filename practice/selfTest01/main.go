package main

// import (
// 	"fmt"
// 	"os"
// 	"sort"
// )

import (
	"encoding/gob"
	"fmt"
	"os"
	"sort"
)

// Level 1 - Hash

// Write a function in Go that accepts a slice of strings as input and returns a hash map where the keys are the unique strings in the slice, and the values are the counts of each string.

// Level 2 - Hash + Sorting

// Modify the function from Level 1 to sort the keys of the hash map in descending order of their corresponding counts. If two keys have the same count, they should appear in lexicographically increasing order. The function should return a slice of the keys in this sorted order.

// Level 3 - Extend: Hash + Sorting + Data Structure

// Modify the function from Level 2 to limit the capacity of the returned slice to a specified number 'n'. The function should accept 'n' as an additional parameter and only return the 'n' most frequent strings. If there are more than 'n' strings with the same frequency, the function should include the lexicographically smallest ones.

// Level 4 - Extend: Hash + Sorting + Data Structure + Backup and Restore

// Modify the function from Level 3 to add two features:

// The function should backup the full sorted list of strings (not limited by 'n') to a file in the local file system.
// Write a new function that reads the backup file and restores the full sorted list of strings.

type kv struct {
	Key   string
	Value int
}

// func sortHash(slicestring []string) []string {
// 	returnOb := make(map[string]int)
// 	for _, value := range slicestring {
// 		returnOb[value]++
// 	}

// 	var sorted []kv
// 	for k, v := range returnOb {
// 		sorted = append(sorted, kv{k, v})
// 	}

// 	sort.Slice(sorted, func(i, j int) bool {
// 		if sorted[i].Value == sorted[j].Value {
// 			return sorted[i].Key < sorted[j].Key // Lexicographical order when counts are equal
// 		}
// 		return sorted[i].Value > sorted[j].Value // Descending order of counts
// 	})

// 	fmt.Println("sorted:", sorted)
// 	var result []string
// 	for _, kv := range sorted {
// 		result = append(result, kv.Key)
// 	}
// 	return result
// }

func sortHash(slicestring []string, n int, filename string) ([]string, error) {
	returnOb := make(map[string]int)
	for _, value := range slicestring {
		returnOb[value]++
	}

	var sorted []kv
	for k, v := range returnOb {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key // Lexicographical order when counts are equal
		}
		return sorted[i].Value > sorted[j].Value // Descending order of counts
	})

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(sorted)
	if err != nil {
		return nil, err
	}

	var result []string
	// Check if n is greater than the length of sorted, if yes, set n to the length of sorted
	if n > len(sorted) {
		n = len(sorted)
	}
	sorted = sorted[:n]
	for _, kv := range sorted {
		result = append(result, kv.Key)
	}
	return result, nil
}

func main() {
	strings := []string{"apple", "orange", "banana", "apple", "orange", "apple"}
	result, err := sortHash(strings, 2, "backup.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
