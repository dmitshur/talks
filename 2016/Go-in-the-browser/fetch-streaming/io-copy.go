resp, err := http.Get("https://example.com/large.csv")
if err != nil {
	// handle error
}
defer resp.Body.Close()

io.Copy(os.Stdout, resp.Body) // HL
