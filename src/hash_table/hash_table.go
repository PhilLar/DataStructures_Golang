package hash_table

import (
	"fmt"
	"errors"
)


type Node struct {
	Data	interface{}
	Key		int
}

type HashTable struct {
	table	[]*Node
}

func NewHashTable(len int) (*HashTable, error) {
	if len <= 0 {
		return nil, errors.New("Hashtable's size must be over 0!")
	}
	ht := &HashTable{}
	ht.SetSize(len)
	return ht, nil
}

func (ht *HashTable) Table() []*Node {
	return ht.table
}

func (ht *HashTable) SetSize(l int) {
	ht.table = make([]*Node, l)
}


func (ht *HashTable) tableOverflowed() bool {
	for _, i := range ht.table {
		if i == nil {
			return false
		}
	}
	return true
}

func (ht *HashTable) Insert(data interface{}, key int) bool {
	//fmt.Println(ht.table)
	if ht.tableOverflowed() || key < 0{
		return false
	} else {
		node := &Node{Data:data, Key:key}
		hash := node.Key % len(ht.table)
		for ht.table[hash] != nil {
			hash = (hash+1) % len(ht.table)
	}
		ht.table[hash] = node
		return true
	}
}

func (ht *HashTable) Delete(key int) bool {
	if key >= 0{
		hash := key % len(ht.table)
		for hash != len(ht.table) {
			hash = hash % len(ht.table)
			//fmt.Println("hash", hash)
			if ht.table[hash] != nil && ht.table[hash].Key == key {
				ht.table[hash] = nil
				return true
			}
			hash = (hash+1)
		}
	}
	return false
}

func (ht *HashTable) GetNode(key int) *Node {
	hash := key % len(ht.table)
	for hash != len(ht.table) {
		if ht.table[hash] != nil && ht.table[hash].Key == key {
			return ht.table[hash]
		}
		hash = (hash+1) % len(ht.table)
	}
	return nil
}

func (ht *HashTable) Show() {
	arr := make([]Node, len(ht.table))
	//fmt.Println(len(ht.table))
	for ind, v := range ht.table {
		if v != nil {
		arr[ind] = *v
		}
	}
	fmt.Println(ht.table)
	fmt.Println(arr)
}


