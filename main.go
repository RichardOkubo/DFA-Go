package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Richard/dfa/color"
)

type DFA struct {
	Delta map[string]string `json:"delta"`
	Sigma []string          `json:"sigma"`
	Q     []string          `json:"Q"`
	F     []string          `json:"F"`
	Q0    string            `json:"q0"`
}

func main() {

	if len(os.Args) <= 1 {
		log.Fatal(color.Red + "Please enter the filename and string to be parsed" + color.Reset)
	}

	if len(os.Args) > 3 {
		log.Fatal(color.Red + "Too many arguments are not allowed" + color.Reset)
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(color.Red+"Error when opening file: "+color.Reset, err)
	}

	if len(os.Args) == 2 {
		log.Fatal(color.Red + "The string to be parsed is required" + color.Reset)
	}

	var dfa DFA

	err = json.Unmarshal(file, &dfa)
	if err != nil {
		log.Fatal(color.Red+"Error during Unmarshal(): "+color.Reset, err)
	}

	w := os.Args[2]

	unique := []string{}
	keys := make(map[string]bool)

	for _, char := range strings.Split(w, "") {
		if _, value := keys[char]; !value {
			keys[char] = true
			unique = append(unique, char)
		}
	}

	isEqual := false
	for _, char := range unique {
		for _, symbol := range dfa.Sigma {
			if char == symbol {
				isEqual = true
				break
			}
		}
		if !isEqual {
			log.Fatal(color.Red + w + " ∉ Σ*" + color.Reset)
		}
		isEqual = false
	}

	configState := "%s" + color.Green + "[q%v]" + color.Reset + "%s\n"

	i := 0
	for i < len(w) {
		fmt.Printf(configState, string(w[:i]), dfa.Q0, string(w[i:]))
		dfa.Q0 = dfa.Delta[dfa.Q0+","+string(w[i])]
		i++
	}
	fmt.Printf(configState, string(w[:i]), dfa.Q0, string(w[i:]))

	status := false
	for _, f := range dfa.F {
		if f == dfa.Q0 {
			status = true
			break
		}
	}
	fmt.Printf(color.Yellow+"Status: %v\n"+color.Reset, status)
}
