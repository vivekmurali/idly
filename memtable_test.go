package main

import (
	"slices"
	"sort"
	"testing"
)

// Checking to make sure that the redblacktree returns the keys in order
func TestMemtablePut(t *testing.T) {
	mt := initMemTable()
	input := map[string]string{"b": "1", "a": "2", "c": "3", "e": "4", "d": "5"}
	for k, v := range input {
		mt.Put(k, v)
	}

	keyInterface := mt.tree.Keys()
	keys := make([]string, len(keyInterface))

	for i, v := range keyInterface {
		keys[i] = v.(string)
	}

	sortedKeysClone := slices.Clone(keys)
	sort.Strings(sortedKeysClone)

	for i, v := range sortedKeysClone {
		if keys[i] != v {
			t.Errorf("not in order")
		}
	}

}
