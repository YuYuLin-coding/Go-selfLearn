package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cleanup()
	fmt.Println("程式執行(按下ctrl+c)")

MainLoop:
	for {
		s := <-sigs
		switch s {
		case syscall.SIGINT:
			fmt.Println("程式中斷")
			break MainLoop
		case syscall.SIGTERM:
			fmt.Println("程式終止")
			break MainLoop
		}
	}
	fmt.Println("程式結果")
}

func cleanup() {
	fmt.Println("進行清理作業")
	for i := 0; i <= 10; i++ {
		fmt.Printf("刪除檔案 %v...\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
