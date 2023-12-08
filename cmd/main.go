package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	abisDir       = "lib/gmx-interface/src/abis"
	contractsDir  = "out/gmx-contracts"
	syntheticsDir = "out/gmx-synthetics"
)

var (
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
)

func main() {
	fs, err := os.Open(abisDir)
	if err != nil {
		log.Fatal(err)
	}

	files, err := fs.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	// Map of file name to source name
	fileMap := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		jsonFile, err := os.Open(abisDir + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		var sourceName string
		if result["sourceName"] == nil {
			sourceName = ""
		} else {
			sourceName = result["sourceName"].(string)
		}

		fileMap[file.Name()] = sourceName
	}

	for k, v := range fileMap {

		// only use the basename and file extension
		baseName := strings.TrimSuffix(filepath.Base(v), filepath.Ext(v))
		if overrides[baseName] != "" {
			baseName = overrides[baseName]
		}

		// Find solidity file and run abigen
		if baseName != "." {
			file := getCorrectFile(fmt.Sprintf("%s.sol/%s%s", baseName, baseName, ".abi.json"), []string{contractsDir, syntheticsDir})
			if file == nil {
				log.Println("No file found for " + fmt.Sprintf("%s.sol/%s%s", baseName, baseName, ".abi.json"))
				continue
			}
			defer file.Close()

			// Run abigen
			sanitizedName := nonAlphanumericRegex.ReplaceAllString(strings.TrimSuffix(k, filepath.Ext(k)), "")

			cmd := fmt.Sprintf("abigen --abi %s --pkg %s --out %s/%s.go", file.Name(), sanitizedName, sanitizedName, sanitizedName)
			err := os.MkdirAll(sanitizedName, os.ModePerm)
			if err != nil {
				log.Println(err)
				continue
			}
			err = exec.Command("sh", "-c", cmd).Run()
			if err != nil {
				log.Println(err)
			}

		} else {
			// If no solidity file found, run abigen with json file
			sanitizedName := nonAlphanumericRegex.ReplaceAllString(strings.TrimSuffix(k, filepath.Ext(k)), "")

			cmd := fmt.Sprintf("abigen --abi %s --pkg %s --out %s/%s.go", fmt.Sprintf("%s/%s", abisDir, k), sanitizedName, sanitizedName, sanitizedName)
			err := os.MkdirAll(sanitizedName, os.ModePerm)
			if err != nil {
				log.Println(err)
				continue
			}
			err = exec.Command("sh", "-c", cmd).Run()
			if err != nil {
				log.Println(err, " failed to run abigen for ", k)
				continue
			}
		}
	}

}
