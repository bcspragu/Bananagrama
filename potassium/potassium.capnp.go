package potassium

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Server struct{ Client capnp.Client }

func (c Server) Connect(ctx context.Context, params func(ConnectRequest) error, opts ...capnp.CallOption) ConnectResponse_Promise {
	if c.Client == nil {
		return ConnectResponse_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa509b3610d81663d,
			MethodID:      0,
			InterfaceName: "potassium.capnp:Server",
			MethodName:    "connect",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ConnectRequest{Struct: s}) }
	}
	return ConnectResponse_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Server_Server interface {
	Connect(Server_connect) error
}

func Server_ServerToClient(s Server_Server) Server {
	c, _ := s.(server.Closer)
	return Server{Client: server.New(Server_Methods(nil, s), c)}
}

func Server_Methods(methods []server.Method, s Server_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa509b3610d81663d,
			MethodID:      0,
			InterfaceName: "potassium.capnp:Server",
			MethodName:    "connect",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Server_connect{c, opts, ConnectRequest{Struct: p}, ConnectResponse{Struct: r}}
			return s.Connect(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	return methods
}

// Server_connect holds the arguments for a server call to Server.connect.
type Server_connect struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ConnectRequest
	Results ConnectResponse
}

type ConnectRequest struct{ capnp.Struct }

// ConnectRequest_TypeID is the unique identifier for the type ConnectRequest.
const ConnectRequest_TypeID = 0xb26464829e0e9c88

func NewConnectRequest(s *capnp.Segment) (ConnectRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return ConnectRequest{st}, err
}

func NewRootConnectRequest(s *capnp.Segment) (ConnectRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return ConnectRequest{st}, err
}

func ReadRootConnectRequest(msg *capnp.Message) (ConnectRequest, error) {
	root, err := msg.RootPtr()
	return ConnectRequest{root.Struct()}, err
}

func (s ConnectRequest) String() string {
	str, _ := text.Marshal(0xb26464829e0e9c88, s.Struct)
	return str
}

func (s ConnectRequest) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ConnectRequest) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ConnectRequest) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ConnectRequest) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ConnectRequest) Player() Player {
	p, _ := s.Struct.Ptr(1)
	return Player{Client: p.Interface().Client()}
}

func (s ConnectRequest) HasPlayer() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ConnectRequest) SetPlayer(v Player) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// ConnectRequest_List is a list of ConnectRequest.
type ConnectRequest_List struct{ capnp.List }

// NewConnectRequest creates a new list of ConnectRequest.
func NewConnectRequest_List(s *capnp.Segment, sz int32) (ConnectRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return ConnectRequest_List{l}, err
}

func (s ConnectRequest_List) At(i int) ConnectRequest { return ConnectRequest{s.List.Struct(i)} }

func (s ConnectRequest_List) Set(i int, v ConnectRequest) error { return s.List.SetStruct(i, v.Struct) }

// ConnectRequest_Promise is a wrapper for a ConnectRequest promised by a client call.
type ConnectRequest_Promise struct{ *capnp.Pipeline }

func (p ConnectRequest_Promise) Struct() (ConnectRequest, error) {
	s, err := p.Pipeline.Struct()
	return ConnectRequest{s}, err
}

func (p ConnectRequest_Promise) Player() Player {
	return Player{Client: p.Pipeline.GetPipeline(1).Client()}
}

type ConnectResponse struct{ capnp.Struct }

// ConnectResponse_TypeID is the unique identifier for the type ConnectResponse.
const ConnectResponse_TypeID = 0xcf98cd13ab9af903

func NewConnectResponse(s *capnp.Segment) (ConnectResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ConnectResponse{st}, err
}

func NewRootConnectResponse(s *capnp.Segment) (ConnectResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ConnectResponse{st}, err
}

func ReadRootConnectResponse(msg *capnp.Message) (ConnectResponse, error) {
	root, err := msg.RootPtr()
	return ConnectResponse{root.Struct()}, err
}

func (s ConnectResponse) String() string {
	str, _ := text.Marshal(0xcf98cd13ab9af903, s.Struct)
	return str
}

func (s ConnectResponse) Status() ConnectResponse_Status {
	return ConnectResponse_Status(s.Struct.Uint16(0))
}

func (s ConnectResponse) SetStatus(v ConnectResponse_Status) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s ConnectResponse) Game() Game {
	p, _ := s.Struct.Ptr(0)
	return Game{Client: p.Interface().Client()}
}

func (s ConnectResponse) HasGame() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ConnectResponse) SetGame(v Game) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// ConnectResponse_List is a list of ConnectResponse.
type ConnectResponse_List struct{ capnp.List }

// NewConnectResponse creates a new list of ConnectResponse.
func NewConnectResponse_List(s *capnp.Segment, sz int32) (ConnectResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return ConnectResponse_List{l}, err
}

func (s ConnectResponse_List) At(i int) ConnectResponse { return ConnectResponse{s.List.Struct(i)} }

func (s ConnectResponse_List) Set(i int, v ConnectResponse) error {
	return s.List.SetStruct(i, v.Struct)
}

// ConnectResponse_Promise is a wrapper for a ConnectResponse promised by a client call.
type ConnectResponse_Promise struct{ *capnp.Pipeline }

func (p ConnectResponse_Promise) Struct() (ConnectResponse, error) {
	s, err := p.Pipeline.Struct()
	return ConnectResponse{s}, err
}

func (p ConnectResponse_Promise) Game() Game {
	return Game{Client: p.Pipeline.GetPipeline(0).Client()}
}

type ConnectResponse_Status uint16

// Values of ConnectResponse_Status.
const (
	ConnectResponse_Status_success          ConnectResponse_Status = 0
	ConnectResponse_Status_nameTooLong      ConnectResponse_Status = 1
	ConnectResponse_Status_nameAlreadyTaken ConnectResponse_Status = 2
	ConnectResponse_Status_gameIsFull       ConnectResponse_Status = 3
	ConnectResponse_Status_yaDoneGoofed     ConnectResponse_Status = 4
)

// String returns the enum's constant name.
func (c ConnectResponse_Status) String() string {
	switch c {
	case ConnectResponse_Status_success:
		return "success"
	case ConnectResponse_Status_nameTooLong:
		return "nameTooLong"
	case ConnectResponse_Status_nameAlreadyTaken:
		return "nameAlreadyTaken"
	case ConnectResponse_Status_gameIsFull:
		return "gameIsFull"
	case ConnectResponse_Status_yaDoneGoofed:
		return "yaDoneGoofed"

	default:
		return ""
	}
}

// ConnectResponse_StatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ConnectResponse_StatusFromString(c string) ConnectResponse_Status {
	switch c {
	case "success":
		return ConnectResponse_Status_success
	case "nameTooLong":
		return ConnectResponse_Status_nameTooLong
	case "nameAlreadyTaken":
		return ConnectResponse_Status_nameAlreadyTaken
	case "gameIsFull":
		return ConnectResponse_Status_gameIsFull
	case "yaDoneGoofed":
		return ConnectResponse_Status_yaDoneGoofed

	default:
		return 0
	}
}

