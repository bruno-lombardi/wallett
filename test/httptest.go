package test

import (
	"bytes"
	"testing"
)

type HttpTestCase struct {
	Name         string
	WhenURL      string
	WhenBody     string
	WhenMethod   string
	ExpectStatus int
	ExpectBody   func(t *testing.T, body *bytes.Buffer) error
}
