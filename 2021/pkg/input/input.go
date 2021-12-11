package input

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"github.com/sanity-io/litter"
)

// ReadStrings - return string array from given input file and separator
func ReadStrings(path, separator string) (values []string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	values = strings.Split(string(file), separator)
	return
}

// ReadInts - return int array from given input file and separator
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

func ReflectAll(lines []string, separator string, arg interface{}) error {

	fmt.Printf("Ptr: %p\n", arg)

	litter.Dump(arg)

	rVal := reflect.ValueOf(arg).Elem()
	// rVal := reflect.Indirect(reflect.ValueOf(arg))
	fmt.Printf("rval Ptr: %p\n", rVal.Pointer)
	// rElem := reflect.ValueOf(rVal.Interface()).Elem()

	kind := rVal.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return fmt.Errorf("not an array: %s", kind)
	}

	typ := reflect.TypeOf(rVal.Interface()).Elem()
	// t := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)
	// fmt.Println(typ.String())

	for _, line := range lines {
		// create new
		newResElem := reflect.New(typ).Elem()
		newResStruct := newResElem.Addr().Interface()

		// litter.Dump(newResElem.Kind())
		// litter.Dump(newResStruct)

		// reflect
		err := Reflect(line, separator, newResStruct)
		if err != nil {
			return err
		}

		// add to array
		rVal = reflect.Append(rVal, newResElem)
	}
	fmt.Printf("Ptr: %p\n", arg)
	litter.Dump(arg)

	// tInt := .Interface()
	// rValInt := rVal.Interface()
	// arg = &rValInt
	litter.Dump(rVal.Interface())
	litter.Dump(arg)

	fmt.Printf("Ptr: %p\n", arg)

	return nil
}

func Reflect(data string, separator string, result interface{}) (err error) {

	e := reflect.ValueOf(result).Elem()
	splits := strings.Split(data, separator)
	// fmt.Println(e.Kind())

	if e.NumField() != len(splits) {
		fmt.Println("data does not match struct", e.NumField(), len(splits))
		return fmt.Errorf("data does not match struct")
	}

	for idx, strVal := range splits {
		varType := e.Type().Field(idx).Type.Kind()

		switch varType {
		case reflect.Int:
			// Parse int
			intVal, err := strconv.Atoi(strVal)
			if err != nil {
				return err
			}
			e.Field(idx).SetInt(int64(intVal))
		case reflect.String:
			// do nothing
			e.Field(idx).SetString(strVal)
		default:
			return fmt.Errorf("unsupported data type: %s", varType)
		}
	}

	return nil
}
