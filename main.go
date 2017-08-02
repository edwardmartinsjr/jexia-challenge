package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	// Load partial Klingon alphabet in correspondent hexadecimal number which is a valid Unicode code point.
	loadKlingonAlphabet()

	// Load args - [Usage example: ./main uhura"]
	convertKlingonAlphabet()

}

func convertKlingonAlphabet() {

	// Load args
	args := os.Args[1:]

	//For each of the inputs
	if len(args) != 0 {

		// set spaces rules
		var space bool

		// Work on phonetic and grammatical rules
		for i := 0; i < len(args); i++ {
			word := strings.ToUpper(args[i])
			if !space {
				space = true
			} else {
				fmt.Printf("%s ", PIqaDLetters["|"])
			}

			for j := 0; j < len(word); j++ {
				switch word[j] {
				case 'C':
					fmt.Printf("%s ", "0xF8D2")
				case 'G':
					fmt.Printf("%s ", "0xF8D5")
				case 'N':
					if j < len(word)-1 {
						if word[j+1] == 'G' {
							fmt.Printf("%s ", "0xF8DC")
							j = j + 1
						} else {
							fmt.Printf("%s ", PIqaDLetters[string(word[j])])
						}
					} else {
						fmt.Printf("%s ", PIqaDLetters[string(word[j])])
					}

				case 'T':
					if j < len(word)-1 {
						if word[j+1] == 'L' {
							fmt.Printf("%s ", "0xF8E4")
							j = j + 1
						} else {
							fmt.Printf("%s ", PIqaDLetters[string(word[j])])
						}
					} else {
						fmt.Printf("%s ", PIqaDLetters[string(word[j])])
					}

				default:
					fmt.Printf("%s ", PIqaDLetters[string(word[j])])
				}
			}

		}

		fmt.Println("")

		// Find out its species using http://stapi.co
		var name string
		for i := 0; i < len(args); i++ {

			name = name + args[i] + " "

		}

		// Get Character
		characterUID, err := getCharacter(name)
		if err != nil {
			log.Fatal("Get Species Err: ", err)
		}

		var specie string
		if len(characterUID) > 0 {
			// Get Specie
			specie, err = getSpecies(characterUID)
			if err != nil {
				log.Fatal("Get Species Err: ", err)
			}
		} else {
			specie = "Character Not Found"
		}

		fmt.Printf("%s ", specie)

		fmt.Println("")

	} else {
		fmt.Println("Usage example: ./main Uhura")
	}
}

func getCharacter(name string) (specie string, err error) {

	// Full character, returned when queried using UID
	body := []byte(fmt.Sprintf("title=%s&name=%s", name, name))

	// Build the request
	resp, err := http.Post("http://stapi.co/api/v1/rest/character/search", "x-www-form-urlencoded", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("%q\n", err.Error())
		return
	}

	// defer set the last action to be taken in the program
	defer resp.Body.Close()

	// Checks if the status code is 200, that indicating the success of the request
	if resp.StatusCode != 200 {
		println(resp.StatusCode)
		return
	}

	// Read All content
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%q\n", err.Error())
		return
	}

	// Unmarshal body to map interface
	var characterBaseResult map[string]interface{}
	err = json.Unmarshal(bodyBytes, &characterBaseResult)

	// map characters
	var character = characterBaseResult["characters"]

	// Go to interface
	var itfce = character.([]interface{})

	// go to type level
	var uid string
	if len(itfce) > 0 {
		var typ = itfce[0].(map[string]interface{})
		uid = typ["uid"].(string)
	} else {
		uid = ""
	}

	return uid, nil
}

func getSpecies(characterUID string) (specie string, err error) {

	// Full character, returned when queried using UID
	url := fmt.Sprintf("http://stapi.co/api/v1/rest/character/?uid=%s", characterUID)

	// Build the request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%q\n", err.Error())
		return
	}

	// defer set the last action to be taken in the program
	defer resp.Body.Close()

	// Checks if the status code is 200, that indicating the success of the request
	if resp.StatusCode != 200 {
		println(resp.StatusCode)
		return
	}

	// Read All content
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%q\n", err.Error())
		return
	}

	// Unmarshal body to struct
	var characterResult CharacterResult
	err = json.Unmarshal(bodyBytes, &characterResult)

	// Get characterSpecies node from mappping
	characterSpeciesResult := characterResult.Character["characterSpecies"]

	// Go to interface
	var itfce = characterSpeciesResult.([]interface{})

	// go to type level
	var name string
	if len(itfce) > 0 {
		var typ = itfce[0].(map[string]interface{})
		name = typ["name"].(string)
	} else {
		name = "Specie Not Found"
	}

	return name, nil
}

// CharacterResult -
type CharacterResult struct {
	Character map[string]interface{}
}

// PIqaDLetters - pIqaD to letters map
var PIqaDLetters map[string]string

func loadKlingonAlphabet() {
	PIqaDLetters = map[string]string{
		"A": "0xF8D0",
		"B": "0xF8D1",
		//"CH":"0xF8D2",
		"D": "0xF8D3",
		"E": "0xF8D4",
		//"GH":"0xF8D5",
		"H": "0xF8D6",
		"I": "0xF8D7",
		"J": "0xF8D8",
		"L": "0xF8D9",
		"M": "0xF8DA",
		"N": "0xF8DB",
		//"NG":"0xF8DC",
		"O": "0xF8DD",
		"P": "0xF8DE",
		"Q": "0xF8DF",
		//"QH":"0xF8E0",
		"R": "0xF8E1",
		"S": "0xF8E2",
		"T": "0xF8E3",
		//"TLH":"0xF8E4",
		"U":  "0xF8E5",
		"V":  "0xF8E6",
		"W":  "0xF8E7",
		"Y":  "0xF8E8",
		"0":  "0xF8F0",
		"1":  "0xF8F1",
		"2":  "0xF8F2",
		"3":  "0xF8F3",
		"4":  "0xF8F4",
		"5":  "0xF8F5",
		"6":  "0xF8F6",
		"7":  "0xF8F7",
		"8":  "0xF8F8",
		"9":  "0xF8F9",
		",":  "0xF8FD",
		".":  "0xF8FE",
		"\"": "0xF8E9",
		"|":  "0x0020",
	}
}
