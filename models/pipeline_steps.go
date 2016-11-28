package models

import "database/sql"

type PipelineStep struct {
	Id              int
	StepClass       string
	State           string
	AsyncResult     []byte
	AsyncResultData map[string]interface{}
}

func QueryPipelineSteps(pipelineId int, db *sql.DB) ([]*PipelineStep, error) {
	rows, err := db.Query("SELECT id, step_class, state, async_result from pipeline_steps WHERE pipeline_id = $1 ORDER BY id", pipelineId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	steps := make([]*PipelineStep, 0)

	for rows.Next() {
		step := new(PipelineStep)

		err := rows.Scan(&step.Id, &step.StepClass, &step.State, &step.AsyncResult)

		if err != nil {
			return nil, err
		}

		steps = append(steps, step)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return steps, nil
}
