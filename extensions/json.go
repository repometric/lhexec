package extensions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//GetObjectInFile - get file and convert contains to object
func GetObjectInFile(path string, object interface{}) interface{} {
	PathValidate(path)
	raw, err := ioutil.ReadFile(path)

	CheckError(err)

	json.Unmarshal(raw, &object)
	return object
}

//ConverObjectToJSON - Convert object to JSON,and print or not
func ConverObjectToJSON(object interface{}, println bool) string {
	bytesJSON, _ := json.MarshalIndent(object, "", "	")
	var result = string(bytesJSON)
	if println {
		fmt.Println(result)
	}
	return result
}

//SaveObjectToFileJSON - Save some object to file in JSON format
func SaveObjectToFileJSON(path string, object interface{}) bool {
	err := os.Remove(path)
	CheckError(err)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	CheckError(err)
	f.WriteString(ConverObjectToJSON(object, false))
	defer f.Close()
	return true
}
