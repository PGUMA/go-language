package main

import (
	"fmt"

	"github.com/PGUMA/go-language/starting-go-language/chapter2-2/animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeeds())
	fmt.Println(animals.RabbitFeed())
}