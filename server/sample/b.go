package sample

type SS struct {
}

func New() Sample {
	s := &SS{}
	return s
}

func (s SS) A() {}

func (s SS) B() {}

func (s SS) sa() {}

func (s SS) mustEmbedUnimplementedSampleV1Server() {}

func mmm() {
	var s Sample
	s = &SS{}
	s.mustEmbedUnimplementedSampleV1Server()
}
