package main

import (
	"fmt"
	"github.com/yddeng/chatbot"
)

func main() {
	conf := chatbot.LoadConfig("config/config.json")

	robot, err := chatbot.Chatbot(conf)
	if err != nil {
		panic(err)
	}

	var q string
	for {
		fmt.Print("=>")
		fmt.Scan(&q)

		ans := robot.GetAnswer(q)
		fmt.Println(ans)
	}
}
