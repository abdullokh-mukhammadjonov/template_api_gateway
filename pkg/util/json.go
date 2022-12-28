package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func PrintJson(v interface{}) {
	beautifulJsonByte, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(beautifulJsonByte))
}

func ToJsonFile(path string, v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", " ")

	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Saved the data as json file at " + path)
}

func JSONStringify(obj interface{}) string {
	beautifulJsonByte, err := json.MarshalIndent(obj, "", "  ")
	body := ""
	if err != nil {
		body = fmt.Sprintf("%v", obj)
	} else {
		body = string(beautifulJsonByte)
	}
	return body
}
