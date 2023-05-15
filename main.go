package main

import (
	"fmt"
	"os"
	"flag"
	"log"
)

const usage = `Usage csv2json [option] <csvFile>
	-s, --separator: Specify the separtor the csv uses, defaults to ","
	-p, --pretty: Specify whether or not to prettify the json output
`

const colorReset = "\033[0m"
const colorRed = "\033[31m"

func usageAndPanic(msg string) {
	flag.Usage()
	log.Println(string(colorRed), msg, string(colorReset))
	os.Exit(1)
}

func main() {
	var sep string
	var pretty bool
	
	flag.Usage = func() { fmt.Print(usage) }
	
	flag.StringVar(&sep, "separator", ",", "Specify the separtor the csv uses")
	flag.StringVar(&sep, "s", ",", "Specify the separtor the csv uses")
	
	flag.BoolVar(&pretty, "pretty", false, "Specify whether or not to prettify the json output")
	flag.BoolVar(&pretty, "p", false, "Specify whether or not to prettify the json output")
	
	flag.Parse()
	
	restArgs := flag.Args()
	
	if len(restArgs) == 0 {
		usageAndPanic("A csv file is required to carry out the operation")
	}
	
	// Make sure the file exists
	if _, err := os.Stat(restArgs[0]); err != nil {
		usageAndPanic("Specified csv file does not exist")
	}
	
	if pretty {
		fmt.Println(convertPretty(restArgs[0], sep, 4))
	} else {
		fmt.Println(convert(restArgs[0], sep))
	}
}
