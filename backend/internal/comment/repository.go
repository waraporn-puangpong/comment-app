package comment

import (
	"backend/internal/db"
	"backend/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	SaveComment(ctx context.Context, comment models.Comment) error
	GetComments() ([]*models.Comment, error)
}

const (
	commentCollectionName = "comment"
	queryTimeout          = 5 * time.Second
)

type repository struct {
	dbconn db.IMongo
}

func NewRepository(dbconn db.IMongo) IRepository {
	return &repository{dbconn: dbconn}
}

func (r *repository) SaveComment(ctx context.Context, comment models.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	if _, err := r.dbconn.Collection(commentCollectionName).InsertOne(ctx, comment); err != nil {
		return fmt.Errorf("save comment: %w", err)
	}

	return nil
}

func (r *repository) GetComments() ([]*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.dbconn.Collection(commentCollectionName).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("find comments: %w", err)
	}
	defer cursor.Close(ctx)

	comments := make([]*models.Comment, 0)
	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, fmt.Errorf("decode comment: %w", err)
		}
		comments = append(comments, &comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("iterate comments cursor: %w", err)
	}

	return comments, nil
}
