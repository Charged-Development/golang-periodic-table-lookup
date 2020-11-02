package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/stretchr/objx"
)

func checkvalue(obj []byte, thing string) {
	fmt.Print("\n\n")
	var lmao map[string]interface{}
	test := json.Unmarshal([]byte(obj), &lmao)
	if test != nil {
		fmt.Print(test)
		return
	}
	asd123, kekw := json.Marshal(lmao[thing])
	if kekw != nil {
		fmt.Print(kekw)
		return
	}
	m := objx.MustFromJSON(string(asd123))
	for key, value := range m {
		if value == nil {
			fmt.Printf("%s: None\n", key)
		} else {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}
func main() {
	content, err := ioutil.ReadFile("./periodic.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	for {
		var input string
		fmt.Print("\n\nGolang Periodic Table> ")
		fmt.Scanln(&input)
		input1 := strings.ToLower(input)
		checkvalue(content, input1)
	}
}
