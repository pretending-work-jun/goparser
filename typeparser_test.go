package goparser_test

import (
	"testing"

	"github.com/jjkoh95/goparser"
)

const testYAMLParserFile = "./typeparser_test.yaml"

func TestYAMLParser(t *testing.T) {
	var typeParser goparser.TypeParser

	typeParser.GetParserFromYAML(testYAMLParserFile)

	if typeParser.ParseFunc["testString"]("test") != "test" {
		t.Error("Expected to parse string correctly")
	}

	if typeParser.ParseFunc["testInt"]("123") != 123 {
		t.Error("Expected to parse int correctly")
	}

	if typeParser.ParseFunc["testFloat"]("0.001") != 0.001 {
		t.Error("Expected to parse float correctly")
	}

	if typeParser.ParseFunc["testBool"]("true") != true {
		t.Error("Expected to parse boolean correctly")
	}
}

func TestSmartParser(t *testing.T) {
	var testVal = [][]string{
		{"testString", "testInt", "testBool", "testNull", "testFloat"},
		{"testString1", "1", "true", "", "0.001"},
		{"testString2", "11", "false", "null", "0.002"},
	}

	var typeParser goparser.TypeParser
	typeParser.GetSmartParser(testVal)

	if typeParser.ParseFunc["testString"]("test") != "test" {
		t.Error("Expected to parse string correctly")
	}

	if typeParser.ParseFunc["testInt"]("123") != 123 {
		t.Error("Expected to parse int correctly")
	}

	if typeParser.ParseFunc["testNull"]("0.001") != "0.001" {
		t.Error("Expected to parse null correctly")
	}

	if typeParser.ParseFunc["testBool"]("true") != true {
		t.Error("Expected to parse boolean correctly")
	}
}
