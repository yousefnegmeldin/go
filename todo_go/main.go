package main

import (
	"bytes"
	"fmt"
)

func main() {
	// var itemsList []string
	// for {
	// 	var input string
	// 	fmt.Println("what do you want to do")
	// 	fmt.Scan(&input)
	// 	switch input {
	// 	case "view":
	// 		ViewList(itemsList)
	// 		break
	// 	case "add":
	// 		AddToList(itemsList, "hello")
	// 		break
	// 	case "check":
	// 	case "delete":
	// 	}
	// 	if input == "break" {
	// 		break
	// 	}
	// }
}

func ViewList(writer *bytes.Buffer, items []string) {
	for i, item := range items {
		if i == len(items)-1 {
			fmt.Fprintf(writer, item)
		} else {
			fmt.Fprintf(writer, item+"\n")
		}
	}
}

func AddToList(items []string, item string) {
	items = append(items, item)
}
