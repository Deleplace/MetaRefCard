package common

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// MockSet is a map masquerading as a set
type MockSet map[string]string

// Keys returns a MockSet as an array
func (m MockSet) Keys() []string {
	array := make([]string, 0, len(m))
	for k := range m {
		array = append(array, k)
	}
	return array
}

// LoadYaml loads Yaml file and prints any errors
func LoadYaml(filename string, out interface{}, label string) {
	loadYaml(filename, out, label)
}

func loadYaml(filename string, out interface{}, label string) {
	yamlData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: yaml ioutil.ReadFile %v ", err)
	}
	err = yaml.Unmarshal([]byte(yamlData), out)
	if err != nil {
		log.Fatalf("error: yaml.Unmarshal %v", err)
	}
	debugOutput := false
	if debugOutput {
		PrintYamlObject(out, label)
	}
}

// PrintYamlObject outputs contents of yaml object with a label
func PrintYamlObject(in interface{}, label string) {
	d, err := yaml.Marshal(in)
	if err != nil {
		log.Fatalf("error: yaml.Marshal %v", err)
	}
	fmt.Printf("=== %s ===\n%s\n\n", label, string(d))

}
