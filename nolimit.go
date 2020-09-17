package parallel

type NoLimit struct {
}

func (NoLimit) Lock() {
}
func (NoLimit) Unlock() {
}
