package interfaces

import (
	"io"

	"github.com/himanshuchhangani/theporter/models"
)

type FileDecoder interface {
	Read(*io.Reader) (map[string]models.Port, error)
}
