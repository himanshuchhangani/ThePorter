package dataprocessor

import (
	"encoding/json"
	"log"

	"github.com/himanshuchhangani/theporter/interfaces"
	"github.com/himanshuchhangani/theporter/models"
)

type Processor struct {
	DecoderEngine *json.Decoder
	StorageEngine interfaces.StorageHandler
}

func (p *Processor) Process() error {

	// read open bracket
	_, err := p.DecoderEngine.Token()
	if err != nil {
		log.Fatal(err)
	}
	// while the array contains values
	for p.DecoderEngine.More() {
		var portData map[string]models.Port
		if err := p.DecoderEngine.Decode(&portData); err != nil {
			log.Fatal("error parsing json " + err.Error())
		}
		for key, value := range portData {
			if err := p.StorageEngine.Save(key, &value); err != nil {
				log.Fatalf("error saving json in db " + err.Error())
			}
		}

	}
	// read closing bracket
	_, err = p.DecoderEngine.Token()
	if err != nil {
		log.Fatal(err)
	}

	return nil

}
