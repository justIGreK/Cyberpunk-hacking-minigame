package service

import (
	"context"
	"hacker-service/internal/models"
)

func (s *MatrixService) GetReports(ctx context.Context) ([]models.HackReport, error){
	reports, err := s.repo.GetReports(ctx)
	if err != nil{
		return nil, err
	}
	return reports, nil
}
