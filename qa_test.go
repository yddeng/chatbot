package chatbot

import "testing"

func TestNewMatching(t *testing.T) {
	list, err := LoadQA("config/qa.json")
	if err != nil {
		t.Fatal(err)
	}

	s := "你好"
	res := Matching(list, s)
	t.Log(res)
	for i := range res {
		t.Log(list.GetAns(i))
	}

	s = "你是谁"
	res = Matching(list, s)
	t.Log(res)
	for i := range res {
		t.Log(list.GetAns(i))
	}

}
