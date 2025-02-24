package main

import (
	"os"
)

func saveData(path string, data []byte) error {
	println("save file to disk")
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return nil
	}
	return fp.Sync()
}

func main() {
	println("Hello go world")
}
