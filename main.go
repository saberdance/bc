package main

import (
	"fmt"
	"time"

	usbmonitor "github.com/saberdance/bc/src/UsbMonitor"
)

func main() {
	//Fake_StartListen 的参数为需要Channel返回的值
	usbmonitor.Fake_StartListen(0)
	for {
		x := <-usbmonitor.Channel
		fmt.Printf("Channel Ret: %d\n", x)
		time.Sleep(time.Second * 2)
	}
}
