package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/denniswanjiru/lex/cmd/define"
	"github.com/denniswanjiru/lex/cmd/pronounce"
)

var word string

func main() {
	defineCmd := flag.NewFlagSet("define", flag.ExitOnError)
	defineCmd.StringVar(&word, "w", "", "Word to be defined")
	defineCmd.StringVar(&word, "word", "", "Word to be defined")

	pronounceCmd := flag.NewFlagSet("pronounce", flag.ExitOnError)
	pronounceCmd.StringVar(&word, "w", "", "Word to be pronounced")
	pronounceCmd.StringVar(&word, "word", "", "Word to be pronounced")

	if len(os.Args) < 2 {
		// call help func
		fmt.Println("Expected either 'define' or 'pronounce subcomands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "define":
		define.WordDefination(defineCmd, &word)
		return
	case "pronounce":
		pronounce.WordPronunciation(pronounceCmd, &word)
		return
	default:
		// call help func
		fmt.Println("Unknown command! Expected either 'define' or 'pronounce subcomands")
		os.Exit(1)
	}

	flag.Parse()
}
