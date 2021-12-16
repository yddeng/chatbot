package module

type Interface interface {
	Init(data []byte) error
	Search(txt string) string
}

type Module int

const (
	FAQ_Module Module = 1
)

type creator func() Interface

var modules = map[Module]creator{}

func Register(m Module, f creator) {
	modules[m] = f
}
