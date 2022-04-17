package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func AcessSQL() {
	db, err := sql.Open("mysql",
		"tmadmin:tm@pswd123@tcp(127.0.0.1:3306)/rtc?parseTime=true")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("link sucess!\n")
	}
	defer db.Close()
	taskId := "127365aa-fab1-4086-b2fa-00dcefa07b39"
	// var fileds map[string]interface{}
	fileds := make(map[string]interface{})
	fileds["task_status"] = 0
	fileds["worker_id"] = "test_id"
	fileds["create_time"] = time.Now()
	err = UpdateTaskInfo(taskId, fileds, db)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateTaskInfo(taskId string, fields map[string]interface{}, db *sql.DB) (err error) {
	var builder strings.Builder
	sql := fmt.Sprintf("UPDATE rtc SET")
	builder.WriteString(sql)
	var args []interface{}
	for field, value := range fields {
		builder.WriteString(field)
		builder.WriteString(" = ?,")
		args = append(args, value)
	}
	builder.WriteString(" WHERE task_id = ?")
	args = append(args, taskId)

	updateSQL := builder.String()

	stmt, err := db.Prepare(updateSQL)
	if err != nil {
		fmt.Printf("failed to do UpdateTaskInfoOnlyFlag")
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(args); err != nil {
		fmt.Printf("failed to do UpdateTaskInfoOnlyFlag")
		return err
	}
	return nil
}
