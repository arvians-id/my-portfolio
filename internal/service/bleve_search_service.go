package service

import (
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/blevesearch/bleve"
)

type BleveSearchServiceContract interface {
	Search(query string) ([]string, error)
	InsertOrUpdate(data *entity.SearchItem) error
	Clear() error
}

type BleveSearchService struct {
	bleveIndex bleve.Index
}

func NewBleveSearchService(bleveIndex bleve.Index) BleveSearchServiceContract {
	return &BleveSearchService{bleveIndex: bleveIndex}
}

func (s *BleveSearchService) Search(query string) ([]string, error) {
	matchQuery := bleve.NewMatchQuery(query)
	matchQuery.SetFuzziness(1)
	matchQuery.SetBoost(1.9)

	searchQuery := bleve.NewQueryStringQuery(query)
	disjunction := bleve.NewDisjunctionQuery(matchQuery, searchQuery)
	searchRequest := bleve.NewSearchRequest(disjunction)
	searchResult, err := s.bleveIndex.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, hit := range searchResult.Hits {
		result = append(result, hit.ID)
	}

	return result, nil
}

func (s *BleveSearchService) InsertOrUpdate(data *entity.SearchItem) error {
	return s.bleveIndex.Index(data.ID, data)
}

func (s *BleveSearchService) Clear() error {
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	searchResponse, err := s.bleveIndex.Search(searchRequest)
	if err != nil {
		return err
	}

	batch := s.bleveIndex.NewBatch()
	for _, hit := range searchResponse.Hits {
		batch.Delete(hit.ID)
	}

	return s.bleveIndex.Batch(batch)
}
