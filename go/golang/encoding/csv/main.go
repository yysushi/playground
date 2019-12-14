package main

import (
	// "bufio"
	"encoding/csv"
	// "io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hoge.csv")
	if err != nil {
		return
	}
	// r := bufio.NewReader(f)
	r := csv.NewReader(f)
	// for {
	// 	record, err := r.Read()
	// 	if err != nil && err != io.EOF {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	if len(record) != 0 {
	// 		records = append(records, record)
	// 	}
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(len(records), records)
}
