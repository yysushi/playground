package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// $ mkpasswd --method=bcrypt 123
	// $2b$05$yCNQdRUckvGvqmCHluq./.SS6rUI6Fqm0f7Sqa7fVH2CPs8tS.By2
	// $ mkpasswd --method=bcrypt 123
	// $2b$05$EP3i/1lRb87WQMnq79Irj.r.v7TS6Cjp7CN69w7Iulir/.IBM9g0i
	hasheds := [][]byte{
		[]byte("$2b$05$yCNQdRUckvGvqmCHluq./.SS6rUI6Fqm0f7Sqa7fVH2CPs8tS.By2"),
		[]byte("$2b$05$EP3i/1lRb87WQMnq79Irj.r.v7TS6Cjp7CN69w7Iulir/.IBM9g0i"),
	}
	for _, hashed := range hasheds {
		if err := bcrypt.CompareHashAndPassword(hashed, []byte("123")); err != nil {
			panic(err)
		}
	}
	fmt.Println("matched")
}
