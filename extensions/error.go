package extensions

import (
	"log"
	"os"
)

const errorCode = 1

//PathValidate - check path
func PathValidate(path string) bool {
	if len(path) == 0 {
		ShowError("File path doesn't find")
	}
	return true
}

//CheckError - check Error
func CheckError(e error) {
	if e != nil {
		ShowError(e.Error())
	}
}

//ShowError - show Error
func ShowError(err string) {
	log.Fatal("Error: " + err)
	os.Exit(errorCode)
}
