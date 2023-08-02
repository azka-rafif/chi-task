package run

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Run struct {
	Id            string      `db:"id"`
	UserId        string      `db:"user_id"`
	RunTime       string      `db:"run_time"`
	RunKilometers float64     `db:"run_kilometers"`
	RunCity       string      `db:"run_city"`
	CreatedAt     time.Time   `db:"created_at"`
	UpdatedAt     time.Time   `db:"updated_at"`
	DeletedAt     null.Time   `db:"deleted_at"`
	CreatedBy     string      `db:"created_by"`
	UpdatedBy     string      `db:"updated_by"`
	DeletedBy     null.String `db:"deleted_by"`
}

type RunPayload struct {
	RunTime       string  `json:"runTime"`
	RunKilometers float64 `json:"runKilometers"`
	RunCity       string  `json:"runCity"`
}

type RunResponseFormat struct {
	Id            string    `json:"id"`
	UserId        string    `json:"userId"`
	RunTime       string    `json:"runTime"`
	RunKilometers float64   `json:"runKilometers"`
	RunCity       string    `json:"runCity"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     null.Time `json:"deletedAt"`
	CreatedBy     string    `json:"createdBy"`
	UpdatedBy     string    `json:"updatedBy"`
	DeletedBy     string    `json:"deletedBy"`
}

func (r Run) NewFromPayload(payload RunPayload) (res Run) {
	runId, _ := uuid.NewV4()
	userId, _ := uuid.NewV4()
	res = Run{
		Id:            runId.String(),
		UserId:        userId.String(),
		RunTime:       payload.RunTime,
		RunKilometers: payload.RunKilometers,
		RunCity:       payload.RunCity,
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     userId.String(),
		UpdatedAt:     time.Now().UTC(),
		UpdatedBy:     userId.String(),
	}
	return
}

func (r *Run) Update(payload RunPayload) {
	r.UpdatedAt = time.Now().UTC()
	r.RunCity = payload.RunCity
	r.RunKilometers = payload.RunKilometers
	r.RunTime = payload.RunTime
}

func (r *Run) ToResponseFormat() RunResponseFormat {
	res := RunResponseFormat{
		Id:            r.Id,
		UserId:        r.UserId,
		RunTime:       r.RunTime,
		RunKilometers: r.RunKilometers,
		RunCity:       r.RunCity,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
		DeletedAt:     r.DeletedAt,
		CreatedBy:     r.CreatedBy,
		UpdatedBy:     r.UpdatedBy,
		DeletedBy:     r.DeletedBy.String,
	}
	return res
}

func (r Run) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ToResponseFormat())
}
