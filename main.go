package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

/*
  Execute a shell command
*/
func shell(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", "msgcat "+command+" || true")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

/*
  Extract a few lines inside a file
  for the match one we append >>> before to add some emphasis for this one
*/
func LinesInFile(fileName string, from int, to int, match int) []string {
	f, err := os.Open(fileName)
	result := []string{}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	n := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n++
		if n < from {
			continue
		}
		if n > to {
			break
		}
		line := scanner.Text()

		if n == match {
			result = append(result, ">>>  "+strconv.Itoa(n)+": "+line)
		} else {
			result = append(result, "     "+strconv.Itoa(n)+": "+line)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return result
}

func main() {
	input := flag.String("input", "", "filePath to the pot file we want to lint")
	flag.Parse()

	if len(*input) == 0 {
		log.Fatal("You must specify a file via -input")
	}

	fmt.Printf("[+] lint file %s\n", *input)
	err, out, errout := shell(*input)
	if err != nil {
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
		log.Fatal(err)
	}

	// No error detected -> we're free to go
	if len(errout) == 0 {
		return
	}

	// We look for the line with the format: <linenumber>:<columnnumber> as it's the most important one
	re := regexp.MustCompile(`[0-9]+:[0-9]+`)
	match := re.Find([]byte(errout))

	config := strings.Split(string(match), ":")
	lineNumber, err := strconv.Atoi(config[0])

	if err != nil {
		log.Fatal(err)
	}

	beforeMatchNumber := lineNumber - 5
	afterMatchNumber := lineNumber + 5

	fmt.Printf("Error inside the file:\n  %q\n\nDetails:\n", strings.Split(errout, "\n")[0])
	for _, line := range LinesInFile(*input, beforeMatchNumber, afterMatchNumber, lineNumber) {
		fmt.Println(line)
	}

	os.Exit(1)
}
