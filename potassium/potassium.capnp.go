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
func (c Player) GameOver(ctx context.Context, params func(Player_gameOver_Params) error, opts ...capnp.CallOption) Player_gameOver_Results_Promise {
	if c.Client == nil {
		return Player_gameOver_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      3,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "gameOver",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Player_gameOver_Params{Struct: s}) }
	}
	return Player_gameOver_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Player_Server interface {
	Split(Player_split) error

	NewTile(Player_newTile) error

	DumpNotice(Player_dumpNotice) error

	GameOver(Player_gameOver) error
}

func Player_ServerToClient(s Player_Server) Player {
	c, _ := s.(server.Closer)
	return Player{Client: server.New(Player_Methods(nil, s), c)}
}

func Player_Methods(methods []server.Method, s Player_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
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

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc03e473f96f00098,
			MethodID:      3,
			InterfaceName: "potassium.capnp:Player",
			MethodName:    "gameOver",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Player_gameOver{c, opts, Player_gameOver_Params{Struct: p}, Player_gameOver_Results{Struct: r}}
			return s.GameOver(call)
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

// Player_gameOver holds the arguments for a server call to Player.gameOver.
type Player_gameOver struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Player_gameOver_Params
	Results Player_gameOver_Results
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

type Player_gameOver_Params struct{ capnp.Struct }

// Player_gameOver_Params_TypeID is the unique identifier for the type Player_gameOver_Params.
const Player_gameOver_Params_TypeID = 0xd3cfce73dc30fef5

func NewPlayer_gameOver_Params(s *capnp.Segment) (Player_gameOver_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_gameOver_Params{st}, err
}

func NewRootPlayer_gameOver_Params(s *capnp.Segment) (Player_gameOver_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_gameOver_Params{st}, err
}

func ReadRootPlayer_gameOver_Params(msg *capnp.Message) (Player_gameOver_Params, error) {
	root, err := msg.RootPtr()
	return Player_gameOver_Params{root.Struct()}, err
}

func (s Player_gameOver_Params) String() string {
	str, _ := text.Marshal(0xd3cfce73dc30fef5, s.Struct)
	return str
}

// Player_gameOver_Params_List is a list of Player_gameOver_Params.
type Player_gameOver_Params_List struct{ capnp.List }

// NewPlayer_gameOver_Params creates a new list of Player_gameOver_Params.
func NewPlayer_gameOver_Params_List(s *capnp.Segment, sz int32) (Player_gameOver_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Player_gameOver_Params_List{l}, err
}

func (s Player_gameOver_Params_List) At(i int) Player_gameOver_Params {
	return Player_gameOver_Params{s.List.Struct(i)}
}

func (s Player_gameOver_Params_List) Set(i int, v Player_gameOver_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Player_gameOver_Params_Promise is a wrapper for a Player_gameOver_Params promised by a client call.
type Player_gameOver_Params_Promise struct{ *capnp.Pipeline }

func (p Player_gameOver_Params_Promise) Struct() (Player_gameOver_Params, error) {
	s, err := p.Pipeline.Struct()
	return Player_gameOver_Params{s}, err
}

type Player_gameOver_Results struct{ capnp.Struct }

// Player_gameOver_Results_TypeID is the unique identifier for the type Player_gameOver_Results.
const Player_gameOver_Results_TypeID = 0x9c8f2cdfe0da3d04

func NewPlayer_gameOver_Results(s *capnp.Segment) (Player_gameOver_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_gameOver_Results{st}, err
}

func NewRootPlayer_gameOver_Results(s *capnp.Segment) (Player_gameOver_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Player_gameOver_Results{st}, err
}

func ReadRootPlayer_gameOver_Results(msg *capnp.Message) (Player_gameOver_Results, error) {
	root, err := msg.RootPtr()
	return Player_gameOver_Results{root.Struct()}, err
}

func (s Player_gameOver_Results) String() string {
	str, _ := text.Marshal(0x9c8f2cdfe0da3d04, s.Struct)
	return str
}

// Player_gameOver_Results_List is a list of Player_gameOver_Results.
type Player_gameOver_Results_List struct{ capnp.List }

// NewPlayer_gameOver_Results creates a new list of Player_gameOver_Results.
func NewPlayer_gameOver_Results_List(s *capnp.Segment, sz int32) (Player_gameOver_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Player_gameOver_Results_List{l}, err
}

func (s Player_gameOver_Results_List) At(i int) Player_gameOver_Results {
	return Player_gameOver_Results{s.List.Struct(i)}
}

func (s Player_gameOver_Results_List) Set(i int, v Player_gameOver_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Player_gameOver_Results_Promise is a wrapper for a Player_gameOver_Results promised by a client call.
type Player_gameOver_Results_Promise struct{ *capnp.Pipeline }

func (p Player_gameOver_Results_Promise) Struct() (Player_gameOver_Results, error) {
	s, err := p.Pipeline.Struct()
	return Player_gameOver_Results{s}, err
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

func (s NewTileRequest) Letters() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s NewTileRequest) HasLetters() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NewTileRequest) SetLetters(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewLetters sets the letters field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s NewTileRequest) NewLetters(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s NewTileRequest) Peelers() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s NewTileRequest) HasPeelers() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s NewTileRequest) SetPeelers(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewPeelers sets the peelers field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s NewTileRequest) NewPeelers(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
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
	PeelResponse_Status_success        PeelResponse_Status = 0
	PeelResponse_Status_invalidWord    PeelResponse_Status = 1
	PeelResponse_Status_detachedBoard  PeelResponse_Status = 2
	PeelResponse_Status_notAllLetters  PeelResponse_Status = 3
	PeelResponse_Status_extraLetters   PeelResponse_Status = 4
	PeelResponse_Status_invalidBoard   PeelResponse_Status = 5
	PeelResponse_Status_gameNotStarted PeelResponse_Status = 6
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
	case PeelResponse_Status_gameNotStarted:
		return "gameNotStarted"

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
	case "gameNotStarted":
		return PeelResponse_Status_gameNotStarted

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
	DumpResponse_Status_malformedRequest DumpResponse_Status = 3
	DumpResponse_Status_gameNotStarted   DumpResponse_Status = 4
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
	case DumpResponse_Status_malformedRequest:
		return "malformedRequest"
	case DumpResponse_Status_gameNotStarted:
		return "gameNotStarted"

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
	case "malformedRequest":
		return DumpResponse_Status_malformedRequest
	case "gameNotStarted":
		return DumpResponse_Status_gameNotStarted

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

type Peel struct{ capnp.Struct }

// Peel_TypeID is the unique identifier for the type Peel.
const Peel_TypeID = 0xcc92cfb307b2d7cf

func NewPeel(s *capnp.Segment) (Peel, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Peel{st}, err
}

func NewRootPeel(s *capnp.Segment) (Peel, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Peel{st}, err
}

func ReadRootPeel(msg *capnp.Message) (Peel, error) {
	root, err := msg.RootPtr()
	return Peel{root.Struct()}, err
}

func (s Peel) String() string {
	str, _ := text.Marshal(0xcc92cfb307b2d7cf, s.Struct)
	return str
}

func (s Peel) ValidBoards() (Peel_PlayerBoard_List, error) {
	p, err := s.Struct.Ptr(0)
	return Peel_PlayerBoard_List{List: p.List()}, err
}

func (s Peel) HasValidBoards() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Peel) SetValidBoards(v Peel_PlayerBoard_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewValidBoards sets the validBoards field to a newly
// allocated Peel_PlayerBoard_List, preferring placement in s's segment.
func (s Peel) NewValidBoards(n int32) (Peel_PlayerBoard_List, error) {
	l, err := NewPeel_PlayerBoard_List(s.Struct.Segment(), n)
	if err != nil {
		return Peel_PlayerBoard_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Peel) NewTiles() (Peel_PlayerTile_List, error) {
	p, err := s.Struct.Ptr(1)
	return Peel_PlayerTile_List{List: p.List()}, err
}

func (s Peel) HasNewTiles() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Peel) SetNewTiles(v Peel_PlayerTile_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewNewTiles sets the newTiles field to a newly
// allocated Peel_PlayerTile_List, preferring placement in s's segment.
func (s Peel) NewNewTiles(n int32) (Peel_PlayerTile_List, error) {
	l, err := NewPeel_PlayerTile_List(s.Struct.Segment(), n)
	if err != nil {
		return Peel_PlayerTile_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Peel) Timestamp() uint64 {
	return s.Struct.Uint64(0)
}

func (s Peel) SetTimestamp(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Peel_List is a list of Peel.
type Peel_List struct{ capnp.List }

// NewPeel creates a new list of Peel.
func NewPeel_List(s *capnp.Segment, sz int32) (Peel_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Peel_List{l}, err
}

func (s Peel_List) At(i int) Peel { return Peel{s.List.Struct(i)} }

func (s Peel_List) Set(i int, v Peel) error { return s.List.SetStruct(i, v.Struct) }

// Peel_Promise is a wrapper for a Peel promised by a client call.
type Peel_Promise struct{ *capnp.Pipeline }

func (p Peel_Promise) Struct() (Peel, error) {
	s, err := p.Pipeline.Struct()
	return Peel{s}, err
}

type Peel_PlayerBoard struct{ capnp.Struct }

// Peel_PlayerBoard_TypeID is the unique identifier for the type Peel_PlayerBoard.
const Peel_PlayerBoard_TypeID = 0xf2aacc8390e9079d

func NewPeel_PlayerBoard(s *capnp.Segment) (Peel_PlayerBoard, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Peel_PlayerBoard{st}, err
}

func NewRootPeel_PlayerBoard(s *capnp.Segment) (Peel_PlayerBoard, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Peel_PlayerBoard{st}, err
}

func ReadRootPeel_PlayerBoard(msg *capnp.Message) (Peel_PlayerBoard, error) {
	root, err := msg.RootPtr()
	return Peel_PlayerBoard{root.Struct()}, err
}

func (s Peel_PlayerBoard) String() string {
	str, _ := text.Marshal(0xf2aacc8390e9079d, s.Struct)
	return str
}

func (s Peel_PlayerBoard) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Peel_PlayerBoard) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Peel_PlayerBoard) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Peel_PlayerBoard) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Peel_PlayerBoard) Board() (Board, error) {
	p, err := s.Struct.Ptr(1)
	return Board{Struct: p.Struct()}, err
}

