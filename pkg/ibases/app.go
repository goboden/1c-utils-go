package ibases

import (
	"os"
	"strings"
)

func Run() {
	fileName := "./../../examples/ibases.v8i"
	fileData, err := readFile(fileName)
	if err != nil {
		println(err.Error(), err)
		return
	}
	readIBases(fileData)
}

func splitStrings(d []byte) []string {
	s := strings.Split(strings.Trim(string(d), "\r\n"), "\n")
	for i, l := range s {
		s[i] = strings.Trim(l, "\r")
	}
	return s
}

func readFile(name string) (str []string, err error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return
	}
	str = splitStrings(data)
	return str, nil
}
