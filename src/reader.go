package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanFile() {
	//Open file
	var fileOpen string
	fmt.Scanln(&fileOpen)
	file, ferr := os.Open(fileOpen)
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		new := strings.Split(line, ".")

		for _, sentence := range new {
			fmt.Println(sentence)
			fmt.Println("______________")
		}
	}

}