func (s Peel_PlayerBoard) HasBoard() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Peel_PlayerBoard) SetBoard(v Board) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewBoard sets the board field to a newly
// allocated Board struct, preferring placement in s's segment.
func (s Peel_PlayerBoard) NewBoard() (Board, error) {
	ss, err := NewBoard(s.Struct.Segment())
	if err != nil {
		return Board{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Peel_PlayerBoard_List is a list of Peel_PlayerBoard.
type Peel_PlayerBoard_List struct{ capnp.List }

// NewPeel_PlayerBoard creates a new list of Peel_PlayerBoard.
func NewPeel_PlayerBoard_List(s *capnp.Segment, sz int32) (Peel_PlayerBoard_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Peel_PlayerBoard_List{l}, err
}

func (s Peel_PlayerBoard_List) At(i int) Peel_PlayerBoard { return Peel_PlayerBoard{s.List.Struct(i)} }

func (s Peel_PlayerBoard_List) Set(i int, v Peel_PlayerBoard) error {
	return s.List.SetStruct(i, v.Struct)
}

// Peel_PlayerBoard_Promise is a wrapper for a Peel_PlayerBoard promised by a client call.
type Peel_PlayerBoard_Promise struct{ *capnp.Pipeline }

func (p Peel_PlayerBoard_Promise) Struct() (Peel_PlayerBoard, error) {
	s, err := p.Pipeline.Struct()
	return Peel_PlayerBoard{s}, err
}

func (p Peel_PlayerBoard_Promise) Board() Board_Promise {
	return Board_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Peel_PlayerTile struct{ capnp.Struct }

// Peel_PlayerTile_TypeID is the unique identifier for the type Peel_PlayerTile.
const Peel_PlayerTile_TypeID = 0x848d7f1593ddbc01

func NewPeel_PlayerTile(s *capnp.Segment) (Peel_PlayerTile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Peel_PlayerTile{st}, err
}

func NewRootPeel_PlayerTile(s *capnp.Segment) (Peel_PlayerTile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Peel_PlayerTile{st}, err
}

func ReadRootPeel_PlayerTile(msg *capnp.Message) (Peel_PlayerTile, error) {
	root, err := msg.RootPtr()
	return Peel_PlayerTile{root.Struct()}, err
}

func (s Peel_PlayerTile) String() string {
	str, _ := text.Marshal(0x848d7f1593ddbc01, s.Struct)
	return str
}

func (s Peel_PlayerTile) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Peel_PlayerTile) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Peel_PlayerTile) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Peel_PlayerTile) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Peel_PlayerTile) Letters() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Peel_PlayerTile) HasLetters() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Peel_PlayerTile) SetLetters(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewLetters sets the letters field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Peel_PlayerTile) NewLetters(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Peel_PlayerTile_List is a list of Peel_PlayerTile.
type Peel_PlayerTile_List struct{ capnp.List }

// NewPeel_PlayerTile creates a new list of Peel_PlayerTile.
func NewPeel_PlayerTile_List(s *capnp.Segment, sz int32) (Peel_PlayerTile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Peel_PlayerTile_List{l}, err
}

func (s Peel_PlayerTile_List) At(i int) Peel_PlayerTile { return Peel_PlayerTile{s.List.Struct(i)} }

func (s Peel_PlayerTile_List) Set(i int, v Peel_PlayerTile) error {
	return s.List.SetStruct(i, v.Struct)
}

// Peel_PlayerTile_Promise is a wrapper for a Peel_PlayerTile promised by a client call.
type Peel_PlayerTile_Promise struct{ *capnp.Pipeline }

func (p Peel_PlayerTile_Promise) Struct() (Peel_PlayerTile, error) {
	s, err := p.Pipeline.Struct()
	return Peel_PlayerTile{s}, err
}

type Dump struct{ capnp.Struct }

// Dump_TypeID is the unique identifier for the type Dump.
const Dump_TypeID = 0xfec3c6a7865ad9d1

func NewDump(s *capnp.Segment) (Dump, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Dump{st}, err
}

func NewRootDump(s *capnp.Segment) (Dump, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Dump{st}, err
}

func ReadRootDump(msg *capnp.Message) (Dump, error) {
	root, err := msg.RootPtr()
	return Dump{root.Struct()}, err
}

func (s Dump) String() string {
	str, _ := text.Marshal(0xfec3c6a7865ad9d1, s.Struct)
	return str
}

func (s Dump) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Dump) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Dump) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Dump) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Dump) Dump() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Dump) HasDump() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Dump) SetDump(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewDump sets the dump field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Dump) NewDump(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Dump) Timestamp() uint64 {
	return s.Struct.Uint64(0)
}

