package config

import (
	"encoding/json"
	"log"
	"os"
	"io"
	"io/ioutil"

	"github.com/evanfeenstra/circuitSocket/app/wrappers/server"

)

// config the settings variable
var Config = &configuration{}

type configuration struct {
	Server    server.Server   `json:"Server"`
	Client    string          `json:"Client"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}


// Parser implements ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file
func Load(configFile string, p Parser) {
	var err error
	var input = io.ReadCloser(os.Stdin)
	if input, err = os.Open(configFile); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
