package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogstashFormatter(t *testing.T) {
	lf := Logstash{Type: "abc"}

	fields := logrus.Fields{
		"message": "def",
		"level":   "ijk",
		"type":    "lmn",
		"one":     1,
		"pi":      3.14,
		"bool":    true,
		"error":   &url.Error{Op: "Get", URL: "http://example.com", Err: fmt.Errorf("The error")},
	}

	entry := logrus.WithFields(fields)
	entry.Message = "msg"
	entry.Level = logrus.InfoLevel

	b, _ := lf.Format(entry)

	var data map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	dec.Decode(&data)

	// base fields
	if expected := json.Number("1"); expected != data["@version"] {
		t.Fatalf("Expected version to be %s but got %s", expected, data["@version"])
	}
	if data["@timestamp"] == "" {
		t.Fatal("Expected timestamp not to be empty")
	}
	if expected := "abc"; expected != data["type"] {
		t.Fatalf("Expected type to be %s but got %s", expected, data["type"])
	}
	if expected := "msg"; expected != data["message"] {
		t.Fatalf("Expected message to be %s but got %s", expected, data["message"])
	}
	if expected := "info"; expected != data["level"] {
		t.Fatalf("Expected level to be %s but got %s", expected, data["level"])
	}
	if expected := "Get http://example.com: The error"; expected != data["error"] {
		t.Fatalf("Expected error to be %s but got %s", expected, data["error"])
	}

	// substituted fields
	if expected := "def"; expected != data["fields.message"] {
		t.Fatalf("Expected fields.message to be %s but got %s", expected, data["fields.message"])
	}
	if expected := "ijk"; expected != data["fields.level"] {
		t.Fatalf("Expected fields.level to be %s but got %s", expected, data["fields.level"])
	}
	if expected := "lmn"; expected != data["fields.type"] {
		t.Fatalf("Expected fields.type to be %s but got %s", expected, data["fields.type"])
	}

	// formats
	if expected := json.Number("1"); expected != data["one"] {
		t.Fatalf("Expected one to be %s but got %s", expected, data["one"])
	}
	if expected := json.Number("3.14"); expected != data["pi"] {
		t.Fatalf("Expected pi to be %s but got %s", expected, data["pi"])
	}
	if data["bool"] != true {
		t.Fatal("Expected bool to be true but got false")
	}
}
