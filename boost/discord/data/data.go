package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// returns whether the given file exists
func DataExists(filePath string) bool {
	_, err := os.Stat(fmt.Sprintf("./discord/db/%v", filePath))
	return !os.IsNotExist(err)
}

// loads the data from the given json file from db/ folder to the given data structure
func LoadData(filePath string, data interface{}) error {
	f, err := os.Open(fmt.Sprintf("./discord/db/%v", filePath))
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// saves the given data to a json file from db/ folder
func SaveData(filePath string, data interface{}) error {
	f, err := os.Create(fmt.Sprintf("./discord/db/%v", filePath))
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func DeleteData(filePath string) error {
	err := os.Remove(fmt.Sprintf("./discord/db/%v", filePath))
	if err != nil {
		return err
	}

	return nil
}
