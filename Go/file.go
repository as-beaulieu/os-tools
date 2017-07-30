package main

import (
	"fmt"
	"os"
	"time"
)

func writeToNewFile(outs []byte) string {
	layout := "20060102_1504"
	//If you want to create tmp/test directories at the root, use /tmp/test
	err := os.MkdirAll("tmp/test", 0755)
	if err != nil {
		fmt.Println("Error making folder:", err)
		return ""
	}

	newPath := "tmp/test/" + time.Now().Format(layout) + ".log"
	logFile, err := os.Create(newPath)
	if err != nil {
		fmt.Println("Cannot create logfile:", err)
		return ""
	}

	_, err = logFile.Write(outs)
	if err != nil {
		fmt.Println("Cannot write to logfile: ", err)
	}

	// save changes
	err = logFile.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return "" //same as above
	}

	return newPath

}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	checkError(err)
}

func readFile() {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// read file
	var text = make([]byte, 1024)
	n, err := file.Read(text)
	if n > 0 {
		fmt.Println(string(text))
	}
	//if there is an error while reading
	//just print however much was read if any
	//at return file will be closed
}
