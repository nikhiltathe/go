package wrapper

type Wrap interface {
	Say()
	Give(v int) Wrap2
}

type Wrap2 interface {
	SayA1()
}
