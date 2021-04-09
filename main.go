package main

func main() {
	server := NewServer(":3000")
	server.Get("/", HandleRoot)
	server.Handle("/api", "POST", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
