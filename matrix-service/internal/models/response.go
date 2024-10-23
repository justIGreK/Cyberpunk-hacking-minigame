package models

type HackMatrix struct {
	ID int	`bson:"matrix_id"`
	Matrix [][]int `bson:"matrix_map"`
	Sequence [][]int `bson:"sequences"`
}

type HackAttempts struct{
	MatrixID int `json:"matrix_id"`
	Attempts string `json:"attempts"`
}
