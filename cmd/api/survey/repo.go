package survey

import (
	"database/sql"
	"errors"
	"fmt"
)

func findAll(db *sql.DB) ([]*Survey, error) {
	const query = `
select id
, description
from survey_api_survey
`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	surveys := make([]*Survey, 0)
	for rows.Next() {
		var survey Survey
		if err := rows.Scan(&survey.Id, &survey.Description); err != nil {
			return nil, err
		}
		surveys = append(surveys, &survey)
	}
	return surveys, nil
}

func insertSurvey(db *sql.DB, survey *Survey) (*Survey, error) {
	const query = `
insert into survey_api_survey
( description
) values
( $1
) returning id
`
	err := db.QueryRow(query, survey.Description).Scan(&survey.Id)
	return survey, err
}

func updateSurvey(db *sql.DB, survey *Survey) (*Survey, error) {
	const query = `
update survey_api_survey
set description = $1
where id = $2
returning description
`
	err := db.QueryRow(query, survey.Description, survey.Id).Scan(&survey.Description)
	return survey, err
}

func deleteSurvey(db *sql.DB, id int) error {
	var result sql.Result
	var err error
	var tx *sql.Tx
	var count int64

	tx, err = db.Begin()

	if err != nil {
		return err
	}

	result, err = tx.Exec(`
delete from survey_api_answer
where question_id in (
  select id
  from question
  where survey_id = $1
)
`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	result, err = tx.Exec(`
delete from survey_api_question
where survey_id = $1;
`, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	result, err = tx.Exec(`
delete from survey_api_survey
where id = $1;
`, id)
	count, err = result.RowsAffected()
	if count != 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("delete survey modified inappropriate rows (count: %v)", count))
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return err
}
