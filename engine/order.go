package engine

type Order struct {
	Price    float32
	Amount   uint64
	IsSeller bool
	ID       uint64
}
