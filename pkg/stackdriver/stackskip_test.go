package stackdriver

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/TV4/logrus-stackdriver-formatter/test"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestStackSkip(t *testing.T) {
	var out bytes.Buffer

	logger := logrus.New()
	logger.Out = &out
	logger.Formatter = NewFormatter(
		WithService("test"),
		WithVersion("0.1"),
		WithStackSkip("github.com/TV4/logrus-stackdriver-formatter/test"),
	)

	mylog := test.LogWrapper{
		Logger: logger,
	}

	mylog.Error("my log entry")

	var got map[string]interface{}
	_ = json.Unmarshal(out.Bytes(), &got)

	want := map[string]interface{}{
		"severity": "ERROR",
		"message":  "my log entry",
		"serviceContext": map[string]interface{}{
			"service": "test",
			"version": "0.1",
		},
		"context": map[string]interface{}{
			"reportLocation": map[string]interface{}{
				"filePath":     "github.com/jenkins-x/logrus-stackdriver-formatter/pkg/stackdriver/stackskip_test.go",
				"lineNumber":   28.0,
				"functionName": "TestStackSkip",
			},
		},
	}

	assert.Equal(t, want, got)
}
