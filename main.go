package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type LicenseStruct struct {
	Name string
	Url  string
}

type Phonetic struct {
	Text      string
	Audio     string
	SourceUrl string
	License   LicenseStruct
}

type DefinationStruct struct {
	Definition string
	Example    string
}

type Meaning struct {
	PartOfSpeech string
	Definitions  []DefinationStruct
}

type MeaningStruct struct {
	Word       string
	Phonetics  []Phonetic
	Meanings   []Meaning
	License    LicenseStruct
	SourceUrls []string
}

func getDefination(cmd *flag.FlagSet, w *string) {
	cmd.Parse(os.Args[2:])

	if *w == "" {
		fmt.Println("The word to be defined is required, use '-w=<your_word>' flag")
		os.Exit(1)
	}

	uri := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%v", *w)

	resp, err := http.Get(uri)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)

	var jsonBytes []MeaningStruct

	err = json.Unmarshal([]byte(string(data)), &jsonBytes)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s %s \n", jsonBytes[0].Word, jsonBytes[0].Phonetics[0].Text)

	for i := 0; i < len(jsonBytes[0].Meanings); i++ {
		res := jsonBytes[0].Meanings[i]

		fmt.Printf("\n %s \n", string(res.PartOfSpeech))

		for j := 0; j < len(res.Definitions); j++ {
			fmt.Printf("\n   %d. %s \n", j+1, res.Definitions[j].Definition)

			if res.Definitions[j].Example != "" {
				fmt.Printf("\n \t- Example: %s \n", res.Definitions[j].Example)
			}

			if j == 2 {
				j = len(res.Definitions)
			}
		}
	}

	fmt.Println()

}

func getPronunciation(cmd *flag.FlagSet, w *string) {
	fmt.Printf("%s \n", string(*w))
	fmt.Println("Coming soon!")
}

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
		getDefination(defineCmd, &word)
		return
	case "pronounce":
		getPronunciation(pronounceCmd, &word)
		return
	default:
		// call help func
		fmt.Println("Unknown command! Expected either 'define' or 'pronounce subcomands")
		os.Exit(1)
	}

	flag.Parse()
}
