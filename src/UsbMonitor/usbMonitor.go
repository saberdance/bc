package usbmonitor

import (
	"sync"
	"time"
)

//证书监控模块
type Monitor struct {
	State int
}

var instance *Monitor
var once sync.Once
var Channel chan int

//取USB证书单例
func Instance() *Monitor {
	once.Do(func() {
		instance := &Monitor{}
		instance.State = -1
	})
	return instance
}

func Fake_ListenFunc(needret int) {
	for {
		Channel <- needret
		time.Sleep(time.Second * 2)
	}
}

func Fake_StartListen(needret int) {
	once.Do(func() {
		Channel = make(chan int)
		go Fake_ListenFunc(needret)
	})
}
