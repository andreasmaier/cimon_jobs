package jobs

import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	"fmt"
	"errors"
)

type JenkinsJob struct {
	Id int32 `json:"id"`
	Server string `json:"server"`
	Path string `json:"path"`
	Status string `json:"status"`
	Alias string `json:"alias"`
}

func createJobInDb(job *JenkinsJob) (*JenkinsJob, error) {
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

func getAllJobs() ([]*JenkinsJob, error) {
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

//func UpdateJob(id int, status string) {
//	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
//	defer con.Close()
//
//	if err != nil {
//		panic(err)
//	}
//
//	if _, err := con.Exec("UPDATE jobs SET status=? WHERE id=?", status, id); err != nil {
//		panic(err)
//	} else {
//		fmt.Printf("Updated job %d to status '%s'", id, status)
//	}
//}
//
//func isJobWatched(path string) bool {
//	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
//	defer con.Close()
//
//	if err != nil {
//		panic(err)
//	}
//
//	return (con.QueryRow("SELECT path from jobs where path=?", path).Scan(&path) == nil)
//}
//
//func GetByPath(path string) *JenkinsJob {
//	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
//	defer con.Close()
//
//	if err != nil {
//		panic(err)
//	}
//
//	job := new(JenkinsJob)
//
//	if err := con.QueryRow("SELECT * from jobs where path=?", path).
//			Scan(&job.Id, &job.Path, &job.Status, &job.Alias); err != nil {
//		panic(err)
//	}
//
//	return job
//}