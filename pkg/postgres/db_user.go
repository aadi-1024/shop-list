package postgres

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) AddUser(username, passHash string) error {
	query := `insert into users (username, pass_hash) values ($1, $2)`
	//queryUniq := `select 0 < (select count(*) from users where username = $1)`

	//row := d.Db.QueryRow(queryUniq, username)
	//
	//var count bool
	//if err := row.Scan(&count); err != nil {
	//	return err
	//}
	//
	//if !count {
	//	return errors.New("username already exists")
	//}

	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, username, passHash)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Database) VerifyPassHash(username, password string) (int, error) {
	query := `select id, pass_hash from users where username = $1`

	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(query, username)
	var id int
	var originalHash string

	if err = row.Scan(&id, &originalHash); err != nil {
		return -1, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(originalHash), []byte(password)); err != nil {
		return -1, err
	}
	return id, nil
}
