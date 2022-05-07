package objects

import (
	"testing"
    "time"
    "os"
    "bufio"
)

func TestLogCanvasPutSingleLog(t *testing.T) {
    lg := NewLogger()
    lg.LogCanvasPut(1, 2, LAYER_DRAW, 3)

    file, err := os.Open("../../logs/" + lg.filename)

    if err != nil {
        t.Errorf("reading log files does not succeeded!")
        t.Error(err)
    }

    scanner := bufio.NewScanner(file)
    
    proper_log := "Tile 3 put on (1, 2), on layer 0"
    proper_prefix := "CANVAS:"

    for scanner.Scan() {
        line := scanner.Text()

        if line[:LOG_PREFIX_N ] != proper_prefix {
            t.Errorf("bad prefix! want %s get: %s", proper_prefix, line[:LOG_PREFIX_N ])
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
    err = os.Remove("../../logs/" + lg.filename)
    if err != nil { t.Error(err) }

}

func TestLogCanvasPutMultipleLog(t *testing.T) {
    lg := NewLogger()

    // some drawing and clearing on the end
    lg.LogCanvasPut(1, 1, LAYER_DRAW, 0)
    lg.LogCanvasPut(2, 2, LAYER_DRAW, 0)
    lg.LogCanvasPut(3, 3, LAYER_DRAW, 0)
    lg.LogCanvasPut(1, 1, LAYER_DRAW, -1)

    file, err := os.Open("../../logs/" + lg.filename)

    if err != nil {
        t.Errorf("reading log files does not succeeded!")
        t.Error(err)
    }
    
    expected_logs_content := make(map[int]string)

    expected_logs_content[0] = "Tile 0 put on (1, 1), on layer 0"
    expected_logs_content[1] = "Tile 0 put on (2, 2), on layer 0"
    expected_logs_content[2] = "Tile 0 put on (3, 3), on layer 0"
    expected_logs_content[3] = "Tile -1 put on (1, 1), on layer 0"

    proper_prefix := "CANVAS:"

    n := 0

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if line[:LOG_PREFIX_N ] != proper_prefix {
            t.Errorf("bad prefix! want %s get: %s", proper_prefix, line[:LOG_PREFIX_N ])
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
    err = os.Remove("../../logs/" + lg.filename)
    if err != nil { t.Error(err) }

}

func TestLogTileStackPutMultiple(t *testing.T) {
    lg := NewLogger()

    // some drawing and clearing on the end
    lg.LogTileStackPut(0, LAYER_DRAW, 0, 0, 0) 
    lg.LogTileStackPut(1, LAYER_DRAW, 1, 1, 0) 
    lg.LogTileStackPut(2, LAYER_DRAW, 2, 2, 0) 
    lg.LogTileStackPut(3, LAYER_ENTITY, 0, 0, 2) 

    file, err := os.Open("../../logs/" + lg.filename)

    if err != nil {
        t.Errorf("reading log files does not succeeded!")
        t.Error(err)
    }
    
    expected_logs_content := make(map[int]string)
    
    expected_logs_content[0] = "Tile from tileset 0 (0, 0) on layer 0 added to stack on pos: 0, stack has now len:1"
    expected_logs_content[1] = "Tile from tileset 0 (1, 1) on layer 0 added to stack on pos: 1, stack has now len:2"
    expected_logs_content[2] = "Tile from tileset 0 (2, 2) on layer 0 added to stack on pos: 2, stack has now len:3"
    expected_logs_content[3] = "Tile from tileset 2 (0, 0) on layer 2 added to stack on pos: 3, stack has now len:4"

    proper_prefix := "STACKK:"

    n := 0

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if line[:LOG_PREFIX_N ] != proper_prefix {
            t.Errorf("bad prefix! want %s get: %s", proper_prefix, line[:LOG_PREFIX_N ])
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
    err = os.Remove("../../logs/" + lg.filename)
    if err != nil { t.Error(err) }

}

func TestLogMultipleDifferentFunctions(t *testing.T) {
    lg := NewLogger()

    // this represents sstarting point for a game!
    lg.LogTileStackPut(0, LAYER_LIGHT, 0, 0, 1) 
    lg.LogTileStackPut(1, LAYER_ENTITY, 0, 0, 2) 
    lg.LogTileStackPut(2, LAYER_DRAW, 0, 0, 0) 
    
    // lets draw something
    lg.LogCanvasPut(5, 5, LAYER_DRAW, 2)
    lg.LogCanvasPut(4, 5, LAYER_LIGHT, 0)
    lg.LogCanvasPut(6, 5, LAYER_LIGHT, 0)
    lg.LogCanvasPut(5, 4, LAYER_LIGHT, 0)
    lg.LogCanvasPut(5, 6, LAYER_LIGHT, 0)
    lg.LogCanvasPut(5, 5, LAYER_ENTITY, 1)

    file, err := os.Open("../../logs/" + lg.filename)

    if err != nil {
        t.Errorf("reading log files does not succeeded!")
        t.Error(err)
    }
    
    expected_logs_content := make(map[int]string)
    
    expected_logs_content[0] = "Tile from tileset 1 (0, 0) on layer 1 added to stack on pos: 0, stack has now len:1"
    expected_logs_content[1] = "Tile from tileset 2 (0, 0) on layer 2 added to stack on pos: 1, stack has now len:2"
    expected_logs_content[2] = "Tile from tileset 0 (0, 0) on layer 0 added to stack on pos: 2, stack has now len:3"
    expected_logs_content[3] = "Tile 2 put on (5, 5), on layer 0"
    expected_logs_content[4] = "Tile 0 put on (4, 5), on layer 1"
    expected_logs_content[5] = "Tile 0 put on (6, 5), on layer 1"
    expected_logs_content[6] = "Tile 0 put on (5, 4), on layer 1"
    expected_logs_content[7] = "Tile 0 put on (5, 6), on layer 1"
    expected_logs_content[8] = "Tile 1 put on (5, 5), on layer 2"

    proper_prefix := make(map[int]string)

    proper_prefix[0] = "STACKK:"
    proper_prefix[1] = "STACKK:"
    proper_prefix[2] = "STACKK:"
    proper_prefix[3] = "CANVAS:"
    proper_prefix[4] = "CANVAS:"
    proper_prefix[5] = "CANVAS:"
    proper_prefix[6] = "CANVAS:"
    proper_prefix[7] = "CANVAS:"
    proper_prefix[8] = "CANVAS:"

    n := 0

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if line[:LOG_PREFIX_N] != proper_prefix[n] {
            t.Errorf("bad prefix! want %s get: %s", proper_prefix[n], line[:LOG_PREFIX_N])
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
    err = os.Remove("../../logs/" + lg.filename)
    if err != nil { t.Error(err) }

}
func TestNewLogger(t *testing.T) {
    log := NewLogger()
    
    dt := time.Now().Format("2006-01-02_15:04:05")

    // cutting is needed as logger filename contains seconds, and it is hard to grasp so
    if dt != log.filename[:19] {
        t.Errorf("adding logger filename does not succeeded!")
    }

    if log.file == nil {
        t.Errorf("creating logger file does not succeeded!")
    }

    if _, err := os.Stat("/logs/" + log.filename); err == nil || os.IsExist(err) {
        t.Errorf("creating logger file does not succeeded!")
    }
}

