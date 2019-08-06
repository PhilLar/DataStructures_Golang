package hash_table

import (
	"fmt"
)

// type Patient struct {
// 	ID 			string
//     FIO 		string
//     bd_year 	int
//     address 	string
//     workPlace 	string
// 	key 		int
// }

type Node struct {
	Data	interface{}
	Key		int
}

type HashTable struct {
	Table	[5]*Node
}

func (ht *HashTable) tableOverflowed() bool {
	for _, i := range ht.Table {
		if i == nil {
			return false
		}
	}
	return true
}

func (ht *HashTable) Insert(node *Node) bool {
	if ht.tableOverflowed() {
		return false
	} else {
		hash := node.Key % len(ht.Table)
		for ht.Table[hash] != nil {
			hash = (hash+1) % len(ht.Table)
	}
		ht.Table[hash] = node
		return true
	}
}

func (ht *HashTable) Delete(key int) bool {
	hash := key % len(ht.Table)
	for hash != len(ht.Table) {
		if ht.Table[hash] != nil && ht.Table[hash].Key == key {
			ht.Table[hash] = nil
			return true
		}
		hash = (hash+1) % len(ht.Table)
	}
	return false
}

func (ht *HashTable) GetNode(key int) *Node {
	hash := key % len(ht.Table)
	for hash != len(ht.Table) {
		if ht.Table[hash] != nil && ht.Table[hash].Key == key {
			return ht.Table[hash]
		}
		hash = (hash+1) % len(ht.Table)
	}
	return nil
}

func (ht *HashTable) Show() {
	arr := make([]Node, len(ht.Table))
	fmt.Println(ht.Table)
	for ind, v := range ht.Table {
		if v != nil {
		arr[ind] = *v
		}
	}
	fmt.Println(arr)
}


