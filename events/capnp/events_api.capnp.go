// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Event struct{ capnp.Struct }

// Event_TypeID is the unique identifier for the type Event.
const Event_TypeID = 0x9c032508b61d1d09

func NewEvent(s *capnp.Segment) (Event, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Event{st}, err
}

func NewRootEvent(s *capnp.Segment) (Event, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Event{st}, err
}

func ReadRootEvent(msg *capnp.Message) (Event, error) {
	root, err := msg.RootPtr()
	return Event{root.Struct()}, err
}

func (s Event) String() string {
	str, _ := text.Marshal(0x9c032508b61d1d09, s.Struct)
	return str
}

func (s Event) Type() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Event) HasType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Event) TypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Event) SetType(v string) error {
	return s.Struct.SetText(0, v)
}

// Event_List is a list of Event.
type Event_List struct{ capnp.List }

// NewEvent creates a new list of Event.
func NewEvent_List(s *capnp.Segment, sz int32) (Event_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Event_List{l}, err
}

func (s Event_List) At(i int) Event { return Event{s.List.Struct(i)} }

func (s Event_List) Set(i int, v Event) error { return s.List.SetStruct(i, v.Struct) }

func (s Event_List) String() string {
	str, _ := text.MarshalList(0x9c032508b61d1d09, s.List)
	return str
}

// Event_Promise is a wrapper for a Event promised by a client call.
type Event_Promise struct{ *capnp.Pipeline }

func (p Event_Promise) Struct() (Event, error) {
	s, err := p.Pipeline.Struct()
	return Event{s}, err
}

const schema_fc8938b535319bfe = "x\xda\x12Hu`2d\x15gb`\x08\x94`e" +
	"\xfb\xcf)+\xbb\x8dC\x95y\x0e\x83\xa0\x1c\xe3\xff\x7f" +
	"\xb3\x0dM\xb7Zt\xfea`edg`0\xfc\xa8" +
	"\xc4(\xcc\xc8\xc8.\xcc\xc8(/\xac\xc8h\xcf`\xf1" +
	"?\xb5,5\xaf\xa4X?\x999\xb1 \xaf@\x1f\xc2" +
	"\x8bO,\xc8\xd4K\x06\x09X\xb9\x96\xb1\xa7\xe6\x95\x04" +
	"02\x06\xb202\xfd\x8f\x9b<?p\xef\xb5\xae\xa3" +
	"\x0c\x81,L\x8c\x8e\x02\x8c\x8c<\x0c\x0c\x82\x8c\\\x0c" +
	"\x8c\x81,\xcc,\x0c\x0c,\x8c\x0c\x0c\x82\xbcZ\x0c\x0c" +
	"\x81\x1c\xcc\x8c\x81\"L\x8c\xfc%\x95\x05\xa9\x8c<\x0c" +
	"L\x8c<\x0c\x8c\x80\x00\x00\x00\xff\xff\x9f[%\x81"

func init() {
	schemas.Register(schema_fc8938b535319bfe,
		0x9c032508b61d1d09)
}