package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	usersR         *UsersRepository
	segmentsR      *SegmentsRepository
	usersSegmentsR *UsersSegmentsRepository
}

func New(cfg *Config) *Store {
	return &Store{
		config: cfg,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UsersRepository {
	if s.usersR != nil {
		return s.usersR
	}

	s.usersR = &UsersRepository{
		store: s,
	}

	return s.usersR
}

func (s *Store) Segment() *SegmentsRepository {
	if s.segmentsR != nil {
		return s.segmentsR
	}

	s.segmentsR = &SegmentsRepository{
		store: s,
	}

	return s.segmentsR
}

func (s *Store) UsersSegments() *UsersSegmentsRepository {
	if s.usersSegmentsR != nil {
		return s.usersSegmentsR
	}

	s.usersSegmentsR = &UsersSegmentsRepository{
		store: s,
	}

	return s.usersSegmentsR
}
