package tree

import (
	"encoding/json"
	"fmt"

	"pro-assister/helpers/project"
)

// treeModel model
type treeModel struct { //nolint:golint
	Model string   `json:"model"`
	Cols  []string `json:"col"`
}

type TreeModels []*treeModel //nolint:golint

func parseJSONToTreeModel(args string) (treeModel treeModel, err error) {
	err = json.Unmarshal([]byte(args), &treeModel)
	if err != nil {
		return treeModel, err
	}
	return treeModel, err
}

func (s *treeModel) getTableAndCols() string {
	schema := project.SchemasLib.GetSchema(s.Model)
	return fmt.Sprintf("%s.%s", schema.GetTableName(), schema.GetCol(s.Cols[0]))
}