func (s Dump) SetTimestamp(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Dump_List is a list of Dump.
type Dump_List struct{ capnp.List }

// NewDump creates a new list of Dump.
func NewDump_List(s *capnp.Segment, sz int32) (Dump_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Dump_List{l}, err
}

func (s Dump_List) At(i int) Dump { return Dump{s.List.Struct(i)} }

func (s Dump_List) Set(i int, v Dump) error { return s.List.SetStruct(i, v.Struct) }

// Dump_Promise is a wrapper for a Dump promised by a client call.
type Dump_Promise struct{ *capnp.Pipeline }

func (p Dump_Promise) Struct() (Dump, error) {
	s, err := p.Pipeline.Struct()
	return Dump{s}, err
}

type Replay struct{ capnp.Struct }

// Replay_TypeID is the unique identifier for the type Replay.
const Replay_TypeID = 0x87c008d7bed714b1

func NewReplay(s *capnp.Segment) (Replay, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return Replay{st}, err
}

func NewRootReplay(s *capnp.Segment) (Replay, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return Replay{st}, err
}

func ReadRootReplay(msg *capnp.Message) (Replay, error) {
	root, err := msg.RootPtr()
	return Replay{root.Struct()}, err
}

func (s Replay) String() string {
	str, _ := text.Marshal(0x87c008d7bed714b1, s.Struct)
	return str
}

func (s Replay) InitialTiles() (Replay_TileFreq_List, error) {
	p, err := s.Struct.Ptr(0)
	return Replay_TileFreq_List{List: p.List()}, err
}

func (s Replay) HasInitialTiles() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Replay) SetInitialTiles(v Replay_TileFreq_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewInitialTiles sets the initialTiles field to a newly
// allocated Replay_TileFreq_List, preferring placement in s's segment.
func (s Replay) NewInitialTiles(n int32) (Replay_TileFreq_List, error) {
	l, err := NewReplay_TileFreq_List(s.Struct.Segment(), n)
	if err != nil {
		return Replay_TileFreq_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Replay) Peels() (Peel_List, error) {
	p, err := s.Struct.Ptr(1)
	return Peel_List{List: p.List()}, err
}

func (s Replay) HasPeels() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Replay) SetPeels(v Peel_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewPeels sets the peels field to a newly
// allocated Peel_List, preferring placement in s's segment.
func (s Replay) NewPeels(n int32) (Peel_List, error) {
	l, err := NewPeel_List(s.Struct.Segment(), n)
	if err != nil {
		return Peel_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Replay) Dumps() (Dump_List, error) {
	p, err := s.Struct.Ptr(2)
	return Dump_List{List: p.List()}, err
}

func (s Replay) HasDumps() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Replay) SetDumps(v Dump_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewDumps sets the dumps field to a newly
// allocated Dump_List, preferring placement in s's segment.
func (s Replay) NewDumps(n int32) (Dump_List, error) {
	l, err := NewDump_List(s.Struct.Segment(), n)
	if err != nil {
		return Dump_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Replay) StartTime() uint64 {
	return s.Struct.Uint64(0)
}

func (s Replay) SetStartTime(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Replay) EndTime() uint64 {
	return s.Struct.Uint64(8)
}

func (s Replay) SetEndTime(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Replay) FinalScores() (Replay_Score_List, error) {
	p, err := s.Struct.Ptr(3)
	return Replay_Score_List{List: p.List()}, err
}

func (s Replay) HasFinalScores() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Replay) SetFinalScores(v Replay_Score_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewFinalScores sets the finalScores field to a newly
// allocated Replay_Score_List, preferring placement in s's segment.
func (s Replay) NewFinalScores(n int32) (Replay_Score_List, error) {
	l, err := NewReplay_Score_List(s.Struct.Segment(), n)
	if err != nil {
		return Replay_Score_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

// Replay_List is a list of Replay.
type Replay_List struct{ capnp.List }

// NewReplay creates a new list of Replay.
func NewReplay_List(s *capnp.Segment, sz int32) (Replay_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4}, sz)
	return Replay_List{l}, err
}

func (s Replay_List) At(i int) Replay { return Replay{s.List.Struct(i)} }

func (s Replay_List) Set(i int, v Replay) error { return s.List.SetStruct(i, v.Struct) }

// Replay_Promise is a wrapper for a Replay promised by a client call.
type Replay_Promise struct{ *capnp.Pipeline }

func (p Replay_Promise) Struct() (Replay, error) {
	s, err := p.Pipeline.Struct()
	return Replay{s}, err
}

type Replay_TileFreq struct{ capnp.Struct }

// Replay_TileFreq_TypeID is the unique identifier for the type Replay_TileFreq.
const Replay_TileFreq_TypeID = 0xe2e3f234b6ca5d2e

func NewReplay_TileFreq(s *capnp.Segment) (Replay_TileFreq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Replay_TileFreq{st}, err
}

func NewRootReplay_TileFreq(s *capnp.Segment) (Replay_TileFreq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Replay_TileFreq{st}, err
}

func ReadRootReplay_TileFreq(msg *capnp.Message) (Replay_TileFreq, error) {
	root, err := msg.RootPtr()
	return Replay_TileFreq{root.Struct()}, err
}

func (s Replay_TileFreq) String() string {
	str, _ := text.Marshal(0xe2e3f234b6ca5d2e, s.Struct)
	return str
}

func (s Replay_TileFreq) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Replay_TileFreq) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Replay_TileFreq) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Replay_TileFreq) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Replay_TileFreq) Frequency() (capnp.Int32List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Int32List{List: p.List()}, err
}

