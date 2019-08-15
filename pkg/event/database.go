package event

import (
	"strings"

	"github.com/wexel-nath/wexel-email/pkg/database"
)

const (
	columnEventID      = "event_id"
	columnEmailID      = "email_id"
	columnEventTypeID  = "event_type_id"
	columnEventContent = "event_content"
	columnEventCreated = "event_created"
)

var (
	insertColumns = []string{
		columnEmailID,
		columnEventTypeID,
		columnEventContent,
	}

	selectColumns = []string{
		columnEventID,
		columnEmailID,
		columnEventTypeID,
		columnEventContent,
		columnEventCreated,
	}
)

func insert(emailID int64, typeID int64, content string) (map[string]interface{}, error) {
	query := `
		INSERT INTO event (
			` + strings.Join(insertColumns, ", ") + `
		)
		VALUES (
			$1, $2, $3
		)
		RETURNING
			` + strings.Join(selectColumns, ", ")

	db := database.GetConnection()
	row := db.QueryRow(query, emailID, typeID, content)
	return database.ScanRowToMap(row, selectColumns)
}
