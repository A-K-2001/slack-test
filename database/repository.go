package database

import "github.com/A-K-2001/slack-test/database/models"

type Repository struct {
	*models.Queries
	DB models.DBTX
}

func NewRepository(db models.DBTX) *Repository {
	return &Repository{
		Queries: models.New(db),
		DB:      db,
	}
}
