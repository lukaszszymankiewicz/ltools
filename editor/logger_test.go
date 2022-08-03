package editor

import (
	"bufio"
	"os"
	"testing"
	"time"
)

const (
	TEST_LOGGER_DIR = "../logs/"
)

func TestLogCanvasPutSingleLog(t *testing.T) {
	lg := NewLogger(TEST_LOGGER_DIR)
	lg.LogCanvasPut(1, 2, 0)

	file, err := os.Open(TEST_LOGGER_DIR + lg.filename)

	if err != nil {
		t.Errorf("reading log files does not succeeded!")
		t.Error(err)
	}

	scanner := bufio.NewScanner(file)

	proper_log := "Tile put on (1, 2), on layer 0"
	proper_prefix := "CANVAS:"

	for scanner.Scan() {
		line := scanner.Text()

		if line[:LOG_PREFIX_N] != proper_prefix {
			t.Errorf("bad prefix! want %s get: %s", proper_prefix, line[:LOG_PREFIX_N])
		}
		if line[LOG_CONTENT_N:] != proper_log {
			t.Errorf("bad log! want:\n %s \nget:\n %s\n", proper_log, line[LOG_CONTENT_N:])
		}
	}

	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	file.Close()

	// just for sanity
	err = os.Remove(TEST_LOGGER_DIR + lg.filename)
	if err != nil {
		t.Error(err)
	}

}

func TestLogCanvasPutMultipleLog(t *testing.T) {
	lg := NewLogger(TEST_LOGGER_DIR)

	// some drawing and clearing on the end
	lg.LogCanvasPut(1, 1, 0)
	lg.LogCanvasPut(2, 2, 0)
	lg.LogCanvasPut(3, 3, 0)
	lg.LogCanvasPut(1, 1, 0)

	file, err := os.Open(TEST_LOGGER_DIR + lg.filename)

	if err != nil {
		t.Errorf("reading log files does not succeeded!")
		t.Error(err)
	}

	expected_logs_content := make(map[int]string)

	expected_logs_content[0] = "Tile put on (1, 1), on layer 0"
	expected_logs_content[1] = "Tile put on (2, 2), on layer 0"
	expected_logs_content[2] = "Tile put on (3, 3), on layer 0"
	expected_logs_content[3] = "Tile put on (1, 1), on layer 0"

	proper_prefix := "CANVAS:"

	n := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line[:LOG_PREFIX_N] != proper_prefix {
			t.Errorf("bad prefix! want %s get: %s", proper_prefix, line[:LOG_PREFIX_N])
		}
		if line[LOG_CONTENT_N:] != expected_logs_content[n] {
			t.Errorf("bad log! want:\n %s \nget:\n %s\n", expected_logs_content[n], line[LOG_CONTENT_N:])
		}
		n++
	}

	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	file.Close()

	// just for sanity
	err = os.Remove(TEST_LOGGER_DIR + lg.filename)
	if err != nil {
		t.Error(err)
	}

}

func TestNewLogger(t *testing.T) {
	log := NewLogger(TEST_LOGGER_DIR)

	dt := time.Now().Format("2006-01-02_15:04:05")

	// cutting is needed as logger filename contains seconds, and it is hard to grasp so
	if dt != log.filename[:19] {
		t.Errorf("adding logger filename does not succeeded!")
	}

	if log.file == nil {
		t.Errorf("creating logger file does not succeeded!")
	}

	if _, err := os.Stat("../logs" + log.filename); err == nil || os.IsExist(err) {
		t.Errorf("creating logger file does not succeeded!\n")
		t.Errorf("%s\n", "logs/"+log.filename)
	}
}
