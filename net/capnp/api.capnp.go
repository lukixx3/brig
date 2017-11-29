// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Sync struct{ Client capnp.Client }

// Sync_TypeID is the unique identifier for the type Sync.
const Sync_TypeID = 0xf5692a07c5cf7872

func (c Sync) FetchStore(ctx context.Context, params func(Sync_fetchStore_Params) error, opts ...capnp.CallOption) Sync_fetchStore_Results_Promise {
	if c.Client == nil {
		return Sync_fetchStore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf5692a07c5cf7872,
			MethodID:      0,
			InterfaceName: "api.capnp:Sync",
			MethodName:    "fetchStore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Sync_fetchStore_Params{Struct: s}) }
	}
	return Sync_fetchStore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Sync_Server interface {
	FetchStore(Sync_fetchStore) error
}

func Sync_ServerToClient(s Sync_Server) Sync {
	c, _ := s.(server.Closer)
	return Sync{Client: server.New(Sync_Methods(nil, s), c)}
}

func Sync_Methods(methods []server.Method, s Sync_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf5692a07c5cf7872,
			MethodID:      0,
			InterfaceName: "api.capnp:Sync",
			MethodName:    "fetchStore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Sync_fetchStore{c, opts, Sync_fetchStore_Params{Struct: p}, Sync_fetchStore_Results{Struct: r}}
			return s.FetchStore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Sync_fetchStore holds the arguments for a server call to Sync.fetchStore.
type Sync_fetchStore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Sync_fetchStore_Params
	Results Sync_fetchStore_Results
}

type Sync_fetchStore_Params struct{ capnp.Struct }

// Sync_fetchStore_Params_TypeID is the unique identifier for the type Sync_fetchStore_Params.
const Sync_fetchStore_Params_TypeID = 0xdc63044e67499411

func NewSync_fetchStore_Params(s *capnp.Segment) (Sync_fetchStore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Sync_fetchStore_Params{st}, err
}

func NewRootSync_fetchStore_Params(s *capnp.Segment) (Sync_fetchStore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Sync_fetchStore_Params{st}, err
}

func ReadRootSync_fetchStore_Params(msg *capnp.Message) (Sync_fetchStore_Params, error) {
	root, err := msg.RootPtr()
	return Sync_fetchStore_Params{root.Struct()}, err
}

func (s Sync_fetchStore_Params) String() string {
	str, _ := text.Marshal(0xdc63044e67499411, s.Struct)
	return str
}

// Sync_fetchStore_Params_List is a list of Sync_fetchStore_Params.
type Sync_fetchStore_Params_List struct{ capnp.List }

// NewSync_fetchStore_Params creates a new list of Sync_fetchStore_Params.
func NewSync_fetchStore_Params_List(s *capnp.Segment, sz int32) (Sync_fetchStore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Sync_fetchStore_Params_List{l}, err
}

func (s Sync_fetchStore_Params_List) At(i int) Sync_fetchStore_Params {
	return Sync_fetchStore_Params{s.List.Struct(i)}
}

func (s Sync_fetchStore_Params_List) Set(i int, v Sync_fetchStore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Sync_fetchStore_Params_List) String() string {
	str, _ := text.MarshalList(0xdc63044e67499411, s.List)
	return str
}

// Sync_fetchStore_Params_Promise is a wrapper for a Sync_fetchStore_Params promised by a client call.
type Sync_fetchStore_Params_Promise struct{ *capnp.Pipeline }

func (p Sync_fetchStore_Params_Promise) Struct() (Sync_fetchStore_Params, error) {
	s, err := p.Pipeline.Struct()
	return Sync_fetchStore_Params{s}, err
}

type Sync_fetchStore_Results struct{ capnp.Struct }

// Sync_fetchStore_Results_TypeID is the unique identifier for the type Sync_fetchStore_Results.
const Sync_fetchStore_Results_TypeID = 0xf834409e30e8009c

func NewSync_fetchStore_Results(s *capnp.Segment) (Sync_fetchStore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Sync_fetchStore_Results{st}, err
}

func NewRootSync_fetchStore_Results(s *capnp.Segment) (Sync_fetchStore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Sync_fetchStore_Results{st}, err
}

func ReadRootSync_fetchStore_Results(msg *capnp.Message) (Sync_fetchStore_Results, error) {
	root, err := msg.RootPtr()
	return Sync_fetchStore_Results{root.Struct()}, err
}

func (s Sync_fetchStore_Results) String() string {
	str, _ := text.Marshal(0xf834409e30e8009c, s.Struct)
	return str
}

func (s Sync_fetchStore_Results) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Sync_fetchStore_Results) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Sync_fetchStore_Results) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Sync_fetchStore_Results_List is a list of Sync_fetchStore_Results.
type Sync_fetchStore_Results_List struct{ capnp.List }

