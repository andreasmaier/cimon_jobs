package jobs

import (
	"golang.org/x/net/context"
)

type JobsServer struct{}

func (s *JobsServer) CreateJob(ctx context.Context, in *CreateJobRequest) (*Job, error) {
	job := in.toJenkinsJob()

	createdJob, err := createJobInDb(job)

	if err != nil {
		return nil, err
	}

	return &Job{
		Id: createdJob.Id,
		Path: createdJob.Path,
		Status: createdJob.Status,
		Alias: createdJob.Alias,
	}, nil
}

func (s *JobsServer) GetAllJobs(ctx context.Context, in *Empty) (*Jobs, error) {
	jjobs, err := getAllJobs()

	if err != nil {
		panic(err)
	}

	jobs := make([]*Job, len(jjobs))

	for idx, jj := range jjobs {
		jobs[idx] = jj.tojobMessage()
	}

	return &Jobs{Jobs: jobs}, nil
}

func (r *CreateJobRequest) toJenkinsJob() (*JenkinsJob) {
	return &JenkinsJob{
		Path: r.Path,
		Status: r.Status,
		Alias: r.Alias,
	}
}

func (jj *JenkinsJob) tojobMessage() (*Job) {
	return &Job{
		Id: jj.Id,
		Path: jj.Path,
		Alias: jj.Alias,
		Status: jj.Alias,
	}
}
