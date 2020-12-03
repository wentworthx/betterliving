package record

import (
	"fmt"
	"time"
)

// Where to record
func Where() (filePath string, datePoint string, err error) {
	dataDir := time.Now().Format("2006")
	dataFileName := time.Now().Format("200601")
	dataDatePoint := time.Now().Format("20060102")
	file := fmt.Sprintf("../data/%s/%s.json", dataDir, dataFileName)

	return file, dataDatePoint, nil
}
