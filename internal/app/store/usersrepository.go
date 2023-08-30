package store

import (
	"avitotest/internal/app/models"

	"github.com/lib/pq"
)

type UsersRepository struct {
	store *Store
}

func (r *UsersRepository) GetAllUsers() ([]*models.User, error) {
	rows, err := r.store.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UsersRepository) GetUser(id int) (*models.User, error) {
	user := models.User{
		Id: id,
	}
	err := r.store.db.QueryRow("select users.name from users where users.id=($1)", id).Scan(&user.Name)
	if err != nil {
		return nil, err
	}
	rows, err := r.store.db.Query("select segments.id, segments.name from users_segments as us join segments on us.segment_id=segments.id where us.user_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		segment := models.Segment{}
		err = rows.Scan(&segment.Id, &segment.Name)
		if err != nil {
			return nil, err
		}
		user.Segments = append(user.Segments, &segment)
	}
	return &user, nil
}

func (r *UsersRepository) Create(u *models.User) (*models.User, error) {
	q := "insert into users (name) values ('" + u.Name + "') returning id"
	if err := r.store.db.QueryRow(q).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UsersRepository) Delete(id int) error {
	if err := r.store.db.QueryRow(
		"delete from users where users.id=$1", id).Err(); err != nil {
		return err
	}
	return nil
}

func (r *UsersRepository) ChangeSegments(id int, add []string, remove []string) (*models.User, error) {
	rows, err := r.store.db.Query("select segments.id from segments where name=any($1)", pq.Array(remove))
	if err != nil {
		return nil, err
	}
	var ids []int
	for rows.Next() {
		var i int
		err := rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if err := r.store.db.QueryRow("delete from users_segments as us where us.user_id=$1 and us.segment_id = any($2)",
		id, pq.Array(ids)).Err(); err != nil {
		return nil, err
	}
	rows2, err := r.store.db.Query("select segments.id from segments where name=any($1)", pq.Array(add))
	if err != nil {
		return nil, err
	}
	for rows2.Next() {
		var i int
		err := rows2.Scan(&i)
		if err != nil {
			return nil, err
		}
		if err := r.store.db.QueryRow("insert into users_segments (user_id, segment_id) values ($1, $2)", id, i).Err(); err != nil {
			return nil, err
		}
	}
	user, err := r.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
