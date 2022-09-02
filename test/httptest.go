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
	BeforeTest   func() error
	ExpectBody   func(t *testing.T, body *bytes.Buffer) error
	AfterTest    func() error
}
