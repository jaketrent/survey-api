package survey

import (
	"database/sql"
)

func insert(db *sql.DB, survey *Survey) (*Survey, error) {
	const query = `
insert into survey
( id
, desc
) values
( $1
, $2
) returning id
`
	err := db.QueryRow(survey.Id, survey.Desc).Scan(&survey.Id)
	return survey, err
}
