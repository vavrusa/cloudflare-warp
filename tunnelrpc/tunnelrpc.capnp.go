// Code generated by capnpc-go. DO NOT EDIT.

package tunnelrpc

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Authentication struct{ capnp.Struct }

// Authentication_TypeID is the unique identifier for the type Authentication.
const Authentication_TypeID = 0xc082ef6e0d42ed1d

func NewAuthentication(s *capnp.Segment) (Authentication, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Authentication{st}, err
}

func NewRootAuthentication(s *capnp.Segment) (Authentication, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Authentication{st}, err
}

func ReadRootAuthentication(msg *capnp.Message) (Authentication, error) {
	root, err := msg.RootPtr()
	return Authentication{root.Struct()}, err
}

func (s Authentication) String() string {
	str, _ := text.Marshal(0xc082ef6e0d42ed1d, s.Struct)
	return str
}

func (s Authentication) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Authentication) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Authentication) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Authentication) SetKey(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Authentication) Email() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Authentication) HasEmail() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Authentication) EmailBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Authentication) SetEmail(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Authentication) OriginCAKey() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Authentication) HasOriginCAKey() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Authentication) OriginCAKeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Authentication) SetOriginCAKey(v string) error {
	return s.Struct.SetText(2, v)
}

// Authentication_List is a list of Authentication.
type Authentication_List struct{ capnp.List }

// NewAuthentication creates a new list of Authentication.
func NewAuthentication_List(s *capnp.Segment, sz int32) (Authentication_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Authentication_List{l}, err
}

func (s Authentication_List) At(i int) Authentication { return Authentication{s.List.Struct(i)} }

func (s Authentication_List) Set(i int, v Authentication) error { return s.List.SetStruct(i, v.Struct) }

// Authentication_Promise is a wrapper for a Authentication promised by a client call.
type Authentication_Promise struct{ *capnp.Pipeline }

func (p Authentication_Promise) Struct() (Authentication, error) {
	s, err := p.Pipeline.Struct()
	return Authentication{s}, err
}

type TunnelRegistration struct{ capnp.Struct }

// TunnelRegistration_TypeID is the unique identifier for the type TunnelRegistration.
const TunnelRegistration_TypeID = 0xf41a0f001ad49e46

func NewTunnelRegistration(s *capnp.Segment) (TunnelRegistration, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return TunnelRegistration{st}, err
}

func NewRootTunnelRegistration(s *capnp.Segment) (TunnelRegistration, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return TunnelRegistration{st}, err
}

func ReadRootTunnelRegistration(msg *capnp.Message) (TunnelRegistration, error) {
	root, err := msg.RootPtr()
	return TunnelRegistration{root.Struct()}, err
}

func (s TunnelRegistration) String() string {
	str, _ := text.Marshal(0xf41a0f001ad49e46, s.Struct)
	return str
}

func (s TunnelRegistration) Err() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s TunnelRegistration) HasErr() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelRegistration) ErrBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s TunnelRegistration) SetErr(v string) error {
	return s.Struct.SetText(0, v)
}

func (s TunnelRegistration) Url() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s TunnelRegistration) HasUrl() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TunnelRegistration) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s TunnelRegistration) SetUrl(v string) error {
	return s.Struct.SetText(1, v)
}

func (s TunnelRegistration) LogLines() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s TunnelRegistration) HasLogLines() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s TunnelRegistration) SetLogLines(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLogLines sets the logLines field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s TunnelRegistration) NewLogLines(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s TunnelRegistration) PermanentFailure() bool {
	return s.Struct.Bit(0)
}

func (s TunnelRegistration) SetPermanentFailure(v bool) {
	s.Struct.SetBit(0, v)
}

// TunnelRegistration_List is a list of TunnelRegistration.
type TunnelRegistration_List struct{ capnp.List }

// NewTunnelRegistration creates a new list of TunnelRegistration.
func NewTunnelRegistration_List(s *capnp.Segment, sz int32) (TunnelRegistration_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return TunnelRegistration_List{l}, err
}

func (s TunnelRegistration_List) At(i int) TunnelRegistration {
	return TunnelRegistration{s.List.Struct(i)}
}

func (s TunnelRegistration_List) Set(i int, v TunnelRegistration) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelRegistration_Promise is a wrapper for a TunnelRegistration promised by a client call.
type TunnelRegistration_Promise struct{ *capnp.Pipeline }

func (p TunnelRegistration_Promise) Struct() (TunnelRegistration, error) {
	s, err := p.Pipeline.Struct()
	return TunnelRegistration{s}, err
}

type RegistrationOptions struct{ capnp.Struct }

