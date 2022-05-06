package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// enc, err := e2e.Encode(e2e.Key, 4, socket.SendBroadcast{
	// 	Event: "hoài nguyễn",
	// 	Date:  time.Now(),
	// 	Data: map[string]interface{}{
	// 		"abc": "abc",
	// 	},
	// })
	// if err != nil {
	// 	return
	// }
	// rs := &socket.SendBroadcast{}
	// err = e2e.Decode(e2e.Key, 4, enc, rs)
	// if err != nil {
	// 	return
	// }
	srv.Start()
}
