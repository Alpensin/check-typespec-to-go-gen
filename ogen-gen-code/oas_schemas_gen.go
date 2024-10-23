// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

// Ref: #/components/schemas/Pet
type Pet struct {
	Name string  `json:"name"`
	Age  int32   `json:"age"`
	Kind PetKind `json:"kind"`
}

// GetName returns the value of Name.
func (s *Pet) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *Pet) GetAge() int32 {
	return s.Age
}

// GetKind returns the value of Kind.
func (s *Pet) GetKind() PetKind {
	return s.Kind
}

// SetName sets the value of Name.
func (s *Pet) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *Pet) SetAge(val int32) {
	s.Age = val
}

// SetKind sets the value of Kind.
func (s *Pet) SetKind(val PetKind) {
	s.Kind = val
}

type PetKind string

const (
	PetKindDog  PetKind = "dog"
	PetKindCat  PetKind = "cat"
	PetKindFish PetKind = "fish"
)

// AllValues returns all PetKind values.
func (PetKind) AllValues() []PetKind {
	return []PetKind{
		PetKindDog,
		PetKindCat,
		PetKindFish,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s PetKind) MarshalText() ([]byte, error) {
	switch s {
	case PetKindDog:
		return []byte(s), nil
	case PetKindCat:
		return []byte(s), nil
	case PetKindFish:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *PetKind) UnmarshalText(data []byte) error {
	switch PetKind(data) {
	case PetKindDog:
		*s = PetKindDog
		return nil
	case PetKindCat:
		*s = PetKindCat
		return nil
	case PetKindFish:
		*s = PetKindFish
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}