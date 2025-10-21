package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)


type User struct {
	ID      int
	Name    string // must be Name not name :D, won't work
	Age     int
	Address Address
}


	// Count how many files are in users_saved directory
func getNumOfFiles() int {
	files, err := ioutil.ReadDir("users_saved")
	if err != nil {
		// if dir doesn’t exist, create it and return 0
		if os.IsNotExist(err) {
			os.Mkdir("users_saved", os.ModePerm)
			return 0
		}
		panic(err)
	}
	return len(files)
}

var nextID = getNumOfFiles() + 1


func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++

	// create file 
	filename := fmt.Sprintf("users_saved/%d.txt", u.ID)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// convert user to JSON
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	// write to file
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}

	return u, nil
}

func GetUserByID(id int) (User, error) {

	filename := fmt.Sprintf("users_saved/%d.txt", id)

	// read file
	stream, err := ioutil.ReadFile(filename)
	if err != nil {
		return User{}, fmt.Errorf("user with ID '%v' not found", id)
	}

	// unmarshal JSON
	var u User
	err = json.Unmarshal(stream, &u)
	if err != nil {
		return User{}, err
	}



	return u, nil
}
