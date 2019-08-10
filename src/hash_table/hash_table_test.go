package hash_table_test

import (
	"testing"
	ht "hash_table"
	"reflect"
	"errors"
	//"fmt"
)

// testing correctness of passed HashTable's length
func TestNew(t *testing.T) {
	tables := []struct {
			lenInput		int
			objectType 	string
			tableSize	int
			err 		error
	}{
			{5, "*hash_table.HashTable", 5, nil},
			{0, "*hash_table.HashTable", 0, errors.New("Hashtable's size must be over 0!")},
			{-1, "*hash_table.HashTable", 1, errors.New("Hashtable's size must be over 0!")},
	}

	for _, table := range tables {
		actual, err := ht.NewHashTable(table.lenInput)
		if err != nil {

			if err.Error() != table.err.Error() {
				t.Errorf("expected error: %v, actual: %v",
					table.err.Error(), err.Error())
			}
		} else {

			if len(actual.Table()) != table.lenInput {
				t.Errorf("expected TableSize: %d, actual: %d",
					table.lenInput, len(actual.Table()))
			}
	
			if reflect.TypeOf(actual).String() != table.objectType {
				t.Errorf("expected type: %s, actual: %s",
					table.objectType, reflect.TypeOf(actual).String())
			}
		}
	}
}

func TestInsert(t *testing.T) {
	tables := []struct {
		lenInput	int
		data 		interface{}
		key			int
		ok			bool
	}{
		{5, "data_test", 0, true},
		{5, nil, 0, true},
		{5, 2, 0, true},
		{5, struct{
			ID 		int
			name 	string
		}{12, "Philipp"}, 0, true},
		{5, &[]string{"1", "two"}, 0, true},
		{5, &[]string{"1", "two"}, -1, false},
	}

	for _, table := range tables {
		actual, _ := ht.NewHashTable(table.lenInput)
		if ok := actual.Insert(table.data, table.key); ok == table.ok && ok{
			if actual.GetNode(table.key).Data != table.data {
				t.Errorf("expected data: %d, data: %d",
					table.data, actual.GetNode(table.key).Data)
			}
		}
	}
}

func TestInsertOverflow(t *testing.T) {
	tables := []struct {
		lenInput	int
		data 		[]interface{}
		overflow 	bool
	}{
		{3, []interface{}{"try", 1, 4.32}, false},
		{3, []interface{}{"try", 1, 4.32, &[]float64{}}, true},
	}
	for _, table := range tables {
		actual, _ := ht.NewHashTable(table.lenInput)
		var ok bool
		for i, data := range table.data {
			ok = actual.Insert(data, i)
		}
		if !ok != table.overflow {
			t.Error("overflow not handled!")
		}
	}
}

func TestDelete(t *testing.T) {
	tables := []struct {
		data		interface{}
		keyInsert	int
		keyDelete	int
		ok			bool
	}{
		{"data_test1", 0, 0, true},
		{"data_test2", 0, 1, false},
		{"data_test3", 2, 2, true},
	}

	for _, table := range tables {
		actual, _ := ht.NewHashTable(3)
		_ = actual.Insert(table.data, table.keyInsert)
		ok := actual.Delete(table.keyDelete)
		if table.ok != ok {
			t.Errorf("expected deletion succes: %v, actual %v", table.ok, ok)
		}
	}
}

