package services

import (
	"context"

	"github.com/go-snowflake/internal/entities"
	repository "github.com/go-snowflake/internal/repositories"
)

type ColumnListService struct {
	repo *repository.ColumnListRepository
}

type IColumnListService interface {
	GetAllColumnsInfo(ctx context.Context, tableName string) ([]*entities.ColumnInfo, error)
}

func NewColumnListService(repo *repository.ColumnListRepository) IColumnListService {
	return &ColumnListService{
		repo: repo,
	}
}

func (s *ColumnListService) GetAllColumnsInfo(ctx context.Context, tableName string) ([]*entities.ColumnInfo, error) {
	return s.repo.FetchAllColumns(ctx, tableName)
}
