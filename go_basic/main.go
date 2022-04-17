package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TimeDuration() {
	stringTime := "2022-03-25 18:32:58"

	loc, _ := time.LoadLocation("Local")

	t0, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	t1 := time.Now()

	duration := t1.Sub(t0)
	fmt.Println(duration)
	if duration < time.Second*time.Duration(60) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

func main() {
	//AcessSQL()
	fmt.Println(len("bd5797b4577c19838dec848a36d28834"))
	fmt.Printf("%v\n", time.Now().Unix())
}

func AcessSQL() {
	db, err := sql.Open("mysql",
		"tmadmin:tm@pswd123@tcp(127.0.0.1:3306)/rtc?parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("link sucess!\n")
	}
	defer db.Close()
	taskId := "127365aa-fab1-4086-b2fa-00dcefa07b39"
	// var fileds map[string]interface{}

	// fileds := make(map[string]interface{})
	// fileds["task_status"] = 0
	// fileds["worker_id"] = "test_id"
	// fileds["create_time"] = time.Now()
	// err = UpdateTaskInfo(taskId, fileds, db)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// time.Sleep(time.Second * 1)
	t1, err := GetTaskInfo(taskId, db)
	fmt.Println(t1)
	t0 := time.Now()
	fmt.Println(t0)
	duration := t0.Sub(t1)
	fmt.Println(duration)
}

func GetTaskInfo(taskId string, db *sql.DB) (time.Time, error) {
	querySQL := fmt.Sprintf("SELECT update_time FROM %s WHERE task_id = ?", "task_info_list")

	stmt, _ := db.Prepare(querySQL)

	defer stmt.Close()

	row := stmt.QueryRow(taskId)

	var update_time time.Time

	row.Scan(&update_time)

	return update_time, nil
}

func UpdateTaskInfo(taskId string, fields map[string]interface{}, db *sql.DB) (err error) {
	var builder strings.Builder
	sql := fmt.Sprintf("UPDATE task_info_list SET ")
	builder.WriteString(sql)
	var args []interface{}
	fieldLength := len(fields)
	for field, value := range fields {
		fieldLength--
		builder.WriteString(field)
		if fieldLength == 0 {
			builder.WriteString(" = ?")
		} else {
			builder.WriteString(" = ?,")
		}

		args = append(args, value)
	}

	builder.WriteString(" WHERE task_id = ?")
	args = append(args, taskId)

	updateSQL := builder.String()
	fmt.Printf(updateSQL)
	stmt, err := db.Prepare(updateSQL)
	if err != nil {
		fmt.Printf("failed to do UpdateTaskInfoOnlyFlag")
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(args...); err != nil {
		fmt.Printf("failed to do UpdateTaskInfoOnlyFlag")
		return err
	}
	return nil
}
