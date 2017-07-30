package main

import (
	"bufio"
	"log"
	"os"
)

//readLines() reads the file, and is used to return the file as one large string
//Inefficient for large files. Better to use a scanner to organize line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//fileScanner is a scanner that scans the file, moving each line to a slice
func fileScanner(pathName string) []string {
	var fileScan []string
	if file, err := os.Open(pathName); err == nil {

		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//log.Println(scanner.Text())
			fileScan = append(fileScan, scanner.Text())
		}

		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}

	return fileScan

}
