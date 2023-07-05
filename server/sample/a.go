package sample

type Sample interface {
	A()
	B()
	sa()
	mustEmbedUnimplementedSampleV1Server()
}
