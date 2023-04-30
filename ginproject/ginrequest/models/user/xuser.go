package user

import (
	"database/sql"
	"fmt"
	"log"
)

type XUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func (user *XUser) Get(db *sql.DB) (u XUser, err error) {
	row := db.QueryRow("select id, name, age, phone from user where id=?", user.Id)
	err = row.Scan(&u.Id, &u.Name, &u.Age, &u.Phone)

	if err != nil {
		fmt.Println("---XUser get---", err)
	}

	return u, err
}

func (user *XUser) GetAll(db *sql.DB) (users []XUser, err error) {
	rows, err := db.Query("select id,name,age,phone from user")
	fmt.Println(rows)
	if err != nil {
		return
	}
	for rows.Next() {
		var u XUser
		rows.Scan(&u.Id, &u.Name, &u.Age, &u.Phone)
		users = append(users, u)
	}
	defer rows.Close()
	return users, err
}

func (user *XUser) Add(db *sql.DB) (Id int, err error) {
	stmt, err := db.Prepare("INSERT into user(name,age,phone) values (?,?,?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(user.Name, user.Age, user.Phone)
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	Id = int(id)
	defer stmt.Close()
	return Id, err
}

func (user *XUser) Update(db *sql.DB) (rows int, err error) {
	stmt, err := db.Prepare("update user set name=?,age=?,phone=? where id=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(user.Name, user.Age, user.Phone, user.Id)
	if err != nil {
		log.Fatalln(err)
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	rows = int(row)
	defer stmt.Close()
	return rows, err
}

func (user *XUser) Del(db *sql.DB) (rows int, err error) {
	stmt, err := db.Prepare("delete from user where id=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(user.Id)
	if err != nil {
		log.Fatalln(err)
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	rows = int(row)
	defer stmt.Close()
	return rows, err
}
