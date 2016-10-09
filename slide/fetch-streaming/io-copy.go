resp, err := http.Get("https://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()

io.Copy(w, resp.Body) // HL
