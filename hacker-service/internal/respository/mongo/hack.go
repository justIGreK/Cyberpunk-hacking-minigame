package mongorepo

import (
	"context"
	"hacker-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatrixRepo struct {
	ReportsCollection *mongo.Collection
	AnswersCollection *mongo.Collection
}

func NewMatrixRepo(client *mongo.Client) *MatrixRepo {
	return &MatrixRepo{
		ReportsCollection: client.Database(dbname).Collection(reportsCollection),
		AnswersCollection: client.Database(dbname).Collection(answersCollection),
	}
}

func (r *MatrixRepo) AddHackReport(ctx context.Context, hackReport models.HackReport) error {
	_, err := r.ReportsCollection.InsertOne(ctx, hackReport)
	if err != nil {
		return err
	}
	return nil
}

func (r *MatrixRepo) AddReadyAnswers(ctx context.Context, readyAnswers models.ReadyAnswers) error{
	_, err := r.AnswersCollection.InsertOne(ctx, readyAnswers)
	if err != nil {
		return err
	}
	return nil
}

func (r *MatrixRepo) GetReadyAnswers(ctx context.Context, id int) (models.ReadyAnswers, error) {
	filter := bson.M{"matrix_id": id}
	var answers models.ReadyAnswers
	err := r.AnswersCollection.FindOne(ctx, filter).Decode(&answers)
	if err != nil {
		return answers, err
	}
	return answers, nil
}

func (r *MatrixRepo) GetMatrix(ctx context.Context, id int) (models.HackMatrix, error) {
	filter := bson.M{"matrix_id": id}
	var matrix models.HackMatrix
	err := r.ReportsCollection.FindOne(ctx, filter).Decode(&matrix)
	if err != nil {
		return matrix, err
	}
	return matrix, nil
}