type ConnectResponse_Status_List struct{ capnp.List }

func NewConnectResponse_Status_List(s *capnp.Segment, sz int32) (ConnectResponse_Status_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ConnectResponse_Status_List{l.List}, err
}

func (l ConnectResponse_Status_List) At(i int) ConnectResponse_Status {
	ul := capnp.UInt16List{List: l.List}
	return ConnectResponse_Status(ul.At(i))
}

func (l ConnectResponse_Status_List) Set(i int, v ConnectResponse_Status) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Player struct{ Client capnp.Client }

func (c Player) Split(ctx context.Context, params func(SplitRequest) error, opts ...capnp.CallOption) Player_split_Results_Promise {
	if c.Client == nil {
		return Player_split_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      0,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "split",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SplitRequest{Struct: s}) }
	}
	return Player_split_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Player) NewTile(ctx context.Context, params func(NewTileRequest) error, opts ...capnp.CallOption) Player_newTile_Results_Promise {
	if c.Client == nil {
		return Player_newTile_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      1,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "newTile",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(NewTileRequest{Struct: s}) }
	}
	return Player_newTile_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Player) DumpNotice(ctx context.Context, params func(DumpNoticeRequest) error, opts ...capnp.CallOption) Player_dumpNotice_Results_Promise {
	if c.Client == nil {
		return Player_dumpNotice_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      2,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "dumpNotice",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(DumpNoticeRequest{Struct: s}) }
	}
	return Player_dumpNotice_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Player_Server interface {
	Split(Player_split) error

	NewTile(Player_newTile) error

	DumpNotice(Player_dumpNotice) error
}

func Player_ServerToClient(s Player_Server) Player {
	c, _ := s.(server.Closer)
	return Player{Client: server.New(Player_Methods(nil, s), c)}
}

