package config

import (
	"encoding/json"
	"fmt"
	"github.com/captainhook-go/captainhook/ch"
	"io/ioutil"
)

func CreateConfiguration(path string, settings map[string]string) {
	if path == "" {
		path = ch.CONFIG
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened captainhook.json")

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", payload)
}
