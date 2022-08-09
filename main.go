package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	_ "github.com/go-sql-driver/mysql"
)

type task struct {
	Id       int
	Task     string
	Assignee string
	Deadline string
	Status   string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "freedb_usman:?Zmp67zTwEp7$%J@tcp(sql.freedb.tech:3306)/freedb_db_task")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var id = 1
	// rows, err := db.Query("select id, task, assignee, deadline, status from tbl_tasking where id = ?", id)
	rows, err := db.Query("select Id, Task, Assignee, Deadline, Status from tbl_tasking")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []task

	for rows.Next() {
		var each = task{}
		var err = rows.Scan(&each.Id, &each.Task, &each.Assignee, &each.Deadline, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// for _, each := range result {
	// 	fmt.Println(each.task)
	// }

}

func main() {

	sqlQuery()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Tasking using Golang",
			"name":  "Usman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/get_data", ActionData)
	http.HandleFunc("/save", handleSave)
	http.HandleFunc("/update", handleUpdate)

	fmt.Println("server started at localhost:80")
	http.ListenAndServe(":80", nil)
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Id       int    `json:"id"`
			Task     string `json:"task"`
			Assignee string `json:"assignee"`
			Deadline string `json:"deadline"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		if payload.Id > 0 {
			_, err = db.Exec("update tbl_tasking set Task = ?, Assignee = ? , Deadline = ?  where id = ?", payload.Task, payload.Assignee, payload.Deadline, payload.Id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("update success!")
		} else {
			_, err = db.Exec("INSERT INTO `tbl_tasking` (`Task`, `Assignee`, `Deadline`) VALUES (?, ?, ?)", payload.Task, payload.Assignee, payload.Deadline)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("insert success!")
		}

		var message = `{"status": "success"}`

		w.Write([]byte(message))
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Id     int    `json:"id"`
			Status string `json:"status"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		if payload.Status != "delete" {
			_, err = db.Exec("update tbl_tasking set Status = ? where id = ?", payload.Status, payload.Id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("update success!")
		} else {
			_, err = db.Exec("delete from tbl_tasking where id = ?", payload.Id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("delete success!")
		}

		var message = `{"status": "success"}`

		w.Write([]byte(message))
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}

func ActionData(w http.ResponseWriter, r *http.Request) {
	// data := []struct {
	// 	Id       int
	// 	Task     string
	// 	Assignee string
	// 	Deadline string
	// 	Status   string
	// }{
	// 	{1, "membuat Database", "uuh", "2022-08-08", "close"},
	// 	{1, "membuat table master", "buni", "2022-08-09", "cancel"},
	// 	{1, "membuat table dan input data master", "yaya", "2022-08-10", "open"},
	// }

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var id = 1
	// rows, err := db.Query("select id, task, assignee, deadline, status from tbl_tasking where id = ?", id)
	rows, err := db.Query("select Id, Task, Assignee, Deadline, Status from tbl_tasking")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []task

	for rows.Next() {
		var each = task{}
		var err = rows.Scan(&each.Id, &each.Task, &each.Assignee, &each.Deadline, &each.Status)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonInBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)

}
