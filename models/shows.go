package models

type Show struct {
	ID      int32
	Name    string
	Season  int32
	Episode int32
}

func (db *DB) AllShows() ([]*Show, error) {
	rows, err := db.Query("SELECT * FROM shows")
	if err != nil {
		_, _ = db.Exec("create table shows (ID integer PRIMARY KEY, Name string not null, Season integer not null, Episode integer not null)")
		return nil, err
	}
	defer rows.Close()

	shows := make([]*Show, 0)
	for rows.Next() {
		sh := new(Show)
		err := rows.Scan(&sh.ID, &sh.Name, &sh.Season, &sh.Episode)
		if err != nil {
			return nil, err
		}
		shows = append(shows, sh)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shows, nil
}

func (db *DB) LoadShow(ID int32) (*Show, error) {
	sh := new(Show)
	r := db.QueryRow("SELECT * FROM shows WHERE ID = ?", ID)
	err := r.Scan(&sh.ID, &sh.Name, &sh.Season, &sh.Episode)
	if err != nil {
		return nil, err
	}

	return sh, nil

}

func (db *DB) AddShow(Name string, Season int32, Episode int32) error {
	stmt, err := db.Prepare("insert into shows (Name, Season, Episode) values (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(Name, Season, Episode)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateShow(ID int32, Name string, Season int32, Episode int32) error {
	stmt, err := db.Prepare("UPDATE shows SET Name = ?, Season = ?, Episode = ? WHERE ID = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(Name, Season, Episode, ID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteShow(ID int32) error {
	stmt, err := db.Prepare("DELETE FROM shows WHERE ID = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}
