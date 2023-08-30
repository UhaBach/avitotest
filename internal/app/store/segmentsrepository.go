package store

import (
	"avitotest/internal/app/models"
	"strconv"
)

type SegmentsRepository struct {
	store *Store
}

func (r *SegmentsRepository) GetAllSegments() ([]*models.Segment, error) {
	rows, err := r.store.db.Query("select * from segments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var segments []*models.Segment
	for rows.Next() {
		segment := models.Segment{}
		err := rows.Scan(&segment.Id, &segment.Name)
		if err != nil {
			continue
		}
		segments = append(segments, &segment)
	}
	return segments, nil
}

func (r *SegmentsRepository) GetSegment(name string) (*models.Segment, error) {
	segment := models.Segment{
		Name: name,
	}
	q := "select segments.id from segments where segments.name='" + name + "'"
	err := r.store.db.QueryRow(q).Scan(&segment.Id)
	if err != nil {
		return nil, err
	}
	rows, err := r.store.db.Query("select users.id, users.name from users_segments as us join users on us.user_id=users.id where us.segment_id=$1", segment.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		segment.Users = append(segment.Users, &user)
	}
	return &segment, nil
}

func (r *SegmentsRepository) Create(s *models.Segment) (*models.Segment, error) {
	q := "insert into segments (name) values ('" + s.Name + "') returning id"
	if err := r.store.db.QueryRow(q).Scan(&s.Id); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *SegmentsRepository) Delete(name string) error {
	q := "delete from segments where segments.name='" + name + "'"
	if err := r.store.db.QueryRow(q).Err(); err != nil {
		return err
	}
	return nil
}

func (r *SegmentsRepository) ChangeUsers(name string, add []int, remove []int) (*models.Segment, error) {
	var id int
	q := "select segments.id from segments where segments.name='" + name + "'"
	err := r.store.db.QueryRow(q).Scan(&id)
	if err != nil {
		return nil, err
	}
	arr := ""
	for index, value := range remove {
		if index != 0 {
			arr += ","
		}
		arr += strconv.Itoa(value)
	}
	q2 := "delete from users_segments as us where us.user_id=any('{" + arr + "}') and us.segment_id=" + strconv.Itoa(id)
	if err := r.store.db.QueryRow(q2).Err(); err != nil {
		return nil, err
	}

	for _, value := range add {
		q3 := "insert into users_segments (user_id, segment_id) values (" + strconv.Itoa(value) + "," + strconv.Itoa(id) + ")"
		if err := r.store.db.QueryRow(q3).Err(); err != nil {
			return nil, err
		}
	}
	segment, err := r.GetSegment(name)
	if err != nil {
		return nil, err
	}
	return segment, nil
}
