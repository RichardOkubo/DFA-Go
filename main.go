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
	w := strings.Split(os.Args[1], "")

	file, err := ioutil.ReadFile("./dfa.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var dfa DFA

	err = json.Unmarshal(file, &dfa)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	configState := "%s" + color.Green + "[q%v]" + color.Reset + "%s\n"

	i := 0
	for i < len(w) {
		fmt.Printf(configState, strings.Join(w[:i], ""), dfa.Q0, strings.Join(w[i:], ""))
		dfa.Q0 = dfa.Delta[dfa.Q0+","+w[i]]
		i++
	}
	fmt.Printf(configState, strings.Join(w[:i], ""), dfa.Q0, strings.Join(w[i:], ""))

	status := false
	for _, f := range dfa.F {
		if f == dfa.Q0 {
			status = true
			break
		}
	}
	fmt.Println(status)
}