// RegistrationOptions_TypeID is the unique identifier for the type RegistrationOptions.
const RegistrationOptions_TypeID = 0xc793e50592935b4a

func NewRegistrationOptions(s *capnp.Segment) (RegistrationOptions, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return RegistrationOptions{st}, err
}

func NewRootRegistrationOptions(s *capnp.Segment) (RegistrationOptions, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return RegistrationOptions{st}, err
}

func ReadRootRegistrationOptions(msg *capnp.Message) (RegistrationOptions, error) {
	root, err := msg.RootPtr()
	return RegistrationOptions{root.Struct()}, err
}

func (s RegistrationOptions) String() string {
	str, _ := text.Marshal(0xc793e50592935b4a, s.Struct)
	return str
}

func (s RegistrationOptions) ClientId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s RegistrationOptions) HasClientId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) ClientIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetClientId(v string) error {
	return s.Struct.SetText(0, v)
}

func (s RegistrationOptions) Version() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s RegistrationOptions) HasVersion() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) VersionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetVersion(v string) error {
	return s.Struct.SetText(1, v)
}

func (s RegistrationOptions) Os() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s RegistrationOptions) HasOs() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) OsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetOs(v string) error {
	return s.Struct.SetText(2, v)
}

func (s RegistrationOptions) ExistingTunnelPolicy() ExistingTunnelPolicy {
	return ExistingTunnelPolicy(s.Struct.Uint16(0))
}

func (s RegistrationOptions) SetExistingTunnelPolicy(v ExistingTunnelPolicy) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s RegistrationOptions) PoolName() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s RegistrationOptions) HasPoolName() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) PoolNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetPoolName(v string) error {
	return s.Struct.SetText(3, v)
}

func (s RegistrationOptions) Tags() (Tag_List, error) {
	p, err := s.Struct.Ptr(4)
	return Tag_List{List: p.List()}, err
}

func (s RegistrationOptions) HasTags() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) SetTags(v Tag_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewTags sets the tags field to a newly
// allocated Tag_List, preferring placement in s's segment.
func (s RegistrationOptions) NewTags(n int32) (Tag_List, error) {
	l, err := NewTag_List(s.Struct.Segment(), n)
	if err != nil {
		return Tag_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s RegistrationOptions) ConnectionId() uint8 {
	return s.Struct.Uint8(2)
}

func (s RegistrationOptions) SetConnectionId(v uint8) {
	s.Struct.SetUint8(2, v)
}

func (s RegistrationOptions) OriginLocalIp() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s RegistrationOptions) HasOriginLocalIp() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) OriginLocalIpBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetOriginLocalIp(v string) error {
	return s.Struct.SetText(5, v)
}

func (s RegistrationOptions) IsAutoupdated() bool {
	return s.Struct.Bit(24)
}

func (s RegistrationOptions) SetIsAutoupdated(v bool) {
	s.Struct.SetBit(24, v)
}

// RegistrationOptions_List is a list of RegistrationOptions.
type RegistrationOptions_List struct{ capnp.List }

// NewRegistrationOptions creates a new list of RegistrationOptions.
func NewRegistrationOptions_List(s *capnp.Segment, sz int32) (RegistrationOptions_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	return RegistrationOptions_List{l}, err
}

func (s RegistrationOptions_List) At(i int) RegistrationOptions {
	return RegistrationOptions{s.List.Struct(i)}
}

func (s RegistrationOptions_List) Set(i int, v RegistrationOptions) error {
	return s.List.SetStruct(i, v.Struct)
}

// RegistrationOptions_Promise is a wrapper for a RegistrationOptions promised by a client call.
type RegistrationOptions_Promise struct{ *capnp.Pipeline }

func (p RegistrationOptions_Promise) Struct() (RegistrationOptions, error) {
	s, err := p.Pipeline.Struct()
	return RegistrationOptions{s}, err
}

type Tag struct{ capnp.Struct }

// Tag_TypeID is the unique identifier for the type Tag.
const Tag_TypeID = 0xcbd96442ae3bb01a

func NewTag(s *capnp.Segment) (Tag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Tag{st}, err
}

func NewRootTag(s *capnp.Segment) (Tag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Tag{st}, err
}

func ReadRootTag(msg *capnp.Message) (Tag, error) {
	root, err := msg.RootPtr()
	return Tag{root.Struct()}, err
}

func (s Tag) String() string {
	str, _ := text.Marshal(0xcbd96442ae3bb01a, s.Struct)
	return str
}

func (s Tag) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Tag) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Tag) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Tag) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Tag) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Tag) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Tag) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Tag) SetValue(v string) error {
	return s.Struct.SetText(1, v)
}

