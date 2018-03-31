package survey

import (
	"database/sql"
)

func insertSurvey(db *sql.DB, survey *Survey) (*Survey, error) {
	const query = `
insert into survey
( desc
) values
( $1
) returning id
`
	err := db.QueryRow(survey.Desc).Scan(&survey.Id)
	return survey, err
}
