package models

import "database/sql"

type Pipeline struct {
	Id    int
	State string
	Type  string
}

func QueryPipelines(db *sql.DB) ([]*Pipeline, error) {
	rows, err := db.Query("SELECT id, state, pipeline_type from pipelines ORDER BY id")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pipelines := make([]*Pipeline, 0)

	for rows.Next() {
		pipeline := new(Pipeline)

		err = rows.Scan(&pipeline.Id, &pipeline.State, &pipeline.Type)
		if err != nil {
			return nil, err
		}

		pipelines = append(pipelines, pipeline)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pipelines, nil
}
