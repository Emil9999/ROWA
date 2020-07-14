package settings

import (
	"os"
	"encoding/json"
)
var Debug = true

var ArduinoOn = false

type Config struct {
    ArduinoOn bool `json:"arduino"`
    Debug bool `json:"debug"`
}


func LoadConfiguration(file string) (err error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	
	if err != nil {
	return 
	}
	
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	Debug = config.Debug
	ArduinoOn = config.ArduinoOn
	return 
}