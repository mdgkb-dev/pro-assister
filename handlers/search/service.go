package search

import (
	"context"

	"github.com/pro-assistance/pro-assister/models"
)

func (s *Service) Search(c context.Context, searchModel *models.SearchModel) (err error) {
	searchGroup, err := R.GetGroupByKey(c, searchModel.Key)
	if err != nil {
		return err
	}
	searchModel.SearchGroup = searchGroup
	err = R.Search(c, searchModel)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SearchMain(c context.Context, searchModel *models.SearchModel) (err error) {
	searchModel.SearchGroups, err = R.GetGroups(c, searchModel.SearchGroupID)
	if err != nil {
		return err
	}
	for i := range searchModel.SearchGroups {
		err = R.Search(c, searchModel)
		if err != nil {
			return err
		}
		searchModel.SearchGroups[i].BuildRoutes()
	}
	return nil
}
