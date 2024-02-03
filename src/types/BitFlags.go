package types

import "strings"

type BitFlagType interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type BitFlag[T BitFlagType] struct {
	Label string
	Value T
}

type BitFlags[T BitFlagType] map[T]BitFlag[T]

func NewBitFlags[T BitFlagType](flag T, labels map[T]string) BitFlags[T] {
	flags := BitFlags[T]{}
	for value, label := range labels {
		if (value & flag) != 0 {
			flags[value] = BitFlag[T]{
				Label: label,
				Value: value,
			}
		}
	}
	return flags
}

func (flags BitFlags[T]) Add(flag BitFlag[T]) {
	flags[flag.Value] = flag
}

func (flags BitFlags[T]) Remove(flag BitFlag[T]) {
	delete(flags, flag.Value)
}

func (flags BitFlags[T]) Contains(flag BitFlag[T]) bool {
	_, ok := flags[flag.Value]
	return ok
}

func (flags BitFlags[T]) Flags() []T {
	values := make([]T, 0, len(flags))
	for value := range flags {
		values = append(values, value)
	}
	return values
}

func (flags BitFlags[T]) Values() []BitFlag[T] {
	values := make([]BitFlag[T], 0, len(flags))
	for _, value := range flags {
		values = append(values, value)
	}

	return values
}

func (flags BitFlags[T]) Labels() []string {
	values := make([]string, 0, len(flags))
	for _, value := range flags {
		values = append(values, value.Label)
	}
	return values
}

func (flags BitFlags[T]) String() string {
	values := make([]string, 0, len(flags))
	for _, value := range flags {
		values = append(values, value.Label)
	}
	return strings.Join(values, "|")
}
