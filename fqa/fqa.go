package fqa

type FQA interface {
	Init()
}

type Module int

const (
	Module_Robot Module = iota
)

func Register(m Module, f FQA) {

}
