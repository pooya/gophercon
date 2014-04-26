func Map(reader io.Reader, writer io.Writer) {
	body, err := ioutil.ReadAll(reader)
	jobutil.Check(err)
	strBody := string(body)
	words := strings.Fields(strBody)
	for _, word := range words {
		_, err := writer.Write([]byte(word + "\n"))
		jobutil.Check(err)
	}
}
