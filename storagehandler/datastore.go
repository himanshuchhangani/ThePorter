package storagehandler

import (
	"encoding/json"

	redis "github.com/go-redis/redis"
	"github.com/himanshuchhangani/theporter/models"
)

type DataStore struct {
	Client *redis.Client
}

// Save method stores the passed Port data in for the presented key.
func (d *DataStore) Save(key string, portData *models.Port) error {
	p, err := json.Marshal(portData)
	if err != nil {
		return err
	}
	d.Client.Set(key, p, 56000)
	return nil
}

// Get method retrieves the stored data corresponding to the passed key.
func (d *DataStore) Get(key string) (*models.Port, error) {
	var p models.Port
	pd, err := d.Client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(pd), &p); err != nil {
		return nil, err
	}
	return &p, nil
}
