package main

import (
	"fmt"
	"github.com/greebear/CronJob/internal"
)

const StudentType int = 1

type StudentFactory struct{}

func (f StudentFactory) Handler() error {
	fmt.Println("")
	return nil
}

func NewStudentCronJob(Id int, jobType int, jobFactory internal.Factory) internal.CronJob {
	return internal.CronJob{
		JobId:      Id,
		JobType:    jobType,
		JobFactory: jobFactory,
	}
}

func main() {

	jobTypeList := make(map[int]internal.Factory)
	var jobList []internal.TaskModel

	for _, eachJob := range jobList {
		implFactory := jobTypeList[eachJob.BusinessType]

		newCronJob := NewStudentCronJob(eachJob.Id, eachJob.BusinessType, implFactory)
		// processing
		newCronJob.Run()
		err := newCronJob.JobFactory.Handler()
		// fail
		if err != nil {
			newCronJob.Fail()
			continue
		}
		// success
		newCronJob.Success()
		continue

	}
}
