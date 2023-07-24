package fileManager

import (
	"io/ioutil"
)

func WriteDataFile(pathFile string, dataJson []byte) error {
	// Create file
	err := ioutil.WriteFile(pathFile, dataJson, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadDataFile(pathFile string) ([]byte, error) {
	data, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}
	return data, nil
}
