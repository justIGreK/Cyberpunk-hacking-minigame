package service

import (
	"context"
	"errors"
	"fmt"
	"matrix-service/internal/models"
	"strconv"
	"strings"
)

func (s *MatrixService) HackMatrix(ctx context.Context, attempt models.HackAttempt) (bool, error) {
	matrix, err := s.repo.GetMatrix(ctx, attempt.MatrixID)
	if err != nil {
		return false, errors.New("Matrix is not found")
	}

	hackCords, err := s.convertInputToCoordinates(attempt.Path)
	if err != nil {
		return false, err
	}
	finalSeq := s.createSequence(matrix.Matrix, hackCords)
	if s.verifySequence(finalSeq, matrix.Sequence) {
		return true, nil
	}
	return false, nil
}

func (s *MatrixService) verifySequence(userSequence []int, targetSequences [][]int) bool {
	for _, target := range targetSequences {
		if len(userSequence) != len(target) {
			continue
		}
		match := true
		for i := range userSequence {
			if userSequence[i] != target[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func (s *MatrixService) createSequence(matrix [][]int, cords [][2]int) []int {
	var sequence []int
	for i := 0; i < len(cords); i++ {
		sequence = append(sequence, matrix[cords[i][0]][cords[i][1]])
	}
	return sequence
}

func (s *MatrixService) convertInputToCoordinates(input string) ([][2]int, error) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	isOkay := s.isAnyEqualSteps(parts)
	if !isOkay {
		return nil, errors.New("Duplicate step detected, violation of the 3nd rule")
	}
	var coordinates [][2]int
	count, lastRow, lastCol := 0, 0, 0
	for _, part := range parts {
		if len(part) != 2 {
			return nil, errors.New("invalid coordinate format: each coordinate must be two digits")
		}
		row, err := strconv.Atoi(string(part[0]))
		if err != nil || row < 0 || row >= matrixSize {
			return nil, errors.New("invalid row coordinate: must be a digit within the matrix bounds")
		}
		col, err := strconv.Atoi(string(part[1]))
		if err != nil || col < 0 || col >= matrixSize {
			return nil, errors.New("invalid column coordinate: must be a digit within the matrix bounds")
		}
		if count%2 == 0 {
			if lastRow != row {
				return nil, fmt.Errorf("invalid row on %v step, violation of the 2nd rule'", count+1)
			}
		} else {
			if lastCol != col {
				return nil, fmt.Errorf("invalid col on %v step, violation of the 2nd rule", count+1)
			}
		}
		lastRow, lastCol = row, col
		count++
		coordinates = append(coordinates, [2]int{row, col})
	}
	return coordinates, nil
}

func (s *MatrixService) isAnyEqualSteps(steps []string) bool {
	for i := 0; i < len(steps); i++ {
		for j := i + 1; j < len(steps); j++ {
			if steps[i] == steps[j] {
				return false
			}
		}
	}
	return true
}
