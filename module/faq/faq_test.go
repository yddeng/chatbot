package faq

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestAnswer_String(t *testing.T) {
	s := "step1&&step21||step22||step23&&step3&&step4"
	a := Json2Answer(s)
	t.Log(a.Answer())
	t.Log(a.String())
}
