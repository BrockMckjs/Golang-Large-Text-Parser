package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	//Specify Directory
	fmt.Println("Enter a Directory to store the files:")
	var directFin string
	fmt.Scanln(&directFin)

	//Variables
	//String conversion of text file
	text := string(file)

	//File name starting value
	var ctr int = 0

	//Sentence tokenizer
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	//Tokenize each sentence in the text varibale string
	sentences := tokenizer.Tokenize(text)
	//For loop to create files for each sentence
	for _, s := range sentences {
		//For every file made, increase title value +1
		ctr += 1
		pageNumber := fmt.Sprint(ctr)
		//Create file in user specified directory
		f, err := os.Create(filepath.Join(directFin, filepath.Base(pageNumber)))
		if err != nil {
			log.Fatal(err)
		}

		//Conversion of s var to a string from byte value
		convertedBytes := string(s.Text)

		//Write converted string into the newly created file
		l, err := f.WriteString(convertedBytes)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		//Close file
		defer f.Close()
		//Print to console
		fmt.Printf("%s\n", convertedBytes)
		fmt.Printf("Character count:%d\n", l)
		fmt.Println("________________________________")
	}
}
