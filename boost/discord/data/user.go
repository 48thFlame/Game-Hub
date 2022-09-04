package data

import "fmt"

type User struct {
	Stats struct {
		Mastermind struct {
			Wins   int `json:"wins"`
			Losses int `json:"losses"`
			Rounds int `json:"rounds"`
		} `json:"mastermind"`
	} `json:"stats"`
	Feedback bool `json:"feedback"` // whether is banned from using the feedback command
}

func GetUserFileName(id string) string {
	return fmt.Sprintf("users/%v.json", id)
}

func LoadUser(id string) (*User, error) {
	user := &User{}

	fileName := GetUserFileName(id)
	if DataExists(fileName) {
		err := LoadData(fileName, user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
