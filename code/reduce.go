func Reduce(reader io.Reader, writer io.Writer) {
	sreader := jobutil.Sorted(reader)
	grouper := jobutil.Grouper(sreader)

	for grouper.Scan() {
		word, count := grouper.Text()
		_, err := writer.Write([]byte(fmt.Sprintf("%d %s\n", count, word)))
		jobutil.Check(err)
	}
	if err := grouper.Err(); err != nil {
		jobutil.Check(err)
	}
	sreader.Close()
}
