package model

func CreateToDo(name, todo string) error {
	insert, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)
	defer insert.Close()
	if err != nil {
		return err
	}
	return nil

}
