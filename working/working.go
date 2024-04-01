package working

import (
	"embed"
	"fmt"
)

//go:embed data/*.json
var jsonFiles embed.FS

func PrintFile(name string) error {
	data, err := jsonFiles.ReadFile("data/" + name)
	if err != nil {
		return fmt.Errorf("unsupported file name: %s", name)
	}
	fmt.Println(string(data))
	return nil
}