func Player_Methods(methods []server.Method, s Player_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      0,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "split",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Player_split{c, opts, SplitRequest{Struct: p}, Player_split_Results{Struct: r}}
			return s.Split(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      1,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "newTile",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Player_newTile{c, opts, NewTileRequest{Struct: p}, Player_newTile_Results{Struct: r}}
			return s.NewTile(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      2,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "dumpNotice",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Player_dumpNotice{c, opts, DumpNoticeRequest{Struct: p}, Player_dumpNotice_Results{Struct: r}}
			return s.DumpNotice(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Player_split holds the arguments for a server call to Player.split.
type Player_split struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SplitRequest
	Results Player_split_Results
}

// Player_newTile holds the arguments for a server call to Player.newTile.
type Player_newTile struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  NewTileRequest
	Results Player_newTile_Results
}

// Player_dumpNotice holds the arguments for a server call to Player.dumpNotice.
type Player_dumpNotice struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  DumpNoticeRequest
	Results Player_dumpNotice_Results
}

type Player_split_Results struct{ capnp.Struct }

// Player_split_Results_TypeID is the unique identifier for the type Player_split_Results.
const Player_split_Results_TypeID = 0x87747dd4e5141ec8

func NewPlayer_split_Results(s *capnp.Segment) (Player_split_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_split_Results{st}, err
}

func NewRootPlayer_split_Results(s *capnp.Segment) (Player_split_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_split_Results{st}, err
}

func ReadRootPlayer_split_Results(msg *capnp.Message) (Player_split_Results, error) {
	root, err := msg.RootPtr()
	return Player_split_Results{root.Struct()}, err
}

func (s Player_split_Results) String() string {
	str, _ := text.Marshal(0x87747dd4e5141ec8, s.Struct)
	return str
}

// Player_split_Results_List is a list of Player_split_Results.
type Player_split_Results_List struct{ capnp.List }

// NewPlayer_split_Results creates a new list of Player_split_Results.
func NewPlayer_split_Results_List(s *capnp.Segment, sz int32) (Player_split_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Player_split_Results_List{l}, err
}

func (s Player_split_Results_List) At(i int) Player_split_Results {
	return Player_split_Results{s.List.Struct(i)}
}

func (s Player_split_Results_List) Set(i int, v Player_split_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Player_split_Results_Promise is a wrapper for a Player_split_Results promised by a client call.
type Player_split_Results_Promise struct{ *capnp.Pipeline }

func (p Player_split_Results_Promise) Struct() (Player_split_Results, error) {
	s, err := p.Pipeline.Struct()
	return Player_split_Results{s}, err
}

type Player_newTile_Results struct{ capnp.Struct }

// Player_newTile_Results_TypeID is the unique identifier for the type Player_newTile_Results.
const Player_newTile_Results_TypeID = 0x8e62748365be5a21

func NewPlayer_newTile_Results(s *capnp.Segment) (Player_newTile_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_newTile_Results{st}, err
}

func NewRootPlayer_newTile_Results(s *capnp.Segment) (Player_newTile_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_newTile_Results{st}, err
}

func ReadRootPlayer_newTile_Results(msg *capnp.Message) (Player_newTile_Results, error) {
	root, err := msg.RootPtr()
	return Player_newTile_Results{root.Struct()}, err
}

func (s Player_newTile_Results) String() string {
	str, _ := text.Marshal(0x8e62748365be5a21, s.Struct)
	return str
}

// Player_newTile_Results_List is a list of Player_newTile_Results.
type Player_newTile_Results_List struct{ capnp.List }

// NewPlayer_newTile_Results creates a new list of Player_newTile_Results.
func NewPlayer_newTile_Results_List(s *capnp.Segment, sz int32) (Player_newTile_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Player_newTile_Results_List{l}, err
}

func (s Player_newTile_Results_List) At(i int) Player_newTile_Results {
	return Player_newTile_Results{s.List.Struct(i)}
}

func (s Player_newTile_Results_List) Set(i int, v Player_newTile_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Player_newTile_Results_Promise is a wrapper for a Player_newTile_Results promised by a client call.
type Player_newTile_Results_Promise struct{ *capnp.Pipeline }

func (p Player_newTile_Results_Promise) Struct() (Player_newTile_Results, error) {
	s, err := p.Pipeline.Struct()
	return Player_newTile_Results{s}, err
}

type Player_dumpNotice_Results struct{ capnp.Struct }

// Player_dumpNotice_Results_TypeID is the unique identifier for the type Player_dumpNotice_Results.
const Player_dumpNotice_Results_TypeID = 0xa8b89fd092566231

func NewPlayer_dumpNotice_Results(s *capnp.Segment) (Player_dumpNotice_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_dumpNotice_Results{st}, err
}

func NewRootPlayer_dumpNotice_Results(s *capnp.Segment) (Player_dumpNotice_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_dumpNotice_Results{st}, err
}

func ReadRootPlayer_dumpNotice_Results(msg *capnp.Message) (Player_dumpNotice_Results, error) {
	root, err := msg.RootPtr()
	return Player_dumpNotice_Results{root.Struct()}, err
}

func (s Player_dumpNotice_Results) String() string {
	str, _ := text.Marshal(0xa8b89fd092566231, s.Struct)
	return str
}

// Player_dumpNotice_Results_List is a list of Player_dumpNotice_Results.
type Player_dumpNotice_Results_List struct{ capnp.List }

// NewPlayer_dumpNotice_Results creates a new list of Player_dumpNotice_Results.
func NewPlayer_dumpNotice_Results_List(s *capnp.Segment, sz int32) (Player_dumpNotice_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Player_dumpNotice_Results_List{l}, err
}

func (s Player_dumpNotice_Results_List) At(i int) Player_dumpNotice_Results {
	return Player_dumpNotice_Results{s.List.Struct(i)}
}

func (s Player_dumpNotice_Results_List) Set(i int, v Player_dumpNotice_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Player_dumpNotice_Results_Promise is a wrapper for a Player_dumpNotice_Results promised by a client call.
type Player_dumpNotice_Results_Promise struct{ *capnp.Pipeline }

func (p Player_dumpNotice_Results_Promise) Struct() (Player_dumpNotice_Results, error) {
	s, err := p.Pipeline.Struct()
	return Player_dumpNotice_Results{s}, err
}

type SplitRequest struct{ capnp.Struct }

// SplitRequest_TypeID is the unique identifier for the type SplitRequest.
const SplitRequest_TypeID = 0xc39499f04a0c4c79

func NewSplitRequest(s *capnp.Segment) (SplitRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SplitRequest{st}, err
}

func NewRootSplitRequest(s *capnp.Segment) (SplitRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SplitRequest{st}, err
}

func ReadRootSplitRequest(msg *capnp.Message) (SplitRequest, error) {
	root, err := msg.RootPtr()
	return SplitRequest{root.Struct()}, err
}

func (s SplitRequest) String() string {
	str, _ := text.Marshal(0xc39499f04a0c4c79, s.Struct)
	return str
}

func (s SplitRequest) Tiles() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s SplitRequest) HasTiles() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SplitRequest) SetTiles(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTiles sets the tiles field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s SplitRequest) NewTiles(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s SplitRequest) Players() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s SplitRequest) HasPlayers() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SplitRequest) SetPlayers(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewPlayers sets the players field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s SplitRequest) NewPlayers(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// SplitRequest_List is a list of SplitRequest.
type SplitRequest_List struct{ capnp.List }

// NewSplitRequest creates a new list of SplitRequest.
func NewSplitRequest_List(s *capnp.Segment, sz int32) (SplitRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SplitRequest_List{l}, err
}

func (s SplitRequest_List) At(i int) SplitRequest { return SplitRequest{s.List.Struct(i)} }

func (s SplitRequest_List) Set(i int, v SplitRequest) error { return s.List.SetStruct(i, v.Struct) }

// SplitRequest_Promise is a wrapper for a SplitRequest promised by a client call.
type SplitRequest_Promise struct{ *capnp.Pipeline }

func (p SplitRequest_Promise) Struct() (SplitRequest, error) {
	s, err := p.Pipeline.Struct()
	return SplitRequest{s}, err
}

type NewTileRequest struct{ capnp.Struct }

// NewTileRequest_TypeID is the unique identifier for the type NewTileRequest.
const NewTileRequest_TypeID = 0xd8aa147e9dd40792

func NewNewTileRequest(s *capnp.Segment) (NewTileRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return NewTileRequest{st}, err
}

func NewRootNewTileRequest(s *capnp.Segment) (NewTileRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return NewTileRequest{st}, err
}

func ReadRootNewTileRequest(msg *capnp.Message) (NewTileRequest, error) {
	root, err := msg.RootPtr()
	return NewTileRequest{root.Struct()}, err
}

func (s NewTileRequest) String() string {
	str, _ := text.Marshal(0xd8aa147e9dd40792, s.Struct)
	return str
}

func (s NewTileRequest) Letter() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s NewTileRequest) HasLetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NewTileRequest) LetterBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s NewTileRequest) SetLetter(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s NewTileRequest) Peeler() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s NewTileRequest) HasPeeler() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s NewTileRequest) PeelerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s NewTileRequest) SetPeeler(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// NewTileRequest_List is a list of NewTileRequest.
type NewTileRequest_List struct{ capnp.List }

// NewNewTileRequest creates a new list of NewTileRequest.
func NewNewTileRequest_List(s *capnp.Segment, sz int32) (NewTileRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return NewTileRequest_List{l}, err
}

func (s NewTileRequest_List) At(i int) NewTileRequest { return NewTileRequest{s.List.Struct(i)} }

func (s NewTileRequest_List) Set(i int, v NewTileRequest) error { return s.List.SetStruct(i, v.Struct) }

// NewTileRequest_Promise is a wrapper for a NewTileRequest promised by a client call.
type NewTileRequest_Promise struct{ *capnp.Pipeline }

func (p NewTileRequest_Promise) Struct() (NewTileRequest, error) {
	s, err := p.Pipeline.Struct()
	return NewTileRequest{s}, err
}

type DumpNoticeRequest struct{ capnp.Struct }

// DumpNoticeRequest_TypeID is the unique identifier for the type DumpNoticeRequest.
const DumpNoticeRequest_TypeID = 0xd410a36248b8782e

func NewDumpNoticeRequest(s *capnp.Segment) (DumpNoticeRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return DumpNoticeRequest{st}, err
}

func NewRootDumpNoticeRequest(s *capnp.Segment) (DumpNoticeRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return DumpNoticeRequest{st}, err
}

func ReadRootDumpNoticeRequest(msg *capnp.Message) (DumpNoticeRequest, error) {
	root, err := msg.RootPtr()
	return DumpNoticeRequest{root.Struct()}, err
}

func (s DumpNoticeRequest) String() string {
	str, _ := text.Marshal(0xd410a36248b8782e, s.Struct)
	return str
}

func (s DumpNoticeRequest) Dumped() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DumpNoticeRequest) HasDumped() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DumpNoticeRequest) DumpedBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DumpNoticeRequest) SetDumped(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s DumpNoticeRequest) Received() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s DumpNoticeRequest) HasReceived() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DumpNoticeRequest) SetReceived(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewReceived sets the received field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s DumpNoticeRequest) NewReceived(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s DumpNoticeRequest) Dumper() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s DumpNoticeRequest) HasDumper() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s DumpNoticeRequest) DumperBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s DumpNoticeRequest) SetDumper(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// DumpNoticeRequest_List is a list of DumpNoticeRequest.
type DumpNoticeRequest_List struct{ capnp.List }

// NewDumpNoticeRequest creates a new list of DumpNoticeRequest.
func NewDumpNoticeRequest_List(s *capnp.Segment, sz int32) (DumpNoticeRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return DumpNoticeRequest_List{l}, err
}

func (s DumpNoticeRequest_List) At(i int) DumpNoticeRequest {
	return DumpNoticeRequest{s.List.Struct(i)}
}

func (s DumpNoticeRequest_List) Set(i int, v DumpNoticeRequest) error {
	return s.List.SetStruct(i, v.Struct)
}

// DumpNoticeRequest_Promise is a wrapper for a DumpNoticeRequest promised by a client call.
type DumpNoticeRequest_Promise struct{ *capnp.Pipeline }

func (p DumpNoticeRequest_Promise) Struct() (DumpNoticeRequest, error) {
	s, err := p.Pipeline.Struct()
	return DumpNoticeRequest{s}, err
}

type Game struct{ Client capnp.Client }

func (c Game) Peel(ctx context.Context, params func(Game_peel_Params) error, opts ...capnp.CallOption) PeelResponse_Promise {
	if c.Client == nil {
		return PeelResponse_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa7f47872907c494c,
			MethodID:      0,
			InterfaceName: "potassium.capnp:Game",
			MethodName:    "peel",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Game_peel_Params{Struct: s}) }
	}
	return PeelResponse_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Game) Dump(ctx context.Context, params func(Game_dump_Params) error, opts ...capnp.CallOption) DumpResponse_Promise {
	if c.Client == nil {
		return DumpResponse_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa7f47872907c494c,
			MethodID:      1,
			InterfaceName: "potassium.capnp:Game",
			MethodName:    "dump",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Game_dump_Params{Struct: s}) }
	}
	return DumpResponse_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Game_Server interface {
	Peel(Game_peel) error

	Dump(Game_dump) error
}

