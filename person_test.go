package test

import (
	"reflect"
	"testing"
)

func TestID(t *testing.T) {
	p := &Person{10, "coco"}
	if !reflect.DeepEqual(10, p.ID()) {
		t.Errorf("required int value: 10")
	}
}

func TestName(t *testing.T) {
	p := &Person{10, "coco"}
	if !reflect.DeepEqual("coco", p.Name()) {
		t.Errorf("required string value: coco")
	}
}

func TestString(t *testing.T) {
	p := &Person{10, "coco"}
	if !reflect.DeepEqual("10coco", p.String()) {
		t.Errorf("required string value: 10coco")
	}
}
