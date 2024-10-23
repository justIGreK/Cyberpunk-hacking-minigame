package service

import (
	"context"
	"errors"
	"math/rand"
	"matrix-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MatrixRepository interface {
	AddMatrix(ctx context.Context, id int, hackMatrix models.HackMatrix) error
	GetMatrix(ctx context.Context, id int) (models.HackMatrix, error)
}

type MatrixService struct {
	repo MatrixRepository
}

func NewMatrixService(matrixRepo MatrixRepository) *MatrixService {
	return &MatrixService{repo: matrixRepo}
}

const (
	matrixSize        = 5
	sequencesCount    = 3
	symbolsCount      = 4
	maxSequenceLength = 7
	minSequenceLength = 4
)

func (s *MatrixService) GetMatrix(ctx context.Context, id int) (*models.HackMatrix, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}
	matrix, err := s.repo.GetMatrix(ctx, id)
	if err == mongo.ErrNoDocuments {
		matrix = models.HackMatrix{
			ID:       id,
			Matrix:   s.GenerateMatrix(),
			Sequence: s.GenerateSequences(),
		}
		err = s.repo.AddMatrix(ctx, id, matrix)
		if err != nil {
			return nil, err
		}
		return &matrix, nil
	} else if err != nil {
		return nil, err
	}

	return &matrix, nil

}

func (s *MatrixService) GenerateMatrix() [][]int {
	matrix := make([][]int, matrixSize)
	for i := range matrix {
		matrix[i] = make([]int, matrixSize)
	}

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			matrix[i][j] = rand.Intn(symbolsCount) + 1
		}
	}

	return matrix
}

func (s *MatrixService) GenerateSequences() [][]int {
	sequences := make([][]int, sequencesCount)
	for i := 0; i < sequencesCount; i++ {
		length := rand.Intn(maxSequenceLength-minSequenceLength+1) + minSequenceLength
		sequence := make([]int, length)

		for j := 0; j < length; j++ {
			sequence[j] = rand.Intn(symbolsCount) + 1
		}

		sequences[i] = sequence
	}
	return sequences
}
