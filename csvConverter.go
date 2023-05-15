package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fetchHeader(scanner *bufio.Scanner, separator string) []string {
	// Assumes that the scanner already advanced by the first scanner call
	// to check that the csv is not empty
	var splittedStr []string
	
	if scanner.Scan() {
		line := scanner.Text()
		splittedStr = strings.Split(line, separator)
	}
	return splittedStr
}

func checkErorr(e error) {
	if e != nil {
		panic(e)
	}
}

func convertPrettyDepth(headers []string, splittedText []string, spacing int, level int) string {
	// Meant to be reusable if somehow the json is more nested
	var outputText string
	
	// Handles the spaces
	outputText += strings.Repeat(" ", spacing * level)
	outputText += "{\n"
	
	// Stuff inside needs level + 1
	for index, col := range splittedText {
		outputText += strings.Repeat(" ", spacing * (level + 1))
		outputText += fmt.Sprintf("\"%s\": \"%s\",\n", headers[index], col)
	}
	outputText = outputText[:len(outputText) - 2] + "\n"
	outputText += strings.Repeat(" ", spacing * level)
	outputText += "},\n"
	
	return outputText
}

func convertPretty(filename string, separator string, spacing int) string {
	var outputText string
	
	f, err := os.Open(filename)
	checkErorr(err)
	scanner := bufio.NewScanner(f)
	
	outputText += "[\n"
	
	headers := fetchHeader(scanner, separator)
	
	if len(headers) == 0 {
		return "[]"
	}
	
	for scanner.Scan() {
		line := scanner.Text()
		splittedText := strings.Split(line, separator)
		
		outputText += convertPrettyDepth(headers, splittedText, spacing, 1)
	}
	
	outputText = outputText[:len(outputText) - 2]
	outputText += "\n]"
	return outputText
}

func convert(filename string, separator string) string {
	var outputText string
	outputText += "["
	
	f, err := os.Open(filename)
	checkErorr(err)
	scanner := bufio.NewScanner(f)
	
	headers := fetchHeader(scanner, separator)
	
	if len(headers) == 0 {
		return "[]"
	}
	
	for scanner.Scan() {
		line := scanner.Text()
		splittedText := strings.Split(line, separator)
		outputText += "{"
		
		for index, col := range splittedText {
			outputText += fmt.Sprintf("\"%s\":\"%s\",", headers[index], col)
		}
		outputText = outputText[:len(outputText) - 1] + "},"
	}
	
	outputText = outputText[:len(outputText) - 1] + "]"
	return outputText
}
