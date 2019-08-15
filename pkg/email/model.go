package email

import (
	"encoding/json"
	"fmt"
)

// Email represents a single record of the email table
type Email struct {
	ID      int64   `json:"email_id"`
	To      string  `json:"email_to"`
	From    string  `json:"email_from"`
	Subject string  `json:"email_subject"`
	Content Content `json:"email_content"`
	Service string  `json:"service_name"`
}

type Content struct {
	Product Product `json:"product"`
	Body    Body    `json:"body"`
}

func newEmail(row map[string]interface{}) (Email, error) {
	email := Email{}
	var ok bool

	if email.ID, ok = row[columnEmailID].(int64); !ok {
		return email, fmt.Errorf("row[%v] does not contain field[%s] type[int64]", row, columnEmailID)
	}
	if email.To, ok = row[columnEmailTo].(string); !ok {
		return email, fmt.Errorf("row[%v] does not contain field[%s] type[string]", row, columnEmailTo)
	}
	if email.From, ok = row[columnEmailFrom].(string); !ok {
		return email, fmt.Errorf("row[%v] does not contain field[%s] type[string]", row, columnEmailFrom)
	}
	if email.Subject, ok = row[columnEmailSubject].(string); !ok {
		return email, fmt.Errorf("row[%v] does not contain field[%s] type[string]", row, columnEmailSubject)
	}
	if content, ok := row[columnEmailContent].([]byte); !ok {
		return email, fmt.Errorf("row[%v] does not contain field[%s] type[[]byte]", row, columnEmailContent)
	} else {
		err := json.Unmarshal(content, &email.Content)
		if err != nil {
			return email, err
		}
	}
	if email.Service, ok = row[columnServiceName].(string); !ok {
		return email, fmt.Errorf("row[%v] does not contain field[%s] type[string]", row, columnServiceName)
	}

	return email, nil
}
