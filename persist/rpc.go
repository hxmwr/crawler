package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(s.Client, s.Index, item)
	if err == nil  {
		*result = "ok"
	}
	return err
}