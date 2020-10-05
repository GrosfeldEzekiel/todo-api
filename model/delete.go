package model

import "fmt"

func Delete(name string) error {
	insertQ, err := con.Query("delete from TODO where name=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}