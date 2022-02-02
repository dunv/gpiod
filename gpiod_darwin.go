//go:build darwin
// +build darwin

package gpiod

import "errors"

func NewChip(chipName string) (FakeChip, error) {
	return FakeChip{}, nil
}

type FakeChip struct {
	ChipName string
}

func (FakeChip) Close() error {
	return nil
}

func (FakeChip) RequestLine(gpioOffset int, inputOption int) (FakeInput, error) {
	return FakeInput{}, nil
}

type FakeInput struct {
}

func (FakeInput) Value() (int, error) {
	return 0, errors.New("not implemented for platform darwin")
}

const AsInput int = 0
