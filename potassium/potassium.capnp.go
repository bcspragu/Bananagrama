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

const schema_f2768a70df34b76a = "x\xda\x8cXml\x14\xe7\x11\x9e\xd9\xbd\xf3\xde\xd9>" +
	"\xdf\xbdYS\xc2\x0f\xd7%\"Jl\x05\x0b\xec\xa4*" +
	"n\xe9\x19\x03\xe1C\x0exmS\x84\x15\xaa\xac\xef^" +
	"\xc3\xc1\xdd\xeeyw\xcf\xd84\x89\x9b\x86\x16\x82\x946" +
	"\xd0\xa6\x85\x0a\xf7C\x11M\xe5\xb6\x8a\x024\xa1\xa8\x89" +
	"@Q\xa2\x10\xe1\x00*\x11\x10\x11%\xb4\xa0\x06\xaaT" +
	"\x81\xc2\x0fPa\xab\xf7\xdd\xdb\x0f\x9f\xcf\xc5\xbf|\xbb" +
	";\xfb\xce3\xcf\xcc<3\xeby\x97\xc2m\xc2\xfc\xf0" +
	"\x9bU\x00\xca\x8ep\x85\x8d\x7f\xfd\xe4\xa73F^\xdc" +
	"\x06d\x06\xda'\xcf\x1e\x90\x0e\x9e\xdc}\x02\xc2\x82\x04" +
	"\xd0r,,\xa0<\x1e\x96\x00\xe4\xe3\xe1-\x80\xf6\xfb" +
	"_\xad\xbd|\xe6\x19k;\x90\x99\x08\x10b6\x8fU" +
	"4\"\x84\xec\xd7k\xcf\xbe}6rt;(\x09\x14" +
	"\xecMo>\xfai~\xe7\xe0u\x083\x1b\xb9\xae\xe2" +
	"\x88\xfc`\xc5C\x00-\x0b+\xd6\"\xa0=\xbb\xf7m" +
	"\xfa\xbc\xd5\xf7\xe3\xc09\x1fI\xad\xec\x9c\xad/T\xd6" +
	"\xae}\xf8\xda\x1e 3\x04;\x92\xe8_\xb3\xe6\x95\xf1" +
	"\xcf\x00\xb0\xe5-\xa9\x12\xe5q\x89c\x91V\x03\xda\xa1" +
	"\x85\x1f\x7f\xf6\xe9#?\xd9\x178\xe3\xb2\xd4\xce\xceX" +
	"\xd8\xff\\L=\x18\xdd\x0f$!\xfaP\x80\xbd~D" +
	"\xfe\x88\x1fqZZ&\xdfb\xbf\xec;\xe7\x07\xfe\xbd" +
	"\xf6\x98\xb1\x1fH\x9d`\x9f\xfe\xca\x89\xf7\x8e\x7f\xfd\x89" +
	"+\xcc\xe1E\xe9\x01\x94\xafq\xeb/\xa4_\x00\xda\xab" +
	"\x9f\xfd\xee\xbc\x9a\xef-\xfe\x1d\x0b\x11\x03!\"3Y" +
	"\x1a\xb9*+\x91\x99\x00\xf2\xba\x08c\xaac\xc5\xd3/" +
	"\x19C7^\x9d\x84\xe1\xdd\xc8\x98<\x1e\xe1aD\xb6" +
	"\xcb\xb3\xa2\x0c\xc3\xfc\xbe\xef\xec>\xf5\xeb\xc3\xbf\x0f\x84" +
	"\x82\xd1.\x16\xca\x8e}5\xbf\xfaA:}\x00H\"" +
	"\xe8R\xe0\xa8\"\xb7\xe5[\xfc\xa4\x9b\xcc\xe5\x9d=_" +
	"\xfe<\xb9\xec\xdbGK\x1d*\xd1#\xf2:\xe6\xa6e" +
	"MTB\xf9\x0e\xf78\xdcQ\xbd\xf2\xcb\xbd?{\xa7" +
	"\xdc\xb9\x97\xa3W\xe5kQ\xee!\xcaB\xf1\x8a\xa2$" +
	"nn\x9c\xa9\x1c\x93\x07*\x1f\x02\x90_\xac|\x0d\xd0" +
	"\x16o\xfd\xf2\x0f\xf2\xf8\x9e\x93\xe5HjQ\xaa\x04\x94" +
	"\xd7W1\x96h\x15;z\x96rp\xf1\x7fn\x8c\x9e" +
	"b\xc5\x07E\xa3cU\x95\x08(\x1f\xafJ\x02\xda7" +
	"\xef\xce\xbb`~x\xf2o\x01f\xaeT\xf1Bi\x1a" +
	":\xbc\xbc\xef\x95\xc4\x99\x92\x08Dfr\xba\xea>\x94" +
	"/V1|\x9fT1T\xbb\xa53\xa3\xcf\xd6\x8e\x9d" +
	"+\x17\xee\x0b\xd5\xb7\xe5\x97\xab\xd9\xaf]\xd5\x0cS\xd3" +
	"|\xeb\xd0\x89\x0f??\x07\xe4k\x82\x1f\x0e`K," +
	"\xd6\x8ar]\x8cY\xce\x8a\x8d\x00\xda\xb1\xa7\xce\x9e\xce" +
	"\x8cn\xfe\x98\x97\x8eW\x1d\x80-j\xec\x01\x94\x07\xb8" +
	"e.\xb6\x01\xd0/\xe4\x92\xee\xe0\xa5\xb3?6&\xff" +
	")\xc6H\xf9K\xec\x9f\x0c\xc0\xfa\x0f\xdex\xf4\xfa?" +
	"\xfe\xceb\xf6\xfb\xca\xe9\xc8]5\x02\xca\xa35\xec\xb5" +
	"\xbd5\x0c\xadW\xb1\x13\xf9^\x8aR\x08Q\x8e\xc6\xaf" +
	"\xca3\xe23\x01Z\xea\xe2\xf5\xac\xf16\x91\xb7\xee?" +
	";\xf6\xc6\x17A\xc23\x09N\xf8@\x82\x11>*]" +
	"y\xe9\xf9\x13c\xd7\xcb\xc9\xc1\xaeD%\xca\xbfM0" +
	"\xe7\xa3\x09\xe6\xfc\xe8\xec\x86o=\xb9w\xc1\x0dPf" +
	"b\x10*\x0f\xebV\xe2\xaa\x1c&\xec\x17\x12f\xfc\xaf" +
	"\xf1\x9b\xe4\x83\xa1o\xde.\xc9\x01\xb7]O\x0e\xc8\x94" +
	"\xdb\xaa\x84\xa18}\xbe\xf7G\xaf\xbe\xf7\xce\xdd\xb2%" +
	"\xb7\x8b\x8c\xc9{\xb9\xf1\xcb\xe45\x98k\xe7uK5" +
	"\xcdLA\xcc5\xa5\xd4\xbc\x96o\xed\xa44\xdb\xd4\x99" +
	"U\x87\xa9\xd1\x93\xc9R\xe8DT\"b\x08 \x84\x00" +
	"\xa4\xa1\x15@\x99#\xa22O@\x82X\x8b\xec\xe6\xdc" +
	"v\x00\xe5a\x11\x95%\x02&\xf3\xfcU\xac\x06\x01\xab" +
	"\x01G\xb2\xd4\xb2\xa8ab\x0d`\xa7\x88\xfcv\x0d`" +
	"\x19\xb7\xfc\xb5&3\x9f\xcdXs\xba\xa8\x19/d-" +
	"\xd33C\xd7,\xd9E\x99\x03%\x82\xc1TGW\xfa" +
	"l\x92h\xb3\xcd\x80?n\xd0\x01\x00\xa8\xefN\xe9\x06" +
	"U\xee\xf7\"\xd8\xbb\x09@\xd9#\xa2r(\x10\xc1\xeb" +
	"\xcd\x00\xca\x1fET\xde\x17\x90\x08B-\x0a\x00\xe4]" +
	"v\xf3\xa8\x88\xca\x05\x01Q\xacE\x11\x80\x9c\xef\x02P" +
	"\xce\x89\xa8\\\x12\x90\x84\xb0\x16C\x00\xe4\"\x8b\xff\x82" +
	"\x88\xca\xe7\x02\x92\xb0X\x8ba\x00r\xb9\x0f@\xb9$" +
	"b\x17\x0ahg\xb4\x8c\x95Q\xb3=\x10\xcfd\xa9G" +
	"F\xc2\x0f\x01\x90\xdd\xac\xcfS\x9a\x0d<\xf6J\xa8\xf8" +
	"8]\xc8\xe5\x03\x8f\xbdD;\x8fm\xd3R\x0d\xab'" +
	"\x93\x03\xa4\x18\x05\x01\xa3\x80#TK\xf7dr\xde\xb5" +
	"\xdd\x9f\xd1\xd4lwJ\x07\xc9\x08\x02\xf1\xe8+\x1e5" +
	"U~4\xba\x85\x91;\xa7+I\xcd\x09\x19\xf2\x0c\xd7" +
	"\xeaF\xbai\xb5\x91\xa1\x9a\xa5Z\x19]\x03^A\xd5" +
	"\x9c\xd1\xbavv>\x99\xd1\x0b\x80\x02!+\x01F\x0a" +
	"\xdafM\xdf\xa2\xd9\x1bu#\xb3U\xd7,\x10\xd5\xac" +
	"=H\x0d+\x93R\xb3\x00\xe09\xa8(A\xb2A\xcd" +
	"\xd1\xd5\x83\xd4`\xc5\xc2\x90\x80k\xe8\x95J75\x06" +
	"\xa9\xc1\xbc\x87\xc40\x807\x0e\xd0\xd5%B\xdaA " +
	"ai$\xa5k\x1aMYm\xd8\x89X\xbe#\xba\xa8" +
	"\x99\xd75\x936u[\xaad\x15Lv\xea\xfd<\xa6" +
	"ENL\x0b\xfaxL\x8f\x19\x00(\x92\xf9\xecO\x88" +
	"\xcc\xdd\x04\x80a\xd2\xb0\x09`\xc4,\xa4R\xd44\xed" +
	"\x8c6\xa8f3\xe9\xb5 \xe9F\xdaNSKMm" +
	"\xa4i\xa8o\xd7U#mk\xba\xb5(\x9b\xed\xa0P" +
	"\xcf\xdb\xc6\xa6C\x96\xa1vP\x0b\xe2\xfc\xb2\xf8r;" +
	"\xc4\xb9\xb5\x0bUp\xa1.)\xe4\xf2\x0cj\x9caU" +
	"B\x18\x94YlMv[\xaaU0\xa7h\xe7\xb2\xdd" +
	"l\xf270\xee\x9f\x03\x88\xf1{\xf7\xb5\xd7\xb0\xf1e" +
	"j\x8e:\x1a\xc2r\xe0J(\xba\xd2K\xe67\x82@" +
	"\x1e\x94\xd0\x9fg\xe8N\x032\x8b=\x8bIq\xd6\x16" +
	"m\x18g\xe5?1G\xa1\x92\xa2`\x16\xabt+\x93" +
	"\xa2\xc5\xb2@s2I\x8b\x9dlw%\xe9@\x81\x9a" +
	"V\x89\xc05\x96\x13\xb8\xd6\"%\x8f\x0a\x18\xd7\xd4\x1c" +
	"u\xe5\xcdU;\xe2\xee\x0f\xac\x14\xca\xd1\x90t\x002" +
	"_\x09N\x84\xbbC\xa0\xbb\x15\x92\x81f\x10\x08eD" +
	"\xb8\x03\x17\xddM\x8f\xacc\x85\xfa\x84\x84\x827\xb9\xd1" +
	"]{\xc8\xa2^\x10\xc8\x02\x09Eo\xde\xa3\xbb\xdd\x91" +
	"\xb9+9\xb9\xf5\\V\xdbp\xa4\xd8\xbemh\xbbL" +
	"\x81\x98b\x97n7\x01\xc0D\x86=\xd6\xba\xd9\x11]" +
	"t ^\x86\xb3\xe6\"gm\x01\xce\x16\xb22\xfa\x86" +
	"\x88J\x8f\x80\xf5VP\xf6\x8a\xb52\xe2p7\x8d\x12" +
	"b\x0d\xc8\x15\xdf\x9f\xaf\xd1\xbe\xc0\xee\x1d\xed\xb5\x1dz" +
	"\xdbu\x90X_\xb8\xe3\x0b\xc4,U\xaa=\x9cK\x99" +
	"$/\x11Qy*\x80s\xfdJ\x00\xe5I\x11\x95!" +
	"\x01\xb1\xa8\xfc\x05\xa6\xf2\x96\x88\xca\xf7\x05\xb4\x9d\x8e\xd3" +
	"U\x90\x8ct@1=,E\xc5,Rk\x02\x80o" +
	"\xe4a,\x1aY\x99\x1c5-5\x07\x98\xf7\x14y\x92" +
	"\xde\xb8\xf5Y\x94\x1c\xe0}\xec/V\xd3\xee\xe3F\xbf" +
	"h\xfd>\xf6\xceq\xfa8\xce\x12\x8f\xc4\xdf\xb8K\x0a" +
	"\xd8\x03\xc5\xda\x98\xb7\xd7\x9cN\xd5Ps&\x80\x12\xf2" +
	"\xdc\xc7\x98\xfb\x88\x88J\xad\x80IG\x19\xdc\x0e\x99r" +
	"\x92x\xfa\x9dt\x0e\x9cl\xb8\xc4k\xe6.\xde\xa9\xc8" +
	"\xcb.\x90N\xe6\xb5MD\xa5#\x90\xce\x15,\x9d\xcb" +
	"ET\xd2\x81I\xae\xb6\x16s\xbcQ\xc0$\x0b\x82\xa6" +
	"=|\x06M\xd1\xcc M\x07\xf2V\xacD\xc7rr" +
	"$^K\xacrR>\x85\x90\xb4\xdf\xa3)\xa6\x90\xd0" +
	"\x11&v\xd3\xd9\x98J\xca\xa4\xa9X\x15\x0c\x85\x13\xf5" +
	"c\xced\x9a\xebL\xa6\x86\x9d|25\xf4\xf2\xc9\xf4" +
	"`p$1E\xeb\xd1\xf5\x0e\x90tm\x03\xbfZ\x94" +
	"5(\xaa\xe9\xe1\x1eu3\xd5\x00\xb8<\xac0\x1f/" +
	"\x80\x98\xcd\xda\xc3\xea\x12]\xa3\xcbt\x88\xeb\xfd4]" +
	">o\xe5\xe6e\x82\xa3jpP\xcd\xde\xc9Q\xcd\xde" +
	"\xcaQ\xd5\xed\x0c\xc0qxY\xa5\xa3\xb5B+\xb6\x14" +
	"\x9b\x8bK5\xbd\xb0\x01\x92\x1b\xf9-;\xa7f\xfbu" +
	"#G1\xed\x14\x87\x15X\x15|\xe5`\xcb\x08o\x1f" +
	"\xff\xcb\x18\xfblw;\x01)\xa3kJ\xc2K\x99\xda" +
	"\xe8\x97\x89\x9b1\xca4#-\xa2\x92g\xf5\x84\x0e\xb3" +
	"\xb9\xfb\x00\x94\x8d\"*\x96\x80D\x14\x9c\xd5p\x80\xdd" +
	"\xcc:B\x12\xb7\xe8\x90\xe5\x15\x8e\x1e\xf4\x87q\x1f\x8c" +
	"\xd3\x838\x84\x11\x100\x02\x88\xc3\xee\xaf\xc9\xac:\x8b" +
	"o\x93\xb7\xdaNg/g2\xf6\x88\x88\xca\xf2I{" +
	"\xb9\xddo0\xda\xb4\x14\xe0\xb0[i\xa1\x92J\x13J" +
	"\x17 \x7f\xab\xf0\xbf\xfb}5\xaa\xf5\xc0<\xc3\xc0\x0c" +
	"\x89\xa8l\x130\x86\xb6\xed\xa0y\x8eM\x89\xa7ET" +
	"v\x08\x18\x13\xee\xda\x0e\x95?d\xfcn\x13Q\xf9\x8d" +
	"\x801\xf1\x8e\xedp9j\x00(\xfbDT\x0e\x0b\x18" +
	"\x0b\xfd\xd7v\xf6\xec?\xb3\xd5\xfd\x90\x88\xca\xa9\xa0\xa2" +
	"yX\x1c6\xeb\xb9^C\xc5\xc4E\xab\xb4\x9bJ\xf6" +
	"\xac\xd2\xc7\x13\xd7\xae{\xb6\"\x17G\xd6\xb7e\xc5\xb1" +
	"\xd9\x17\xc7\xfa>\xb6\xb7a\xc2\xff\x9e\x03\xc4\x04\xe0\xff" +
	"\xfd\x0c\xe3\x9b!L+\xe1\xcd\x01\xc9\x9f\x98\xf0{z" +
	"\x16J\x0a\xad;\xce>\x9c\xa6v:\x1d\x9f&;b" +
	"RI{\xed\xe9\xec\xbc\xce\x82>\x89\xad9\x02\xd6o" +
	"\xd1'\x8c\xdc\xc0\x7f\xb4\xa6\xd8\x14\x98\xf4LgH4" +
	"\xfa\x8b\x80;\xf3\xd7w\x05f\xc4\xc48\xf8\xe2YZ" +
	"\x04eF\xf9\xff\x02\x00\x00\xff\xff%_\x83\xa2"

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
