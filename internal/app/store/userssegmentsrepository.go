package store

import "avitotest/internal/app/models"

type UsersSegmentsRepository struct {
	store *Store
}

func (r *UsersSegmentsRepository) Create(us *models.UsersSegments) (*models.UsersSegments, error) {
	if err := r.store.db.QueryRow(
		"insert into users_segments (user_id, segment_id) values ($1, $2) returnig id",
		us.User_id, us.Segment_id).Scan(&us.Id); err != nil {
		return nil, err
	}

	return us, nil
}
