package model

func CreateTodo (name, todo string) error {
	insertQ, err := con.Query("insert into TODO values (?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
