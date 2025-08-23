package domain

import "reflect"

type EntityInterface interface {
	GetName() string
	GetID() uint
}

type Entity struct {
	ID uint
}

func (e Entity) GetName() string {
	return reflect.TypeOf(e).Name()
}
func (e Entity) GetID() uint {
	return e.ID
}
