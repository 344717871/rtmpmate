package main

import (
	"fmt"
	"rtmpmate.com/net/http/HTTPListener"
	"rtmpmate.com/net/rtmp/RTMPListener"
)

const _NAME_ string = "rtmpmate"
const _VERSION_ string = "0.0.01"

func main() {
	fmt.Printf("SERVER: %s\n", _NAME_)
	fmt.Printf("VERSION: %s\n\n", _VERSION_)

	httpln, err := HTTPListener.New()
	if err != nil {
		fmt.Printf("Failed to create HTTPListener: %v.\n", err)
		return
	}

	go httpln.Listen("tcp4", 80)

	rtmpln, err := RTMPListener.New()
	if err != nil {
		fmt.Printf("Failed to create RTMPListener: %v.\n", err)
		return
	}

	rtmpln.Listen("tcp4", 1935)

	fmt.Printf("Server exiting...\n")
}
