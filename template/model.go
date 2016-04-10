package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/josephspurrier/gowebapi/shared/database"

	"github.com/boltdb/bolt"
)

// Name of the bucket
const bucketName = "entity"

// Entity information
type Entity struct{}

// Group list
type Group []Entity

// Errors
var (
	ErrNoResult = errors.New("no results")
	ErrNoChange = errors.New("no change")
)

// New entity
func New() *Entity {
	i := &Entity{}
	return i
}

// List all entities
func ListAll() (Group, error) {
	var u Group
	e := database.BoltDB.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		// If bucket is not found
		if b == nil {
			return ErrNoResult
		}

		// Get the iterator
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var single Entity

			// Decode the record
			err := json.Unmarshal(v, &single)
			if err != nil {
				log.Println(err)
				continue
			}

			u = append(u, single)
		}

		return nil
	})

	return u, e
}

// Delete all entities
func DeleteAll() (int, error) {
	// Number of items
	count := 0

	e := database.BoltDB.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		// If bucket is not found
		if b == nil {
			return ErrNoChange
		}

		// Number of current level items
		count = b.Stats().KeyN

		err := tx.DeleteBucket([]byte(bucketName))
		// If bucket is not found
		if err == bolt.ErrBucketNotFound {
			return ErrNoResult
		} else if err != nil {
			// Set the number of items deleted
			count = 0

			return err
		}

		return nil
	})

	return count, e
}

// Create entity
func (u *Entity) Create() (int, error) {
	// Number of items
	count := 0

	e := database.BoltDB.Update(func(tx *bolt.Tx) error {
		// Create the bucket
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		// Set the time
		now := time.Now()
		u.Created_at = now

		if len(u.Id) == 0 {
			id, err := bucket.NextSequence()
			if err != nil {
				return err
			}
			u.Id = fmt.Sprintf("%d", id)
		}

		// Encode the record
		bytes, err := json.Marshal(u)
		if err != nil {
			return err
		}

		// Store the record
		if err = bucket.Put([]byte(u.Id), bytes); err != nil {
			return err
		}

		// Set then number of items added
		count = 1

		return nil
	})

	return count, e
}
