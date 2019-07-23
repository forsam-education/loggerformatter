package loggerformatter

import (
	"github.com/forsam-education/simplelogger"
	"gotest.tools/assert"
	"testing"
)

func TestJSONFormatter_Format(t *testing.T) {
	formatter := JSONFormatter{ServiceName: "TestService", CorrelationID: "abcd"}

	_, err := formatter.Format(simplelogger.DEBUG, "Hello this is my message", nil)
	assert.NilError(t, err)

	_, err := formatter.Format(simplelogger.DEBUG, "Hello this is my message", nil)
	assert.NilError(t, err)
}
