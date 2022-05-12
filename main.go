package main

import (
	"fmt"
	"time"
)

func main() {
	colors := []string{"red", "blue", "orange", "pink", "purple"}

	firstChannel := make(chan string)

	go func(arr []string) {
		for _, color := range arr {
			firstChannel <- color // Kanala veriyi gönderiyoruz.
			time.Sleep(time.Second)
		}
	}(colors)

	go func() {
		for i := 0; i < 5; i++ {
			data := <-firstChannel // Kanaldan alınan veri data değişkenine atanıyor.
			fmt.Println(data)
		}
	}()
	<-time.After(time.Second * 7)
}
