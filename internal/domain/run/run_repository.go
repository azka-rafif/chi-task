package run

import (
	"fmt"

	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/logger"
)

type RunRepositoryMySQL struct {
	DB *infras.MySQLConn
}

type RunRepository interface {
	Create(load Run) (err error)
	GetAll(limit, offset int, sort, field string, location string) (res []Run, err error)
	Update(load Run) (err error)
	GetByID(id string) (res Run, err error)
	Delete(id string) (err error)
}

func NewRunRepositoryMySQL(db *infras.MySQLConn) *RunRepositoryMySQL {
	return &RunRepositoryMySQL{DB: db}
}

func (r *RunRepositoryMySQL) Create(load Run) (err error) {
	query := `INSERT INTO run (id,user_id,run_time,run_kilometers,run_city,created_at,updated_at,created_by,updated_by)
    VALUES (:id,:user_id,:run_time,:run_kilometers,:run_city,:created_at,:updated_at,:created_by,:updated_by)`

	stmt, err := r.DB.Write.PrepareNamed(query)

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(load)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}

func (r *RunRepositoryMySQL) GetAll(limit, offset int, sort, field string, location string) (res []Run, err error) {
	query := `SELECT * FROM run `

	if location != "" {
		query += fmt.Sprintf(`WHERE run_city = '%s' `, location)
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT %d OFFSET %d", field, sort, limit, offset)
	err = r.DB.Read.Select(&res, query)
	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
	}
	return
}

func (r *RunRepositoryMySQL) Update(load Run) (err error) {
	query := `UPDATE run
    SET
        run_time = :run_time,
        run_kilometers = :run_kilometers,
        run_city = :run_city,
        updated_at = :updated_at,
        updated_by = :updated_by,
        deleted_at = :deleted_at,
        deleted_by = :deleted_by
    WHERE id = :id
    `

	stmt, err := r.DB.Write.PrepareNamed(query)

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(load)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}

func (r *RunRepositoryMySQL) GetByID(id string) (res Run, err error) {
	err = r.DB.Read.Get(&res, "SELECT * FROM run WHERE id = ?", id)
	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *RunRepositoryMySQL) Delete(id string) (err error) {
	_, err = r.DB.Write.Exec("DELETE FROM run WHERE id = ?", id)
	return
}
