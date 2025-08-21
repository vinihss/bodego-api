package domain

import "reflect"

type EntityInterface interface {
	GetName() string
}

type Entity struct {
}

func (e Entity) GetName() string {
	return reflect.TypeOf(e).Name()
}
