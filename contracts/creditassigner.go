package contracts

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
	Assignments(investment int32) ([]map[int32]int32, error)
}
