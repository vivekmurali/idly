package main

import (
	"bufio"
	"encoding/binary"
	"os"
	"sort"
)

type SSTable struct {
	filename string
	index    map[string]int64
}

type KeyValue struct {
	Key   string
	Value string
}

func Write(filename string, kvs []KeyValue) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Key < kvs[j].Key
	})

	index := make(map[string]int64)

	writer := bufio.NewWriter(file)

	for _, kv := range kvs {
		offset, err := file.Seek(0, os.SEEK_CUR)
		if err != nil {
			return err
		}
		index[kv.Key] = offset

		binary.Write(writer, binary.LittleEndian, int32(len(kv.Key)))
		writer.WriteString(kv.Key)
		binary.Write(writer, binary.LittleEndian, int32(len(kv.Value)))
		writer.WriteString(kv.Value)
	}
	writer.Flush()

	indexOffset, err := file.Seek(0, os.SEEK_CUR)
	if err != nil {
		return err
	}

	for _, kv := range kvs {
		binary.Write(writer, binary.LittleEndian, int32(len(kv.Key)))
		writer.WriteString(kv.Key)
		binary.Write(writer, binary.LittleEndian, index[kv.Key])
	}
	writer.Flush()
	binary.Write(writer, binary.LittleEndian, indexOffset)
	writer.Flush()

	return nil
}
