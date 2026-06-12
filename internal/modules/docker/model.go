package docker

type Container struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Image  string `json:"image"`
}