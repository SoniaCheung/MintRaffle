package main

import (
	"fmt"
	"math/rand"
	"soniacheung/mint-raffle/cmd/mint-raffle/app/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

func main() {
	var engine *xorm.Engine
	engine, err := xorm.NewEngine("mysql", "root:Root1234@tcp(127.0.0.1:3306)/MintRaffle?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer engine.Close()

	fmt.Println("MySQL DB connection established")

	session := engine.NewSession()
	defer session.Close()

	// update expired project status to pending
	var project models.Project
	if _, err = session.Table(project.TableName()).Where("due_time <= NOW()").Update(map[string]interface{}{"status": "pending"}); err != nil {
		logrus.Fatalf("Update status to pending error: %s", err.Error())
		return
	}

	// select all pending projects
	peningProjects := make([]models.Project, 0)
	if err = session.Table(project.TableName()).Where("status = ?", "pending").Find(&peningProjects); err != nil {
		logrus.Fatalf("Select pending projects error: %s", err.Error())
		return
	}

	for _, p := range peningProjects {
		// get submissions
		var submission models.Submission
		submissions := make([]models.Submission, 0)
		if err = session.Table(submission.TableName()).Where("project_id = ?", p.Id).Find(&submissions); err != nil {
			logrus.Fatalf("Select submissions error: %s", err.Error())
			return
		}
		// pick winners
		shuffle(submissions)
		max := min(len(submissions), p.MaxWinner)
		submissions = submissions[0:max]
		selectedIds := make([]int, 0)
		for _, s := range submissions {
			selectedIds = append(selectedIds, s.Id)
		}

		// update winners
		if _, err = session.Table(submission.TableName()).In("id", selectedIds).Update(map[string]interface{}{"winner": true}); err != nil {
			logrus.Fatalf("Update submission winner error: %s", err.Error())
			return
		}

		// update project status
		if _, err = session.Table(project.TableName()).Where("id = ?", p.Id).Update(map[string]interface{}{"status": "closed"}); err != nil {
			logrus.Fatalf("Update project status to closed error: %s", err.Error())
			return
		}
	}

	fmt.Println("Raffle finished")

}

func shuffle(slice []models.Submission) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
