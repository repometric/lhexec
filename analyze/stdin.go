package analyze

import (
	"io/ioutil"
	"log"
	"bytes"
	"os"
)

//This function performs the stdin analysis by the engine
func AnalyzeStdin(engine string, content []rune) (string, error)  {
	tmpfile, err := ioutil.TempFile("./", "stdin_temp")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	var buf bytes.Buffer
	for i := 0; i < len(content); i++ {
		buf.WriteRune(content[i])
	}
	if _, fileError := tmpfile.Write(buf.Bytes()); err != nil {
		log.Fatal(fileError)
	}
	out, err := Execute([]string{engine, tmpfile.Name()})
	if fileError := tmpfile.Close(); fileError != nil {
		log.Fatal(fileError)
	}
	return out, err
}
