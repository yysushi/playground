package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// config ...
type Config struct {
	Debug bool
}

func main() {
	var config Config
	// 1.
	f, err := os.Open("hoge.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", config)
	// 2.
	b, err := ioutil.ReadFile("hoge2.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", config)
}
