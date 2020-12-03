package record

import (
	"fmt"
	"io/ioutil"
)

// WriteRecords from data file
func WriteRecords(file string, recordsByte []byte) (done bool, err error) {
	errWrite := ioutil.WriteFile(file, recordsByte, 0666)
	if err != nil {
		fmt.Println("write fail:", file)
		return false, errWrite
	}
	fmt.Println("write success:", file)
	return true, nil
}
