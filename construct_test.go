package serr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicConstruction(t *testing.T) {
	errx := New("Test error message")

	// expected
	exPath := "/serr/construct_test.go:11"
	exMsg := "Test error message"

	if !assert.Contains(t, fmt.Sprintf("%s:%d", errx.File(), errx.Line()), exPath, "Error trace") {
		t.FailNow()
	}

	if !assert.Equal(t, errx.Error(), exMsg, "Error message") {
		t.FailNow()
	}
}

func TestBasicWithCommentConstruction(t *testing.T) {
	errx := Newc("Test error message", "Test comments")

	// expected
	exPath := "/serr/construct_test.go:27"
	exMsg := "Test error message"
	exComment := "Test comments"

	if !assert.Contains(t, fmt.Sprintf("%s:%d", errx.File(), errx.Line()), exPath, "Error trace") {
		t.FailNow()
	}

	if !assert.Equal(t, errx.Comments(), exComment, "Error comment") {
		t.FailNow()
	}

	if !assert.Equal(t, errx.Error(), exMsg, "Error message") {
		t.FailNow()
	}
}
