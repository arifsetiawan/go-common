package persistence

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"os"
	"path/filepath"
)

// Dump is
func Dump(dataPath string, name string, data interface{}) error {
	return dumpJSON(dataPath, name, data)
}

// Restore is
func Restore(dataPath string, name string, data interface{}) error {
	return restoreJSON(dataPath, name, data)
}

// dump whole data as JSON
func dumpJSON(dataPath string, name string, data interface{}) error {

	absPath, _ := filepath.Abs(dataPath + "/" + name + ".json")
	fi, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer fi.Close()

	encoder := json.NewEncoder(fi)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// restore whole data from JSON
func restoreJSON(dataPath string, name string, data interface{}) error {

	absPath, _ := filepath.Abs(dataPath + "/" + name + ".json")
	fi, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer fi.Close()

	decoder := json.NewDecoder(fi)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// dump whole data as gob
func dumpGob(dataPath string, name string, data interface{}) error {

	absPath, _ := filepath.Abs(dataPath + "/" + name)
	fi, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer fi.Close()

	fz := gzip.NewWriter(fi)
	defer fz.Close()

	encoder := gob.NewEncoder(fz)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// restore whole data from gob
func restoreGob(dataPath string, name string, data interface{}) error {

	absPath, _ := filepath.Abs(dataPath + "/" + name)
	fi, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return err
	}
	defer fz.Close()

	decoder := gob.NewDecoder(fz)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}
