package repositoryUser

import (
	"context"

	"github.com/xLeSHka/mentorLink/internal/models"
)

func (r *UsersRepository) Login(ctx context.Context, person *models.User) (*models.User, error) {
	//err := r.DB.WithContext(ctx).FirstOrCreate(person).Error
	err := r.DB.WithContext(ctx).Create(person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}
