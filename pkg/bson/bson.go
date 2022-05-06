package bson

import (
	"github.com/fatih/structs"

	"go.mongodb.org/mongo-driver/bson"
)

func E(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: value}
}

func EsToMap(Es ...bson.E) map[string]interface{} {
	rs := map[string]interface{}{}
	for _, e := range Es {
		if e.Value != nil {
			rs[e.Key] = e.Value
		}
	}
	return rs
}

type InterfaceStruct interface{}

type NullStruct map[interface{}]string

var (
	nullStruct = NullStruct{
		nil: "nil",
		0:   "0",
		"":  "empty",
		// false: "false",
	}
)

func StructToMap(SI InterfaceStruct) map[string]interface{} {
	sm := structs.Map(SI)
	rs := map[string]interface{}{}
	for key, value := range sm {
		_, ok := nullStruct[value]
		if !ok {
			rs[key] = value
		}
	}
	return rs
}
