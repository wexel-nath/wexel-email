package email

import (
	"strings"

	"github.com/wexel-nath/wexel-email/pkg/database"
)

const (
	columnEmailID      = "email_id"
	columnEmailTo      = "email_to"
	columnEmailFrom    = "email_from"
	columnEmailSubject = "email_subject"
	columnEmailContent = "email_content"
	columnServiceID    = "service_id"
	columnServiceName  = "service_name"
)

var (
	insertColumns = []string{
		columnEmailTo,
		columnEmailFrom,
		columnEmailSubject,
		columnEmailContent,
		columnServiceID,
	}

	selectColumns = []string{
		columnEmailID,
		columnEmailTo,
		columnEmailFrom,
		columnEmailSubject,
		columnEmailContent,
		columnServiceName,
	}
)

func insert(
	to string,
	from string,
	subject string,
	content []byte,
	service string,
) (map[string]interface{}, error) {
	query := `
		INSERT INTO users (
			` + strings.Join(insertColumns, ", ") + `
		)
		VALUES (
			$1, $2, $3, $4, $5
		)
		RETURNING
			` + strings.Join(selectColumns, ", ")

	db := database.GetConnection()
	row := db.QueryRow(query, to, from, subject, content, service)
	return database.ScanRowToMap(row, selectColumns)
}
