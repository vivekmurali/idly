package main

import (
	"bufio"
	"encoding/binary"
	"errors"
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

	var curOffset int64
	for _, kv := range kvs {
		index[kv.Key] = curOffset
		binary.Write(writer, binary.LittleEndian, int32(len(kv.Key)))
		writer.WriteString(kv.Key)
		binary.Write(writer, binary.LittleEndian, int32(len(kv.Value)))
		writer.WriteString(kv.Value)
		curOffset += 8 + int64(len(kv.Key)) + int64(len(kv.Value))
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

func Read(filename string) (*SSTable, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	file.Seek(-8, os.SEEK_END)
	var indexOffset int64
	err = binary.Read(file, binary.LittleEndian, &indexOffset)
	if err != nil {
		return nil, err
	}
	file.Seek(indexOffset, os.SEEK_SET)
	index := make(map[string]int64)

	for {
		var keyLen int32
		err = binary.Read(file, binary.LittleEndian, &keyLen)
		if err != nil {
			break
		}
		key := make([]byte, keyLen)
		file.Read(key)
		var offset int64
		binary.Read(file, binary.LittleEndian, &offset)
		index[string(key)] = offset
	}

	return &SSTable{filename, index}, nil
}

func (sst *SSTable) Get(key string) (string, error) {
	val, ok := sst.index[key]
	if !ok {
		return "", errors.New("Key not found")
	}
	file, err := os.Open(sst.filename)
	if err != nil {
		return "", err
	}
	file.Seek(val, os.SEEK_SET)
	var keyLen int32
	binary.Read(file, binary.LittleEndian, &keyLen)
	keyByte := make([]byte, keyLen)
	file.Read(keyByte)

	var valLen int32
	binary.Read(file, binary.LittleEndian, &valLen)
	value := make([]byte, valLen)
	file.Read(value)

	return string(value), nil
}
