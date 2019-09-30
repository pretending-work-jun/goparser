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

// GetSmartParser - infer parser by looping through [][]string
func (parser *TypeParser) GetSmartParser(val [][]string) {
	keys := val[0] // header is always key
	parser.ParseFunc = make(map[string]ParseFuncType)
	for col := 0; col < len(keys); col++ {
		var hasString, hasFloat, hasInt, hasBool bool
		for row := 1; row < len(val); row++ {
			if (hasBool && hasInt) || (hasBool && hasFloat) {
				break
			}
			if val[row][col] == "" || strings.ToLower(val[row][col]) == "null" {
				continue
			}
			if _, err := strconv.Atoi(val[row][col]); err == nil {
				hasInt = true
				continue
			}
			if _, err := strconv.ParseFloat(val[row][col], 64); err == nil {
				hasFloat = true
				continue
			}
			if strings.ToLower(val[row][col]) == "true" || strings.ToLower(val[row][col]) == "false" {
				hasBool = true
				continue
			}
			hasString = true
			break
		}
		if hasString || (hasBool && hasInt) || (hasBool && hasFloat) {
			parser.ParseFunc[keys[col]] = parser.parseString
		} else if hasFloat {
			parser.ParseFunc[keys[col]] = parser.parseFloat
		} else if hasInt {
			parser.ParseFunc[keys[col]] = parser.parseInt
		} else if hasBool {
			parser.ParseFunc[keys[col]] = parser.parseBool
		}
	}
}