// NewSync_fetchStore_Results creates a new list of Sync_fetchStore_Results.
func NewSync_fetchStore_Results_List(s *capnp.Segment, sz int32) (Sync_fetchStore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Sync_fetchStore_Results_List{l}, err
}

func (s Sync_fetchStore_Results_List) At(i int) Sync_fetchStore_Results {
	return Sync_fetchStore_Results{s.List.Struct(i)}
}

func (s Sync_fetchStore_Results_List) Set(i int, v Sync_fetchStore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Sync_fetchStore_Results_List) String() string {
	str, _ := text.MarshalList(0xf834409e30e8009c, s.List)
	return str
}

// Sync_fetchStore_Results_Promise is a wrapper for a Sync_fetchStore_Results promised by a client call.
type Sync_fetchStore_Results_Promise struct{ *capnp.Pipeline }

func (p Sync_fetchStore_Results_Promise) Struct() (Sync_fetchStore_Results, error) {
	s, err := p.Pipeline.Struct()
	return Sync_fetchStore_Results{s}, err
}

type Meta struct{ Client capnp.Client }

// Meta_TypeID is the unique identifier for the type Meta.
const Meta_TypeID = 0xb02d2ba0578cc7ff

func (c Meta) Ping(ctx context.Context, params func(Meta_ping_Params) error, opts ...capnp.CallOption) Meta_ping_Results_Promise {
	if c.Client == nil {
		return Meta_ping_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xb02d2ba0578cc7ff,
			MethodID:      0,
			InterfaceName: "api.capnp:Meta",
			MethodName:    "ping",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Meta_ping_Params{Struct: s}) }
	}
	return Meta_ping_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Meta_Server interface {
	Ping(Meta_ping) error
}

func Meta_ServerToClient(s Meta_Server) Meta {
	c, _ := s.(server.Closer)
	return Meta{Client: server.New(Meta_Methods(nil, s), c)}
}

func Meta_Methods(methods []server.Method, s Meta_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb02d2ba0578cc7ff,
			MethodID:      0,
			InterfaceName: "api.capnp:Meta",
			MethodName:    "ping",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Meta_ping{c, opts, Meta_ping_Params{Struct: p}, Meta_ping_Results{Struct: r}}
			return s.Ping(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Meta_ping holds the arguments for a server call to Meta.ping.
type Meta_ping struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Meta_ping_Params
	Results Meta_ping_Results
}

type Meta_ping_Params struct{ capnp.Struct }

// Meta_ping_Params_TypeID is the unique identifier for the type Meta_ping_Params.
const Meta_ping_Params_TypeID = 0xe1a9fd466eca248c

func NewMeta_ping_Params(s *capnp.Segment) (Meta_ping_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Meta_ping_Params{st}, err
}

func NewRootMeta_ping_Params(s *capnp.Segment) (Meta_ping_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Meta_ping_Params{st}, err
}

func ReadRootMeta_ping_Params(msg *capnp.Message) (Meta_ping_Params, error) {
	root, err := msg.RootPtr()
	return Meta_ping_Params{root.Struct()}, err
}

func (s Meta_ping_Params) String() string {
	str, _ := text.Marshal(0xe1a9fd466eca248c, s.Struct)
	return str
}

// Meta_ping_Params_List is a list of Meta_ping_Params.
type Meta_ping_Params_List struct{ capnp.List }

// NewMeta_ping_Params creates a new list of Meta_ping_Params.
func NewMeta_ping_Params_List(s *capnp.Segment, sz int32) (Meta_ping_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Meta_ping_Params_List{l}, err
}

func (s Meta_ping_Params_List) At(i int) Meta_ping_Params { return Meta_ping_Params{s.List.Struct(i)} }

func (s Meta_ping_Params_List) Set(i int, v Meta_ping_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Meta_ping_Params_List) String() string {
	str, _ := text.MarshalList(0xe1a9fd466eca248c, s.List)
	return str
}

// Meta_ping_Params_Promise is a wrapper for a Meta_ping_Params promised by a client call.
type Meta_ping_Params_Promise struct{ *capnp.Pipeline }

func (p Meta_ping_Params_Promise) Struct() (Meta_ping_Params, error) {
	s, err := p.Pipeline.Struct()
	return Meta_ping_Params{s}, err
}

type Meta_ping_Results struct{ capnp.Struct }

// Meta_ping_Results_TypeID is the unique identifier for the type Meta_ping_Results.
const Meta_ping_Results_TypeID = 0x9a90fde15285e327

func NewMeta_ping_Results(s *capnp.Segment) (Meta_ping_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Meta_ping_Results{st}, err
}

func NewRootMeta_ping_Results(s *capnp.Segment) (Meta_ping_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Meta_ping_Results{st}, err
}

func ReadRootMeta_ping_Results(msg *capnp.Message) (Meta_ping_Results, error) {
	root, err := msg.RootPtr()
	return Meta_ping_Results{root.Struct()}, err
}

func (s Meta_ping_Results) String() string {
	str, _ := text.Marshal(0x9a90fde15285e327, s.Struct)
	return str
}

func (s Meta_ping_Results) Reply() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Meta_ping_Results) HasReply() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Meta_ping_Results) ReplyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Meta_ping_Results) SetReply(v string) error {
	return s.Struct.SetText(0, v)
}

