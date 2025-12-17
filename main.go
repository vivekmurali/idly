package main

import "log"

func main() {
	kvs := []KeyValue{
		{"abc", "123"},
		{"def", "456"},
		{"ghijk", "6789011"},
		{"abcd", "123"},
		{"defasd", "456"},
		{"ghijkafj", "6789011"},
		{"abcvghhhd", "123"},
		{"defhdbnvnnf", "456"},
		{"ghijdddddsk", "6789011"},
		{"abclkpfj", "123"},
		{"defvghudvsh", "456"},
		{"ghijksa", "6789011"},
	}

	mt := initMemTable()

	for _, v := range kvs {
		mt.Put(v.Key, v.Value)
	}

	// err := Write("ss_1.dat", kvs)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Completed writing to SS Table")

	sst, err := Read("new_sstable.dat")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Read from SSTable")

	val1, err := sst.Get("abc")
	if err != nil {
		log.Fatal(err)
	}

	// val1, err := sst.Get("abc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// val2, err := sst.Get("def")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// val3, err := sst.Get("ghijk")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("abc: ", val1)
	// log.Println("def: ", val2)
	// log.Println("ghijk: ", val3)

}
