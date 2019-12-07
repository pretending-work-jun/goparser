package goparser

// ParseInterfaceToString - This makes [][]interface{} much easier to work with.
// Caution - Interface must be `string` (data received from spreadsheet etc).
func ParseInterfaceToString(val [][]interface{}) [][]string {
	resp := make([][]string, len(val))
	for i := 0; i < len(val); i++ {
		resp[i] = make([]string, len(val[i]))
		for j := 0; j < len(val[i]); j++ {
			resp[i][j] = val[i][j].(string)
		}
	}
	return resp
}

// ParseRowToList - Convert [][]string to []map[string] (list<key, val>).
func ParseRowToList(val [][]string, typeParser *TypeParser) []map[string]interface{} {
	keys := val[0] // header is always key

	obj := make([]map[string]interface{}, len(val)-1) // this is the object we want to return
	for row := 1; row < len(val); row++ {
		obj[row-1] = make(map[string]interface{})
		for col := 0; col < len(keys); col++ {
			if val[row][col] == "" || val[row][col] == "null" {
				obj[row-1][keys[col]] = nil
				continue
			}
			parsedVal := typeParser.ParseFunc[keys[col]](val[row][col])
			obj[row-1][keys[col]] = parsedVal
		}
	}

	return obj
}

// ParseRowToMap - Parse rows to map, note that keyCol must be unique otherwise it will be overridden by the last row.
func ParseRowToMap(val [][]string, keyCol string, typeParser *TypeParser) map[string]map[string]interface{} {
	res := make(map[string]map[string]interface{})

	keys := val[0] // header is always key

	var keyColIndex int
	for i, v := range keys {
		if v == keyCol {
			keyColIndex = i
			break
		}
	}

	for row := 1; row < len(val); row++ {
		var key string
		rowVal := make(map[string]interface{})
		for col := 0; col < len(keys); col++ {
			if col == keyColIndex {
				key = val[row][col]
				continue
			}
			if val[row][col] == "" || val[row][col] == "null" {
				rowVal[keys[col]] = nil
				continue
			}
			rowVal[keys[col]] = typeParser.ParseFunc[keys[col]](val[row][col])
		}
		res[key] = rowVal
	}

	return res
}