// Meta_ping_Results_List is a list of Meta_ping_Results.
type Meta_ping_Results_List struct{ capnp.List }

// NewMeta_ping_Results creates a new list of Meta_ping_Results.
func NewMeta_ping_Results_List(s *capnp.Segment, sz int32) (Meta_ping_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Meta_ping_Results_List{l}, err
}

func (s Meta_ping_Results_List) At(i int) Meta_ping_Results {
	return Meta_ping_Results{s.List.Struct(i)}
}

func (s Meta_ping_Results_List) Set(i int, v Meta_ping_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Meta_ping_Results_List) String() string {
	str, _ := text.MarshalList(0x9a90fde15285e327, s.List)
	return str
}

// Meta_ping_Results_Promise is a wrapper for a Meta_ping_Results promised by a client call.
type Meta_ping_Results_Promise struct{ *capnp.Pipeline }

func (p Meta_ping_Results_Promise) Struct() (Meta_ping_Results, error) {
	s, err := p.Pipeline.Struct()
	return Meta_ping_Results{s}, err
}

type API struct{ Client capnp.Client }

// API_TypeID is the unique identifier for the type API.
const API_TypeID = 0xb74958502f92fefd

func (c API) Version(ctx context.Context, params func(API_version_Params) error, opts ...capnp.CallOption) API_version_Results_Promise {
	if c.Client == nil {
		return API_version_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xb74958502f92fefd,
			MethodID:      0,
			InterfaceName: "api.capnp:API",
			MethodName:    "version",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(API_version_Params{Struct: s}) }
	}
	return API_version_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c API) FetchStore(ctx context.Context, params func(Sync_fetchStore_Params) error, opts ...capnp.CallOption) Sync_fetchStore_Results_Promise {
	if c.Client == nil {
		return Sync_fetchStore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf5692a07c5cf7872,
			MethodID:      0,
			InterfaceName: "api.capnp:Sync",
			MethodName:    "fetchStore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Sync_fetchStore_Params{Struct: s}) }
	}
	return Sync_fetchStore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c API) Ping(ctx context.Context, params func(Meta_ping_Params) error, opts ...capnp.CallOption) Meta_ping_Results_Promise {
	if c.Client == nil {
		return Meta_ping_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xb02d2ba0578cc7ff,
			MethodID:      0,
			InterfaceName: "api.capnp:Meta",
			MethodName:    "ping",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Meta_ping_Params{Struct: s}) }
	}
	return Meta_ping_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type API_Server interface {
	Version(API_version) error

	FetchStore(Sync_fetchStore) error

	Ping(Meta_ping) error
}

func API_ServerToClient(s API_Server) API {
	c, _ := s.(server.Closer)
	return API{Client: server.New(API_Methods(nil, s), c)}
}

