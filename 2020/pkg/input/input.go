package input

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadStrings(path, separator string) (values []string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	values = strings.Split(string(file), separator)
	return
}

func ReadInts(path, separator string) (values []int, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	strVals := strings.Split(string(file), separator)

	for _, val := range strVals {
		iVal, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		values = append(values, iVal)
	}

	return
}
