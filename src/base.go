package Base

import (
	"fmt"
	"reflect"
	"strings"
)

type Interface interface {
	/*
		Return the interface kind
	*/
	Kind() string
	/*
		Check if an interface{} respond to a method name
	*/
	RespondTo(string) bool
	/*
		List the object methods
	*/
	Methods() []string
}

type data struct{}

func New() Interface {
	return &data{}
}

func Kind(instance interface{}) string {
	data := reflect.TypeOf(instance).String()
	kind := strings.Split(data, ".")[0]
	return fmt.Sprintf("%s.Interface", kind)
}

/*
Must be implemented by sub-interface of Object.Interface
*/
func (d data) Kind() string {
	return Kind(d)
}

func RespondTo(instance interface{}, methodName string) bool {
	method := reflect.ValueOf(instance).MethodByName(methodName)
	return method.IsValid()
}

/*
Must be implemented by sub-interface of Object.Interface
*/
func (d data) RespondTo(methodName string) bool {
	return RespondTo(d, methodName)
}

func Methods(instance interface{}) []string {
	t := reflect.TypeOf((instance))
	var methods []string
	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}
	return methods
}

func (d data) Methods() []string {
	return Methods(d)
}
