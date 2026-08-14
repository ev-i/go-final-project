package db

import "fmt"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

func UpdateTask(task *Task) error {

	query := `UPDATE scheduler SET date=?, title=?, comment=?, repeat=? WHERE id = ?`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func Tasks(limit int) ([]*Task, error) {
	tasks := []*Task{}

	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler LIMIT ?", limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	t := Task{}
	err := db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id).Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
