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
	DumpResponse_Status_malformedRequest DumpResponse_Status = 3
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Peel{st}, err
}

func NewRootPeel(s *capnp.Segment) (Peel, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
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

func (s Peel) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Peel) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Peel) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Peel) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Peel) Board() (Board, error) {
	p, err := s.Struct.Ptr(1)
	return Board{Struct: p.Struct()}, err
}

func (s Peel) HasBoard() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Peel) SetBoard(v Board) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewBoard sets the board field to a newly
// allocated Board struct, preferring placement in s's segment.
func (s Peel) NewBoard() (Board, error) {
	ss, err := NewBoard(s.Struct.Segment())
	if err != nil {
		return Board{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Peel) NewTiles() (Peel_Entry_List, error) {
	p, err := s.Struct.Ptr(2)
	return Peel_Entry_List{List: p.List()}, err
}

func (s Peel) HasNewTiles() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Peel) SetNewTiles(v Peel_Entry_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewNewTiles sets the newTiles field to a newly
// allocated Peel_Entry_List, preferring placement in s's segment.
func (s Peel) NewNewTiles(n int32) (Peel_Entry_List, error) {
	l, err := NewPeel_Entry_List(s.Struct.Segment(), n)
	if err != nil {
		return Peel_Entry_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
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
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
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

func (p Peel_Promise) Board() Board_Promise {
	return Board_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Peel_Entry struct{ capnp.Struct }

// Peel_Entry_TypeID is the unique identifier for the type Peel_Entry.
const Peel_Entry_TypeID = 0x91d57b12a6c06883

func NewPeel_Entry(s *capnp.Segment) (Peel_Entry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Peel_Entry{st}, err
}

func NewRootPeel_Entry(s *capnp.Segment) (Peel_Entry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Peel_Entry{st}, err
}

func ReadRootPeel_Entry(msg *capnp.Message) (Peel_Entry, error) {
	root, err := msg.RootPtr()
	return Peel_Entry{root.Struct()}, err
}

func (s Peel_Entry) String() string {
	str, _ := text.Marshal(0x91d57b12a6c06883, s.Struct)
	return str
}

func (s Peel_Entry) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Peel_Entry) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Peel_Entry) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Peel_Entry) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Peel_Entry) Letter() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Peel_Entry) HasLetter() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Peel_Entry) LetterBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Peel_Entry) SetLetter(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// Peel_Entry_List is a list of Peel_Entry.
type Peel_Entry_List struct{ capnp.List }

// NewPeel_Entry creates a new list of Peel_Entry.
func NewPeel_Entry_List(s *capnp.Segment, sz int32) (Peel_Entry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Peel_Entry_List{l}, err
}

func (s Peel_Entry_List) At(i int) Peel_Entry { return Peel_Entry{s.List.Struct(i)} }

func (s Peel_Entry_List) Set(i int, v Peel_Entry) error { return s.List.SetStruct(i, v.Struct) }

// Peel_Entry_Promise is a wrapper for a Peel_Entry promised by a client call.
type Peel_Entry_Promise struct{ *capnp.Pipeline }

func (p Peel_Entry_Promise) Struct() (Peel_Entry, error) {
	s, err := p.Pipeline.Struct()
	return Peel_Entry{s}, err
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

func (s Dump) Dump() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Dump) HasDump() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Dump) DumpBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Dump) SetDump(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Replay{st}, err
}

