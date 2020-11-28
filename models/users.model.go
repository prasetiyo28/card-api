package models

import (
	"net/http"

	"github.com/prasetiyo28/card-api/db"
)

type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FetchAllUsers() (Response, error) {
	var obj Users
	var arrayobj []Users
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM USERS"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Email, &obj.Password, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}

		arrayobj = append(arrayobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayobj

	return res, nil

}

func StoreUsers(username string, email string, password string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "INSERT users (username,email,password ,created_at,updated_at) values(?,?,?,NOW(),NOW()) "
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, password)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateUsers(id int, username string, email string, password string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE users SET username = ?, email= ?, password = ? where id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, password, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
