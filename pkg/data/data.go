package data

import (
	_ "embed"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//go:embed data.yaml
var fileData []byte

type Data struct {
	Texts map[int]string `yaml:"texts"`
}

type Service struct {
	texts map[int]string
}

func New() (*Service, error) {
	var dt Data
	if err := yaml.Unmarshal(fileData, &dt); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal yaml data")
	}
	if len(dt.Texts) == 0 {
		return nil, fmt.Errorf("empty array of texts in data yaml file")
	}
	return &Service{
		texts: dt.Texts,
	}, nil
}

func (s *Service) GetTextAfter(lastSentTextID int) (text string, id int) {
	minID := -1
	nextID := -1
	for textID := range s.texts {
		if minID == -1 || textID < minID {
			minID = textID
		}
		if textID > lastSentTextID {
			if nextID == -1 || textID < nextID {
				nextID = textID
			}
		}
	}
	if nextID != -1 {
		// There is a "next" text to post -- choosing it.
		return s.texts[nextID], nextID
	}
	// The 'lastSentTextID' is the last one in our array. Starting posting from the beginning.
	return s.texts[minID], minID
}