func NewRootReplay(s *capnp.Segment) (Replay, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
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

func (s Replay) InitialTiles() (Replay_Entry_List, error) {
	p, err := s.Struct.Ptr(0)
	return Replay_Entry_List{List: p.List()}, err
}

func (s Replay) HasInitialTiles() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Replay) SetInitialTiles(v Replay_Entry_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewInitialTiles sets the initialTiles field to a newly
// allocated Replay_Entry_List, preferring placement in s's segment.
func (s Replay) NewInitialTiles(n int32) (Replay_Entry_List, error) {
	l, err := NewReplay_Entry_List(s.Struct.Segment(), n)
	if err != nil {
		return Replay_Entry_List{}, err
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

// Replay_List is a list of Replay.
type Replay_List struct{ capnp.List }

// NewReplay creates a new list of Replay.
func NewReplay_List(s *capnp.Segment, sz int32) (Replay_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
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

type Replay_Entry struct{ capnp.Struct }

// Replay_Entry_TypeID is the unique identifier for the type Replay_Entry.
const Replay_Entry_TypeID = 0xfe64d60234157407

func NewReplay_Entry(s *capnp.Segment) (Replay_Entry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Replay_Entry{st}, err
}

func NewRootReplay_Entry(s *capnp.Segment) (Replay_Entry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Replay_Entry{st}, err
}

func ReadRootReplay_Entry(msg *capnp.Message) (Replay_Entry, error) {
	root, err := msg.RootPtr()
	return Replay_Entry{root.Struct()}, err
}

func (s Replay_Entry) String() string {
	str, _ := text.Marshal(0xfe64d60234157407, s.Struct)
	return str
}

func (s Replay_Entry) Player() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Replay_Entry) HasPlayer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Replay_Entry) PlayerBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Replay_Entry) SetPlayer(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Replay_Entry) Frequency() (capnp.Int32List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Int32List{List: p.List()}, err
}

func (s Replay_Entry) HasFrequency() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Replay_Entry) SetFrequency(v capnp.Int32List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewFrequency sets the frequency field to a newly
// allocated capnp.Int32List, preferring placement in s's segment.
func (s Replay_Entry) NewFrequency(n int32) (capnp.Int32List, error) {
	l, err := capnp.NewInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int32List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Replay_Entry_List is a list of Replay_Entry.
type Replay_Entry_List struct{ capnp.List }

// NewReplay_Entry creates a new list of Replay_Entry.
func NewReplay_Entry_List(s *capnp.Segment, sz int32) (Replay_Entry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Replay_Entry_List{l}, err
}

func (s Replay_Entry_List) At(i int) Replay_Entry { return Replay_Entry{s.List.Struct(i)} }

func (s Replay_Entry_List) Set(i int, v Replay_Entry) error { return s.List.SetStruct(i, v.Struct) }

// Replay_Entry_Promise is a wrapper for a Replay_Entry promised by a client call.
type Replay_Entry_Promise struct{ *capnp.Pipeline }

func (p Replay_Entry_Promise) Struct() (Replay_Entry, error) {
	s, err := p.Pipeline.Struct()
	return Replay_Entry{s}, err
}

const schema_f2768a70df34b76a = "x\xda\x8cX\x7fl\x13\xf7\x15\x7f\xef\xce\xf69\x8e\x1d" +
	"\xfb\xb8d\xb4\x7f\xb0,\x15H%j\" 0\xadL" +
	"\xcc!\x90\xf1C\x01r\x01FA\xeb\xc4\xc5\xfeB\x0c" +
	"\xe7;sw\x0e1-\xcdX\xdbQ\xd0~\x81\xa6\x8d" +
	"N\xb0\x0dUmG&T\x01e\xa3HTL\x95\xaa" +
	"1A\x01\x8d.-\x9a\xd4\xfe3mc\xda\xd4m\xed" +
	"\x1fC\x83\x9b\xde\xf7\xec\xbb\x8b\xe3\x0c\xfe\xca\xd7w\xcf" +
	"\xdf\xf7\xde\xe7}\xde\xe7=g\xc1\x8ah\xaf\xb00*" +
	"5\x03\xa8\xdf\x8c\xc6\xdc\xdf~\xbe\xf5O\xb7\xf6;\x07" +
	"A\x9e\x8d\x00\x11\x09\xa0\xa7\x10\xedD\x88\xb8gZ'" +
	"\xdf\x9e\x8c_>\x08r\x06\xdd]\xbf^\xfcQ\xe9\xf0" +
	"\xe8\xbf *J\x00\x8a\x1a\xbd\xa8l\x8d\xce\x06P\xb4" +
	"\xe8\x1b\x80n\xc7\xb6\xb7\xd9\xf3\xce\xf0\xf7B\xb7t\xc5" +
	"\x96\xd2-\xcf\x8f\\~m\xd63\xef\x1f\x01\xb9\x0d\xdd" +
	"\xeb\x93g\xa5s\xd7\x8f^\x85\xa8@\xb7\xb4\xc5n+" +
	"\x1d1:\xcd\x89\xed\x05t\xf7\x1dJ\xb4ny\xfc\x9f" +
	"\xc7@n\x13\xdcxf\xc7\xe6\xcd\xaf\\\xfb\x18\x00{" +
	"*\xb1\x04*\x87\xb8\xe5\x8b\xb1\x0d\x80nd\xd9\xed\x8f" +
	"?z\xe2\xfb\xc7C\xfeN\xc6\xfa\xc8\xdf\xb2\x1d\x07R" +
	"\xda\xb9\xa6WA\xce\x88A\xd4@_\xbf\xa8\x1c\xe1W" +
	"|7\xb6J9O'\xf7\xde\x87{\xfe\xb1\xe57\xd6" +
	"\xab \xcf\x11\xdc\x9b\x9f\xbb\xfa\xee\x95/\xae\xfb+9" +
	"<\x11{\x0c\x95\xd3\xdc\xfaT\xec\xc7\x80\xee\x86\xe7\xbe" +
	"\xb1\xa0\xe5\x99\x15\xaf\x81\x9a\xc10\x1aH&(\xddQ" +
	"R\x12\xa1\xd1&Q\x1e\x03k\x9e\xfd\x815\xf6\xe9\xeb" +
	"\xd3b8 M(\x87$\x9e\x86tPy\x9fN\xee" +
	"\xc2\xe1\xaf\x1d\xbd\xf1\xb3\x0b\xbf\x08\xa5rI\x1a\xa2T" +
	"^:\xde\xf2\xd3o\xe5\xf3g\xeb\x0a\xc0\xa1;%\xdd" +
	"U\xce\xf3\x9b\xce\x90\xcb{\xc7>\xf9Qv\xd5W." +
	"\xd7;L\xc5/*mq\xbaS\x8eK\xa8\xbcEG" +
	"\xb72\x90\\\xfb\xc9\xcb?|\xa7\xd1\xbd'\xe3w\x94" +
	"\xd3q\xee!N\xa9\xf8\x15\xab\xcb\x9b\xb3`^\xd3\x84" +
	"\xd2\xd5Dy/i\xfa3\xa0+\xfe\xe7'\xbfT\xae" +
	"\x1d\xbb\xde\x08\xa4\x9e\x8e\x84\x80JW\x82['\xe8\xea" +
	"G\xd5s+\xfe\xfd\xe9\x89\x1b\xc4\x0c\xa8\x1a\x1dI$" +
	"\x10Py9\x91\x05t?\xbb\xbf\xe0\x8f\xf6{\xd7\x7f" +
	"\x1fB\xe6\xad\x04'U\xf7\xd8\x85\xd5\xc3\xafdn5" +
	"\xa0f\xcf\xc9\xc4,T\xce$(\xbe\xd3\x09\xe2\xe6Q" +
	"\xe9\xd6\x89\xe7Z'>h\x94\xee\xd6\xe6\xbb\x0ak\xa6" +
	"\x93\xd6L1u/t\xde\xbc\xfa\xde_>\x00\xf9\x0b" +
	"B\x90\x0e`\xcf\xa5\xe6\xa5\xa8\\\xe3\x96W\x9a\xc7\x01" +
	"\xdd\xd4\xf6\xc9\x9b\x85\x13\xbbos\xea\xf8\xec\x00\xec\x91" +
	"\x93\x8f\xa1\xd2\x91\xe4\xacN\xee\x04\x0c\x88\xacfP\xa8" +
	"\xa7\xce\xe6\xe4\x84\xf2t\x92@)$\x09B\x9f\x84S" +
	"!\xecG)\x82\xa8,O\xddQ\xd6\xa5f\x03\xf4l" +
	"N\xb5#\xa0\xbbK\xbe\xf4\xc8\xe4\xc4\xaf\xfe\x1e\xc6\xf0" +
	"d\x0b\xc7\xf0T\x0ba\xf8\xb7k\x9f\xc9\xbf\x1b\xfb\xf2" +
	"\xdd\xba\xe4\xb9\xef+-g\x95\x9b-t\xba\xc6m%" +
	"\xa7m\xb1\xf0\x87\xfc}\xc2;\xe8~\x0f\xa8{-w" +
	"\x94\xa64\x9d\xa2i\x02\xea\xe6\x87\xdb\xbe\xfd\xfa\xbb\xef" +
	"\xdc\xaf/57.\xa7'\x94\xfd\xdc\xb8\x92~\x03\xba" +
	"\xdc\x92\xe9h\xb6](\x8b\xc5\xee\x9cV2JK\x07" +
	"u\xad\xc2\xacn\xbb\xa4\x17\x9c\xb9C\xccN\x97u\xc7" +
	"\xf6\xcd\xb0f\x96\x1db%]\xab\xa8\x11\x0c\x07\x87\x8b" +
	"\xda\xfb\x0d\xc7\xaa\xa8I1\x02\x10A\x00\xb9\x7f\x17\x80" +
	"\xbaRDu\xbb\x802b+\xd2\xc3\xa7\x17\x01\xa8O" +
	"\x89\xa8:\x02\xca\x82\xd0\x8a\x02\x80\xbc\x87\x1e\xea\"\xaa" +
	"/\x09\xe8\x16\x8c\x82S\xd0\xf4M\x90.\xe8\xcc\xc6\x16" +
	"\xc0A\x111\x13\xf8\x02\xa4\x87\xed%\xc6\xf4\xd0k\xbf" +
	"'\xaa\xaf\xf3\xe5b)\xf4\xda\x87\xc6{=c\xf6\x06" +
	"\xdb\xbb\xa9\xa0\xb3\xb9CYfO\xc9_\xf0\x0d\x19\xd3" +
	"\xbb\xfb\x0dG\xb4*\x83\x88j\xdc\xcfw\xfeR\x00u" +
	"\xae\x88\xea\x82P\xbe]\xf4\xf0q\x11\xd5\xc5\x02fK" +
	"\xdc\x07&A\xc0$`Vg\x8e\x13|\x9c\x1e\xd2\x16" +
	"\xd3\xcawo\xb0\x0a\xccp4\xa7`\x1a\x00\xe40\xc9" +
	"!\x9b\xd3G\x99\xc8m\xdb\x00P\x90\xe5\xb5\x00\xe3e" +
	"c\xb7a\xee5\xdc\x11\xd3*\xec3\x0d\x07DMw" +
	"G\x99\xe5\x14r\x9a\x0e\x00\xbe\x83X]\xce;\xb5\"" +
	"\xdb0\xca,*:\xe5\x0c5C\xbf\xe4\x1b\x995\xca" +
	",\xf2\x1e\x11\xa3\x00\xbe\x08b\xad\x1be\xb9\x0f\x049" +
	"*\x8d\xe7L\xc3`9\xa7\x17\x07\xb1\x11\xc8\x8c\xe9C" +
	"\xcc.\x99\x86\xcd\xba7:\x9a\xe4\x94m\xba\xf5\x11\x9e" +
	"\xd3r/\xa7'\x87yNK,\x00\x14\xe5\x85\xf4'" +
	"\"w\xed\x02\xc0\xa8<\x7f\x17\xc0\xb8]\xce\xe5\x98m" +
	"\xbb\x05cT\xd3\x0b\xf9- \x99V\xde\xcd3G\xcb" +
	"\x8d\xb0<\xb4\xf7\x99\x9a\x95w\x0d\xd3Y\xae\xeb\x03\x0c" +
	"\xda\x09e\xdbec\x8e\xa5\x0d0\x07\xd2\xfcc\xf5\xcb" +
	"}\x90\xe6\xd6\xd3\xca\xbc\xb2\\,Q\xa8i\x8a\x95\x93" +
	"=\x10\x17\\\x9a\xdd\xe8hN\xd9\x9e\xa1\xfa~\xf1\xfb" +
	"\xaa\xc5_)`\xd6\xe6\xdf\xc0tp\x0f \xa6\x01\xc7" +
	"=\x1a\xf8\\%:\x84\x19\xea7^z\x95Vd\x1e" +
	"\xe5\xa8\x065\x95\xc1\x9a:\xc9\x0b;A\x90\xe7I\x18" +
	"\xa88\xd64P~\x94\xde\xa5\xa44uM/\xa6\xa9" +
	";\xa6\xd6(RG\x0a\xb2Xo:\x85\x1c\xab\xd2\x02" +
	"\x1b\xf4\xc2\x0a\xaf\xdaCY\xb6\xa7\xccl\xa7\xae\x1f:" +
	"\x1f\xd0\x0fiC+2\xbf\x1b\xaa\xcd!\xd7\xa6&Q" +
	"\xa1\x11\x0cY/@\xf2\x95\xe1@\xd4&'\xd6\xb6&" +
	"\xd2\x13Af\x04Dm\xcc`m\x17\x92\xb7\x12Q\xd7" +
	"I(\xf8\xf3\x0ak\xc3^^\xbe\x0d\x04\xf9I\x09E" +
	"\x7f\xcaam\xa7\x91\xbb\xd6rp\xdb\xb9<\xf6\xe2x" +
	"U(z\xd1\xad!\x05b\x8e>\xd6\xba\x09\x00\xa6\"" +
	"\xec\xa3\xb6\x91\xae\x18b{\xd2\x0d0[T\xc5\xac7" +
	"\x84\xd92\xa2\xd1\x97DT7\x09\xd8\xee\x84U\xb1\xca" +
	"\x95q\x0f\xbb\x87\xa0\x105 's\xb0\x01\xfa\xca\x9d" +
	"\xf1\xa3\xd0\xa8H_\x17Q\x1d\x09E\xc1(\xb4\xed\"" +
	"\xaazH\xb9\x0bk\x01\xd4\x11\x11\xd5\x17\x04D\xb1\x15" +
	"E\x00\xf9\xc0\x10m\xb1\"\xaa\xdf\x99&y\xed\xc3\xd4" +
	"l\x98\x09&  f\x00\xdd*\x9a6\x00\x04\x9a\xed" +
	"\xc7X\xd5l\xa7Pd\xb6\xa3\x15\x01K\xd8\x04\x026" +
	"5\x12\xcd\x1a%\xab*\x03<\xdb`\x83x\xe8\xd6\xed" +
	"\x0c\xe9\xb6\xdf\xba\xfe=^\xeb\xa6\xa9\xd6(\x07\xabe" +
	"\x1dg\xfd\xa0\xa8syG\xcd\x1d\xd4,\xadh\x03\xa8" +
	"\x11\xdf}\x8a\xdc\xc7ET[\x85\x07\xcf\x84z\xc9\xce" +
	"z\x17N7\\\xe9\xf7\xef\x10oN\xe4L\x0bMg" +
	"\xf2\xda+\xa2:\x10\xaa\xf1\x1a*\xe7j\x11\xd5|\xa8" +
	"\xc6a6d)\x09\x96\xf7\xe3\xb3X\x8e\x15FY>" +
	"T\xb7*\xf9<\xcb\xe9\x99\xf8]\xb0\xde+\xf9\x0c\xda" +
	"\xf1\xc0Y:\x15\xa8,\xe9\xda\xff\xc3\xad\x8e\x16\xddU" +
	"\x16\x90W/\xcb%\xde\xf0\xe9\xf2\x86\xcf\xfc\xc3|\xf8" +
	"\xcc\xdf\xc6\x87\xcf\xbc\xf0\xd4!\xd1\xdad\x9a\x03 \x99" +
	"\xc6N\xfei\xb9n1\xd4\xf2\x95M\xdanf\x00p" +
	"\x05Xc\x7f\xb5\x0c\xa2\xae\xbb\x15m\xa5i\xb0U&" +
	"\xa4\xcd\x1d,\xdf\xb8N\x8dFb\x86G5\xdf\x8b\xaa" +
	"\xe30\x8f\xaac\x1f\x8fj\xce\xe1P8\x1e\x0e\xebM" +
	"t\xd6\x18\xd5\x16\xa2\xd1\xd7o\x98\xe5\x9d\x90\x1d\xe1\x8f" +
	"\xdc\xa2\xa6\xef0\xad\"\xc3\xbcG\x06'\xb4\x0d\x04\xe2" +
	"@\xfb\x06o\x97\xe0'\x1f\x0e\xbb\xb5\x05\x04\xa4\x82i" +
	"\x84E\xa23\xa0\x85\xaf\x11\xc3\x00j^D\xb5D\xfc" +
	"A\x0f\xd9\xe2\xac\xaaF\xd0\xca'\x0a\x9eH\xec\x99U" +
	"]\xf9\xc6\x04L;l\xcc\xf1Kg\x86\xfda:\x08" +
	"\xc6\xeb9\x1c\xc38\x08\x18\x07\xc4J\xed\xd4xI\x9b" +
	":\xbd\x83_\x95\x81\x04\xb4\xfa\xc9\xec'j\x8dy:" +
	"\x96B\xd7\xf5\xd29@\x92\xf7\xac\xb7\x97\xa6\x84\xfb\xae" +
	"\x97\xcf\x8b\x94\xe4\x0b\"\xaa?\x170%\xdes\xbd\x84" +
	"NX\x00\xeaq\x11\xd5\x0b\x02\xa6\"\xffu[1\x02" +
	" \x9f\xa7\x1d\xf8M\x11\xd5\x1ba\x19\xf1c\xf1Rj" +
	"\xe7\x9b\x08\xc4\xa6.4\xf5b^\xb7\xcf\xd4\xbf\x9e\xba" +
	"\xde\xcc4\x09\xa6*\x125MCEZ\x14(\xd2\x8c" +
	"\x92=\x8d?\xde\xde\xe5-\x89\xd3n\x9a+`\xfb^" +
	"\xd3\xca\x87V\xf2\xd0\xff\x12\xa6\xc4\xe8\x17\xd1\xfb\xa1\xd1" +
	"\xdd\x9f\xa6\xf1\xf40\xfa@\x83\xe7\x09\x11\xd5\xd5\xd3\x06" +
	"\x8f\xbb\xc3\"\xe2\x1b9\xc0J-\x82\xc8\x8cS\x92z" +
	"\xf2a\xd4\xb2\xb3\xfa\x03gP@\xac\x8a\xe5:\x8aa" +
	"@D\xf5\xa9i1\xf0\xa5\xcb\x0f\xa8\xc10\xfb_\x00" +
	"\x00\x00\xff\xff{\xd4\xeaP"

func init() {
	schemas.Register(schema_f2768a70df34b76a,
		0x87747dd4e5141ec8,
		0x87c008d7bed714b1,
		0x8e62748365be5a21,
		0x91d57b12a6c06883,
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
		0xe94d36c9c6cc16d1,
		0xedb6aad718bd116a,
		0xfa3b78ca11f5cdeb,
		0xfe64d60234157407,
		0xfec3c6a7865ad9d1)
}
