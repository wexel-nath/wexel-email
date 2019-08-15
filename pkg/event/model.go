package event

import (
	"fmt"
	"time"
)

type Event struct {
	ID      int64     `json:"event_id"`
	EmailID int64     `json:"email_id"`
	TypeID  int64     `json:"event_type_id"`
	Content string    `json:"event_content"`
	Created time.Time `json:"event_created"`
}

func newEvent(row map[string]interface{}) (Event, error) {
	event := Event{}
	var ok bool

	if event.ID, ok = row[columnEventID].(int64); !ok {
		return event, fmt.Errorf("row[%v] does not contain field[%s] type[int64]", row, columnEventID)
	}
	if event.EmailID, ok = row[columnEmailID].(int64); !ok {
		return event, fmt.Errorf("row[%v] does not contain field[%s] type[int64]", row, columnEmailID)
	}
	if event.TypeID, ok = row[columnEventTypeID].(int64); !ok {
		return event, fmt.Errorf("row[%v] does not contain field[%s] type[int64]", row, columnEventTypeID)
	}
	if event.Content, ok = row[columnEventContent].(string); !ok {
		return event, fmt.Errorf("row[%v] does not contain field[%s] type[string]", row, columnEventContent)
	}
	if event.Created, ok = row[columnEventCreated].(time.Time); !ok {
		return event, fmt.Errorf("row[%v] does not contain field[%s] type[time.Time]", row, columnEventCreated)
	}

	return event, nil
}
