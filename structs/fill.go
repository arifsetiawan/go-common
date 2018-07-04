package structs

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

func Fill(dest interface{}, src interface{}) {
	mSrc := structs.Map(src)
	mDest := structs.Map(dest)
	for key, val := range mSrc {
		if _, ok := mDest[key]; ok {
			structs.New(dest).Field(key).Set(val)
		}
	}
}

func MapToStruct(in map[string]interface{}, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}

	return nil
}

func StructToMap(in interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

// camelCase is
func camelCase(s string) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	var prev rune
	for _, curr := range s {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || (prev == 0) {
				buffer = append(buffer, toUpper(curr))
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
		prev = curr
	}

	return strings.Replace(string(buffer), "Id", "ID", -1)
}

// ValueByTag is
func ValueByTag(i interface{}, t string) interface{} {
	return reflect.Indirect(reflect.ValueOf(i)).FieldByName(camelCase(t)).Interface()
}
