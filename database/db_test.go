package database

import (
	"strings"
	"testing"
)

func TestGetDBVersion(t *testing.T) {
	expectedVersion := "PostgreSQL 16.0"
	version, err := GetDBVersion()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if !strings.Contains(version, expectedVersion) {
		t.Errorf("Expected version %q, got: %q", expectedVersion, version)
	}
}
