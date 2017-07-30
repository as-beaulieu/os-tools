// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"os/exec"
// 	"time"
// )

// func main() {
// layout := "20060102_1504"
// err := os.MkdirAll("/tmp/test", 0755)
// if err != nil {
// 	fmt.Println("Error making folder:", err)
// 	return
// }
// logFile, err := os.Create("/tmp/test/" + time.Now().Format(layout) + ".log")
// if err != nil {
// 	fmt.Println("Cannot create logfile:", err)
// 	return
// }

// fmt.Println("Starting to run echo")
// cmd := exec.Command("ping", "8.8.8.8")
// cmd.Stdout = io.MultiWriter(logFile, os.Stdout)
// cmd.Stderr = cmd.Stdout
// if err := cmd.Run(); err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println("Finished running")
// }

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"time"
)

type Ping struct {
	average time.Duration
}

type SystemInformation struct {
	hostName       string
	hostNameValue  string
	osName         string
	osNameValue    string
	osVersion      string
	osVersionValue string
}

var path = "test.txt"

func main() {

	cmd := exec.Command("systeminfo")
	//cmd := exec.Command("ping", "8.8.8.8")
	// Linux version
	//cmd := exec.Command("ping", "-c 4", "8.8.8.8")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	printCommand(cmd)
	err := cmd.Run()
	printError(err)
	output := cmdOutput.Bytes()
	//printOutput(output)
	//ping := Ping{}
	//parsePing(output, &ping)
	pathName := writeToNewFile(output)
	fmt.Println(pathName)

	// open a file, scan each line

	scannedFile := fileScanner(pathName)

	//fileScan[0] is blank
	//fmt.Println(scannedFile[1])
	newSysInfo := SystemInformation{}
	s := strings.Split(scannedFile[1], ":")
	//strings.TrimSpace removes all whitespace on beginning and end of the value
	newSysInfo.hostName = strings.TrimSpace(s[0])
	newSysInfo.hostNameValue = strings.TrimSpace(s[1])

	s = strings.Split(scannedFile[2], ":")
	newSysInfo.osName = strings.TrimSpace(s[0])
	newSysInfo.osNameValue = strings.TrimSpace(s[1])

	s = strings.Split(scannedFile[3], ":")
	newSysInfo.osVersion = strings.TrimSpace(s[0])
	newSysInfo.osVersionValue = strings.TrimSpace(s[1])

	fmt.Println("Inside the struct: ", newSysInfo)

	//fmt.Println(ping)
	//getExeDirectory()

}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func parsePing(outs []byte, ping *Ping) {
	var average = regexp.MustCompile(`Average = (\d+ms)`)
	result := average.FindStringSubmatch(string(outs))

	if len(result) > 0 {
		ping.average, _ = time.ParseDuration(result[1])
	}
	// Linux version
	/*var average = regexp.MustCompile(`min\/avg\/max\/mdev = (0\.\d+)\/(0\.\d+)\/(0\.\d+)\/(0\.\d+) ms`)
	  result := average.FindAllStringSubmatch(string(outs), -1)

	  if len(result) > 0 {
	          ping.average, _ = time.ParseDuration(result[0][2] + "ms")
	  }*/
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func getHomeDirectory() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)
}

//Returns the directory path the main.go file runs from
func getExeDirectory() {
	dir, _ := os.Getwd()
	fmt.Println("Plain exe path: ", dir)
	fmt.Println(strings.Replace(dir, " ", "\\ ", -1))
}
