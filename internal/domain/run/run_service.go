package run

type RunServiceImpl struct {
	Repo RunRepository
}

type RunService interface {
	Create(payload RunPayload) (res Run, err error)
	GetAll(offset, limit int, field, sort, location string) (res []Run, err error)
	Update(id string, load RunPayload) (res Run, err error)
	Delete(id string) (err error)
}

func NewRunServiceImpl(repo RunRepository) *RunServiceImpl {
	return &RunServiceImpl{Repo: repo}
}

func (s *RunServiceImpl) Create(payload RunPayload) (res Run, err error) {
	res = res.NewFromPayload(payload)
	err = s.Repo.Create(res)
	return
}

func (s *RunServiceImpl) GetAll(offset, limit int, field, sort, location string) (res []Run, err error) {
	res, err = s.Repo.GetAll(limit, offset, sort, field, location)
	if err != nil {
		return
	}
	return
}

func (s *RunServiceImpl) Update(id string, load RunPayload) (res Run, err error) {
	res, err = s.Repo.GetByID(id)
	if err != nil {
		return
	}
	res.Update(load)
	err = s.Repo.Update(res)
	return
}

func (s *RunServiceImpl) Delete(id string) (err error) {
	_, err = s.Repo.GetByID(id)
	if err != nil {
		return
	}
	err = s.Repo.Delete(id)
	return
}
