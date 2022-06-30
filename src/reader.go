package main

import (
	"fmt"
	"os"

	"gopkg.in/neurosnap/sentences.v1/english"
)

func scanFile() {
	//Open file
	fmt.Println("Enter file directory:")
	var fileOpen string
	fmt.Scanln(&fileOpen)
	file, err := os.ReadFile(fileOpen)
	if err != nil {
		panic(err)
	}

	text := string(file)

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}
	sentences := tokenizer.Tokenize(text)
	for _, s := range sentences {
		fmt.Println(s.Text)
		fmt.Println("__________")
	}
}
