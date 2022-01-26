package ibases

import (
	"strings"
)

type ibaseList struct {
	ibases  []ibase
	folders []folder
}

func (ibl *ibaseList) Count() (ibases int, folders int) {
	ibases = len(ibl.ibases)
	folders = len(ibl.folders)
	return
}

type ibaseItem struct {
	name, id string
	external bool
	path     string
}

type ibase struct {
	ibaseItem
	connect string
}

type folder struct {
	ibaseItem
}

func (ibd *ibaseList) appendData(name string, data map[string]string) {
	if _, ok := data["Connect"]; ok {
		ibase := new(ibase)
		ibase.fill(name, data)
		ibase.connect = data["Connect"]
		ibd.ibases = append(ibd.ibases, *ibase)
	} else {
		folder := new(folder)
		folder.fill(name, data)
		ibd.folders = append(ibd.folders, *folder)
	}
}

func (itm *ibaseItem) fill(name string, data map[string]string) {
	itm.name = name
	itm.id = data["ID"]
	itm.external = data["External"] != "0"
	itm.path = data["Folder"]
}

func readIBases(data []string) *ibaseList {
	ibases := readData(data)
	// printIBases(ibases)
	printFolders(ibases, "/", 0)
	return ibases
}

func parseIBName(s string) (name string, ok bool) {
	if strings.HasPrefix(s, "[") {
		name = strings.Trim(s, "[]")
		ok = true
		return
	}
	return
}

func readData(lines []string) *ibaseList {
	var cibname string
	ibases := new(ibaseList)
	params := make([]string, 0, 15)
	parsed := make(map[string]string)
	for _, line := range lines {
		if ibname, ok := parseIBName(line); ok {
			if cibname != ibname {
				if cibname != "" {
					parsed = parseParams(params)
					ibases.appendData(cibname, parsed)
					params = params[:0]
				}
				cibname = ibname
			}
		} else {
			params = append(params, line)
		}
	}
	ibases.appendData(cibname, parsed)
	return ibases
}

func parseParams(params []string) map[string]string {
	parsed := make(map[string]string)
	for _, line := range params {
		key, value := parseParam(line)
		parsed[key] = value
	}
	return parsed
}

func parseParam(param string) (key, value string) {
	for i, ch := range param {
		if string(ch) == "=" {
			key = param[:i]
			value = param[i+1:]
			return
		}
	}
	return
}