// Tag_List is a list of Tag.
type Tag_List struct{ capnp.List }

// NewTag creates a new list of Tag.
func NewTag_List(s *capnp.Segment, sz int32) (Tag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Tag_List{l}, err
}

func (s Tag_List) At(i int) Tag { return Tag{s.List.Struct(i)} }

func (s Tag_List) Set(i int, v Tag) error { return s.List.SetStruct(i, v.Struct) }

// Tag_Promise is a wrapper for a Tag promised by a client call.
type Tag_Promise struct{ *capnp.Pipeline }

func (p Tag_Promise) Struct() (Tag, error) {
	s, err := p.Pipeline.Struct()
	return Tag{s}, err
}

type ExistingTunnelPolicy uint16

// ExistingTunnelPolicy_TypeID is the unique identifier for the type ExistingTunnelPolicy.
const ExistingTunnelPolicy_TypeID = 0x84cb9536a2cf6d3c

// Values of ExistingTunnelPolicy.
const (
	ExistingTunnelPolicy_ignore     ExistingTunnelPolicy = 0
	ExistingTunnelPolicy_disconnect ExistingTunnelPolicy = 1
	ExistingTunnelPolicy_balance    ExistingTunnelPolicy = 2
)

// String returns the enum's constant name.
func (c ExistingTunnelPolicy) String() string {
	switch c {
	case ExistingTunnelPolicy_ignore:
		return "ignore"
	case ExistingTunnelPolicy_disconnect:
		return "disconnect"
	case ExistingTunnelPolicy_balance:
		return "balance"

	default:
		return ""
	}
}

// ExistingTunnelPolicyFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ExistingTunnelPolicyFromString(c string) ExistingTunnelPolicy {
	switch c {
	case "ignore":
		return ExistingTunnelPolicy_ignore
	case "disconnect":
		return ExistingTunnelPolicy_disconnect
	case "balance":
		return ExistingTunnelPolicy_balance

	default:
		return 0
	}
}

type ExistingTunnelPolicy_List struct{ capnp.List }

func NewExistingTunnelPolicy_List(s *capnp.Segment, sz int32) (ExistingTunnelPolicy_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ExistingTunnelPolicy_List{l.List}, err
}

func (l ExistingTunnelPolicy_List) At(i int) ExistingTunnelPolicy {
	ul := capnp.UInt16List{List: l.List}
	return ExistingTunnelPolicy(ul.At(i))
}

func (l ExistingTunnelPolicy_List) Set(i int, v ExistingTunnelPolicy) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type ServerInfo struct{ capnp.Struct }

// ServerInfo_TypeID is the unique identifier for the type ServerInfo.
const ServerInfo_TypeID = 0xf2c68e2547ec3866

func NewServerInfo(s *capnp.Segment) (ServerInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ServerInfo{st}, err
}

func NewRootServerInfo(s *capnp.Segment) (ServerInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ServerInfo{st}, err
}

func ReadRootServerInfo(msg *capnp.Message) (ServerInfo, error) {
	root, err := msg.RootPtr()
	return ServerInfo{root.Struct()}, err
}

func (s ServerInfo) String() string {
	str, _ := text.Marshal(0xf2c68e2547ec3866, s.Struct)
	return str
}

func (s ServerInfo) LocationName() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ServerInfo) HasLocationName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ServerInfo) LocationNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ServerInfo) SetLocationName(v string) error {
	return s.Struct.SetText(0, v)
}

// ServerInfo_List is a list of ServerInfo.
type ServerInfo_List struct{ capnp.List }

// NewServerInfo creates a new list of ServerInfo.
func NewServerInfo_List(s *capnp.Segment, sz int32) (ServerInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ServerInfo_List{l}, err
}

func (s ServerInfo_List) At(i int) ServerInfo { return ServerInfo{s.List.Struct(i)} }

func (s ServerInfo_List) Set(i int, v ServerInfo) error { return s.List.SetStruct(i, v.Struct) }

// ServerInfo_Promise is a wrapper for a ServerInfo promised by a client call.
type ServerInfo_Promise struct{ *capnp.Pipeline }

func (p ServerInfo_Promise) Struct() (ServerInfo, error) {
	s, err := p.Pipeline.Struct()
	return ServerInfo{s}, err
}

type TunnelServer struct{ Client capnp.Client }

// TunnelServer_TypeID is the unique identifier for the type TunnelServer.
const TunnelServer_TypeID = 0xea58385c65416035

