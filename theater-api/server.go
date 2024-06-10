package main

import (
	"fmt"

	"github.com/noonyuu/theater-api/db"
)

func main() {
	fmt.Println("Hello, world!")

	db.Setup()
}
