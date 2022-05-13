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

func (*Chip) RequestLine(gpioOffset int, inputOption interface{}) (*Line, error) {
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

type OutputOption []int

func AsOutput(values ...int) OutputOption {
	vv := append([]int(nil), values...)
	return OutputOption(vv)
}

func (l *Line) SetValue(value int) error {
	return errors.New("not implemented for platform darwin")
}
