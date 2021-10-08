package chatbot

import (
	"fmt"
	"testing"
)

func TestCut(t *testing.T) {
	s := Cut("我来到北京清华大学")
	fmt.Println(s)

	m := CutMap("我来到北京清华大学？")
	fmt.Println(m)
}
