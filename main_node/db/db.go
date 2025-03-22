package db

import (
	"database/sql"

	"github.com/Du-bem/dominari_scriptum/main_node/types"
)

type DB struct {
	*sql.DB
}

const CREATE_DATABASE string = `
	CREATE TABLE IF NOT EXISTS satellite_data (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,

		position_x REAL NOT NULL,
		position_y REAL NOT NULL,
		position_z REAL NOT NULL,

		velocity_x REAL NOT NULL,
		velocity_y REAL NOT NULL,
		velocity_z REAL NOT NULL,

		recorded_on DATETIME NOT NULL,
		created_on DATETIME NOT NULL,
		updated_on DATETIME NOT NULL,
		checksum TEXT NOT NULL
	);
`

// NewDatabase returns a new instance of the database
func NewDatabase(fileName string) (*DB, error) {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(CREATE_DATABASE); err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}

func (d *DB) InsertSatelliteData(types.SatelliteData) (int, error) {
	res, err := d.Exec("INSERT INTO satellite_data VALUES(?,?,?,?,?,?,?,?,?,?,?);", activity.Time, activity.Description)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}
