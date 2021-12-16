package robot

import (
	"fmt"
	"github.com/yddeng/chatbot/utils"
	"math/rand"
	"time"
)

/*
	1. 分词, 词向量
	2.
*/

type RobotInfo struct {
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Birthday string `json:"birthday"`
}

type RobotConfig struct {
	RobotInfo     *RobotInfo `json:"robot_info"`
	Cosine        float64    `json:"cosine"`
	DefaultAnswer []string   `json:"default_answer"`
}

type ChatBot struct {
	qaList *QAList
	config *RobotConfig
}

func Chatbot(conf *Config) (*ChatBot, error) {
	chatbot := new(ChatBot)
	rand.Seed(time.Now().UnixNano())

	if err := utils.DecodeJsonFromFile(&chatbot.config, conf.RobotPath); err != nil {
		return nil, err
	}

	qaList, err := loadQAList(conf.QaPath)
	if err != nil {
		return nil, err
	}
	chatbot.qaList = qaList

	return chatbot, nil
}

func (r *ChatBot) GetAnswer(q string) string {
	wm := utils.CutMap(q)
	max := float64(0)
	ans := ""

	for w := range wm {
		if _, ok := r.qaList.WorldSeq[w]; ok {
			for _, id := range r.qaList.WorldSeq[w] {
				co := utils.Cosine_sim(wm, r.qaList.QA[id].WorldTimes)
				if co > r.config.Cosine {
					if co > max {
						max = co
						ans = r.qaList.GetAnswer(id)
						fmt.Println(co)
					}
				}
			}
		}
	}

	if ans != "" {
		return ans
	}

	idx := rand.Int() % len(r.config.DefaultAnswer)
	return r.config.DefaultAnswer[idx]
}
