package goparser_test

import (
	"testing"

	"github.com/jjkoh95/goparser"
)

func TestParseInterfaceToString(t *testing.T) {
	var val = [][]interface{}{
		{"abc", "", "true", "true"},
		{"123", "0.011", "true", "false"},
	}

	result := goparser.ParseInterfaceToString(val)

	if result[0][0] != "abc" {
		t.Error("Expected string `abc`")
	}

	if result[0][1] != "" {
		t.Error("Expected string ``")
	}

	if result[0][2] != "true" {
		t.Error("Expected string `true`")
	}

	if result[0][3] != "true" {
		t.Error("Expected string `true`")
	}

	if result[1][0] != "123" {
		t.Error("Expected string `123`")
	}

	if result[1][1] != "0.011" {
		t.Error("Expected string `0.011`")
	}
}

var rows = [][]string{
	{"testKey", "col1", "col2", "col3"},
	{"A001", "0.001", "1", "true"},
	{"B001", "1", "33", "false"},
	{"C001", "1.1", "88", "null"},
}

func TestRowToList(t *testing.T) {
	var typeParser goparser.TypeParser
	typeParser.GetSmartParser(rows)

	result := goparser.ParseRowToList(rows, &typeParser)

	if result[0]["testKey"] != "A001" {
		t.Error("Expected testKey - A001")
	}

	if result[0]["col1"] != 0.001 {
		t.Error("Expected col1 - 0.001")
	}

	if result[1]["col2"] != 33 {
		t.Error("Expected col2 - 33")
	}

	if result[2]["col3"] != nil {
		t.Error("Expected col3 - null")
	}
}

func TestParseRowToMap(t *testing.T) {
	var typeParser goparser.TypeParser
	typeParser.GetSmartParser(rows)

	result := goparser.ParseRowToMap(rows, "testKey", &typeParser)

	if result["A001"]["col1"] != 0.001 {
		t.Error("Expected [A001][col1] to be 0.001")
	}

	if result["B001"]["col2"] != 33 {
		t.Error("Expected [B001][col2] to be 33")
	}

	if result["B001"]["col3"] != false {
		t.Error("Expected [B001][col3] to be false")
	}

	if result["C001"]["col3"] != nil {
		t.Error("Expected [C001][col3] to be null")
	}
}
