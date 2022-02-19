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

	fmt.Printf("%s \n \n", jsonBytes[0].Word)

	for i := 0; i < len(jsonBytes[0].Meanings); i++ {
		res := jsonBytes[0].Meanings[i]

		fmt.Printf("%s \n", string(res.PartOfSpeech))
		fmt.Printf("\n   %s \n \n", res.Definitions[0].Definition)
	}

}

func getPronunciation(cmd *flag.FlagSet, w *string) {
	fmt.Println("Coming soon!")
}

func main() {
	defineCmd := flag.NewFlagSet("define", flag.ExitOnError)
	word := defineCmd.String("w", "", "Word to be defined")

	pronounceCmd := flag.NewFlagSet("pronounce", flag.ExitOnError)
	pWord := pronounceCmd.String("w", "", "Word to be pronounced")

	if len(os.Args) < 2 {
		// call help func
		fmt.Println("Expected either 'define' or 'pronounce subcomands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "define":
		getDefination(defineCmd, word)
		return
	case "pronounce":
		getPronunciation(pronounceCmd, pWord)
		return
	default:
		// call help func
		fmt.Println("Unknown command! Expected either 'define' or 'pronounce subcomands")
		os.Exit(1)
	}

	flag.Parse()
}
