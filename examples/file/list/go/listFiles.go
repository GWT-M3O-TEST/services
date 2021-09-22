package example

import (
	"fmt"
	"github.com/micro/services/clients/go/file"
	"os"
)

// List files by their project and optionally a path.
func ListFiles() {
	fileService := file.NewFileService(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := fileService.List(&file.ListRequest{
		Project: "examples",
	})
	fmt.Println(rsp, err)
}