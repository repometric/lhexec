package analyze

// RunStdin function performs the stdin analysis by the engine
func RunStdin(context Context) (string, error) {
	/*tmpfile, err := ioutil.TempFile(context.Project, "temp")

	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up
	if _, fileError := tmpfile.WriteString(context.StdinContent); err != nil {
		log.Fatal(fileError)
	}
	out, err := Execute([]string{engine, tmpfile.Name()})
	if fileError := tmpfile.Close(); fileError != nil {
		log.Fatal(fileError)
	}
	return out, err*/
	return "", nil
}
