package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, dbUrl string) (*Store, func(...string)) {
	t.Helper()

	cfg := NewConfig()
	cfg.DatabaseURL = dbUrl
	s := New(cfg)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
