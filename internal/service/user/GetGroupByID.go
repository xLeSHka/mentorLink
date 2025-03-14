package userService

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/xLeSHka/mentorLink/internal/app/httpError"
	"github.com/xLeSHka/mentorLink/internal/models"
	"gorm.io/gorm"
	"net/http"
)

func (s *UsersService) GetGroupByID(ctx context.Context, ID uuid.UUID) (*models.Group, error) {

	group, err := s.usersRepository.GetGroupByID(ctx, ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httpError.New(http.StatusNotFound, err.Error())
		}
		return nil, httpError.New(http.StatusInternalServerError, err.Error())
	}
	return group, nil
}
