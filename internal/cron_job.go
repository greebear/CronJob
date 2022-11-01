package internal

import "errors"

type Factory interface {
	Handler() error
}

type CronJob struct {
	JobId   int
	JobType int
	status  Status

	JobFactory Factory
}

type Status int

const (
	pending    Status = 1
	processing Status = 2
	success    Status = 3
	fail       Status = 4
)

func (s CronJob) GetStatus() Status {
	// get db.tab.status
	return s.status
}

func (s CronJob) Fail() error {
	if s.status == processing {
		// change db.tab.status to fail
		s.status = fail
	}
	return errors.New("wrong status")
}

func (s CronJob) Success() error {
	if s.status == processing {
		// change db.tab.status to success
		s.status = success
	}
	return errors.New("wrong status")
}

func (s CronJob) Run() error {
	if s.status == pending {
		// change db.tab.status to process
		s.status = processing
		return nil
	}
	return errors.New("wrong status")
}
