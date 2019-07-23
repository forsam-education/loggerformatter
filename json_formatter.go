package loggerformatters

import (
	"encoding/json"
	"fmt"
	"github.com/forsam-education/simplelogger"
	"time"
)

// JSONFormatter is a simplelogger formatter that returns a JSON formatted log when called, with some context.
type JSONFormatter struct {
	ServiceName   string
	CorrelationID string
}

// LogMessage is the structure to stringify log message details to JSON.
type LogMessage struct {
	ServiceName   string                    `json:"service_name"`
	Time          string                    `json:"time"`
	Level         string                    `json:"level"`
	Message       string                    `json:"message"`
	CorrelationID string                    `json:"correlation_id"`
	ExtraData     simplelogger.LogExtraData `json:"extra_data"`
}

// Format returns the Log as a stringified JSON, using the format specified above in LogMessage.
func (formatter JSONFormatter) Format(level simplelogger.LogLevel, message string, data simplelogger.LogExtraData) (string, error) {
	logMessage := LogMessage{
		ServiceName:   formatter.ServiceName,
		Time:          time.Now().UTC().Format(time.RFC3339),
		Message:       message,
		Level:         level.String(),
		CorrelationID: formatter.CorrelationID,
		ExtraData:     data,
	}

	jsonData, err := json.Marshal(logMessage)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", string(jsonData)), nil
}
