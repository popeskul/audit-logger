package repository

import (
	"context"
	"github.com/popeskul/audit-logger/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const LogCollectionName = "logs"
const DBName = "audit"

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(ctx context.Context, item domain.LogItem) error {
	_, err := r.db.Database(DBName).Collection(LogCollectionName).InsertOne(ctx, item)

	return err
}
