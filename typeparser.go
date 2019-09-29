package goparser

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseFuncType - wrap parsing function type
type ParseFuncType = func(s string) interface{}

func (parser *TypeParser) parseFloat(s string) interface{} {
	parsedVal, _ := strconv.ParseFloat(s, 64)
	return parsedVal
}

func (parser *TypeParser) parseInt(s string) interface{} {
	parsedVal, _ := strconv.Atoi(s)
	return parsedVal
}

func (parser *TypeParser) parseBool(s string) interface{} {
	return strings.ToLower(s) == "true"
}

func (parser *TypeParser) parseString(s string) interface{} {
	return s
}

// TypeParser - parse type smartly
type TypeParser struct {
	ParseFunc map[string]ParseFuncType
}

// ParseTypeYAML - expected format from YAML file
type ParseTypeYAML struct {
	ColName string
	ColType string
}

// GetParserFromYAML - smartly make the parsing functions from YAML file
func (parser *TypeParser) GetParserFromYAML(filename string) {
	YAMLFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error opening file")
	}

	defer YAMLFile.Close()

	YAMLByte, err := ioutil.ReadAll(YAMLFile)
	if err != nil {
		log.Fatalln("Error reading file")
	}

	var parseType []ParseTypeYAML

	err = yaml.Unmarshal(YAMLByte, &parseType)
	if err != nil {
		log.Fatalln("Error parsing YAML")
	}

	parser.ParseFunc = make(map[string]ParseFuncType)
	for _, v := range parseType {
		switch v.ColType {
		// case "string":
		// 	parser.ParseFunc[v.ColName] = parser.parseString
		case "int":
			parser.ParseFunc[v.ColName] = parser.parseInt
		case "float":
			parser.ParseFunc[v.ColName] = parser.parseFloat
		case "boolean":
			parser.ParseFunc[v.ColName] = parser.parseBool
		default:
			parser.ParseFunc[v.ColName] = parser.parseString
		}
	}
}
