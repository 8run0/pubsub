package vehicle

import (
	"encoding/json"
	"time"
)

type Registration string

type Vehicle struct {
	Registration Registration `json:"registration"`
}

func (v Vehicle) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(v)
	return bytes, err
}

type Event struct {
	Payload any `json:"vehicle_event"`
}

func (e *Event) UnmarshalJSON(b []byte) error {
	ae := &ArrivalEvent{}
	err := json.Unmarshal(b, ae)
	if err == nil && (ae.Time != time.Time{}) {
		e.Payload = ae
		return nil
	}
	if _, ok := err.(*json.UnmarshalTypeError); err != nil && !ok {
		return err
	}

	de := &DepartureEvent{}
	err = json.Unmarshal(b, de)
	if err != nil {
		return err
	}

	e.Payload = de
	return nil
}

func (e Event) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(e)
	return bytes, err
}
func (e *Event) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, e); err != nil {
		return err
	}
	return nil
}

func (e Event) GetPayload() any {
	return e.Payload
}

type VehicleEvent interface {
	GetVehicle() *Vehicle
	GetTime() time.Time
}

type ArrivalEvent struct {
	Vehicle *Vehicle  `json:"vehicle"`
	Time    time.Time `json:"arrival_time"`
}

func (a ArrivalEvent) Compare(b ArrivalEvent) int {
	if a.Vehicle == nil || b.Vehicle == nil {
		return 0
	}
	return 0
}
func (v ArrivalEvent) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(v)
	return bytes, err
}

func (v *ArrivalEvent) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

func (e ArrivalEvent) GetVehicle() *Vehicle {
	return e.Vehicle
}
func (e ArrivalEvent) GetTime() time.Time {
	return e.Time
}

type DepartureEvent struct {
	Vehicle *Vehicle  `json:"vehicle"`
	Time    time.Time `json:"departure_time"`
}

func (a DepartureEvent) Compare(b DepartureEvent) int {
	if a.Vehicle == nil || b.Vehicle == nil {
		return 0
	}
	return 0
}

func (v DepartureEvent) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(v)
	return bytes, err
}
func (v *DepartureEvent) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
func (e DepartureEvent) GetVehicle() *Vehicle {
	return e.Vehicle
}
func (e DepartureEvent) GetTime() time.Time {
	return e.Time
}

var (
	Channel = "vehicles"
)
