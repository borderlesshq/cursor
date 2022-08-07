package hqcursor

import (
	"encoding/base64"
	"errors"
	"strings"
)

var ErrCursorInvalid = errors.New("invalid cursor")

type Direction string

const (
	Next Direction = "next"
	Prev Direction = "prev"
)

func (d Direction) IsValid() bool {
	switch d {
	case Prev, Next:
		return true
	default:
		return false
	}
}

func (d Direction) String() string {
	return string(d)
}

type Cursor struct {
	Direction Direction
	Index     string
}

func DecodeCursor(pageCursor string) (Cursor, error) {
	if pageCursor == "" {
		return Cursor{}, nil
	}

	b, err := base64.StdEncoding.DecodeString(pageCursor)
	if err != nil {
		return Cursor{}, err
	}

	cursor := string(b)
	split := strings.Split(cursor, ":")
	if len(split) != 2 {
		return Cursor{}, ErrCursorInvalid
	}

	direction := Direction(split[0])

	if !direction.IsValid() {
		return Cursor{}, ErrCursorInvalid
	}

	return Cursor{
		Direction: direction,
		Index:     split[1],
	}, nil
}

func EncodeCursor(index string, direction Direction) string {
	if len(index) == 0 {
		return ""
	}

	return base64.StdEncoding.EncodeToString([]byte((direction.String() + ":" + index)))
}
