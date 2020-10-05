package model

import "LearningGo/section4/todo-api/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("select * from TODO")
	if err != nil {
		return nil, err
	}
	var todos []views.PostRequest
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("select * from TODO where name=?", name)
	if err != nil {
		return nil, err
	}
	var todos []views.PostRequest
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
