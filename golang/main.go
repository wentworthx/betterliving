package main

import (
	"betterliving/golang/record"
	"fmt"
	"log"
	"os"
)

var (
	projectHaveDone    string
	amount             int
	whetherMoreDetails string
	details            map[string]string
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// TODO: brief interaction.
	// TODO: webpage on VPS for anywhere record. maybe vue
	fmt.Println("What have you done today?")
	// TODO: check & fuzzy matching
	// TODO: droplist
	fmt.Scanln(&projectHaveDone)
	fmt.Println("What is the amount?")
	fmt.Scanln(&amount)
	fmt.Println("Any more details to record? (Y / N|anykey)")
	fmt.Scanln(&whetherMoreDetails)

	if whetherMoreDetails == "Y" {
		// TODO: collect details and record
	}

	record.Append(projectHaveDone, amount, details)

	fmt.Println("You are so great today! Keep going tomorrow ^_^")
}
