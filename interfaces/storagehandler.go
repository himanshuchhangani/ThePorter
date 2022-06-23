package interfaces

import (
	"github.com/himanshuchhangani/theporter/models"
)

// StorageHandler interface defines the methods which are needed to be implemented
// by structs in order to be used a storage medium.
type StorageHandler interface {
	Save(string, *models.Port) error
	Get(string) (*models.Port, error)
}
