package main

import (
	"fmt"
	"github.com/yddeng/chatbot/robot"
)

func main() {
	conf := robot.LoadConfig("config/config.json")

	r, err := robot.Chatbot(conf)
	if err != nil {
		panic(err)
	}

	var q string
	for {
		fmt.Print("=>")
		fmt.Scan(&q)

		ans := r.GetAnswer(q)
		fmt.Println(ans)
	}
}
