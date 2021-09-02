package models

type Transactions struct {
	Total               int
	Successful          int
	UnSuccessful        int
	SuccessfulAverage   float32
	UnSuccessfulAverage float32
}
