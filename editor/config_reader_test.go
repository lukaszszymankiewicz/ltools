package editor

import (
	"testing"
)

const ( 
    TEST_CONFIG_DIR = "../config.json"
)

func TestReadConfig(t *testing.T) {
	config := ReadConfig(TEST_CONFIG_DIR)

	if config.LevelDir == "" {
		t.Errorf("Config File inproperly read - Level Direction section is empty\n")
		t.Errorf("%s \n", config.LevelDir)
	}
}
