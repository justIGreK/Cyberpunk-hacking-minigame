package mongorepo

import (
	"context"
	"matrix-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatrixRepo struct {
	matrixCollection *mongo.Collection
}

func NewMatrixRepo(client *mongo.Client) *MatrixRepo {
	return &MatrixRepo{
		matrixCollection: client.Database(dbname).Collection(matrixCollection),
	}
}

func (r *MatrixRepo) AddMatrix(ctx context.Context, id int, hackMatrix models.HackMatrix) error {
	
	_, err := r.matrixCollection.InsertOne(ctx, hackMatrix)
	if err!=nil{
		return err
	}
	return nil
}

func (r *MatrixRepo) GetMatrix(ctx context.Context, id int) (models.HackMatrix, error) {
	filter := bson.M{"matrix_id": id}
	var matrix models.HackMatrix
	err := r.matrixCollection.FindOne(ctx, filter).Decode(&matrix)
	if err!=nil{
		return matrix, err 
	}
	return matrix, nil
}
