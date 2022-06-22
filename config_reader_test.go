package main

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	var path string = "config.json"

	config := ReadConfig(path)

	if config.LevelDir == "" {
		t.Errorf("Config File inproperly read - Level Direction section is empty\n")
		t.Errorf("%s \n", config.LevelDir)
	}
}
