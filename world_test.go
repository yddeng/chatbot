package chatbot

import (
	"fmt"
	"testing"
)

func TestCut(t *testing.T) {
	s := Cut("我来到北京清华大学")
	fmt.Println(s)

	s = Cut("how are you")
	fmt.Println(s)
}
