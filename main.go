package main

import (
	"fmt"
	"getcare-notification/internal/model"
	"getcare-notification/pkg/bson"
)

func main() {
	fmt.Println("Ok!")
	msg := &model.SendBroadcast{
		Event: "hi22",
		Data:  nil,
	}
	fmt.Println(bson.StructToMap(msg))
}
