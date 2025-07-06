package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func RunInit() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter node name: ")
	nodeName, _ := reader.ReadString('\n')
	nodeName = cleanString(nodeName)

	fmt.Print("Enter drive name: ")
	driveName, _ := reader.ReadString('\n')
	driveName = cleanString(driveName)

	baseDir := filepath.Join(os.Getenv("HOME"), "shadownet")
	keysDir := filepath.Join(baseDir, "keys")
	configFile := filepath.Join(baseDir, "config.json")
	driveDir := filepath.Join(baseDir, "drives", driveName)
	driveMeta := filepath.Join(driveDir, driveName+".json")

	//create dirs
	dirs := []string{baseDir, keysDir, driveDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", dir, err)
			return
		}
	}

	//create empty files
	files := []string{configFile, driveMeta}
	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			fmt.Println("Error creating file:", file, err)
			return
		}
		defer f.Close()
	}

	fmt.Println("Basic shadownet structure created")
	fmt.Println("Node:", nodeName)
	fmt.Println("Drive: ", driveName)
}

func cleanString(s string) string {
	return string([]byte(s)[:len(s)-1]) // remove newline
}
