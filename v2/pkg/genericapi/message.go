package genericapi

// APIMessage is a struct that is used to communicate between modules internally
type APIMessage struct {
	Author      string
	Content     string
	Destination string
}