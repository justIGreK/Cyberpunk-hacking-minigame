package mongorepo

import (
	"context"
	"hacker-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *MatrixRepo) AddHackReport(ctx context.Context, hackReport models.HackReport) error {
	_, err := r.ReportsCollection.InsertOne(ctx, hackReport)
	if err != nil {
		return err
	}
	return nil
}


func (r *MatrixRepo) GetReports(ctx context.Context) ([]models.HackReport, error) {
	var reports []models.HackReport
	filter := bson.M{}
	cursor, err := r.ReportsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &reports)
	if err!=nil{
		return nil, err
	}
	return reports, nil
}
