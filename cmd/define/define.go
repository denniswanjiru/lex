package define

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

func WordDefination(cmd *flag.FlagSet, w *string) {
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
