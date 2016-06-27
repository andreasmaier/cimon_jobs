package jobs

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
)

func CreateJobInDb(job *JenkinsJob) (*JenkinsJob, error) {
	if job.Path == "" {
		return nil, errors.New("Empty path name for job not allowed")
	}

	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
	defer con.Close()

	if err != nil {
		return nil, err
	}

	var path string
	if err := con.QueryRow("SELECT path from jobs where path=?", job.Path).Scan(&path); err != nil {
		if result, err := con.Exec("INSERT INTO jobs (path, status, alias) values (?, 'undefined', ?)", job.Path, job.Alias); err != nil {
			return nil, err
		} else {
			id, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}

			fmt.Printf("Added new job to database with id [%d] and path '%s'\n", id, job.Path)

			newJob := new(JenkinsJob)
			if err := con.QueryRow("SELECT * from jobs where id=?", id).
				Scan(&newJob.Id, &newJob.Path, &newJob.Status, &newJob.Alias); err != nil {
				return nil, err
			}

			return newJob, nil
		}
	} else {
		return nil, errors.New("A job with that path already exists in the database")
	}
}

func GetAllJobsFromDb() ([]*JenkinsJob, error) {
	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
	defer con.Close()

	if err != nil {
		return nil, err
	}

	var jobs []*JenkinsJob
	rows, err := con.Query("SELECT * from jobs")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		job := new(JenkinsJob)
		if err = rows.Scan(&job.Id, &job.Path, &job.Status, &job.Alias); err != nil {
			panic(err)
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func UpdateJobInDb(path string, status string) error {
	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
	defer con.Close()

	if err != nil {
		return err
	}

	if _, err := con.Exec("UPDATE jobs SET status=? WHERE path=?", status, path); err != nil {
		return err
	} else {
		fmt.Printf("Updated job %d to status '%s'", path, status)

		return nil
	}
}
