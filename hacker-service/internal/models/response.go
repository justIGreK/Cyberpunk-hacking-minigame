package models

import "time"

type HackMatrix struct {
	ID        int     `bson:"matrix_id"`
	Matrix    [][]int `bson:"matrix_map"`
	Sequences [][]int `bson:"sequences"`
}

type HackReport struct {
	ID           int       `bson:"matrix_id"`
	Created      time.Time `bson:"created"`
	ResultOfHack int       `bson:"result_of_hack"`
}

type BruteforceResult struct {
	Results      []bool
	SuccessRoute []string
}

type ReadyAnswers struct {
	ID      int      `bson:"matrix_id"`
	Answers []string `bson:"answers"`
}
