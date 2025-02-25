package sqlite

import (
	"database/sql"
	"log"
	"testing"
)

var storagePath = "../../../storage/storage.db"

var stogare *Storage = &Storage{
	db: func() *sql.DB {
		db, err := sql.Open("sqlite3", storagePath)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}(),
}

func TestStorage_DeleteURL(t *testing.T) {
	type args struct {
		alias string
	}
	tests := []struct {
		name string
		s    *Storage
		args args
		want string
	}{
		{
			name: "1",
			s:    stogare,
			args: args{
				alias: "vk",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.DeleteURL(tt.args.alias)
			if err != nil {
				t.Errorf("Storage.DeleteURL() = err: %s", err)
			}
		})
	}
}
