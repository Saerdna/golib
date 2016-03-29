package cmp

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	a := []string{"1", "2", "5", "-1"}
	b := []string{"3", "2", "4"}
	more := FindMoreItem(a, b)
	morebuffer := strings.Join(more, ",")
	ansbuffer := strings.Join([]string{"1", "5", "-1"}, ",")
	if morebuffer != ansbuffer {
		t.Errorf("less:%s  ans:%s", morebuffer, ansbuffer)
	}
}

func TestTwo(t *testing.T) {
	a := []string{"1", "2", "5", "-1"}
	b := []string{"3", "2", "4"}
	more, less := Compare(a, b)
	morebuffer := strings.Join(more, ",")
	lessbuffer := strings.Join(less, ",")
	ansmorebuffer := strings.Join([]string{"1", "5", "-1"}, ",")
	anslessbuffer := strings.Join([]string{"3", "4"}, ",")
	if lessbuffer != anslessbuffer {
		t.Errorf("less:%s:%s", lessbuffer, anslessbuffer)
	}
	if morebuffer != ansmorebuffer {
		t.Errorf("more:%s:%s", morebuffer, ansmorebuffer)
	}
}
