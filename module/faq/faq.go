package faq

import (
	"encoding/json"
	"github.com/yddeng/chatbot/module"
	"math/rand"
	"regexp"
	"strings"
)

/*
 多个问题引用一个回答，一个回答有多个回复
*/

type Question struct {
	q []string
	a *Answer
}

func Json2Question(s string) *Question {
	qs := strings.Split(s, "||")
	return &Question{q: qs}
}

// 余弦值、正则表达式
func (q *Question) Match(s string) bool {
	match, _ := regexp.MatchString(q.q[0], s)
	if match {
		return true
	}
	return false
}

/*
	多条回复中间用 '&&'，随机回复用 '||' . 优先级前者大于后者
*/
type Answer struct {
	ans  []string // 多个随机一个回复
	next *Answer  // 下一条回复
}

func (a *Answer) Answer() (ans []string) {
	for e := a; e != nil; e = e.next {
		i := rand.Int() % len(e.ans)
		ans = append(ans, e.ans[i])
	}
	return
}

func (a *Answer) String() string {
	s := strings.Join(a.ans, "||")
	for e := a.next; e != nil; e = e.next {
		s += "&&" + strings.Join(e.ans, "||")
	}
	return s
}

func Json2Answer(txt string) (a *Answer) {
	var prev *Answer
	steps := strings.Split(txt, "&&")
	for i, s := range steps {
		ans := strings.Split(s, "||")
		e := &Answer{ans: ans}
		if i == 0 {
			a = e
		} else {
			prev.next = e
		}
		prev = e
	}
	return
}

type FAQ struct {
	qa []*Question
}

func (this *FAQ) Init(data []byte) (err error) {
	var qa map[string]string
	if err = json.Unmarshal(data, &qa); err != nil {
		return err
	}
	this.qa = make([]*Question, 0, len(qa))
	for qStr, aStr := range qa {
		q := Json2Question(qStr)
		q.a = Json2Answer(aStr)
		this.qa = append(this.qa, q)
	}
	return nil
}

func (this *FAQ) Search(txt string) string {

	return ""
}

func init() {
	module.Register(module.FAQ_Module, func() module.Interface {
		return &FAQ{}
	})
}
