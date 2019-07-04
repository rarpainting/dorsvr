package main

import (
	"flag"

	"github.com/djwackey/dorsvr/rtspclient"
)

var (
	rtspURL = flag.String("rtsp-url", "", "")
)

func main() {
	flag.Parse()

	client := rtspclient.New()

	// to connect rtsp server
	if !client.DialRTSP(*rtspURL) {
		return
	}

	// send the options/describe request
	client.SendRequest()

	select {}
}
