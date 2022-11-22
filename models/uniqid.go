package md

import (
	"fmt"
)

type MyUniqId struct {
	Id    int64  `json:"id" xorm:"autoincr unique notnull pk 'id'"`
	Types string `json:"types" xorm:"'types'"`
	Name  string `json:"name" xorm:"'name'"`
}

func CreateId(types, name string) (string, error) {
	str := ""
	id := MyUniqId{}
	id.Types = types
	id.Name = name

	session := localdb.NewSession()
	if err := session.Begin(); err != nil {
		return "", err
	}

	if _, err := session.Insert(&id); err != nil {
		return "", err
	}

	if result, err := session.Exec("SELECT last_insert_rowid()"); err != nil {
		session.Rollback()
		return str, err
	} else {
		if lastinsertid, errres := result.LastInsertId(); errres != nil {
			return str, errres
		} else {
			str = id.Types + "-" + id.Name + "-" + fmt.Sprintf("%0*d", 6, lastinsertid)
		}
	}

	if err := session.Commit(); err != nil {
		return "", err
	}
	return str, nil
}
