package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

type task struct {
	ID    int
	Value string
}

//Init opens or creates the databse
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//GetTasks will return a slice of all tasks in the bucket
func GetTasks() ([]task, error) {
	var tasks []task
	db.View(func(tx *bolt.Tx) error {
		// Assume our events bucket exists and has RFC3339 encoded time keys.
		c := tx.Bucket(taskBucket).Cursor()

		// Iterate over the 90's.
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, task{
				ID:    btoi(k),
				Value: string(v),
			})
		}

		return nil
	})

	return tasks, nil
}

//AddTask will add a task to the bucket
func AddTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil

}

//DeleteTasks will remove the tasks with the ids provided from the bucket
func DeleteTasks(ids []int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		for _, id := range ids {
			err := b.Delete(itob(id))
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
