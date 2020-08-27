package model

import view "../views"

func ReadAll() ([]view.PostRequest, error) {
	rows, err := con.Query("select * from todo")
	if err != nil {
		return nil, err
	}

	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}
func ReadByName(name string) ([]view.PostRequest, error) {
	rows, err := con.Query("select * from todo where name=?", name)
	if err != nil {
		return nil, err
	}

	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	if err==nil {
		return nil,nil
	}

	return todos, nil
}
