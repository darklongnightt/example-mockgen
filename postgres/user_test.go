package postgres_test

import (
	"errors"
	"example-mockgen/models"
	"example-mockgen/postgres"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	tests := []struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		args       *models.User
		want       *models.User
		err        error
	}{
		{
			name: "success create user",
			args: &models.User{Name: "success", Age: 22},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("INSERT INTO users (id, name, age, created_at) VALUES ($1, $2, $3, $4)")).
					WithArgs(sqlmock.AnyArg(), "success", 22, sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &models.User{Name: "success", Age: 22},
			err:  nil,
		},
		{
			name: "db error",
			args: &models.User{Name: "fail", Age: 12},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectExec(regexp.QuoteMeta("INSERT INTO users (id, name, age, created_at) VALUES ($1, $2, $3, $4)")).
					WithArgs(sqlmock.AnyArg(), "fail", 12, sqlmock.AnyArg()).
					WillReturnError(errors.New("db error"))
			},
			want: nil,
			err:  errors.New("db error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			assert.NoError(t, err)
			defer mockDB.Close()

			tc.beforeTest(mockSQL)
			db := sqlx.NewDb(mockDB, "sqlmock")
			repo := postgres.New(db)

			got, err := repo.AddUser(tc.args)
			assert.Equal(t, tc.err, err)
			if tc.want != nil {
				assert.Equal(t, tc.want.Name, got.Name)
				assert.Equal(t, tc.want.Age, got.Age)
			} else {
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name       string
		beforeTest func(sqlmock.Sqlmock)
		want       []*models.User
		err        error
	}{
		{
			name: "success get users",
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "created_at"}).
					AddRow("1", "John", 12, time.Unix(0, 0)).
					AddRow("2", "Doe", 21, time.Unix(0, 0))
				mockSQL.
					ExpectQuery("^SELECT (.+) FROM users ORDER BY created_at DESC$").
					WillReturnRows(rows)
			},
			want: []*models.User{
				{
					ID:        "1",
					Name:      "John",
					Age:       12,
					CreatedAt: time.Unix(0, 0),
				},
				{
					ID:        "2",
					Name:      "Doe",
					Age:       21,
					CreatedAt: time.Unix(0, 0),
				},
			},
			err: nil,
		},
		{
			name: "db error",
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery("^SELECT (.+) FROM users ORDER BY created_at DESC$").
					WillReturnError(errors.New("db error"))
			},
			want: nil,
			err:  errors.New("db error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			assert.NoError(t, err)
			defer mockDB.Close()

			tc.beforeTest(mockSQL)
			db := sqlx.NewDb(mockDB, "sqlmock")
			repo := postgres.New(db)

			got, err := repo.GetUsers()
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
