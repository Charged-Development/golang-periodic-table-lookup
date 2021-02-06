package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type mainStruct struct {
	Name                          string    `json:"name"`
	Appearance                    string    `json:"appearance"`
	AtomicMass                    float64   `json:"atomic_mass"`
	Boil                          int       `json:"boil"`
	Category                      string    `json:"category"`
	Color                         string    `json:"color"`
	Density                       float64   `json:"density"`
	DiscoveredBy                  string    `json:"discovered_by"`
	Melt                          int       `json:"melt"`
	MolarHeat                     float64   `json:"molar_heat"`
	NamedBy                       string    `json:"named_by"`
	Number                        int       `json:"number"`
	Period                        int       `json:"period"`
	Phase                         string    `json:"phase"`
	Source                        string    `json:"source"`
	SpectralImg                   string    `json:"spectral_img"`
	Summary                       string    `json:"summary"`
	Symbol                        string    `json:"symbol"`
	Xpos                          int       `json:"xpos"`
	Ypos                          int       `json:"ypos"`
	Shells                        []int     `json:"shells"`
	ElectronConfiguration         string    `json:"electron_configuration"`
	ElectronConfigurationSemantic string    `json:"electron_configuration_semantic"`
	ElectronAffinity              float64   `json:"electron_affinity"`
	ElectronegativityPauling      float64   `json:"electronegativity_pauling"`
	IonizationEnergies            []float64 `json:"ionization_energies"`
}

type loadToMap map[string]*mainStruct

var (
	mainMap loadToMap
)

func logValue(input string) {
	fmt.Println("\n\n")
	var data = mainMap[input]
	test := reflect.ValueOf(data)
	if test.Kind() == reflect.Ptr {
		test = test.Elem()
	}
	typeOfTest := test.Type()
	for i := 0; i < test.NumField(); i++ {
		key := typeOfTest.Field(i).Name
		value := test.Field(i)
		if fmt.Sprintf("%v", value) == "nil" || fmt.Sprintf("%v", value) == "" {
			fmt.Printf("%v: None\n", key)
		} else {
			fmt.Printf("%v: %v\n", key, value)
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("./periodic.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	err = json.Unmarshal(content, &mainMap)
	for {
		var input string
		fmt.Print("\n\nGolang Periodic Table> ")
		fmt.Scanln(&input)
		input1 := strings.ToLower(input)
		if _, ok := mainMap[input1]; !ok {
			fmt.Println(input1 + " doesn't exist in the periodic table.")
			return
		}
		logValue(input1)
	}
}
