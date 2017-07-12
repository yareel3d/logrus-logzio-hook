package logzio

import (
	"bytes"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Formatter func(*logrus.Entry) ([]byte, error)

type Hook struct {
	client        *http.Client
	address       string
	appName       string
	contextFields logrus.Fields
	formatter     Formatter
}

var defaultFormatter = formatter.Logstash{}

// New creates a new Logz.io logrus hook
func New(address, appName string, contextFields logrus.Fields) *Hook {
	return &Hook{
		client:        &http.Client{},
		address:       address,
		appName:       appName,
		contextFields: contextFields,
		formatter:     defaultFormatter.Format,
	}
}

// SetClient sets the hook client to the given client
func (h *Hook) SetClient(client *http.Client) *Hook {
	h.client = client
	return h
}

// SetFormatter sets the hook formatter to the given formatter
func (h *Hook) SetFormatter(formatter Formatter) *Hook {
	h.formatter = formatter
	return h
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	// Add in context fields.
	for k, v := range h.contextFields {
		// We don't override fields that are already set
		if _, ok := entry.Data[k]; !ok {
			entry.Data[k] = v
		}
	}

	method := http.MethodPost
	if m, ok := entry.Data["HTTP.Method"]; ok {
		method = m
	}
	reader = bytes.NewBuffer(h.formatter(entry))
	req, err := http.NewRequest(method, h.address, reader)
	return err
}

func (h *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
