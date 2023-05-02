package user

import (
	"database/sql"
	"fmt"
	"log"
)

type XUser struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Age    int    `json:"age"`
}

func (user *XUser) GetUserByPhone(db *sql.DB) (res XUser, err error) {
	row := db.QueryRow("select userid, name, phone, age from user where phone=?;", user.Phone)
	err = row.Scan(&res.UserId, &res.Name, &res.Phone, &res.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("---phone:%s not exist!\n", user.Phone)
		} else {
			fmt.Printf("---GetUserByPhone:%v\n", err)
		}
	} else {
		fmt.Printf("---GetUserByPhone res user---:%v\n", res)
	}

	return res, err
}

func (user *XUser) GetUserById(db *sql.DB) (res XUser, err error) {

	row := db.QueryRow("select userid, name, phone, age from user where userid=?;", user.UserId)
	err = row.Scan(&res.UserId, &res.Name, &res.Phone, &res.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("---userid:%s not exist!\n", res.UserId)
		}
	}

	return res, err
}

func (user *XUser) GetAll(db *sql.DB) (users []XUser, err error) {
	rows, err := db.Query("select userid,name,phone,age from user")
	fmt.Println(rows)
	if err != nil {
		return
	}
	for rows.Next() {
		var u XUser
		rows.Scan(&u.UserId, &u.Name, &u.Phone, &u.Age)
		users = append(users, u)
	}
	defer rows.Close()
	return users, err
}

func (user *XUser) AddNewUser(db *sql.DB) (Id int, err error) {
	stmt, err := db.Prepare("INSERT into user(userid,name,phone,age) values (?,?,?,?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(user.UserId, user.Name, user.Phone, user.Age)
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
	// stmt, err := db.Prepare("update user set name=?,age=?,phone=? where id=?")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	sql := fmt.Sprintf("UPDATE user set name='%s', age='%d', phone='%s' where userid='%s';",
		user.Name,
		user.Age,
		user.Phone,
		user.UserId)
	rs, err := db.Exec(sql)

	fmt.Sprintln(sql)

	if err != nil {
		log.Fatalln(err)
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	rows = int(row)
	fmt.Printf("---XUser---Update:%d\n", rows)

	// defer stmt.Close()
	return rows, err
}

func (user *XUser) Del(db *sql.DB) (rows int, err error) {
	stmt, err := db.Prepare("delete from user where userid=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(user.UserId)
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