func (s Replay_TileFreq) HasFrequency() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Replay_TileFreq) SetFrequency(v capnp.Int32List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewFrequency sets the frequency field to a newly
// allocated capnp.Int32List, preferring placement in s's segment.
func (s Replay_TileFreq) NewFrequency(n int32) (capnp.Int32List, error) {
	l, err := capnp.NewInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int32List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Replay_TileFreq_List is a list of Replay_TileFreq.
type Replay_TileFreq_List struct{ capnp.List }

// NewReplay_TileFreq creates a new list of Replay_TileFreq.
func NewReplay_TileFreq_List(s *capnp.Segment, sz int32) (Replay_TileFreq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Replay_TileFreq_List{l}, err
}

func (s Replay_TileFreq_List) At(i int) Replay_TileFreq { return Replay_TileFreq{s.List.Struct(i)} }

func (s Replay_TileFreq_List) Set(i int, v Replay_TileFreq) error {
	return s.List.SetStruct(i, v.Struct)
}

// Replay_TileFreq_Promise is a wrapper for a Replay_TileFreq promised by a client call.
type Replay_TileFreq_Promise struct{ *capnp.Pipeline }

func (p Replay_TileFreq_Promise) Struct() (Replay_TileFreq, error) {
	s, err := p.Pipeline.Struct()
	return Replay_TileFreq{s}, err
}

type Replay_Score struct{ capnp.Struct }

// Replay_Score_TypeID is the unique identifier for the type Replay_Score.
const Replay_Score_TypeID = 0xf439995c3c2921c0

func NewReplay_Score(s *capnp.Segment) (Replay_Score, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Replay_Score{st}, err
}

func NewRootReplay_Score(s *capnp.Segment) (Replay_Score, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Replay_Score{st}, err
}

func ReadRootReplay_Score(msg *capnp.Message) (Replay_Score, error) {
	root, err := msg.RootPtr()
	return Replay_Score{root.Struct()}, err
}

func (s Replay_Score) String() string {
	str, _ := text.Marshal(0xf439995c3c2921c0, s.Struct)
	return str
}

func (s Replay_Score) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Replay_Score) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Replay_Score) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Replay_Score) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Replay_Score) Score() uint32 {
	return s.Struct.Uint32(0)
}