func Game_ServerToClient(s Game_Server) Game {
	c, _ := s.(server.Closer)
	return Game{Client: server.New(Game_Methods(nil, s), c)}
}

func Game_Methods(methods []server.Method, s Game_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa7f47872907c494c,
			MethodID:      0,
			InterfaceName: "potassium.capnp:Game",
			MethodName:    "peel",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Game_peel{c, opts, Game_peel_Params{Struct: p}, PeelResponse{Struct: r}}
			return s.Peel(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa7f47872907c494c,
			MethodID:      1,
			InterfaceName: "potassium.capnp:Game",
			MethodName:    "dump",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Game_dump{c, opts, Game_dump_Params{Struct: p}, DumpResponse{Struct: r}}
			return s.Dump(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	return methods
}

// Game_peel holds the arguments for a server call to Game.peel.
type Game_peel struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Game_peel_Params
	Results PeelResponse
}

// Game_dump holds the arguments for a server call to Game.dump.
type Game_dump struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Game_dump_Params
	Results DumpResponse
}

type Game_peel_Params struct{ capnp.Struct }

// Game_peel_Params_TypeID is the unique identifier for the type Game_peel_Params.
const Game_peel_Params_TypeID = 0xedb6aad718bd116a

func NewGame_peel_Params(s *capnp.Segment) (Game_peel_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Game_peel_Params{st}, err
}

func NewRootGame_peel_Params(s *capnp.Segment) (Game_peel_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Game_peel_Params{st}, err
}

func ReadRootGame_peel_Params(msg *capnp.Message) (Game_peel_Params, error) {
	root, err := msg.RootPtr()
	return Game_peel_Params{root.Struct()}, err
}

func (s Game_peel_Params) String() string {
	str, _ := text.Marshal(0xedb6aad718bd116a, s.Struct)
	return str
}

func (s Game_peel_Params) Board() (Board, error) {
	p, err := s.Struct.Ptr(0)
	return Board{Struct: p.Struct()}, err
}

