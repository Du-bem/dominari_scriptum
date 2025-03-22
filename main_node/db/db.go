package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/Du-bem/dominari_scriptum/main_node/types"
)

type DB struct {
	*sql.DB
	ctx context.Context
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
func NewDatabase(ctx context.Context, fileName string) (*DB, error) {
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

// Insert pushes the satellites raw data into the db.
func (d *DB) Insert(checksum string, data *types.SatelliteRequestData) (int, error) {
	res, err := d.ExecContext(d.ctx, "INSERT INTO satellite_data VALUES(?,?,?,?,?,?,?,?,?,?,?);",
		data.Name, data.Position[0], data.Position[1], data.Position[2],
		data.Velocity[0], data.Velocity[1], data.Velocity[2],
		data.RecordedOn, time.Now().UTC(), time.Now().UTC(), checksum)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

// List returns a batch of 10 records of the Satellite data if they exist
func (d *DB) List(offset int) ([]*types.SatelliteUIData, error) {
	rows, err := d.QueryContext(d.ctx,
		"SELECT * FROM satellite_data ORDER BY id DESC OFFSET ? LIMIT 10", offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := make([]*types.SatelliteUIData, 0)
	for rows.Next() {
		i := types.SatelliteUIData{}
		err = rows.Scan(&i.ID, &i.Name, &i.PositionX, &i.PositionY, &i.PositionZ,
			&i.VelocityX, &i.VelocityY, &i.VelocityZ, &i.RecordedOn,
			&i.CreatedOn, &i.UpdatedOn, &i.CheckSum)
		if err != nil {
			return nil, err
		}
		data = append(data, &i)
	}
	return data, nil
}

// SearchByID return the Satellite UI data identified by the record ID provided.
func (d *DB) SearchByID(recordID int) (*types.SatelliteUIData, error) {
	row := d.QueryRowContext(d.ctx, "SELECT * FROM satellite_data WHERE ID = ?", recordID)

	i := types.SatelliteUIData{}
	err := row.Scan(&i.ID, &i.Name, &i.PositionX, &i.PositionY, &i.PositionZ,
		&i.VelocityX, &i.VelocityY, &i.VelocityZ, &i.RecordedOn,
		&i.CreatedOn, &i.UpdatedOn, &i.CheckSum)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
