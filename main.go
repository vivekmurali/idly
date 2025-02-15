package main

import "log"

func main() {
	kvs := []KeyValue{
		{"abc", "123"},
		{"def", "456"},
		{"ghijk", "6789011"},
	}
	err := Write("ss_1.dat", kvs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Completed writing to SS Table")

	sst, err := Read("ss_1.dat")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Read from SSTable")

	val1, err := sst.Get("abc")
	if err != nil {
		log.Fatal(err)
	}
	val2, err := sst.Get("def")
	if err != nil {
		log.Fatal(err)
	}
	val3, err := sst.Get("ghijk")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("abc: ", val1)
	log.Println("def: ", val2)
	log.Println("ghijk: ", val3)

}