func (s Game_peel_Params) HasBoard() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Game_peel_Params) SetBoard(v Board) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBoard sets the board field to a newly
// allocated Board struct, preferring placement in s's segment.
func (s Game_peel_Params) NewBoard() (Board, error) {
	ss, err := NewBoard(s.Struct.Segment())
	if err != nil {
		return Board{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Game_peel_Params_List is a list of Game_peel_Params.
type Game_peel_Params_List struct{ capnp.List }

// NewGame_peel_Params creates a new list of Game_peel_Params.
func NewGame_peel_Params_List(s *capnp.Segment, sz int32) (Game_peel_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Game_peel_Params_List{l}, err
}

func (s Game_peel_Params_List) At(i int) Game_peel_Params { return Game_peel_Params{s.List.Struct(i)} }

func (s Game_peel_Params_List) Set(i int, v Game_peel_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Game_peel_Params_Promise is a wrapper for a Game_peel_Params promised by a client call.
type Game_peel_Params_Promise struct{ *capnp.Pipeline }

func (p Game_peel_Params_Promise) Struct() (Game_peel_Params, error) {
	s, err := p.Pipeline.Struct()
	return Game_peel_Params{s}, err
}

func (p Game_peel_Params_Promise) Board() Board_Promise {
	return Board_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Game_dump_Params struct{ capnp.Struct }

// Game_dump_Params_TypeID is the unique identifier for the type Game_dump_Params.
const Game_dump_Params_TypeID = 0xd09df4f343b35119

func NewGame_dump_Params(s *capnp.Segment) (Game_dump_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Game_dump_Params{st}, err
}

func NewRootGame_dump_Params(s *capnp.Segment) (Game_dump_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Game_dump_Params{st}, err
}

func ReadRootGame_dump_Params(msg *capnp.Message) (Game_dump_Params, error) {
	root, err := msg.RootPtr()
	return Game_dump_Params{root.Struct()}, err
}

func (s Game_dump_Params) String() string {
	str, _ := text.Marshal(0xd09df4f343b35119, s.Struct)
	return str
}

func (s Game_dump_Params) Letter() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Game_dump_Params) HasLetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Game_dump_Params) LetterBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Game_dump_Params) SetLetter(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Game_dump_Params_List is a list of Game_dump_Params.
type Game_dump_Params_List struct{ capnp.List }

// NewGame_dump_Params creates a new list of Game_dump_Params.
func NewGame_dump_Params_List(s *capnp.Segment, sz int32) (Game_dump_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Game_dump_Params_List{l}, err
}

func (s Game_dump_Params_List) At(i int) Game_dump_Params { return Game_dump_Params{s.List.Struct(i)} }

func (s Game_dump_Params_List) Set(i int, v Game_dump_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Game_dump_Params_Promise is a wrapper for a Game_dump_Params promised by a client call.
type Game_dump_Params_Promise struct{ *capnp.Pipeline }

func (p Game_dump_Params_Promise) Struct() (Game_dump_Params, error) {
	s, err := p.Pipeline.Struct()
	return Game_dump_Params{s}, err
}

type PeelResponse struct{ capnp.Struct }
type PeelResponse_Which uint16

const (
	PeelResponse_Which_valid         PeelResponse_Which = 0
	PeelResponse_Which_invalidWord   PeelResponse_Which = 1
	PeelResponse_Which_notAllLetters PeelResponse_Which = 2
	PeelResponse_Which_extraLetters  PeelResponse_Which = 3
)

func (w PeelResponse_Which) String() string {
	const s = "validinvalidWordnotAllLettersextraLetters"
	switch w {
	case PeelResponse_Which_valid:
		return s[0:5]
	case PeelResponse_Which_invalidWord:
		return s[5:16]
	case PeelResponse_Which_notAllLetters:
		return s[16:29]
	case PeelResponse_Which_extraLetters:
		return s[29:41]

	}
	return "PeelResponse_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// PeelResponse_TypeID is the unique identifier for the type PeelResponse.
const PeelResponse_TypeID = 0xe94d36c9c6cc16d1

func NewPeelResponse(s *capnp.Segment) (PeelResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PeelResponse{st}, err
}

func NewRootPeelResponse(s *capnp.Segment) (PeelResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PeelResponse{st}, err
}

func ReadRootPeelResponse(msg *capnp.Message) (PeelResponse, error) {
	root, err := msg.RootPtr()
	return PeelResponse{root.Struct()}, err
}

func (s PeelResponse) String() string {
	str, _ := text.Marshal(0xe94d36c9c6cc16d1, s.Struct)
	return str
}

func (s PeelResponse) Which() PeelResponse_Which {
	return PeelResponse_Which(s.Struct.Uint16(2))
}
func (s PeelResponse) Status() PeelResponse_Status {
	return PeelResponse_Status(s.Struct.Uint16(0))
}

func (s PeelResponse) SetStatus(v PeelResponse_Status) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s PeelResponse) SetValid() {
	s.Struct.SetUint16(2, 0)

}

func (s PeelResponse) InvalidWord() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s PeelResponse) HasInvalidWord() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PeelResponse) SetInvalidWord(v capnp.TextList) error {
	s.Struct.SetUint16(2, 1)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewInvalidWord sets the invalidWord field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s PeelResponse) NewInvalidWord(n int32) (capnp.TextList, error) {
	s.Struct.SetUint16(2, 1)
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s PeelResponse) NotAllLetters() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s PeelResponse) HasNotAllLetters() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PeelResponse) SetNotAllLetters(v capnp.TextList) error {
	s.Struct.SetUint16(2, 2)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNotAllLetters sets the notAllLetters field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s PeelResponse) NewNotAllLetters(n int32) (capnp.TextList, error) {
	s.Struct.SetUint16(2, 2)
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s PeelResponse) ExtraLetters() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s PeelResponse) HasExtraLetters() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PeelResponse) SetExtraLetters(v capnp.TextList) error {
	s.Struct.SetUint16(2, 3)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewExtraLetters sets the extraLetters field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s PeelResponse) NewExtraLetters(n int32) (capnp.TextList, error) {
	s.Struct.SetUint16(2, 3)
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// PeelResponse_List is a list of PeelResponse.
type PeelResponse_List struct{ capnp.List }

// NewPeelResponse creates a new list of PeelResponse.
func NewPeelResponse_List(s *capnp.Segment, sz int32) (PeelResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PeelResponse_List{l}, err
}

func (s PeelResponse_List) At(i int) PeelResponse { return PeelResponse{s.List.Struct(i)} }

func (s PeelResponse_List) Set(i int, v PeelResponse) error { return s.List.SetStruct(i, v.Struct) }

// PeelResponse_Promise is a wrapper for a PeelResponse promised by a client call.
type PeelResponse_Promise struct{ *capnp.Pipeline }

func (p PeelResponse_Promise) Struct() (PeelResponse, error) {
	s, err := p.Pipeline.Struct()
	return PeelResponse{s}, err
}

type PeelResponse_Status uint16

// Values of PeelResponse_Status.
const (
	PeelResponse_Status_success       PeelResponse_Status = 0
	PeelResponse_Status_invalidWord   PeelResponse_Status = 1
	PeelResponse_Status_detachedBoard PeelResponse_Status = 2
	PeelResponse_Status_notAllLetters PeelResponse_Status = 3
	PeelResponse_Status_extraLetters  PeelResponse_Status = 4
	PeelResponse_Status_invalidBoard  PeelResponse_Status = 5
)

// String returns the enum's constant name.
func (c PeelResponse_Status) String() string {
	switch c {
	case PeelResponse_Status_success:
		return "success"
	case PeelResponse_Status_invalidWord:
		return "invalidWord"
	case PeelResponse_Status_detachedBoard:
		return "detachedBoard"
	case PeelResponse_Status_notAllLetters:
		return "notAllLetters"
	case PeelResponse_Status_extraLetters:
		return "extraLetters"
	case PeelResponse_Status_invalidBoard:
		return "invalidBoard"

	default:
		return ""
	}
}

// PeelResponse_StatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func PeelResponse_StatusFromString(c string) PeelResponse_Status {
	switch c {
	case "success":
		return PeelResponse_Status_success
	case "invalidWord":
		return PeelResponse_Status_invalidWord
	case "detachedBoard":
		return PeelResponse_Status_detachedBoard
	case "notAllLetters":
		return PeelResponse_Status_notAllLetters
	case "extraLetters":
		return PeelResponse_Status_extraLetters
	case "invalidBoard":
		return PeelResponse_Status_invalidBoard

	default:
		return 0
	}
}

type PeelResponse_Status_List struct{ capnp.List }

func NewPeelResponse_Status_List(s *capnp.Segment, sz int32) (PeelResponse_Status_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return PeelResponse_Status_List{l.List}, err
}

func (l PeelResponse_Status_List) At(i int) PeelResponse_Status {
	ul := capnp.UInt16List{List: l.List}
	return PeelResponse_Status(ul.At(i))
}

func (l PeelResponse_Status_List) Set(i int, v PeelResponse_Status) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type DumpResponse struct{ capnp.Struct }

// DumpResponse_TypeID is the unique identifier for the type DumpResponse.
const DumpResponse_TypeID = 0xa6437b0e305e7e4f

func NewDumpResponse(s *capnp.Segment) (DumpResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return DumpResponse{st}, err
}

func NewRootDumpResponse(s *capnp.Segment) (DumpResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return DumpResponse{st}, err
}

func ReadRootDumpResponse(msg *capnp.Message) (DumpResponse, error) {
	root, err := msg.RootPtr()
	return DumpResponse{root.Struct()}, err
}

func (s DumpResponse) String() string {
	str, _ := text.Marshal(0xa6437b0e305e7e4f, s.Struct)
	return str
}

func (s DumpResponse) Status() DumpResponse_Status {
	return DumpResponse_Status(s.Struct.Uint16(0))
}

