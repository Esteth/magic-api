package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // driver for database/sql
	"github.com/pkg/errors"

	pb "esteth.net/magic/api/proto"
)

// DB provides access to values stored in the database.
type DB struct {
	*sql.DB
}

// Dial connects to the database and returns a Conn to query it.
func Dial(connString string) (*DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to database")
	}

	for attempts := 1; attempts < 10; attempts++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("could not ping database: %v", err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("database still not up. exiting")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not ping database")
	}

	return &DB{db}, nil
}

func (db *DB) GetWaitTimes(attractionID string) ([]*pb.WaitTime, error) {
	rows, err := db.Query("SELECT wait_time, timestamp FROM wait_times WHERE attraction_id = $1 ORDER BY timestamp DESC LIMIT 1", attractionID)
	if err != nil {
		return nil, errors.Wrap(err, "could not issue query")
	}
	defer rows.Close()

	results := make([]*pb.WaitTime, 0, 10)

	for rows.Next() {
		var waitTime int32
		var timestamp int64
		err := rows.Scan(&waitTime, &timestamp)
		if err != nil {
			return nil, errors.Wrap(err, "Could not scan row")
		}
		results = append(results, &pb.WaitTime{
			WaitTime: waitTime,
		})
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error encountered in rows object")
	}

	return results, nil
}

func (db *DB) GetAttractions() ([]*pb.Attraction, error) {
	var count int32
	err := db.QueryRow("SELECT count(*) FROM attractions").Scan(&count)
	if err != nil {
		return nil, errors.Wrap(err, "could not get count of attractions")
	}

	rows, err := db.Query("SELECT id, name FROM attractions")
	if err != nil {
		return nil, errors.Wrap(err, "could not issue query")
	}
	defer rows.Close()

	results := make([]*pb.Attraction, 0, count)

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, errors.Wrap(err, "could not scan id and name")
		}
		results = append(results, &pb.Attraction{
			Id:   id,
			Name: name,
		})
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error encountered in rows object")
	}
	return results, nil
}
