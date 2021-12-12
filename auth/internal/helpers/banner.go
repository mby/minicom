package helpers

import (
	"fmt"
)

func PrintBanner(projectName, port string) {
	fmt.Println("project: " + projectName)
	fmt.Printf("running on port %s\n", port)
}