func (s DumpResponse) SetStatus(v DumpResponse_Status) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s DumpResponse) Letters() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s DumpResponse) HasLetters() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DumpResponse) SetLetters(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewLetters sets the letters field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s DumpResponse) NewLetters(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// DumpResponse_List is a list of DumpResponse.
type DumpResponse_List struct{ capnp.List }

// NewDumpResponse creates a new list of DumpResponse.
func NewDumpResponse_List(s *capnp.Segment, sz int32) (DumpResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return DumpResponse_List{l}, err
}

func (s DumpResponse_List) At(i int) DumpResponse { return DumpResponse{s.List.Struct(i)} }

func (s DumpResponse_List) Set(i int, v DumpResponse) error { return s.List.SetStruct(i, v.Struct) }

// DumpResponse_Promise is a wrapper for a DumpResponse promised by a client call.
type DumpResponse_Promise struct{ *capnp.Pipeline }

func (p DumpResponse_Promise) Struct() (DumpResponse, error) {
	s, err := p.Pipeline.Struct()
	return DumpResponse{s}, err
}

type DumpResponse_Status uint16

// Values of DumpResponse_Status.
const (
	DumpResponse_Status_success          DumpResponse_Status = 0
	DumpResponse_Status_letterNotInTiles DumpResponse_Status = 1
	DumpResponse_Status_notEnoughTiles   DumpResponse_Status = 2
)

// String returns the enum's constant name.
func (c DumpResponse_Status) String() string {
	switch c {
	case DumpResponse_Status_success:
		return "success"
	case DumpResponse_Status_letterNotInTiles:
		return "letterNotInTiles"
	case DumpResponse_Status_notEnoughTiles:
		return "notEnoughTiles"

	default:
		return ""
	}
}

// DumpResponse_StatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func DumpResponse_StatusFromString(c string) DumpResponse_Status {
	switch c {
	case "success":
		return DumpResponse_Status_success
	case "letterNotInTiles":
		return DumpResponse_Status_letterNotInTiles
	case "notEnoughTiles":
		return DumpResponse_Status_notEnoughTiles

	default:
		return 0
	}
}

type DumpResponse_Status_List struct{ capnp.List }

func NewDumpResponse_Status_List(s *capnp.Segment, sz int32) (DumpResponse_Status_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return DumpResponse_Status_List{l.List}, err
}

func (l DumpResponse_Status_List) At(i int) DumpResponse_Status {
	ul := capnp.UInt16List{List: l.List}
	return DumpResponse_Status(ul.At(i))
}

func (l DumpResponse_Status_List) Set(i int, v DumpResponse_Status) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Board struct{ capnp.Struct }

// Board_TypeID is the unique identifier for the type Board.
const Board_TypeID = 0xfa3b78ca11f5cdeb

func NewBoard(s *capnp.Segment) (Board, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Board{st}, err
}

func NewRootBoard(s *capnp.Segment) (Board, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Board{st}, err
}

func ReadRootBoard(msg *capnp.Message) (Board, error) {
	root, err := msg.RootPtr()
	return Board{root.Struct()}, err
}

func (s Board) String() string {
	str, _ := text.Marshal(0xfa3b78ca11f5cdeb, s.Struct)
	return str
}

func (s Board) Words() (Word_List, error) {
	p, err := s.Struct.Ptr(0)
	return Word_List{List: p.List()}, err
}

func (s Board) HasWords() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Board) SetWords(v Word_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewWords sets the words field to a newly
// allocated Word_List, preferring placement in s's segment.
func (s Board) NewWords(n int32) (Word_List, error) {
	l, err := NewWord_List(s.Struct.Segment(), n)
	if err != nil {
		return Word_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Board_List is a list of Board.
type Board_List struct{ capnp.List }

// NewBoard creates a new list of Board.
func NewBoard_List(s *capnp.Segment, sz int32) (Board_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Board_List{l}, err
}

func (s Board_List) At(i int) Board { return Board{s.List.Struct(i)} }

func (s Board_List) Set(i int, v Board) error { return s.List.SetStruct(i, v.Struct) }

// Board_Promise is a wrapper for a Board promised by a client call.
type Board_Promise struct{ *capnp.Pipeline }

func (p Board_Promise) Struct() (Board, error) {
	s, err := p.Pipeline.Struct()
	return Board{s}, err
}

type Word struct{ capnp.Struct }

// Word_TypeID is the unique identifier for the type Word.
const Word_TypeID = 0xe0cda35555661008

func NewWord(s *capnp.Segment) (Word, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Word{st}, err
}

func NewRootWord(s *capnp.Segment) (Word, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Word{st}, err
}

func ReadRootWord(msg *capnp.Message) (Word, error) {
	root, err := msg.RootPtr()
	return Word{root.Struct()}, err
}

func (s Word) String() string {
	str, _ := text.Marshal(0xe0cda35555661008, s.Struct)
	return str
}

func (s Word) Text() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Word) HasText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Word) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Word) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Word) Orientation() Word_Orientation {
	return Word_Orientation(s.Struct.Uint16(0))
}

func (s Word) SetOrientation(v Word_Orientation) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s Word) X() uint32 {
	return s.Struct.Uint32(4)
}

func (s Word) SetX(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Word) Y() uint32 {
	return s.Struct.Uint32(8)
}

func (s Word) SetY(v uint32) {
	s.Struct.SetUint32(8, v)
}

// Word_List is a list of Word.
type Word_List struct{ capnp.List }

// NewWord creates a new list of Word.
func NewWord_List(s *capnp.Segment, sz int32) (Word_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Word_List{l}, err
}

func (s Word_List) At(i int) Word { return Word{s.List.Struct(i)} }

func (s Word_List) Set(i int, v Word) error { return s.List.SetStruct(i, v.Struct) }

// Word_Promise is a wrapper for a Word promised by a client call.
type Word_Promise struct{ *capnp.Pipeline }

func (p Word_Promise) Struct() (Word, error) {
	s, err := p.Pipeline.Struct()
	return Word{s}, err
}

type Word_Orientation uint16

// Values of Word_Orientation.
const (
	Word_Orientation_unknown    Word_Orientation = 0
	Word_Orientation_horizontal Word_Orientation = 1
	Word_Orientation_vertical   Word_Orientation = 2
)

// String returns the enum's constant name.
func (c Word_Orientation) String() string {
	switch c {
	case Word_Orientation_unknown:
		return "unknown"
	case Word_Orientation_horizontal:
		return "horizontal"
	case Word_Orientation_vertical:
		return "vertical"

	default:
		return ""
	}
}

// Word_OrientationFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Word_OrientationFromString(c string) Word_Orientation {
	switch c {
	case "unknown":
		return Word_Orientation_unknown
	case "horizontal":
		return Word_Orientation_horizontal
	case "vertical":
		return Word_Orientation_vertical

	default:
		return 0
	}
}

type Word_Orientation_List struct{ capnp.List }

func NewWord_Orientation_List(s *capnp.Segment, sz int32) (Word_Orientation_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Word_Orientation_List{l.List}, err
}

