package comment

import (
	"backend/internal/db"
	"backend/internal/server"
	"backend/models"
)

type IService interface {
	SaveComment(ctx *server.Context, comment models.Comment) error
	GetComments(ctx *server.Context) ([]*models.Comment, error)
}

type service struct {
	repo IRepository
}

func NewService(dbconn db.IMongo) IService {
	return &service{
		repo: NewRepository(dbconn),
	}
}

func (s *service) SaveComment(ctx *server.Context, comment models.Comment) error {
	return s.repo.SaveComment(ctx, comment)
}

func (s *service) GetComments(ctx *server.Context) ([]*models.Comment, error) {
	return s.repo.GetComments(ctx)
}
