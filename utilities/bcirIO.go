package utilities

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func CreateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
		LogInfo("The directory \"%s\" is created", path)
	}

	return nil
}

func FindFileWithExt(root, ext string) []string {
	var files []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			files = append(files, s)
		}
		return nil
	})
	return files
}

func DeletePolicyRecord(policyRecords []string, policyName string) bool {
	for _, value := range policyRecords {
		if strings.Contains(value, policyName) {
			if e := os.Remove(value); e != nil {
				LogFatal("Can not remove policy [%s]: %s", policyName, e.Error())
				return false
			}
		}
	}
	return true
}

func WriteToFile(file string, value ...interface{}) bool {
	fh, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		LogError(fmt.Sprintf("Can not open file %s: %s", file, err))
		return false
	}
	for _, val := range value {
		_, err := fh.WriteString(reflect.ValueOf(val).String())
		if err != nil {
			LogError(fmt.Sprintf("Can not write to file %s: %s", file, err))
			return false
		} else {
			LogInfo("Write to file %s", file)
		}
	}
	return true
}

func ReadFileIntoMap(file string, keyMap map[string]string) error {
	fh, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fh.Close()

	reader := bufio.NewReader(fh)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		s := strings.Split(line, "=")
		if len(s) == 2 {
			if _, ok := keyMap[s[0]]; ok {
				keyMap[s[0]] = s[1]
			}
		}
	}

	return nil
}

func ReadFileByKey(file string, key string) (string, error) {
	fh, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fh.Close()

	reader := bufio.NewReader(fh)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		s := strings.Split(line, "=")
		if len(s) == 2 && s[0] == key {
			return s[1], nil
		}
	}

	return "", nil
}
