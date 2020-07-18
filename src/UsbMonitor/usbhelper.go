// +build windows

package usbmonitor

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func CheckUsbCert() (int, error) {
	dir, _ := os.Getwd()
	dllPath := dir + "\\uc\\uc.dll"
	fmt.Printf("dllPath: %s\n", dllPath)
	handle, err := syscall.LoadLibrary(dllPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return -1, err
	}
	defer syscall.FreeLibrary(handle)
	checker, err := syscall.GetProcAddress(handle, "CheckUSBReady")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return -2, err
	}
	ret, _, _ := syscall.Syscall(checker, 0, 0, 0, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return -3, err
	}
	fmt.Println("检查USB结果:", ret)
	return int(ret), err
}

func ListenFunc() {
	for {
		time.Sleep(time.Second * 2)
		ret, _ := CheckUsbCert()
		Channel <- ret
	}
}

func StartListen() {
	once.Do(func() {
		Channel = make(chan int)
		go ListenFunc()
	})
}
