package stdlib_json

import (
	"encoding/json"
	"os"
)

type Users []User

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func (u *Users) LoadUserData(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) WriteDataToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
