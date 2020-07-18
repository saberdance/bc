package main

import (
	// 	"time"
	"fmt"

	usbmonitor "github.com/saberdance/bc/src/UsbMonitor"
)

func main() {
	//Fake_StartListen 的参数为需要Channel返回的值
	// usbmonitor.Fake_StartListen(0)
	// for {
	// 	x := <-usbmonitor.Channel
	// 	fmt.Printf("Channel Ret: %d\n", x)
	// 	time.Sleep(time.Second * 2)
	// }
	x := usbmonitor.Fake_GetPrivateKey("")
	y := usbmonitor.Fake_GetPrivateKey("1122334455")
	fmt.Printf("fake: %s\n", x)
	fmt.Printf("fake: %s\n", y)
}
