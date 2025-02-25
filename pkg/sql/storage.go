package sql

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	SQLHandler[T ManagedObject] interface {
		Create(T) error
		Read(T) (*T, error)
	}

	SQLHandlerImpl[T ManagedObject] struct {
		*sqlx.DB
	}

	ManagedObject interface {
		Insert() string
		Get() string
	}
)

func NewSQLHandler[T ManagedObject]() (SQLHandler[T], error) {
	conn, err := connect()
	if err != nil {
		return nil, fmt.Errorf("error while connecting to the database, %s", err.Error())
	}

	return &SQLHandlerImpl[T]{conn}, nil
}

func (s *SQLHandlerImpl[T]) Create(obj T) error {
	_, err := s.NamedExec(obj.Insert(), obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLHandlerImpl[T]) Read(obj T) (*T, error) {
	var inst T

	err := s.Get(&inst, obj.Get())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return &inst, err
		}
	}

	return &inst, nil
}
