package service

import (
	"context"
	"fmt"
	"hacker-service/internal/models"
	"hacker-service/pkg/client"
	"log"
	"strings"
	"time"
)

type MatrixRepository interface {
	GetReadyAnswers(ctx context.Context, id int) (models.ReadyAnswers, error)
	GetMatrix(ctx context.Context, id int) (models.HackMatrix, error)
	GetReports(ctx context.Context) ([]models.HackReport, error)
	AddReadyAnswers(ctx context.Context, readyAnswers models.ReadyAnswers) error
	AddHackReport(ctx context.Context, hackReport models.HackReport) error
}

type MatrixService struct {
	repo MatrixRepository
}

func NewMatrixService(matrixRepo MatrixRepository) *MatrixService {
	return &MatrixService{repo: matrixRepo}
}

func (s *MatrixService) HackMatrix(ctx context.Context, id int) ([]string, error) {
	readyAnswers, err := s.repo.GetReadyAnswers(ctx, id)
	if err == nil {
		return readyAnswers.Answers, nil
	}
	hackMatrix, err := client.GetSequence(id)
	if err != nil {
		return nil, err
	}
	bfResults, answers := s.bruteForceHack(hackMatrix)
	s.saveResults(ctx, bfResults, answers, id)

	return answers, nil
}

func (s *MatrixService) saveResults(ctx context.Context, bfResults models.BruteforceResult, answers []string, id int) {
	resultOfHack := 0
	for _, result := range bfResults.Results {
		if result {
			resultOfHack++
		}
	}

	err := s.repo.AddHackReport(ctx, models.HackReport{
		ID:           id,
		Created:      time.Now().UTC(),
		ResultOfHack: resultOfHack,
	})
	if err != nil {
		log.Println(err)
	}
	err = s.repo.AddReadyAnswers(ctx, models.ReadyAnswers{
		ID:      id,
		Answers: answers,
	})
	if err != nil {
		log.Println(err)
	}

}

func (s *MatrixService) bruteForceHack(hackMaterial *models.HackMatrix) (models.BruteforceResult, []string) {
	bfResults := models.BruteforceResult{
		Results:      make([]bool, len(hackMaterial.Sequences)),
		SuccessRoute: make([]string, len(hackMaterial.Sequences)),
	}
	readyAnswers := make([]string, len(hackMaterial.Sequences))

	for seqIndex, sequence := range hackMaterial.Sequences {
		result, route := s.isSequencePossible(hackMaterial.Matrix, sequence)
		bfResults.Results[seqIndex] = result
		bfResults.SuccessRoute[seqIndex] = route
		if result {
			readyAnswers[seqIndex] = fmt.Sprintf("Sequence %v is posible: %s", sequence, route)
		} else {
			readyAnswers[seqIndex] = fmt.Sprintf("Sequence %v is imposible", sequence)
		}

	}

	return bfResults, readyAnswers
}

func (s *MatrixService) isSequencePossible(matrix [][]int, sequence []int) (bool, string) {
	for startCol := 0; startCol < len(matrix[0]); startCol++ {
		path := []string{}
		if s.dfs(0, startCol, matrix, sequence, 0, true, &path) {
			return true, strings.Join(path, " ")
		}
	}

	return false, ""
}

func (s *MatrixService) dfs(row, col int, matrix [][]int, sequence []int, index int, changeRow bool, path *[]string) bool {
	if index == len(sequence) {
		return true
	}

	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] != sequence[index] {
		return false
	}

	*path = append(*path, fmt.Sprintf("%d%d", row, col))

	if changeRow {
		for nextRow := 0; nextRow < len(matrix); nextRow++ {
			if nextRow != row {
				newPath := make([]string, len(*path))
				copy(newPath, *path)
				if s.dfs(nextRow, col, matrix, sequence, index+1, !changeRow, &newPath) {
					*path = newPath
					return true
				}
			}
		}
	} else {
		for nextCol := 0; nextCol < len(matrix[0]); nextCol++ {
			if nextCol != col {
				newPath := make([]string, len(*path))
				copy(newPath, *path)
				if s.dfs(row, nextCol, matrix, sequence, index+1, !changeRow, &newPath) {
					*path = newPath
					return true
				}
			}
		}
	}

	return false
}
