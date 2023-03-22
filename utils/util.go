package utils

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

func PrintLog(err interface{}) {
	if err != nil {
		log.Print(err)
	}
}

func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan:
		//use of IsNil method
		return reflect.ValueOf(val).IsNil()
	}
	if reflect.ValueOf(val).Kind() != reflect.Ptr && reflect.ValueOf(val).Len() == 0 {
		return true
	}
	return false
}

func AutoMap(from interface{}, to interface{}) error {
	jsonFrom, _ := json.Marshal(from)
	err := json.Unmarshal([]byte(string(jsonFrom)), to)
	return err
}

func StringToInt(from string) int {
	intVar, _ := strconv.Atoi(from)
	return intVar
}

func IntToString(from int) string {
	stringVar := strconv.Itoa(from)
	return stringVar
}
