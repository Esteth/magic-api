package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	pb "esteth.net/magic/api/proto"
	wkt_timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/lib/pq" // driver for database/sql
)

// DB provides access to values stored in the database.
type DB struct {
	*sql.DB
}

// Dial connects to the database and returns a Conn to query it.
func Dial(connString string) (*DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
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
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return &DB{db}, nil
}

func (db *DB) GetWaitHistory(attractionID string) ([]*pb.WaitTime, error) {
	rows, err := db.Query("SELECT wait_time, timestamp FROM wait_times WHERE attraction_id = $1 ORDER BY timestamp DESC LIMIT 100", attractionID)
	if err != nil {
		return nil, fmt.Errorf("could not issue query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("could not close rows: %v\n", err)
		}
	}()

	results := make([]*pb.WaitTime, 0, 10)

	for rows.Next() {
		var waitTime int32
		var timestamp time.Time
		err := rows.Scan(&waitTime, &timestamp)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		results = append(results, &pb.WaitTime{
			AttractionId: attractionID,
			WaitTime:     waitTime,
			LastUpdated: &wkt_timestamp.Timestamp{
				Seconds: timestamp.Unix(),
			},
		})
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error encountered in rows after getting wait times: %w", err)
	}

	return results, nil
}

func (db *DB) GetWaitTimes() ([]*pb.WaitTime, error) {
	rows, err := db.Query(
		`SELECT attractions.id, wait_times.wait_time, wait_times.timestamp
				FROM attractions
				CROSS JOIN LATERAL (
					SELECT wait_times.timestamp, wait_times.wait_time
					FROM wait_times
					WHERE attractions.id = wait_times.attraction_id
					ORDER BY wait_times.timestamp DESC NULLS LAST
					LIMIT 1
					) wait_times;`)
	if err != nil {
		return nil, fmt.Errorf("could not issue query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("could not close rows: %v\n", err)
		}
	}()

	results := make([]*pb.WaitTime, 0, 10)

	for rows.Next() {
		var attractionID string
		var waitTime int32
		var timestamp time.Time
		err := rows.Scan(&attractionID, &waitTime, &timestamp)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		results = append(results, &pb.WaitTime{
			AttractionId: attractionID,
			WaitTime:     waitTime,
			LastUpdated: &wkt_timestamp.Timestamp{
				Seconds: timestamp.Unix(),
			},
		})
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error encountered in rows after getting wait times: %w", err)
	}

	return results, nil
}

func (db *DB) GetAttractions() ([]*pb.Attraction, error) {
	var count int32
	err := db.QueryRow("SELECT count(*) FROM attractions").Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("could not get count of attractions: %w", err)
	}

	rows, err := db.Query("SELECT id, name FROM attractions")
	if err != nil {
		return nil, fmt.Errorf("could not issue query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("could not close rows: %v\n", err)
		}
	}()

	results := make([]*pb.Attraction, 0, count)

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, fmt.Errorf("could not scan id and name: %w", err)
		}
		results = append(results, &pb.Attraction{
			Id:   id,
			Name: name,
		})
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error encounted in rows object after reading attractions: %w", err)
	}
	return results, nil
}
