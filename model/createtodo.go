package model

func CreateToDo(name, todo string) error {
	insert, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)
	defer insert.Close()
	if err != nil {
		return err
	}
	return nil

}
func DeleteToDo(name string) error {
	_, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	
	if err != nil {
		return err
	}
	return nil

}
