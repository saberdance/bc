package usbmonitor

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
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
		time.Sleep(time.Second * 2)
		Channel <- needret
	}
}

func Fake_StartListen(needret int) {
	once.Do(func() {
		Channel = make(chan int)
		go Fake_ListenFunc(needret)
	})
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Fake_GetPrivateKey(needret string) string {
	if needret == "" {
		return GetRandomString(32)
	}
	return needret
}
