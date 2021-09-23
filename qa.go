package chatbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type QA struct {
	Q string `json:"q"`
	A string `json:"a"`
	c map[string]int
}

type qaList struct {
	QA     []*QA `json:"qa"`
	manber map[string][]int
}

func (this *qaList) GetAns(i int) string {
	return this.QA[i].A
}

func LoadQA(filename string) (*qaList, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	var list *qaList
	if err = json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	for _, v := range list.QA {
		v.c = CutMap(v.Q)
	}

	list.manber = map[string][]int{}
	for i, v := range list.QA {
		for w := range v.c {
			if _, exist := list.manber[w]; exist {
				list.manber[w] = append(list.manber[w], i)
			} else {
				list.manber[w] = []int{i}
			}
		}
	}

	fmt.Println(list)
	return list, nil
}

func Matching(qa *qaList, s string) map[int]float64 {
	mm := CutMap(s)

	res := map[int]float64{}
	for w := range mm {
		if _, ok := qa.manber[w]; ok {
			for _, id := range qa.manber[w] {
				co := Cosine_sim(mm, qa.QA[id].c)
				res[id] = co
			}
		}
	}

	return res
}

func Cosine_sim(a, b map[string]int) float64 {
	allworlds := make([]string, 0, len(a)+len(b))
	for v := range a {
		allworlds = append(allworlds, v)
	}
	for v := range b {
		allworlds = append(allworlds, v)
	}

	ca := make([]float64, len(allworlds))
	cb := make([]float64, len(allworlds))
	for i, v := range allworlds {
		ca[i] = float64(a[v])
		cb[i] = float64(b[v])
	}

	if ret, err := Cosine(ca, cb); err == nil {
		return ret
	}
	return 0
}
