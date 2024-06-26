package postgres

import (
	"context"
	"shoplist/pkg/models"
)

func (d *Database) Insert(title, desc string, uid int) (int, error) {
	query := `insert into list_item (title, dsc, userid) values ($1, $2, $3) returning id`

	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(query, title, desc, uid)
	if err != nil {
		return -1, err
	}

	var id int

	if err = row.Scan(&id); err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (d *Database) Delete(id, uid int) error {
	query := `delete from list_item where id = $1 and userid = $2`

	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, id, uid)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Update(data models.ListItem) (models.ListItem, error) {
	query := `update list_item set title=$1, dsc=$2 where id = $3 and userid = $4 returning *`

	m := models.ListItem{}

	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return m, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(query, data.Title, data.Description, data.Id, data.UserId)

	if err = row.Scan(&m.Id, &m.Title, &m.Description, &m.UserId); err != nil {
		return m, err
	}

	err = tx.Commit()
	return m, nil
}

func (d *Database) GetAll(uid int) ([]models.ListItem, error) {
	query := `select * from list_item where userid = $1`

	//even though this is a read-only query, transactions are still advised to prevent receiving inconsistent data
	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]models.ListItem, 0)

	for rows.Next() {
		m := models.ListItem{}
		err = rows.Scan(&m.Id, &m.Title, &m.Description, &m.UserId)
		if err != nil {
			return data, nil
		}
		data = append(data, m)
	}
	err = rows.Err()
	return data, err
}

func (d *Database) GetById(id, uid int) (*models.ListItem, error) {
	query := `select * from list_item where id = $1 and userid = $2`

	ret := &models.ListItem{}

	tx, err := d.Db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(query, id, uid)

	if err = row.Scan(&ret.Id, &ret.Title, &ret.Description, &ret.UserId); err != nil {
		return nil, err
	}

	return ret, nil
}
