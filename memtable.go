package main

import (
	"log"

	"github.com/emirpasic/gods/trees/redblacktree"
)

/*
Put
Get
Delete
*/

type memtable struct {
	tree *redblacktree.Tree
}

func initMemTable() *memtable {
	return &memtable{redblacktree.NewWithStringComparator()}
}

func (m *memtable) Put(key, value string) {
	m.tree.Put(key, value)

	if m.tree.Size() > 9 {
		m.Flush()
		m.tree.Clear()
	}
}

func (m *memtable) Get(key string) (string, bool) {
	value, found := m.tree.Get(key)
	return value.(string), found
}

func (m *memtable) Delete(key string) {
	m.tree.Remove(key)
}

// Get all key value pairs and turn it into an SSTable and write the SSTable
func (m *memtable) Flush() {
	iter := m.tree.Iterator()

	iter.Begin()

	kvs := make([]KeyValue, 0)

	for iter.Next() {
		key := iter.Key().(string)
		value := iter.Value().(string)

		kv := KeyValue{Key: key, Value: value}
		kvs = append(kvs, kv)
	}

	if err := Write("new_sstable.dat", kvs); err != nil {
		log.Fatalf("error writing sstable: %v", err)
	}

}