func (l Word_Orientation_List) At(i int) Word_Orientation {
	ul := capnp.UInt16List{List: l.List}
	return Word_Orientation(ul.At(i))
}

func (l Word_Orientation_List) Set(i int, v Word_Orientation) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_f2768a70df34b76a = "x\xda\x8cW]l\x1cW\x15>\xe7\xce\xacg\x13g" +
	"\xb3{;k\xd2<\x84%\x95-\xd5\x96l\xc5q\x83" +
	"\xa8QY\xdb\x89q\x1d9\x89\xaf\xed`\xd5\x02\xc4x" +
	"\xe7&^g<\xb3\x9e\x99\xb5\xd7\xa1\xadU(j\xc9" +
	"\x03B\x15\x82\x80j\x8aT\x99\x1f\xf3\x80\xdc6\x90V" +
	"*\x02!\x01E\x0d\x8d\xa5Hmy\x81\x97\x0a\x08*" +
	"*\xd0>4\"\x0c\xbaw<3\xeb\xf5\x86\xf0\xe4\xbb" +
	"3g\xce\xcfw\xbe\xf3\x9d\xeb#o(\x03\xa47\xf5" +
	"\\\x1a\x80\x99\xa9\x96\xe0\xb7\x1f\xcd\xbfs\xe31\xff)" +
	"\xa0\x07\x10@\xd5\x00\xfa\x8e\xa9]\x08jpx\xe6\xe7" +
	"\xfc\xcb\xfe\xec\xd7\xea\xde\x1cR\xfb\xc5\x9b\x8b_\xdd\x9b" +
	"\x9f\xbe\xff\x1f\x97\x81\xb6\x91 \x9d;w\xf6\xec\xf3\xd7" +
	"\xfe\x04\x80})u/\xeam\xc2T\xa7\xea\x19\xc0\xe0" +
	"\xa1sOd\x8c\x17\xf7\xac\x03\xcd)\xc1\xfc\xcf\x1e\xf8" +
	"c\xe5\xd2\xd2?\x01P\xefV_\xd1\x8fI\xc3^u" +
	"D\x7fD\x9c\x82\xdbo/\xfe}\xfa\x97\xee:\xd0C" +
	"$\xd8\xfa\xc8\xeb\xbf~\xed\xe3\xa7\xfe*\xdc\x0e\xaa\xf7" +
	"\xa1\xce\xa4\xf5)\xf5[\x80\xc1\x99\xc7?\x7fd\xff\x17" +
	"\x8f\x7f\x1fX\x0e1\xf1\x9bBa\xb2\xae\xde\xd47\xd5" +
	"\x03\x00\xfa\xcb\xea2`06\xfa\xe8\xd7\xdd\xda\xfb?" +
	"\xd8\x95CGjC\xefN\x89O:SO\xe9+\xe2" +
	"\x14\xf4\xce~\xe6\x99\xeb\xcf]\xfda]\xd1FjB" +
	"\x14\xfd\xf4\xb3\xfb\xbf\xfb%\xd3|\x01h\xae>$\x91" +
	"Y\xa5n\xe9\x8fHOgS\xcb\x80\xb7/\xbf\xf7\xcd" +
	"\xe2\xc8\xa7~\xd1\x18p3\xf5\x8a\xfe\xb24\xbb\x92\xfa" +
	"\x8d\xde\xdb\"\x02\xae\x8c\xed;\xf9\xde\xb7\xbf\xf1\xabf" +
	"n\x0f\xb6\xdc\xd4;\x84\x95~\xb8ET\xa2|\xf8\x9d" +
	"\x1f\xeb\xd7.\xbf\xd1\xac\xec\xbe\xb5\x16\x82\xfa\x8fZD" +
	"\xdd\x9b\xd2\xfa {\xf1\xf8\xbf\xde_\xbb\x0e\xb4\x0da" +
	"\xdb(\xa3\xedE@\xbdM+\x02\x06=\xb5\xab\x0f\xcf" +
	">\x9f\xbb\xd1\x10Z\x91\x1c\xd0\xeeA}X\x13\xb1\x07" +
	"\xb5\x9f\x00\x06\xcfh7\xd6\x1e\xcfo\xbc\xd5,\xcfw" +
	"\xb5[\xfa\x87\xd2\xf6\x03MD\xee\xe9\xf5_z\xfd\xf7" +
	"\x7fy\x0b\xe8\xc7H\x924`\x1fK\xf7\xa3n\xa4\x85" +
	"\xe5\xe7\xd2\xab\x80A\xe6\x0bon\x95\xd7.\xfcA\xb6" +
	"<\xee*`\xdfz\xfa>\xd4\xafH\xcb\xcd\xb4`R" +
	"L3\x96C\xd2\xd8\xf2\xad\xf4\x86\xfevZ\x94\xfeN" +
	"\xfa\xcf\x80\x09yv\x025\x8c\x9a\x8a\xa8\xbf\xba\xe7\xa6" +
	"\xfe\xda\x9e\x03\x00}[{\x0a\x08\x18\xcc\xd3W\xef}" +
	"s\xe3\xa7\xef\xd6#u\xa8U\"\xd5\xd1*\x90\xfa\xdb" +
	"\xb5\x0f\xe8\xefj\x9f\xbc\xd5P\xbc\x8c=\xdc\xfa\x82~" +
	"\xaaU\x9cF[\x8b\xd0\x1dT\x1c\xdf\xf0\xbcrUY" +
	"\xe8)\x19\x15\xbb\xd2?n\x19+\xdc\xed\xf1*V\xd9" +
	"o\x9f\xe0^\xb6j\xf9\xde\x1d\xcdl\xbe<U\xb6x" +
	"\xfbD\x91{\xcd\x0d\xa7\x1d\xd7\xec9\xe3\x96\xb9\xed\x1b" +
	"~\xd9\xb1\x01\xc6\x11\xd9>$\x00\xf4\xd0\x10\x00\"m" +
	"\x9b\x01@B\xe9I\x80\xd5\xaa}\xc1v\x96\xed`\xce" +
	"q\xcb\x17\x1d\xdb\x07\xc5\xb0\x82%\xee\xfa\xe5\x92a\x01" +
	"@\x1c\x00\xa3\x00\xc5I\xee.qWxU\x95\x14@" +
	"L}\x8czI\xe9\x10\x10\x9a\xd2VK\x8em\xf3\x92" +
	"?\x80\xe3\x88MJ\xe2\xdc\x9a\xe0^\xc5\xb1=\xde3" +
	"\xe9\x1b\x9a_\xf5\x84\xd7{e\xae\x83a\xae\x0f\xce\xca" +
	"\\\x8f\xb9\x00\xa8\xd0^\xf1G\xa5\xdd\xf3\x00\x98\xa2\x9d" +
	"\xf3\x00\xab^\xb5T\xe2\x9e\x17\x94\xed%\xc3*\x9b\xd3" +
	"\xa09\xae\x19\x98\xdc7Js\xdc\x84\xc2\x90c\xb8f" +
	"`;\xfe\xa0e\x8dq(\xf8>w\xbd\x80\xd7|\xd7" +
	"\x18\xe3>d\xe5\xcf\xed\x8f\x87 +\xad\xa3TI\x94" +
	"\xea\x89\xeaBE\xa4\x9a\x15\xb92\x15\xeb\xa9\x89\xfd\xc5" +
	"I\xdf\xf0\xab\x1eK+*\x80\x8a\x00\xb4\xb3\x1f\x80\xb5" +
	"+\xc8\x8e\x10D\xcc\xa3x\xd6=\x04\xc0\xeeW\x90\x9d" +
	" X\xf4\xe4\x17\x98M\xfc\x00b\x16p\xd5\xe22A" +
	"\xdc\x0f8\xae \xee\x03\"\x8e\xbb\xbb\x90\x1d1\x16\xb8" +
	"@+-{\x10q\x14#n\xd3\xde. \xb4C\xc3" +
	"d\xd21\x9a zP\xbc\xcbh\xd9\x0a\xe7\xd6\x00f" +
	"\xcd\xeaBeg\x8f\xd4\x06\xda\x09\x8b\xd3\x8e_.q" +
	"A\xd1\xaa\xe5\xa3\xb7\x1b\xa4\xe3a\xb7'\x8a|\xb1\xca" +
	"=?L.F\xa4+A\x84\xc6\x90\xf4oC\xf2\x00" +
	"\xc1\xacm,pY\xef>\xc0bE\xc6E\x1ai\xa5" +
	"\xa0B3\x18\x8aa\x82\x92\xe2\x12\x88H01\xda\\" +
	"\x94\x1d\x05B\x87\x05\x10\x91Ha\xb4\xbb\xe8\x83\x82\xa8" +
	"\xdd\x1a\x92X\xed0\x92xzx\x06\x08=\xa8\x15\xe4" +
	"d\x0e\xe0\xea\xf6\xe8\x0d`\x10\xa1\x01J\x89\xef\xc4-" +
	"\xc6bR|4\xc1\x17\xb3M\x908\xba\x8d\xc4@\x1d" +
	"\x12\x0f\x09r|BA6E\xb0\xe0\x97-\xde\xc8\x80" +
	"\xd5\x10\x91;\x12Ci\xec\xc2\xf6`\x81dk\"\xb9" +
	"\xff7[\xbb\x92\xd6$l\x8d\xfd\x84l\xcd\x9e\x17=" +
	"\xa3\xc9\x0emhS\x9c\x94 \xab$Q\xfb\xb8\xe1\x1a" +
	"\x0b\x1e\x00S\xe3\xf0\x19\x11>\xad \xcb\x13,\x86\xfc" +
	"\x8fx\xb0\xdb\xd3\x89\x98\x89\x13\x92f\xe8\x87\xbd\x8f\x9c" +
	"\x0d\x0bg\x03\x0a\xb2\xb1:tGO\x02\xb0\x87\x15d" +
	"&AJH^\x0a\x8c!,?\xab \x9b#X\x14" +
	"\xb9q3\x0e\xeb\xf2\x12//q\x13\x00\x1a\x00\x0f-" +
	"w'\x18w\xfetH\x93;LA\xff]\xa6\xa0\xa1" +
	"\xfe\xa2\x98\xd0\xff\x05GC\xb7{\xb6\x9b+\xa2\x86U" +
	"\x1e\x0be\xb4;\x94\xd1\xceKRF;g\xa4\x8cv" +
	"\xd4\xeb\xa7\x18\xbf)\xc7\x19\x03\xcd\xb1\xcf\xcb_\x83\x96" +
	"\xcb\xd10W\xa6\x8c\x0b\xdc\x06\x08D\xb3G\xbdOW" +
	"A\xb1\xac`\xc58\xe1\xd8|\xc4\x81\xacs\x8e\x9b\xcd" +
	"\xfb\xd4L\xdcw.\xa2K2\xab\xb6\x8buy\x84\x00" +
	"\x9cv\xd0\x1f\xb5\x05\x90\x1e\x80P\xefa\xdb\xa9\x9e\x87" +
	"\xe2\x9c|\xd4D\x0e\xc5\xda\x93TOn\x9f8\x1bD" +
	"{\x10\xb4\xb2c\xb3\\\xdc\x07\xa3+\xe9}\xd4\x06>" +
	"+\xee\xbc\x0a\xb2\x8a \x09\x86\xf0-\xdc\x03\xc0\xe6\x14" +
	"d>A\xaa\x90<*\x00tQ<\xb4\x14d5\x82" +
	"Y\x9f\xd7\xfc\xb8?N}<\xcc&\xc9\x84\xf3\x825" +
	"L\x03\xc14 \xaeD\xa7\xdd\x0c\x8a\xf6b\xb2l\x92" +
	"\xabo2\xbe\xf9\xb8\x98\xc7\x04\x7fj\x0a\xb2'\x09f" +
	"0\x08\xc2r\x9e\x102\xf3\xa8\x82\xeci\x82\x19\xf2\x9f" +
	" \xac\xe7+\xa2\xc8'\x15d\xdf#\x98Qn\x07a" +
	"Ak.\x00{VAv\x95`F\xfdw\x90G\x15" +
	"\x80^\x99\x07`/)\xc8\xae\xd7K@\x9cKXR" +
	"A.Nh\xd9\xb9\x7f\x1bU\xaaa\xfd6\xbe\xde\xb9" +
	"\x8d\xef*qRM\xc4d4U\x93\xa3\x89\x9a\x14f" +
	"\xc5:\xc7\\rC\x03\xc4\\\xb3=\x12^\x13\xc2;" +
	"\xcd.O\xed\x04\x0b\xcb\x8ek\xc6\x99\xe5\xea\xff\xad\x11" +
	"\x0f\xff\x1b\x00\x00\xff\xff\xcc\xc7\xa8\x0a"

func init() {
	schemas.Register(schema_f2768a70df34b76a,
		0x87747dd4e5141ec8,
		0x8e62748365be5a21,
		0x98f12857140a897a,
		0xa509b3610d81663d,
		0xa572c157ee71d9fd,
		0xa6437b0e305e7e4f,
		0xa7f47872907c494c,
		0xa8b89fd092566231,
		0xb26464829e0e9c88,
		0xc03e473f96f00098,
		0xc39499f04a0c4c79,
		0xcf98cd13ab9af903,
		0xd09df4f343b35119,
		0xd410a36248b8782e,
		0xd8aa147e9dd40792,
		0xd8e8ceccb474312e,
		0xda6b9d69d1d7600d,
		0xe0cda35555661008,
		0xe94d36c9c6cc16d1,
		0xedb6aad718bd116a,
		0xfa3b78ca11f5cdeb)
}
