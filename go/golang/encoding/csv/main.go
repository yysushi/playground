package main

import (
	// "bufio"
	"encoding/csv"
	// "io/ioutil"
	"log"
	"os"
)

func main() {
	// 1. os.Open, csv.NewReader, r.ReadAll
	f, err := os.Open("hoge.csv")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(len(records), records)
}
