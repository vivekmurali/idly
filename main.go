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

}
