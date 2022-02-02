//go:build darwin
// +build darwin

package gpiod

import "errors"

func NewChip(chipName string) (*Chip, error) {
	return &Chip{}, errors.New("not implemented for platform darwin")
}

type Chip struct {
	ChipName string
}

func (*Chip) Close() error {
	return errors.New("not implemented for platform darwin")
}

func (*Chip) RequestLine(gpioOffset int, inputOption int) (*Line, error) {
	return &Line{}, errors.New("not implemented for platform darwin")
}

type Line struct {
}

func (*Line) Value() (int, error) {
	return 0, errors.New("not implemented for platform darwin")
}

func (*Line) Close() error {
	return errors.New("not implemented for platform darwin")
}

const AsInput int = 0
