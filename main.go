package main

import (
	"log"

	_ "CMAPI/internal/entity/fields" // Register core field types via init()
)

func main() {
	log.Println("CMAPI bootstrapped with registered field types.")
}
