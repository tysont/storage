package storage

import (

	"strings"
	"fmt"
	"reflect"
)

type MemoryStorer struct {
	
	TypeKeyObjects map[string]map[string]fmt.Stringer
}

func (this *MemoryStorer) Get(class string, key string) fmt.Stringer {

	class = normalizeClassName(class)

	if (this.TypeKeyObjects[class] == nil) {
		return nil
	}

	if (this.TypeKeyObjects[class][key] == nil) {
		return nil
	}
	
	//TODO: Should this return a deep copy instead of a reference to the same object?
	return this.TypeKeyObjects[class][key]
}

func (this *MemoryStorer) Put(object fmt.Stringer) {

	class := getShortTypeName(object)
	class = normalizeClassName(class)

	if (this.TypeKeyObjects == nil) {
		this.TypeKeyObjects = make(map[string]map[string]fmt.Stringer)
	}

	if (this.TypeKeyObjects[class] == nil) {
		this.TypeKeyObjects[class] = make(map[string]fmt.Stringer)
	}

	key := object.String()
	this.TypeKeyObjects[class][key] = object
}

func getShortTypeName(object interface{}) string {

	class := reflect.TypeOf(object).String()
	index := strings.LastIndex(class, ".")
	if index > 0 {

		class = class[index + 1:]
	}

	return class
}

func normalizeClassName(class string) string {

	return strings.ToLower(class)
}