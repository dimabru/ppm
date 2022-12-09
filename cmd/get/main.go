package getCmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all path binaries",
	Long:  `Get all path binaries`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute()
	},
}

func New() *cobra.Command {
	return getCmd
}

func execute() error {
	pathValue := os.Getenv("PATH")
	parts := strings.Split(pathValue, ":")

	for _, part := range clearDuplicates(parts) {
		getAllFilesInPath(part)
	}

	return nil
}

func getAllFilesInPath(part string) error {
	var binaries []string

	err := filepath.Walk(part, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			binaries = append(binaries, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Base Directory: ", part, "| Binaries Count: ", len(binaries))

	return nil
}

func clearDuplicates(parts []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range parts {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
