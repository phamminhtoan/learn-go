// Removing key/value pairs with the "delete" function
package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	ranks["brozen"] = 3
	rank, ok := ranks["brozen"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "brozen")
	rank, ok = ranks["brozen"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
}
