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
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type Ping struct {
	average time.Duration
}

func main() {
	cmd := exec.Command("ping", "8.8.8.8")
	// Linux version
	//cmd := exec.Command("ping", "-c 4", "8.8.8.8")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	printCommand(cmd)
	err := cmd.Run()
	printError(err)
	output := cmdOutput.Bytes()
	printOutput(output)
	ping := Ping{}
	parseOutput(output, &ping)
	writeTestFile(output)

	fmt.Println(ping)
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

func parseOutput(outs []byte, ping *Ping) {
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

func writeTestFile(outs []byte) {
	layout := "20060102_1504"
	//If you want to create tmp/test directories at the root, use /tmp/test
	err := os.MkdirAll("tmp/test", 0755)
	if err != nil {
		fmt.Println("Error making folder:", err)
		return
	}
	logFile, err := os.Create("tmp/test/" + time.Now().Format(layout) + ".log")
	if err != nil {
		fmt.Println("Cannot create logfile:", err)
		return
	}

	_, err = logFile.Write(outs)
	if err != nil {
		fmt.Println("Cannot write to logfile: ", err)
	}
}
