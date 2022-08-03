package editor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	LevelDir string
}

func ReadConfig(path string) Config {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	var config Config

	json.Unmarshal(bytes, &config)

	return config
}
