package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/xiaozefeng/gbdc/command"
)

var minutes = flag.Int("m", 15, "minutes")
var f = flag.String("f", "work", "work or rest")

func main() {
	flag.Parse()
	t := time.Duration(*minutes) * time.Minute
	log.Println("staring tomato clock successfully ....")
	total := time.After(t)
	message := getMessage(*f)
	log.Printf("%s\n", message)
	alertMessage := fmt.Sprintf("定时 %d 分钟结束了", *minutes)
	for {
		select {
		case <-total:
			command.Run("noti " + alertMessage)
			return
		case <-time.Tick(1 * time.Minute):
			log.Printf("一分钟过去了哦!\n")
		}
	}
	log.Printf("定时器结束了!\n")
}

func getMessage(f string) string {
	switch f {
	case "rest":
		return "休息一下吧!"
	case "work":
		return "开始工作了!"
	default:
		panic("unsupported flag")
	}
}
