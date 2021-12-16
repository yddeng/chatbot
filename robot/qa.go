package robot

import (
	"encoding/json"
	"github.com/yddeng/chatbot/utils"
	"io/ioutil"
)

type QA struct {
	Q          string         `json:"q"` // question
	A          string         `json:"a"` // answer
	WorldTimes map[string]int // 分词 及 词频
}

type QAList struct {
	QA       []*QA            `json:"qa"`
	WorldSeq map[string][]int // 分词对应的问题ID
}

func (this *QAList) GetAnswer(i int) string {
	return this.QA[i].A
}

func loadQAList(filename string) (*QAList, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(data))
	var list *QAList
	if err = json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	for _, v := range list.QA {
		v.WorldTimes = utils.CutMap(v.Q)
	}

	list.WorldSeq = map[string][]int{}
	for i, v := range list.QA {
		for w := range v.WorldTimes {
			if _, exist := list.WorldSeq[w]; exist {
				list.WorldSeq[w] = append(list.WorldSeq[w], i)
			} else {
				list.WorldSeq[w] = []int{i}
			}
		}
	}

	//fmt.Println(list)
	return list, nil
}
