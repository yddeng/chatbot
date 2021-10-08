package joke

import "fmt"

type Jokes struct {
}

func (this *Jokes) Init() {
	filename := "./fqa/joke/jokes.json"
	fmt.Println(filename)
}
