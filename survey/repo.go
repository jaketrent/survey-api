package survey

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func findAll(db *sql.DB) ([]*Survey, error) {
	log.Print("Find all surveys...")
	const query = `
select id
, description
from survey
`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Find all surveys error (msg: %s)", err.Error())
		return nil, err
	}
	defer rows.Close()

	surveys := make([]*Survey, 0)
	for rows.Next() {
		var survey Survey
		if err := rows.Scan(&survey.Id, &survey.Description); err != nil {
			log.Printf("Find all surveys row scan error (msg: %s)", err.Error())
			return nil, err
		}
		surveys = append(surveys, &survey)
	}
	log.Printf("Find all surveys success (count: %v)", len(surveys))
	return surveys, nil
}

func insertSurvey(db *sql.DB, survey *Survey) (*Survey, error) {
	const query = `
insert into survey
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
update survey
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
delete from answer
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
delete from question
where survey_id = $1;
`, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	result, err = tx.Exec(`
delete from survey
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