func (c TunnelServer) RegisterTunnel(ctx context.Context, params func(TunnelServer_registerTunnel_Params) error, opts ...capnp.CallOption) TunnelServer_registerTunnel_Results_Promise {
	if c.Client == nil {
		return TunnelServer_registerTunnel_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      0,
			InterfaceName: "tunnelrpc/tunnelrpc.capnp:TunnelServer",
			MethodName:    "registerTunnel",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(TunnelServer_registerTunnel_Params{Struct: s}) }
	}
	return TunnelServer_registerTunnel_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c TunnelServer) GetServerInfo(ctx context.Context, params func(TunnelServer_getServerInfo_Params) error, opts ...capnp.CallOption) TunnelServer_getServerInfo_Results_Promise {
	if c.Client == nil {
		return TunnelServer_getServerInfo_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      1,
			InterfaceName: "tunnelrpc/tunnelrpc.capnp:TunnelServer",
			MethodName:    "getServerInfo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(TunnelServer_getServerInfo_Params{Struct: s}) }
	}
	return TunnelServer_getServerInfo_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type TunnelServer_Server interface {
	RegisterTunnel(TunnelServer_registerTunnel) error

	GetServerInfo(TunnelServer_getServerInfo) error
}

func TunnelServer_ServerToClient(s TunnelServer_Server) TunnelServer {
	c, _ := s.(server.Closer)
	return TunnelServer{Client: server.New(TunnelServer_Methods(nil, s), c)}
}

func TunnelServer_Methods(methods []server.Method, s TunnelServer_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      0,
			InterfaceName: "tunnelrpc/tunnelrpc.capnp:TunnelServer",
			MethodName:    "registerTunnel",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := TunnelServer_registerTunnel{c, opts, TunnelServer_registerTunnel_Params{Struct: p}, TunnelServer_registerTunnel_Results{Struct: r}}
			return s.RegisterTunnel(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      1,
			InterfaceName: "tunnelrpc/tunnelrpc.capnp:TunnelServer",
			MethodName:    "getServerInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := TunnelServer_getServerInfo{c, opts, TunnelServer_getServerInfo_Params{Struct: p}, TunnelServer_getServerInfo_Results{Struct: r}}
			return s.GetServerInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// TunnelServer_registerTunnel holds the arguments for a server call to TunnelServer.registerTunnel.
type TunnelServer_registerTunnel struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  TunnelServer_registerTunnel_Params
	Results TunnelServer_registerTunnel_Results
}

// TunnelServer_getServerInfo holds the arguments for a server call to TunnelServer.getServerInfo.
type TunnelServer_getServerInfo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  TunnelServer_getServerInfo_Params
	Results TunnelServer_getServerInfo_Results
}

type TunnelServer_registerTunnel_Params struct{ capnp.Struct }

// TunnelServer_registerTunnel_Params_TypeID is the unique identifier for the type TunnelServer_registerTunnel_Params.
const TunnelServer_registerTunnel_Params_TypeID = 0xb70431c0dc014915

func NewTunnelServer_registerTunnel_Params(s *capnp.Segment) (TunnelServer_registerTunnel_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return TunnelServer_registerTunnel_Params{st}, err
}

func NewRootTunnelServer_registerTunnel_Params(s *capnp.Segment) (TunnelServer_registerTunnel_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return TunnelServer_registerTunnel_Params{st}, err
}

func ReadRootTunnelServer_registerTunnel_Params(msg *capnp.Message) (TunnelServer_registerTunnel_Params, error) {
	root, err := msg.RootPtr()
	return TunnelServer_registerTunnel_Params{root.Struct()}, err
}

func (s TunnelServer_registerTunnel_Params) String() string {
	str, _ := text.Marshal(0xb70431c0dc014915, s.Struct)
	return str
}

func (s TunnelServer_registerTunnel_Params) OriginCert() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s TunnelServer_registerTunnel_Params) HasOriginCert() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Params) SetOriginCert(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s TunnelServer_registerTunnel_Params) Hostname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s TunnelServer_registerTunnel_Params) HasHostname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Params) HostnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s TunnelServer_registerTunnel_Params) SetHostname(v string) error {
	return s.Struct.SetText(1, v)
}

func (s TunnelServer_registerTunnel_Params) Options() (RegistrationOptions, error) {
	p, err := s.Struct.Ptr(2)
	return RegistrationOptions{Struct: p.Struct()}, err
}

