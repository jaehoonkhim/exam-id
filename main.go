package main

import (
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
)

func main() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