func (s Replay_Score) SetScore(v uint32) {
	s.Struct.SetUint32(0, v)
}

// Replay_Score_List is a list of Replay_Score.
type Replay_Score_List struct{ capnp.List }

// NewReplay_Score creates a new list of Replay_Score.
func NewReplay_Score_List(s *capnp.Segment, sz int32) (Replay_Score_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Replay_Score_List{l}, err
}

func (s Replay_Score_List) At(i int) Replay_Score { return Replay_Score{s.List.Struct(i)} }

func (s Replay_Score_List) Set(i int, v Replay_Score) error { return s.List.SetStruct(i, v.Struct) }

// Replay_Score_Promise is a wrapper for a Replay_Score promised by a client call.
type Replay_Score_Promise struct{ *capnp.Pipeline }

func (p Replay_Score_Promise) Struct() (Replay_Score, error) {
	s, err := p.Pipeline.Struct()
	return Replay_Score{s}, err
}

const schema_f2768a70df34b76a = "x\xda\x8cX\x7fl\x1b\xf5\x15\x7f\xef\xce\xce\xd9I\x1c" +
	"\xfb\xcb\xa5+HcYQ\x11$\"Q\x93\xc04\xb2" +
	"eNCKi\x15\xda\\\x12V\x11\xc1\xc4\xc5\xfe\x86" +
	"\xba\xd8w\xce\xdd9M: c\xd0\x01\x95\xd8\x18\x9b" +
	"\xa6v\"\xfb!\xd4\xb1e\x93\x10m\x07m\xb5\xa2V" +
	"\xa8\x88\xa2\x86\xb6Z\xab\xb6\xa8\x15t\xa3\x1at\x02\xd1" +
	"\xae\xfd\xa3\xd5\xda\x9b\xde\x9d\xef|q\x9c5\x7f\xc5\xbe" +
	"{\xdf\xef{\xef\xf3>\xef\xf3\x9e\xb3\xe4\xcbp\x97\xd0" +
	"\x1a~\xbb\x06@y!\\e\xe3\xdf\xce\xfcb\xc1\xc4" +
	"K\xcf\x01[\x80\xf6\xe1\x13\xdb\xa5\x1d\x87_9\x04a" +
	"A\x02h\xdf\x1f\x16P\x9e\x0eK\x00\xf2\xc1\xf0\x06@" +
	"\xfb\xfdo\xd4\x9f;\xf6\x94\xf5<\xb0\x85\x08\x10\"\x9b" +
	"{\xaa\x9a\x10B\xf6\x9b\xf5'\xde9\x11\xd9\xf7<(" +
	"\x09\x14\xec\xf5o\xdf\xfdq~\xf3\xe8E\x08\x93\x8d|" +
	"k\xd5\x1e\xf9\xf6\xaa;\x00\xda;\xab\xd6\"\xa0\xbdh" +
	"\xf0\x1d\xfe\xac5\xf4\xd3\xc0=\xc7\xa5\x0e\xbag\xe3\x8b" +
	"\xd5\xf5k\xef\xbc\xb0\x05\xd8\x02\xc1\x8e$\x86\x1fz\xe8" +
	"\xb5\xe9O\x00\xb0}\xafT\x8d\xf2\xb4\xe4\xc4\"\xad\x01" +
	"\xb4C\x9d\x1f}\xf2\xf1]?{5p\xc79\xa9\x9b" +
	"\xee\xe8\x1c~&\xa6\xee\x88n\x03\x96\x10K\xa1\x00\x1d" +
	"\xdf#\x1fw\xae8*\xad\x90\xaf\xd0'\xfb\xda\xa9\x91" +
	"/\xd7\xee7\xb6\x01\xbbU\xb0\x8f~\xed\xd0{\x07\xbf" +
	"\xf5\xe0\xe7\xe4\xf0\xact\x1b\xca\x17\x1c\xeb/\xa47\x00" +
	"\xed5O\xff`I\xdd\x0f\xef\xfb\x03\xa5\x88\x81\x14\x91" +
	"L\x1e\x8e\x9c\x97yd!\x80\x9c\x8b\x10R=+\x9f" +
	"|\xd9\x18\xbb\xf4\xfa\xac\x18NE\xa6\xe4\xb3\x11:r" +
	"&\xf2\xbc\xdc\x1c\xa5\x18Z\x87\xbe\xff\xca\x91\xdf\xee\xfa" +
	"c \x95\x05\xd1>J\xe5\x85W\xeb~\xf3\xe3tz" +
	";\xb0D\xd0%UG\xc6\xe8U9F\xe7\xe5ht" +
	"\x03\xe0\xb5-_\xfd*\xb9\xe2{\xfb\xca\x1d\xf2\xe8\x1e" +
	"9Gf\xed\x99\xa8\x842\xab&\x8f\xe3=\xb5\xab\xbe" +
	"\xda\xfa\xcbw+\xdd{%z^\x0eW;\x1e\xaa)" +
	"\x15\x9f\x14ey;\xc6OUO\xc9\x9b\xaa\xef\x00\x90" +
	"\x7f_M \x89W~\xfdgyz\xcb\xe1J \xb5" +
	"\xf3\x1a\x01\xe5\x91\x1aBi\xbc\x86\xae\xbeE\xd9q\xdf" +
	"\x7f.M\x1e!\xf2A\xd1\xe8xM5\x02\xcagj" +
	"\x92\x80\xf6\xe5\xebKN\x9b\x1f\x1e\xfe{\x00\x99k5" +
	"\x0eQZ\xc6v=0\xf4Z\xe2XY\x06\xa2\xc3\x83" +
	"\x9a\x9bP\xbe\\C\xf1]\xa8\xa1\xa8^\x91\x8eM>" +
	"]?u\xb2R\xba\x93\xb5W\xe5?\xd5\xd2\xa7m\xb5" +
	"\x14SK\xab\xb5\xf3\xd0\x87\x9f\x9d\x04\xf6M\xa1\x94\x0e" +
	"`\xfb\xa2X\x07\xca\xad1\xb2l\x8eM\x00\xda\xb1\xc7" +
	"N\x1c\xcdL>\xf1\x91C\x1d\x9f\x1d\x80\xed\x85\xd8m" +
	"(or,\x9fq,}\"\x97u\x87C\x9d\xe9\xd8" +
	"\x94|<F\xa0\x9c\x8d\xfd\x8b\x02x\xf4\x83\xb7\xee\xbe" +
	"\xf8\xcf\x7fP\xce\xa5\xber;rw\x9d\x80\xf2\x81:" +
	":\xb6\xbf\x8e\xa2\xf5\x19;\x13\xef\xe5(\x85\x10\xe5\xd6" +
	"\xf8y\xb93\xbe\x10\xa0}y\xbc\x81\x1ao=\xdb{" +
	"\xf3\x89\xa9\xb7\xbe\x08\x02\xfe\xf3\x84\x03\xf8\xd6\x04\x01>" +
	")}\xfe\xf2\xb3\x87\xa6.V\x92\x83\xdd\x89j\x94\x0f" +
	"&\xc8\xf9\x81\x049\xdf\xb7\xa8\xf1\xbb\x8fl\xbd\xf7\x12" +
	"(\x0b1\x18\xaa\x93\xd6\"v^nf\xf4\xa9\x91\x91" +
	"\xf1\xbf\xa7/\xb3\x0f\xc6\xbes\xb5\xac\x06\x8e\xed&\xb6" +
	"]~\xc9\xb1}\x91Q\x14GO\x0d\xfe\xe4\xf5\xf7\xde" +
	"\xbd^\x91r\xbb\xd9\x94\xbc\xdf1\xde\xcb\xde\x80f;" +
	"\xaf[\xaaif\x0ab\xae%\xa5\xe6\xb5|G/\xe7" +
	"\xd9\x96\xde\xac:\xce\x8d\x81L\x96C/\xa2\x12\x11C" +
	"\x00!\x04`\x8d\x1d\x00\xcab\x11\x95%\x022\xc4z" +
	"\xa4\x87\xcd\xdd\x00\xca\x9d\"*\xcb\x04L\xe6\x9d\xa3X" +
	"\x0b\x02\xd6\x02Nd\xb9eq\xc3\xc4:\xc0^\x11\x9d" +
	"\xc7u\x80\x15\xdc:\xc7Z\xcc|6c-\xee\xe3f" +
	"\xbc\x90\xb5L\xdf\x0c=\xb3d\x1f'\x07J\x04\x83\xa5" +
	"\x8e\xae*\xa1\xc9\xa2m6\x05~\xbf\xc1G\x00\xa0\xa1" +
	"?\xa5\x1b\\\xb9\xd9\xcf`\xebz\x00e\x8b\x88\xca\xce" +
	"@\x06o\xb6\x01(\x7f\x11Qy_@&\x08\xf5(" +
	"\x00\xb0\x03\xf4p\x9f\x88\xcai\x01Q\xacG\x11\x80\x9d" +
	"\xea\x03PN\x8a\xa8|* \x0ba=\x86\x00\xd8Y" +
	"\xca\xff\xb4\x88\xcag\x02\xb2\xb0X\x8fa\x00vn\x08" +
	"@\xf9T\xc4>\x14\xd0\xceh\x19+\xa3f\x07 \x9e" +
	"\xc9r\x1f\x8cD)\x05@z\xd8\x90\xe7<\x1bx\xed" +
	"S\xa8\xf8:]\xc8\xe5\x03\xaf\xfdB\xbb\xafm\xd3R" +
	"\x0dk \x93\x03\xe4\x18\x05\x01\xa3\x80\x13\\K\x0fd" +
	"r\xfew{8\xa3\xa9\xd9\xfe\x94\x0e\x92\x11\x0c\xc4\x87" +
	"\xafx\xd5\\\xf5\xd1\xf8\x06\x02wq_\x92\x9b3*" +
	"\xe4\x1b\xae\xd5\x8dt\xcb\x1a#\xc35K\xb52\xba\x06" +
	"\x0e\x83j\x1dDo\xed\xa6\xfb\xd9\x82A\x00\x14\x18[" +
	"\x050Q\xd0\x9e\xd0\xf4\x0d\x9a\xbdN72\x1bu\xcd" +
	"\x02Q\xcd\xda\xa3\xdc\xb02)5\x0b\x00\xbe\x83\xaa\xb2" +
	"H\x1eWs|\xcd(7\x88,\x14\x09x\x86>U" +
	"\xfa\xb91\xca\x0d\xf2\x1e\x12\xc3\x00\xfe8@O\x97\x18" +
	"\xeb\x06\x81\x85\xa5\x89\x94\xaei<eua/b\xe5" +
	"\x8e\xe8\xe3f^\xd7L\xde\xd2o\xa9\x92U0\xe9\xd6" +
	"\xaf;9=\xe8\xe6\xb4|\xc8\xc9i\xa9\x01\x80\"\xeb" +
	"\xa4?!v\xefz\x00\x0c\xb3{\xe8O\x15k\xdd\x08" +
	"0a\x16R)n\x9avF\x1bU\xb3\x99\xf4Z\x90" +
	"t#m\xa7\xb9\xa5\xa6\xd6\xf144t\xeb\xaa\x91\xb6" +
	"5\xddZ\x9a\xcd\xf6php\xba\xc7\xe6c\x96\xa1\xf6" +
	"p\x0b\xe2\xce\xd7\xe2\xe1n\x88;\xd6\x04\xc5j\xdd\xea" +
	"\x87$\x11\x80\xa7\xfd\x14\x04/\x85e\x85\\\x9eR\x88" +
	"S\x0eJ\x08\x83\xf2\x8b\x1d\xc9~K\xb5\x0a\xe6\x1cm" +
	"^\xb1\xcbM\xe7\x04\xc6K\xf7\x00b\xfc\xc6\xfd\xee7" +
	"r|\x85\x9a\xe3\xae\xb6Pm<iEO\x92Yk" +
	"\x13\x08\xecv\x09Ks\x0e\xbd)\xc1n\xa1w1)" +
	"N\xed\xd2\x85qj\x8b\x99\xb5\x0b\x95\x91\x85,V\xeb" +
	"V&\xc5\x8btAs6H\xf7\xb9,\xe8K\xf2\x91" +
	"\x027\xad2\xe1k\xaa$|\x1dEH\xee\x160\xae" +
	"\xa99\xee\xc9\x9e\xa7\x82\xcc\xdb+\x88\"\x95`H\xba" +
	"\x01\x92\xaf\x84\x03\x84\xb7[\xa0\xb7-\xb2\x916\x10\x18" +
	"' \xbcA\x8c\xde\x06\xc8\x1e&\x02?(\xa1\xe0O" +
	"t\xf4\xd6!\xb6t\x10\x04v\xaf\x84\xa2\xbf\x07\xa0\xb7" +
	"\xf5\xb1\xe6U\x0e\xb8\x0d\x8e\xdcv\xe1D\xb1\xad\xbb\xd0" +
	"\xf6\x90\x021E_\xbd.\x03\x80\x99\x08\xfb\xa8\xf5\xd3" +
	"\x15}|$^\x01\xb3\xb6\"f]\x01\xcc:\x89F" +
	"\xdf\x16Q\x19\x10\xb0\xc1\x0a\xcaa\x91+\x13.v\xf3" +
	"\xa0\x105\xa63\x09Js7:\x14\xd8\xc9\xa3\x83\xb6" +
	"\x0bo\xb7\x0e\x125\x8a7\xd6@\xccr\xa5\xd6\x8f\x93" +
	"zWY&\xa2\xf2X \xceGW\x01(\x8f\x88\xa8" +
	"\x8c\x09\x88\xc5\x89P \xf5\xb7DT~$\xa0\xed\xb6" +
	"\xa0\xae\x82d\xa4\x03J\xea\xc7RT\xd2\"\xb4&\x00" +
	"\x94\x8c\xfc\x18\x8bFV&\xc7MK\xcd\x01\xe6}\xa5" +
	"\x9e\xa5C\x1e?\x8bR\x04N\x1f\x97\x16\xaey\xf7q" +
	"S\x89\xb4\xa5>\xf6\xefq\xfb8N\x85GV\xda\xc4" +
	"\xcb\x08\xec\x07Em\xec\xb4\xd7\xe2^\xd5Ps&\x80" +
	"\x12\xf2\xdd\xc7\xc8}DD\xa5^\xc0\xa4\xab\x0c^\x87" +
	"\xcc9a|]O\xba\x17\xce6\\\xe67s\x9f\xd3" +
	"\xa9\xe8\xd0.PN\xf2\xda%\xa2\xd2\x13(\xe7J*" +
	"\xe7\x03\"*\xe9\xc0\x84W;\x8a5^'`\x92\x92" +
	"\xe0i?>\x83\xa7xf\x94\xa7\x03u+2\xd1\xb5" +
	"\x9c\x9d\x89\xdf\x12\xab\xdd\x92\xcf!$\xdd7h\x8a9" +
	"$t\x82\xc4n>\x9bT\x19MZ\x8a\xac\xa0(\xdc" +
	"\xac\xefq'V\xb3;\xb1\x1a7;\x13\xabq\xd0\x99" +
	"X\xb7\xaf\x0f\xcc(R\xb4\x01]\xef\x01I\xd7\x1ew" +
	"\xbe-\xcd\x1a\x1c\xd5\xf4\xf8\x80\xfa\x04\xd7\x00\x1cyX" +
	"i\xde_\x001\x9b\xb5\xc7\xd5e\xba\xc6W\xe8\x10\xd7" +
	"\x87\x03CH,\x1fB\xe5stfT\x9b\x9d\xa8\x9a" +
	"7\xbaQmv\xa2j\x0cNN\x17\x9e\xd5:Z+" +
	"\xb5bg\xd1\xbc\\\xae\xe9\x85\xc7!\xb9\xceyd\xe7" +
	"\xd4\xec\xb0n\xe48\xa6]\x8eX\xc5`+\x8e\xc9\x92" +
	"\xa2\xd0\xf2\xe2\xb4U\xe9\x974\x0e\xd9\xde6\x03RF" +
	"\xd7\x94\x84_J\xb5\xa9D\x1f\xaf\x92\x9c\xb4$-\xa2" +
	"\x92'\x9e\xa1\x9b[\xee&\x00e\x9d\x88\x8a% \x13" +
	"\x05w\x95\x1c\xa1\x87YW`\xe2\x16\x1f\xb3|B\xe9" +
	"A\x7f\x18/\x05\xe3\xf6&\x8ea\x04\x04\x8c\x00\xe2\xb8" +
	"\xf7i6\xda\xee\xa2\xdc\xe2\xaf\xc2\xf3\xd9\xe3I\xde\xee" +
	"\x12Qy`\xd6\x1eo\x0f\x1b\x84\xa3\x96\x02\x1c\xf7\x18" +
	"\x18*c\xa0P\xbe0\x95\xb6\x8d\xd2\xff\x09J*U" +
	"\xef\x07\xf3\x14\x053&\xa2\xf2\x9c\x801\xb4m7\x9a" +
	"ghz<)\xa2\xf2\x82\x801\xe1\xba\xedB\xb9\x89" +
	"\xf0}ND\xe5w\x02\xc6\xc4k\xb6\x8b\xe5\xa4\x01\xa0" +
	"\xbc*\xa2\xb2K\xc0X\xe8\xbf\xb6\xbb\x97\xff\x95V\xfd" +
	"\x9d\"*G\x82J\xe7\xc7\xe2\xa2\xd9\xe0\xe88T\xcd" +
	"\xdc\xc8\xca\xbb\xacl!+\x7f=s?\xbba\x8b:" +
	"\xa2I\xfd\\Q4\xdbJ\xa2\xd90D\x0b\x1e&J" +
	"\xbf\xff\x001\x01\xf8\x7f\x7f\xb69+$\xcc\xab\xe0m" +
	"\x81Q0\xb3\xe07\xf4,\x94\x11\xad?N?\xb4\xe6" +
	"v:\x1f\x9f&]1\x8b\xd2~{\xba\xcb\xb1\xbb\xd0" +
	"\xcfBk\xb1\x80\x0d\x1b\xf4\x19\xa38\xf0\x1f\xb096" +
	"\x08\x92\xa4\xf9\x0c\x8f\xa6\xd2\x82\xe0\xed\x02\x8f\xf6\x05f" +
	"\xc7\xcc<\x9c\x85\xb4\x9c\x04\x15F\xfc\xff\x02\x00\x00\xff" +
	"\xff\x0e\x9e\x93C"

func init() {
	schemas.Register(schema_f2768a70df34b76a,
		0x848d7f1593ddbc01,
		0x87747dd4e5141ec8,
		0x87c008d7bed714b1,
		0x8e62748365be5a21,
		0x98f12857140a897a,
		0x9c8f2cdfe0da3d04,
		0xa509b3610d81663d,
		0xa572c157ee71d9fd,
		0xa6437b0e305e7e4f,
		0xa7f47872907c494c,
		0xa8b89fd092566231,
		0xb26464829e0e9c88,
		0xc03e473f96f00098,
		0xc39499f04a0c4c79,
		0xcc92cfb307b2d7cf,
		0xcf98cd13ab9af903,
		0xd09df4f343b35119,
		0xd3cfce73dc30fef5,
		0xd410a36248b8782e,
		0xd8aa147e9dd40792,
		0xd8e8ceccb474312e,
		0xda6b9d69d1d7600d,
		0xe0cda35555661008,
		0xe2e3f234b6ca5d2e,
		0xe94d36c9c6cc16d1,
		0xedb6aad718bd116a,
		0xf2aacc8390e9079d,
		0xf439995c3c2921c0,
		0xfa3b78ca11f5cdeb,
		0xfec3c6a7865ad9d1)
}