func (s TunnelServer_registerTunnel_Params) HasOptions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Params) SetOptions(v RegistrationOptions) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewOptions sets the options field to a newly
// allocated RegistrationOptions struct, preferring placement in s's segment.
func (s TunnelServer_registerTunnel_Params) NewOptions() (RegistrationOptions, error) {
	ss, err := NewRegistrationOptions(s.Struct.Segment())
	if err != nil {
		return RegistrationOptions{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// TunnelServer_registerTunnel_Params_List is a list of TunnelServer_registerTunnel_Params.
type TunnelServer_registerTunnel_Params_List struct{ capnp.List }

// NewTunnelServer_registerTunnel_Params creates a new list of TunnelServer_registerTunnel_Params.
func NewTunnelServer_registerTunnel_Params_List(s *capnp.Segment, sz int32) (TunnelServer_registerTunnel_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return TunnelServer_registerTunnel_Params_List{l}, err
}

func (s TunnelServer_registerTunnel_Params_List) At(i int) TunnelServer_registerTunnel_Params {
	return TunnelServer_registerTunnel_Params{s.List.Struct(i)}
}

func (s TunnelServer_registerTunnel_Params_List) Set(i int, v TunnelServer_registerTunnel_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_registerTunnel_Params_Promise is a wrapper for a TunnelServer_registerTunnel_Params promised by a client call.
type TunnelServer_registerTunnel_Params_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_registerTunnel_Params_Promise) Struct() (TunnelServer_registerTunnel_Params, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_registerTunnel_Params{s}, err
}

func (p TunnelServer_registerTunnel_Params_Promise) Options() RegistrationOptions_Promise {
	return RegistrationOptions_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type TunnelServer_registerTunnel_Results struct{ capnp.Struct }

// TunnelServer_registerTunnel_Results_TypeID is the unique identifier for the type TunnelServer_registerTunnel_Results.
const TunnelServer_registerTunnel_Results_TypeID = 0xf2c122394f447e8e

func NewTunnelServer_registerTunnel_Results(s *capnp.Segment) (TunnelServer_registerTunnel_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_registerTunnel_Results{st}, err
}

func NewRootTunnelServer_registerTunnel_Results(s *capnp.Segment) (TunnelServer_registerTunnel_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_registerTunnel_Results{st}, err
}

func ReadRootTunnelServer_registerTunnel_Results(msg *capnp.Message) (TunnelServer_registerTunnel_Results, error) {
	root, err := msg.RootPtr()
	return TunnelServer_registerTunnel_Results{root.Struct()}, err
}

func (s TunnelServer_registerTunnel_Results) String() string {
	str, _ := text.Marshal(0xf2c122394f447e8e, s.Struct)
	return str
}

func (s TunnelServer_registerTunnel_Results) Result() (TunnelRegistration, error) {
	p, err := s.Struct.Ptr(0)
	return TunnelRegistration{Struct: p.Struct()}, err
}

func (s TunnelServer_registerTunnel_Results) HasResult() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Results) SetResult(v TunnelRegistration) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResult sets the result field to a newly
// allocated TunnelRegistration struct, preferring placement in s's segment.
func (s TunnelServer_registerTunnel_Results) NewResult() (TunnelRegistration, error) {
	ss, err := NewTunnelRegistration(s.Struct.Segment())
	if err != nil {
		return TunnelRegistration{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// TunnelServer_registerTunnel_Results_List is a list of TunnelServer_registerTunnel_Results.
type TunnelServer_registerTunnel_Results_List struct{ capnp.List }

// NewTunnelServer_registerTunnel_Results creates a new list of TunnelServer_registerTunnel_Results.
func NewTunnelServer_registerTunnel_Results_List(s *capnp.Segment, sz int32) (TunnelServer_registerTunnel_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TunnelServer_registerTunnel_Results_List{l}, err
}

func (s TunnelServer_registerTunnel_Results_List) At(i int) TunnelServer_registerTunnel_Results {
	return TunnelServer_registerTunnel_Results{s.List.Struct(i)}
}

func (s TunnelServer_registerTunnel_Results_List) Set(i int, v TunnelServer_registerTunnel_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_registerTunnel_Results_Promise is a wrapper for a TunnelServer_registerTunnel_Results promised by a client call.
type TunnelServer_registerTunnel_Results_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_registerTunnel_Results_Promise) Struct() (TunnelServer_registerTunnel_Results, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_registerTunnel_Results{s}, err
}

func (p TunnelServer_registerTunnel_Results_Promise) Result() TunnelRegistration_Promise {
	return TunnelRegistration_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type TunnelServer_getServerInfo_Params struct{ capnp.Struct }

// TunnelServer_getServerInfo_Params_TypeID is the unique identifier for the type TunnelServer_getServerInfo_Params.
const TunnelServer_getServerInfo_Params_TypeID = 0xdc3ed6801961e502

func NewTunnelServer_getServerInfo_Params(s *capnp.Segment) (TunnelServer_getServerInfo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TunnelServer_getServerInfo_Params{st}, err
}

func NewRootTunnelServer_getServerInfo_Params(s *capnp.Segment) (TunnelServer_getServerInfo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TunnelServer_getServerInfo_Params{st}, err
}

func ReadRootTunnelServer_getServerInfo_Params(msg *capnp.Message) (TunnelServer_getServerInfo_Params, error) {
	root, err := msg.RootPtr()
	return TunnelServer_getServerInfo_Params{root.Struct()}, err
}

func (s TunnelServer_getServerInfo_Params) String() string {
	str, _ := text.Marshal(0xdc3ed6801961e502, s.Struct)
	return str
}

// TunnelServer_getServerInfo_Params_List is a list of TunnelServer_getServerInfo_Params.
type TunnelServer_getServerInfo_Params_List struct{ capnp.List }

// NewTunnelServer_getServerInfo_Params creates a new list of TunnelServer_getServerInfo_Params.
func NewTunnelServer_getServerInfo_Params_List(s *capnp.Segment, sz int32) (TunnelServer_getServerInfo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return TunnelServer_getServerInfo_Params_List{l}, err
}

func (s TunnelServer_getServerInfo_Params_List) At(i int) TunnelServer_getServerInfo_Params {
	return TunnelServer_getServerInfo_Params{s.List.Struct(i)}
}

func (s TunnelServer_getServerInfo_Params_List) Set(i int, v TunnelServer_getServerInfo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_getServerInfo_Params_Promise is a wrapper for a TunnelServer_getServerInfo_Params promised by a client call.
type TunnelServer_getServerInfo_Params_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_getServerInfo_Params_Promise) Struct() (TunnelServer_getServerInfo_Params, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_getServerInfo_Params{s}, err
}

type TunnelServer_getServerInfo_Results struct{ capnp.Struct }

// TunnelServer_getServerInfo_Results_TypeID is the unique identifier for the type TunnelServer_getServerInfo_Results.
const TunnelServer_getServerInfo_Results_TypeID = 0xe3e37d096a5b564e

func NewTunnelServer_getServerInfo_Results(s *capnp.Segment) (TunnelServer_getServerInfo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_getServerInfo_Results{st}, err
}

func NewRootTunnelServer_getServerInfo_Results(s *capnp.Segment) (TunnelServer_getServerInfo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_getServerInfo_Results{st}, err
}

func ReadRootTunnelServer_getServerInfo_Results(msg *capnp.Message) (TunnelServer_getServerInfo_Results, error) {
	root, err := msg.RootPtr()
	return TunnelServer_getServerInfo_Results{root.Struct()}, err
}

func (s TunnelServer_getServerInfo_Results) String() string {
	str, _ := text.Marshal(0xe3e37d096a5b564e, s.Struct)
	return str
}

func (s TunnelServer_getServerInfo_Results) Result() (ServerInfo, error) {
	p, err := s.Struct.Ptr(0)
	return ServerInfo{Struct: p.Struct()}, err
}

func (s TunnelServer_getServerInfo_Results) HasResult() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelServer_getServerInfo_Results) SetResult(v ServerInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResult sets the result field to a newly
// allocated ServerInfo struct, preferring placement in s's segment.
func (s TunnelServer_getServerInfo_Results) NewResult() (ServerInfo, error) {
	ss, err := NewServerInfo(s.Struct.Segment())
	if err != nil {
		return ServerInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// TunnelServer_getServerInfo_Results_List is a list of TunnelServer_getServerInfo_Results.
type TunnelServer_getServerInfo_Results_List struct{ capnp.List }

// NewTunnelServer_getServerInfo_Results creates a new list of TunnelServer_getServerInfo_Results.
func NewTunnelServer_getServerInfo_Results_List(s *capnp.Segment, sz int32) (TunnelServer_getServerInfo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TunnelServer_getServerInfo_Results_List{l}, err
}

func (s TunnelServer_getServerInfo_Results_List) At(i int) TunnelServer_getServerInfo_Results {
	return TunnelServer_getServerInfo_Results{s.List.Struct(i)}
}

func (s TunnelServer_getServerInfo_Results_List) Set(i int, v TunnelServer_getServerInfo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_getServerInfo_Results_Promise is a wrapper for a TunnelServer_getServerInfo_Results promised by a client call.
type TunnelServer_getServerInfo_Results_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_getServerInfo_Results_Promise) Struct() (TunnelServer_getServerInfo_Results, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_getServerInfo_Results{s}, err
}

func (p TunnelServer_getServerInfo_Results_Promise) Result() ServerInfo_Promise {
	return ServerInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_db8274f9144abc7e = "x\xda\x9cUM\x88\x14W\x17\xbd\xe7\xbd\xee\xaa\x11z" +
	"\xbe\x9e\xa2z@\x1b\xa4AF>\x95\x8c:Q\x83\x99" +
	"H\xe6'j\xe8\xc9\x8c\xf6\x9b1\xc1\xbf\x80e\xf7\xb3" +
	"\xa7&\xd5UMU\xb5\x89\x82\x9a\x88\x10\x08(1\x92" +
	"E6Y\x08n\x02\xf9\xd9\x85@\x16f\x93,$d" +
	"\x13\x02\x12\xb3\x89H@\x14\xc3` BB\x85W5" +
	"\xd5]\x8eD\x83\xbb\xd7\xa7\xee\xbb\xf7\xdc{\xcf;\xbd" +
	"\xf92\x1bg#\xf9My\"\xb1#\xafE;Z?" +
	"\\~\xee\xc3k\xe7\xc8(\xb3\xe8\xf4\xd7S\xa5\x07\xe1" +
	"\xd9\x9f\x89\xb0\xe5\x14;\x09\xf3\"\xd3\x89\xcc\xf3l/" +
	"!\x1a\xac\xe2\xc6\xd5\x91\xdc\x97d\xfc\x1fDy\xae\x13" +
	"m\xb9\xc2n\x83`~\xc5>'D\xab\xefN\xf6\xbb" +
	"\xf7\xce^%\xa3\x8c^\xaa$\xf0\x00\x9f\x82\xd9RG" +
	"\xd3\xe6*x\xea\xd0\xa5\x0f\xf2\xb7.}G\xa2\x8cl" +
	"\xb4\xa6\xa2\xf39\x1f\xe6\xaa\x9c:\x0e\xe6\"\x10\xa2\xf2" +
	"\x17/|6\xd9\xb8~mY\xee\x98\xde\x84\xb6h\xce" +
	"\xa8{fU{\x93\x10\xb1[\xd6\xaa\xb7\x7fz\xf1F" +
	"B4\xcerE\xfb\x15\x94\x8b\xf6\xbcvha\xc5\xa9" +
	"\x9b7\x97Z\x80\xfa\xf4\x91\x16\xb7\xf0\xa96F\x88\xb6" +
	"\x1d\x99\x90\x87\xb7\xef\xbfMF\x99?4\x8d\xef\xb5Q" +
	"\x98\xbf\xc4E\xaek\xef\x9a#\xbaN\x14]8\xbds" +
	"\xef\xf3k\xbeY\xcc\xa6[\xa5/\xaat\xc3\xbaJw" +
	"l\xfb\x9d\x97\xd7^\xf8vq\x19\xeb8pF\xdf\x00" +
	"\xf3u\x95\xc7<\xa0\x82\xef\xed\xfe\xf8\xc7r\xb1|\x7f" +
	"\xd9<\xe2\xe9\x9d\xd2\x17`^\x8cc\xcf\xeb\xbf\xd1p" +
	"\x14v\\W:~;W\xdf\x94\x1e\xeb\x1b\xebV\xdb" +
	"m\x8f\xeez\xcb\x0eB\xdbm\xee\x8b\xf1\xb1\x9a\xe7\xd8" +
	"\xf5\x135@\x14\xc0\x88\x8c\xd5\xa3D\x801x\x90\x08" +
	"\xcc0&\x89\xc6\xec\xa6\xeb\xf92j\xd8A\xdds]" +
	"I\xbc\x1e\x9e9j9\x96[\x97\xddB\xda\xa3\x85\x92" +
	"\x02s\xd2?.\xfd\x8d\xbel\xdaA(\xfd\x04\x1c\xaa" +
	"YE\xdfj\x05\xa2\xc0sD9\x10\x19\xbb\x0e\x12\x89" +
	"\x9d\x1c\xa2\xc6`\x00%(pf\x8aHLs\x88\xfd" +
	"\x0c\x06c\xa5\x98\xe1\xab\x93D\xa2\xc6!\x0e3D\x9e" +
	"o7m\xf7%I\xdc\x0f\xd1O\x0c\xfd\x84h\xde\x0b" +
	"B\xd7jI\"B\x81\x18\x0a\x843^;\xb4=7" +
	"\xc0@OY\x04\x0c\x10\x1e7\xab\x89N8/\xdd\xd0" +
	"\xae[\xea2Q<\xa6\x1e\xe55Db\x9cCLg" +
	"(W\x9f\xcd\xf4\x91R\x9e9\xda\xebC\x7fC\x9eH" +
	"YUd\xcb\xb2\x9d\xf4W\xda\xcc\x04\xe9\xaf\xf4b\x1e" +
	"\xc7o6\x9e\xaa\x1f\xb3\xdb\xdb\xae\xc4\x1d*\x8eC]" +
	"\x8ew\xd5\x04\xefp\x88?3\x1c\xffP\x13\xfc\x9dC" +
	"\xfc\x95\xe1\xf8\xa0L$\xees\xcc\x82\x01\xbc\x04Nd" +
	"\xfc\xfd\x09\xd1,8\xe6\x0a`0r\xbc\x84\x1c\x91\xb9" +
	"\x02SDs}\x0a/)<\x9f+!Od\x1a\xd8" +
	"@4WP\xf8:\x85k\xac\x04\x8d\xc8\\\x8b\x05\xa2" +
	"\xb9!\x85oV\xb8\x9e/)\x85\x9b\xc3\xf0\x89\xe6\x9e" +
	"Q\xf8v\x85\xf7\xad,\xa1\x8f\xc8\xdc\x16\xe3[\x15>" +
	"\x0e\x86\xa8\xee\xd8\xd2\x0d\xab\x8d\xec:\x8fK?\xb0=" +
	"7\xfd\xcd\xbd\xa0;/\xb9$p$Z\xabyE\xa5" +
	"p\x14{VF@\x91\x10\xb5=\xcf\xd9\xf3\xb0L\x8a" +
	"\xa1\xd5\x0c\xf0?B\x8d\x03\x03=W!(0\x8a\xf5" +
	"_\x0fm*zn\xb5\x01\x8d\x18\xb4\xee\xde\xa6=\xaa" +
	"\xd4-\xa7\xda\xee2\xb1\x83\x89N\xe8u\xdaTiX" +
	"\xa1l\x00\xc4\x80\xccF\xd9\xf2\x8dV\xda\xa3\xfb\xac\xa6" +
	"\xda`_w\x83\xeb7\x10\x89!\x0e\xb19\xb3\xc1a" +
	"\xa5\xb2u\x1cb+CQI\xbd\xab\xa8\xe3\x96\xd3\x91" +
	"\x8fh\xe7I\xcf\xb3)\xc3\xe4Tu\x8fyC5\xcb" +
	"\xd7\xadV\xf0\x94\xb7geP\xec8a r\xdd\x1e" +
	"\xfaG\x89D\x1f\x87(1\x8c\xf92\xe88!\x06z" +
	"\xee\xb7\xec!\xf2\x7f+7\x96TI\xe6\x93'\xea\xfe" +
	"\xe5 uZc\xe4$1c\xbd\x8e\x9e\xcb#5u" +
	"c\xb5O\xcc\x18\xd4\xa3\xd4\x8bh,I;\x8e(\xed" +
	"\x80*q\x0f\xe3\xa8\x01O\xebm\xb3\xb2\x12\xfc\x97\xfe" +
	"SC\x7fr\xf7I\x9d\xa2b\xa6z\xcf\xe4] \x12" +
	"\x05\x0e\xb1\x92!r\xbc%\x9b*\xee\xc9\x08\xe2q\xf6" +
	"\x91\x10NM\xa4\xa8.\xab\xfc\x03\xdd\xfc\x96r\xb8\xc3" +
	"\x1cb>\xa3=\xa9\xc0#\x1c\xc2\xc9\xb8\x87\xad|f" +
	"\x9eC\x9c\xeb\xb9\xc7;\xef\x11\x89s\x1c\xe2}\x06]" +
	"\xfa~JI\xef\xf8=\xcfs\xbc\xe6\xb4\xed\xca@=" +
	"\xc2\xa5w\xa7>\xa9\xd7\xd6\x96~\xcbr\xa5\x8bp\xb7" +
	"e;\x1d_*\xa1$O\xe8\x9f\x00\x00\x00\xff\xff\x0d" +
	"v6\xc6"

func init() {
	schemas.Register(schema_db8274f9144abc7e,
		0x84cb9536a2cf6d3c,
		0xb70431c0dc014915,
		0xc082ef6e0d42ed1d,
		0xc793e50592935b4a,
		0xcbd96442ae3bb01a,
		0xdc3ed6801961e502,
		0xe3e37d096a5b564e,
		0xea58385c65416035,
		0xf2c122394f447e8e,
		0xf2c68e2547ec3866,
		0xf41a0f001ad49e46)
}
