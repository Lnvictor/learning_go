package main

import (
	"fmt"
	"time"
)

func main(){
	channel := make(chan string)
	go escrever("hello world", channel)

	for m := range channel{
		fmt.Println(m)
	}
}

func escrever(text string, channel chan string){
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
	close(channel)
}
