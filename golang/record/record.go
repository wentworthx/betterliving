package record

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Record contains information need to store and display
type Record struct {
	Project string            `json:"project"`
	Amount  string            `json:"amount"`
	Time    string            `json:"time"`
	Details map[string]string `json:"details"`
}

// Append project, amount and details
func Append(project string, amount int, details map[string]string) (finished bool, err error) {
	file, date, _ := Where()

	var records map[string][]*Record
	recordsByte, _ := ReadRecords(file)
	json.Unmarshal(recordsByte, &records)

	currentRecord := records[date]
	if currentRecord == nil {
		records[date] = []*Record{}
	}
	records[date] = append(records[date], &Record{
		Project: project,
		Amount:  strconv.Itoa(amount),
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	})
	for date, record := range records {
		fmt.Println(date, record)
	}
	recordsNewByte, _ := json.MarshalIndent(&records, "", "    ")
	fmt.Println(string(recordsNewByte))

	writeDone, errWrite := WriteRecords(file, recordsNewByte)
	return writeDone, errWrite
}
