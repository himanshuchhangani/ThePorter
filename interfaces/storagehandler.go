package interfaces

import (
	"github.com/himanshuchhangani/theporter/models"
)

type StorageHandler interface {
	Save(string, *models.Port) error
	Get(string) (*models.Port, error)
}
