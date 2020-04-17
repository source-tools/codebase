package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var triggers = []string{"void", "def", "function", "func", "int", "var", "double", "float", "string", "String", "bool", "boolean"}
var commentSyntaxes = []string{"#", "//", "/*"}
var varSyntaxes = []string{"int", "var", "double", "float", "string", "String", "bool", "boolean"}

func getFileLines(filepath string) []string {
	currentfile, err := os.Open(filepath)

	if err != nil {

	}

	fileScanner := bufio.NewScanner(currentfile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	currentfile.Close()

	return lines
}

func containsTrigger(line string) bool {

	var i int = 0
	for i < len(triggers) {
		if strings.Contains(line, triggers[i]) {
			return true
		}
		i++
	}

	return false
}

func writeLineToFile(filepath string, str string) {

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if _, err := file.WriteString(str + "\n"); err != nil {
		log.Fatal(err)
	}
	if err != nil {

	}

	file.Close()
}

func startsWith(line string, term string) bool {

	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, "")

	if len(lineArr) < 1 {
		return false
	}

	termArr := strings.Split(term, "")

	var starts bool = true
	var i int = 0
	for i < len(termArr) {
		if termArr[i] != lineArr[i] {
			starts = false
		}

		i++
	}

	return starts
}

func startsWithComment(line string) bool {

	var i int = 0
	for i < len(commentSyntaxes) {

		if strings.Contains(line, commentSyntaxes[i]) {
			return true
		}

		i++
	}

	return false
}

func createFile(path string) {
	os.Create(path)
}

func containsVar(line string) bool {

	var i int = 0
	for i < len(varSyntaxes) {
		if strings.Contains(line, varSyntaxes[i]) {
			return true
		}
		i++
	}

	return false
}

func getVarValue(line string) string {

	if !strings.Contains(line, "=") {
		return ""
	}

	arr := strings.Split(line, "=")
	return strings.TrimSpace(arr[1])
}

func isFunction(line string) bool {

	if strings.Contains(line, "(") {
		if strings.Contains(line, ")") {
			return true
		}
	}

	return false
}

func getRawParameters(line string) string {

	var parameters string = ""
	var run bool = false
	arr := strings.Split(line, "")
	var i int = 0

	for i < len(arr) {

		if arr[i] == ")" {
			run = false
			return parameters
		}

		if run {
			parameters = parameters + arr[i]
		}

		if arr[i] == "(" {
			run = true
		}
		i++
	}

	return parameters
}
