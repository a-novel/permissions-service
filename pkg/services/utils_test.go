package services_test

import (
	"errors"
	"time"
)

var (
	baseTime = time.Date(2020, time.May, 4, 8, 0, 0, 0, time.UTC)
)

var (
	fooErr = errors.New("foo")
)
