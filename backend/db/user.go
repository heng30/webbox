package db

type UserInfo struct {
    Name string
    Passwd string
}

func QueryUsers() ([]UserInfo, error) {
	var users []UserInfo
	s := "SELECT * FROM users"
	if rows, err := sdb.Query(s); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var uid int
			var name, passwd string
			if err := rows.Scan(&uid, &name, &passwd); err != nil {
				continue
			}
            users = append(users, UserInfo{name, passwd})
		}
	}
	return users, nil
}

func AddUser(name, passwd string) error {
	s := "INSERT INTO users(name, passwd) values(?, ?)"
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(name, passwd); err != nil {
			return err
		}
	}
	return nil
}

func DelUser(name string) error {
	s := "DELETE FROM users where name=?"
	if stmt, err := sdb.Prepare(s); err != nil {
		return err
	} else {
		if _, err := stmt.Exec(name); err != nil {
			return err
		}
	}
	return nil
}