func API_Methods(methods []server.Method, s API_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb74958502f92fefd,
			MethodID:      0,
			InterfaceName: "api.capnp:API",
			MethodName:    "version",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := API_version{c, opts, API_version_Params{Struct: p}, API_version_Results{Struct: r}}
			return s.Version(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf5692a07c5cf7872,
			MethodID:      0,
			InterfaceName: "api.capnp:Sync",
			MethodName:    "fetchStore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Sync_fetchStore{c, opts, Sync_fetchStore_Params{Struct: p}, Sync_fetchStore_Results{Struct: r}}
			return s.FetchStore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb02d2ba0578cc7ff,
			MethodID:      0,
			InterfaceName: "api.capnp:Meta",
			MethodName:    "ping",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Meta_ping{c, opts, Meta_ping_Params{Struct: p}, Meta_ping_Results{Struct: r}}
			return s.Ping(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// API_version holds the arguments for a server call to API.version.
type API_version struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  API_version_Params
	Results API_version_Results
}

type API_version_Params struct{ capnp.Struct }

// API_version_Params_TypeID is the unique identifier for the type API_version_Params.
const API_version_Params_TypeID = 0xfbab528dd0716804

func NewAPI_version_Params(s *capnp.Segment) (API_version_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return API_version_Params{st}, err
}

func NewRootAPI_version_Params(s *capnp.Segment) (API_version_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return API_version_Params{st}, err
}

func ReadRootAPI_version_Params(msg *capnp.Message) (API_version_Params, error) {
	root, err := msg.RootPtr()
	return API_version_Params{root.Struct()}, err
}

func (s API_version_Params) String() string {
	str, _ := text.Marshal(0xfbab528dd0716804, s.Struct)
	return str
}

// API_version_Params_List is a list of API_version_Params.
type API_version_Params_List struct{ capnp.List }

// NewAPI_version_Params creates a new list of API_version_Params.
func NewAPI_version_Params_List(s *capnp.Segment, sz int32) (API_version_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return API_version_Params_List{l}, err
}

func (s API_version_Params_List) At(i int) API_version_Params {
	return API_version_Params{s.List.Struct(i)}
}

func (s API_version_Params_List) Set(i int, v API_version_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s API_version_Params_List) String() string {
	str, _ := text.MarshalList(0xfbab528dd0716804, s.List)
	return str
}

// API_version_Params_Promise is a wrapper for a API_version_Params promised by a client call.
type API_version_Params_Promise struct{ *capnp.Pipeline }

func (p API_version_Params_Promise) Struct() (API_version_Params, error) {
	s, err := p.Pipeline.Struct()
	return API_version_Params{s}, err
}

type API_version_Results struct{ capnp.Struct }

// API_version_Results_TypeID is the unique identifier for the type API_version_Results.
const API_version_Results_TypeID = 0xebdd19e3dba3370b

func NewAPI_version_Results(s *capnp.Segment) (API_version_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return API_version_Results{st}, err
}

func NewRootAPI_version_Results(s *capnp.Segment) (API_version_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return API_version_Results{st}, err
}

func ReadRootAPI_version_Results(msg *capnp.Message) (API_version_Results, error) {
	root, err := msg.RootPtr()
	return API_version_Results{root.Struct()}, err
}

func (s API_version_Results) String() string {
	str, _ := text.Marshal(0xebdd19e3dba3370b, s.Struct)
	return str
}

func (s API_version_Results) Version() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s API_version_Results) SetVersion(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// API_version_Results_List is a list of API_version_Results.
type API_version_Results_List struct{ capnp.List }

// NewAPI_version_Results creates a new list of API_version_Results.
func NewAPI_version_Results_List(s *capnp.Segment, sz int32) (API_version_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return API_version_Results_List{l}, err
}

func (s API_version_Results_List) At(i int) API_version_Results {
	return API_version_Results{s.List.Struct(i)}
}

func (s API_version_Results_List) Set(i int, v API_version_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s API_version_Results_List) String() string {
	str, _ := text.MarshalList(0xebdd19e3dba3370b, s.List)
	return str
}

// API_version_Results_Promise is a wrapper for a API_version_Results promised by a client call.
type API_version_Results_Promise struct{ *capnp.Pipeline }

func (p API_version_Results_Promise) Struct() (API_version_Results, error) {
	s, err := p.Pipeline.Struct()
	return API_version_Results{s}, err
}

const schema_9bcb07fb35756ee6 = "x\xdal\x92\xbfk\x13a\x18\xc7\xbf\xcfso\xbcZ" +
	"\xdb\x86\xd7sq\xea\x12)F\x8c\xd1V\x84,M\x0b" +
	"\x1an\xa8\xdc\xa5\x83Vp8\xe2\xd9\x1e$\xd7\xeb\xdd" +
	"ELA\x8aC\xb7\x0a\"\xbah\x11A\x1c\xc4\xc5\xcd" +
	"\xbf@\x04\x1dt\xf6\xc7P\x0b\"\xba\x09\x82P\xe2\xc9" +
	"\xa5\xbd\\\x92\xba\xbd\xf0\xfc\xf8~\xbe\xdf\xe7-f\xa9" +
	"\xcc\xa73k\x020\xa72\x07\xa2\x89\xed\xf5\xeaV\xfb" +
	"\xeeC\xc8,\x01\x19R\x01\xed\x18\x7f\x05i\xc7y\x1a" +
	"\x14Eo6.=9q\xf2%\xe4\xb0\x12}s\x9b" +
	"gw\xd4w\x8f\x00\xd2t^\xd5\xe68n\xd7\xb9\xa2" +
	"\xad\xc4\xaf\xa8\xfd\xf7\xde)\xe3\xb2\xfej_\xf3\x02\xfb" +
	"\xda\xd5N\xf3\x02W\xb4\xdb<\x01D\xf2\xbe\xbexQ" +
	"\xd4>\xef*\x0b\x15\x98l\xf10AD\x1b\xb9\xb7\xee" +
	"\x85\xf6\xf3\xad\xb4\xa2Y\xfc\x11\":t\xee\xe9\xa7\xed" +
	"\xa3_~\xc2\x1c\xa3\xa4r\x9e\x7fuhbX\xff\xe6" +
	"\xfb\xd7j\xde\xf9\xbdO\xdf\xe1U\xad\xd1\xd1w\xb8\xa2" +
	"=\x88_\xed\xcd\xef\xc5\xc7\xe5\xa9?\xa9\xef\xc9[|" +
	"\x98@\xdazg\x97XZ\xf9p\xa7\xfab\x07r," +
	"\x91z\xc6?P\x8c,\xcf)\xd4,\xcfe\xaf4g" +
	"\x87V\xc1s\xdc\xc5\\\xd5\x0e\x9aj=\x0cL\xa1\x08" +
	"@\x10 G\xcf\x00\xe6\x90B\xe6\x11\xa6q\xdf\xf6\xea" +
	"-\x1a\x01\xd3\x08\xa8\xbb\x02\xd3\xbb;\x0c\"S(\x19" +
	"\xa0k\x9d\x92\xbbH\x99\x07\xcb\x8c\x9a\x8de\xcadP" +
	"\xcf\xf4\xb8W\x9a1\xf4t8A\xa6$')g;" +
	"\xc3k7l?p\x96\xdd2\x99C\xd4\x13\x13\x90\xde" +
	"\x17\xe8.V\xbc\xd2|\xcb\xad\x15\xae\xdbami>" +
	"\\\xf6\xed\x9ca\xf9V#\xe8\xe9\xe9s\x1fW\x95F" +
	"\xd0[\x9c1\xf4\xc2\x9ej\xaej\x8f\x07\xcd\x81pf" +
	"\xd3p\x12:\x12`\x12\x03\xf1\xc4 \xa9\xc3\xe4\xcf\xd0" +
	"&\xf6\xae'\xaf\x80\xe5A5JX\xa1\xf8v\x7fL" +
	"\xffq\x13_\xab\x1eR\x1fP>\x05\xca^\xb3B\x8b" +
	"F\xc14\xdaC3`\xca\xb0\xb2q$\xff\x02\x00\x00" +
	"\xff\xff\xbe\xb6\xf0v"

func init() {
	schemas.Register(schema_9bcb07fb35756ee6,
		0x9a90fde15285e327,
		0xb02d2ba0578cc7ff,
		0xb74958502f92fefd,
		0xdc63044e67499411,
		0xe1a9fd466eca248c,
		0xebdd19e3dba3370b,
		0xf5692a07c5cf7872,
		0xf834409e30e8009c,
		0xfbab528dd0716804)
}
