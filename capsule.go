package capsuleio

import (
	"io/ioutil"
	"os"
	"strings"
)

var (
	storage map[string]string
	source  string
)

func init() {
	storage = make(map[string]string)
}

//Sets the source file for the key value store (aka the capsule)
func Open(capsule string) {
	File, _ := ioutil.ReadFile(capsule)
	source = string(File)
	load()
}

//Gets the string value from the current capsule
func Get(key string) string {
	if len(storage) < 1 {
		//load local file from directory
		basepath, _ := os.Getwd()

		fileinfo, _ := ioutil.ReadDir(basepath)
		for _, file := range fileinfo {
			if !file.IsDir() && strings.Contains(file.Name(), ".capsule") {
				File, _ := ioutil.ReadFile(basepath + "/" + file.Name())
				source = string(File)
				
				break
			}
		}
		
		load()
		return storage[key]
	}
	
	return storage[key]
}

func load() {
	lines := strings.Split(source, "\n")
	for i := range lines {
		group := strings.SplitAfterN(lines[i], "=", 2)
		if len(group) >= 2 {
			storage[strings.TrimSpace(strings.Replace(group[0], "=", "", 1))] = strings.TrimSpace(group[1])
		}
	}
}
