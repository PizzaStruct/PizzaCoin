package helpers

import (
	"encoding/json"
	"io/ioutil"
)

func UpdateJsonFile(data interface{}, file string) {
	res, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(file, res, 0)
	if err != nil {
		panic(err)
	}
}
