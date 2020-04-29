package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/ogier/pflag"
)

// Flags
var options struct {
	format  string
	version bool
	source  string
	help    bool
}

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s[format]\n", os.Args[0])
		pflag.PrintDefaults()
	}

	// Bind flags to variables
	pflag.StringVarP(&options.format, "format", "f", "text", "Enter the output format for campaign information.")
	pflag.BoolVarP(&options.version, "version", "v", false, "View application version information.")
	pflag.StringVarP(&options.source, "source", "i", "gloomhaven.json", "Enter the source save file for campaign.")
	pflag.BoolVar(&options.help, "help", false, "Usage help for tool.")
	pflag.Parse()

	var strFormat *string
	strFormat = &options.format

	if strFormat != nil {
		if *strFormat != "text" {
			fmt.Printf("%s\n", "Only text format is supported at this time.\n")
			os.Exit(0)
		}

		fmt.Printf("Text export selected.\n")

		data, err := ioutil.ReadFile(options.source)
		if err != nil {
			panic(err)
		}

		saveDate, err := jsonparser.GetString(data, "Date")

		if err != nil {
			panic(err)
		}

		outputFile := fmt.Sprintf("Save Date: %s\n", saveDate)

		//scriptState, err := jsonparser.GetString(data, "LuaScriptState")
		//scriptState, err := jsonparser.GetString(data, "ObjectStates", "[1]", "Name")

		/*
			var handler func([]byte, []byte, jsonparser.ValueType, int) error
			handler = func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				fmt.Printf("Key: '%s'\n", string(key))
				switch string(key) {
				case "ObjectStates":
					//fmt.Printf("Value: %s\n", string(value)[:16])
					jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err error) {
						fmt.Printf("    Array Value: '%s'\n", string(value2)[:8]))

					}
				}
				return nil
			}

			jsonparser.ObjectEach(data, handler)
		*/

		jsonparser.ArrayEach(data, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err error) {
			aVal, err := jsonparser.GetString(value2, "Description")
			//fmt.Printf("    Array Name: '%s'\n", aVal)
			if aVal == "Party Sheet" {
				partyName, err := jsonparser.GetString(value2, "Nickname")

				if err != nil {
					panic(err)
				}

				outputFile += "\n=== Party Sheet ===\n"
				outputFile += fmt.Sprintf("Party Name: %s\n", partyName)

				scriptState, err := jsonparser.GetString(value2, "LuaScriptState")

				if err != nil {
					panic(err)
				}

				cleanScriptState, err := strconv.Unquote(scriptState)
				outputFile += fmt.Sprintf("Script State: %s\n", cleanScriptState)
			}
		}, "ObjectStates")

		if err != nil {
			panic(err)
		}

		//fmt.Printf("Script State : %s\n", scriptState)

		f, err := os.Create("output.txt")
		if err != nil {
			f.Close()
			panic(err)
		}

		l, err := f.WriteString(outputFile)
		if err != nil {
			f.Close()
			panic(err)
		}

		fmt.Printf("%d bytes written successfully.\n", l)
		err = f.Close()
		if err != nil {
			panic(err)
		}

	}

	if options.version {
		fmt.Printf("%s", "gloomhaven-exporter version 0.0.1\n")
		os.Exit(0)
	}

	if options.help {
		Usage()
		os.Exit(0)
	}
}
