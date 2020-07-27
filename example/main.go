package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gdbu/uuid"
)

func main() {
	// Generate user id
	userID := uuid.New()

	// You can print it as a string
	fmt.Println(userID.String())

	// Additionally, a pointer to a UUID is a stringer interface
	fmt.Println(&userID)

	// You can marshal it as JSON
	if bs, err := json.Marshal(userID); err != nil {
		log.Fatal("Error marshaling!", err)
	} else {
		fmt.Println("JSON bytes!", string(bs))
	}
}
