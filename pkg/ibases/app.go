package ibases

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Run() {
	filePath, err := findIBasesFile()
	if err == nil {
		fileData, err := readFile(filePath)
		if err != nil {
			println(err.Error(), err)
			return
		}
		readIBases(fileData)
	} else {
		fmt.Printf("IBases file is not found: %s", err.Error())
	}
}

func findIBasesFile() (string, error) {
	var filePath string

	userProfile, found := os.LookupEnv("USERPROFILE")
	if found {
		filePath = filePath + userProfile + "/AppData/Roaming/1C/1CEStart/ibases.v8i"
		return filePath, nil
	}
	return filePath, errors.New("Enviroment variable [USERPROFILE] is not found.")
}

func readFile(name string) (str []string, err error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return
	}
	str = splitStrings(data)
	return str, nil
}

func splitStrings(d []byte) []string {
	s := strings.Split(strings.Trim(string(d), "\r\n"), "\n")
	for i, l := range s {
		s[i] = strings.Trim(l, "\r")
	}
	return s
}
