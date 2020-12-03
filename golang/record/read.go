package record

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadRecords from data file
func ReadRecords(file string) (recordsByte []byte, err error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("read file(%s) failed", file, err)
		return []byte{}, err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd failed")
		return []byte{}, err
	}

	return fd, nil
}
