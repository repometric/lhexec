package analyze

import (
	"log"
	"io/ioutil"
	"go/scanner"
)
//This function performs the context(Folder or Files) analysis by the engine
func Analyze(engine string, context Context) (string, error) {
	input := []string{engine}
	var projectPath = ""
	if context.Project!="" {
		projectPath = context.Project+"/"
	}
	if context.Folder!="" {
		files, err := ioutil.ReadDir(projectPath+context.Folder)
		if err != nil {
			log.Fatal(err)
		}
		for _, file:=range files {
			if !file.IsDir(){
				input = append(input,projectPath+context.Folder+"/"+file.Name())
			}
		}

	} else if context.File!="" {
			input = append(input, projectPath+context.File)
	} else {
		return "", scanner.Error{}
	}
	out, err := Execute(input)
	return out, err
}
