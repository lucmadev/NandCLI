package docker

import (
	"encoding/json"
)

func ConvertJsonToContainer(data []byte) Container {
	// data := []byte(`
	// {
	// 	"name":"nginx",
	// 	"status":"running",
	// 	"image":"nginx:latest"
	// }
	// `)

	var c Container

	err := json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	return c 
	
}