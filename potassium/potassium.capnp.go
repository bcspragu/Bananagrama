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

const schema_f2768a70df34b76a = "x\xda\x8cXml\x14\xe7\x11\x9e\xd9\xbd\xcb\xde\xd9>" +
	"\xdfm\xd6\x94\xf0\x83^\x89\x88\x1a\xa3\xd8\x02;T\x0d" +
	"-=\xe3@\xc0\xc8\x01\xd66EE\xa1e}\xf7\x1a" +
	"\x1f\xec\xed\x9ew\xd7\xc6&!.\x0d)\x09\xea\x17\xa8" +
	"i\xa1\x82~(\"i\xdcF\x11_)A\"\x02E" +
	"Di\x04\x01T\x10\x1f\xa2J\xd2\x065\xa1j\x95\xa6" +
	"\xe4GQa\xaby\xf7nw}g\x17~y?\xe6" +
	"\xde\x99y\xe6\x99gf=\xfb\xd5h\x9b0'\xfa\x9d" +
	"Z\x00uk\xf4\x1e\xf7\x9d/6\\;\xbf\xd9\xd9\x06" +
	"\xf2T\x04\x88H\x00\xad/Dg!D\xdc\xfd\x0d\x17" +
	"\xdf\xbc\x18;\xbe\x0d\xd4\x14\x0a\xee\xfa?<\xfc~q" +
	"\xfb\xd0g\x10%\x1bes\xf4\xa8\xf2l\xf4\xcb\x00\xad" +
	"{\xa3\xab\x10\xd0\x9d\xb1\xfaM\xf6\x8c\xd3\xfb\xa3\xd09" +
	"s\xa5yt\xce3\xfd\xc7_\xba\xf7\xc9\x0b;@\x9e" +
	"\x82\xee\x99\x8b\x07\xa4\x83gv\x9e\x82\xa8@\xc7L\x97" +
	"\xae(\x8d\x12]= m\x04t7=_\xd3\xb0\xea" +
	"\xc1\x7f\xed\x02y\x8a\xe0\xc6R}+W\xbex\xfa\x03" +
	"\x00l\xdd\"\xd5\xa0\xb2\x83[\xfePZ\x0e\xe8F\xe6" +
	"_\xf9\xe0\xfd\x87~\xbc'\xe4\xef\x15\xa9\x9d\xfc\xcd\xef" +
	"\xdb\x92\xd0\x0e\xc6\xf7\x81\x9c\x12\x83\xb0\x81~~T\xd9" +
	"\xcd\x8fxAZ\xac\x1c\xa3+\xf7\xd6\xe5\x81\x7f\xae:" +
	"a\xed\x03y\xba\xe0\x9e\xfb\xc2\xa9\xb7\xff\xf8\x95\xc7?" +
	"!\x87\xfb\xa4\xfbQ9\xcc\xad\xf7K?\x07t\x97?" +
	"\xfd\xed\xd9\xf5O>\xfa\x12\xc1\x81!8\x90L\xe2\xb1" +
	"\xeb\xca\x94\xd8T\xca(Fytv<\xf5\x13k\xf8" +
	"\xc6\xcbU1<\x1f\x1bSv\xc4x\x1a\xb1m\xca\x9f" +
	"\xe9\xca\x9d\xd3\xfb\xcd\x9dg\x7fu\xe4\xb7\xa1TN\xc6" +
	"\xba(\x95\xe7\xf6\xd4\xff\xf2{\xb9\xdc\x01\x90Sa\x97" +
	"\x1c\xba\xfd\xb1\x9b\xca1~\xd2\x1b\xe4\xf2\xd6\xaeO\x7f" +
	"\x96Y\xfc\x8d\xe3\x95\x0e\xa7\xc4\x8f*\xd3\xe3t\xe6\xb4" +
	"\xb8\x84\xca\x09\xbatG:\xeb\x96~\xba\xfb\xa7oM" +
	"t\xee+\xf1\xeb\xca\xe18\xf7\x10\xa7T\xfc\x8aU\xe4" +
	"-\x92IS\xcd\x982\xb7\x86\xf2\x9e_\xf37@W" +
	"\xfc\xcf/~\xa7\x9c\xdeuf\"\x90Z\x1bk\x05T" +
	"\xe6\xd6r\xebZ:z\x9az\xf0\xd1\x7f\xdf\xd8{\x96" +
	"\x98\x01%\xa3\xdd\xb55\x08\xa8\xfc\xa66\x03\xe8~~" +
	"{\xf6U\xfb\xbd3\x7f\x0a!s\xa2\x96\x93\xaay\xf8" +
	"\xc8\x92\xde\x17S\xe7+2\x109\x0fj\xefE\xe5\x8d" +
	"Z\x8a\xefp\xedk\x80\xeeN\xe9\xfc\xde\xa7\x1b\xc6." +
	"M\x94\xaeVwS)\xd4\xd1U\xbe\x8ebj\x9e\xe3" +
	"\x1c:\xf5\xde\xc7\x97@\xfe\x92\x10\xa4\x03\xd8z\xb2n" +
	"\x1e*\x17\xb8\xe5\xb9\xbaQ@7\xb1\xf6\xe2\xb9\xfc\xde" +
	"\x0dW8u|v\x00\xb6NK\xdc\x8fJc\x82\xb3" +
	":\xb1\x0e0 rE'q\xea\xacI\x8c),A" +
	"\xa0\x0c$\x08\xc2\xe65\xef\xbe\xfe\xf0g\x7f\xfd\x0b\xe5" +
	"\x1c\xf4 \x0f\xb6\xb5\xa3^@ee=\xfdL\xad\xa7" +
	"h}\xc6\x8e\xc7{\x11J\x11D\xe5D\xfdu\xe5t" +
	"\xfdT\x80\xd6\x0b\xf5ij\xd2\xf5\xf2\xb1\xfb.\x8e\xbd" +
	"\xfe\x8f0\xe03R\x1c\xf0\xc6\x14\x01~|F\xe3\xd7" +
	"\x9f\xd8\xfd\xc8\x0dP\xa7b\xd8;\x8f\xb4#u]Y" +
	"\x99\xe2\xceS\xe4\xfc\xef\xa7?\x97\xdf\x1d\xfe\xda\xcd\x0a" +
	"X\xb9\xed\xab\xa9\x03\xcaan\xbb\x9f\x1f|\xee\xf2\xea" +
	"\xef\xbf\xfc\xf6[\xb7+\x89\xc1k\xf0ajL\xf9\x84" +
	"\x1b_K\xbd\x06Mn\xd1t4\xdb\xce\x0f\x8a\x85\xe6" +
	"\xacV4\x8a\xf3V\xe8\xda\x08\xb3\x9a\xed\xa2\x9ewf" +
	"v1;9\xa8;\xb6o\x86e\xb3L\x17+\xea\xda" +
	"\x88\x1a\xc30\x8c\xf1\xa5AZr\xbc\xc5\xed\xc9\xeb\xec" +
	"1\x8b\x0d\x00@\xba;kZL\xbdO\x8c\x00D\x10" +
	"@\xde\xbd\x1e@\xdd%\xa2zH@\x19\xb1\x01\xe9\xe1" +
	"\xfe\x16\x00\xf5\xf7\"\xaa\xef\x08(\x0bB\x03\x0a\x00\xf2" +
	"Izx\\D\xf5\xaa\x80(6\xa0\x08 _\xee\x02" +
	"P/\x89\xa8~$\xa0\x1c\xc1\x06\x8c\x00\xc8\x1f\xb6\x03" +
	"\xa8WET?\x16P\x8e\x8a\x0d\x18\x05\x90\xaf\xf5\x02" +
	"\xa8\x1f\x89\xd8\x85\x02\xbay#\xef\xe45\xbd\x07\x92y" +
	"\x9d\xd9X\x0f\xb8BDL\x05)\x00\xd2\xc3t\x911" +
	"=\xf4\xdao\xcc\xd2\xeb\xdc`\xa1\x18z\xed#\xee\xbd" +
	"vmG\xb3\x9c\x9e|\x01\x90a\x1c\x04\x8c\x03\x8e2" +
	"#\xd7\x93/\xf8\xf7n_\xde\xd0\xf4\xee\xac\x09\x92\x15" +
	"\x0e\xc4\x87\xaft\xd4d\xf51\xd8F\x02wfW\x86" +
	"\xd9\xe3*$\xf8\x86\x8c\xe9\xcd\x8b\x0cG\xb4FV " +
	"\xaa1\x1f\xf8\xc6y\x00\xeaL\x11\xd5\xd9!\xe0\x9b\xe8" +
	"\xe1\x83\"\xaa\x0f\x0b\x98)r\x1fX\x07\x02\xd6\x01f" +
	"t\xe68\xc1muH\xabL+\xd7\xbc\xdc\xca3\xc3" +
	"\xd1\x9c\xbci\x00\x90\xc3:^\xbb\xe9\xed\x94\x89<e" +
	"5\x00\x0a\xb2\xbc\x14`t\xd0\xd8`\x98\x1b\x0d\xb7\xdf" +
	"\xb4\xf2\x9bL\xc3\x01Q\xd3\xdd!f9\xf9\xac\xa6\x03" +
	"\x80\xef\xe0\x9e\x8a\x9c\xd7i\x05\xb6|\x88YDK\xca" +
	"\x19\xca\x86>)\xbb\x995\xc4,\xf2\x1e\x11\xa3\x00\xbe" +
	"\xa8cY]d\xb9\x1d\x049*\x8dfM\xc3`Y" +
	"\xa7\x0dW\xe0D 3\xa6w1\xbbh\x1a6k\xee" +
	"v4\xc9\x19\xb4\xe9\xd4\xfbxN\x0b\xbc\x9c\x1e\xe9\xe5" +
	"9\xcd\xb5\x00P\x94\xe7\xd0\x9f\x88\xdc\xb4\x1e\x00\xa3r" +
	"\xe3z\x80Q{0\x9be\xb6\xed\xe6\x8d!M\xcf\xe7" +
	"V\x81dZ97\xc7\x1c-\xdb\xcfr\x90n75" +
	"+\xe7\x1a\xa6\xb3@\xd7;\x19\xa4\x09e\xdbe\xc3\x8e" +
	"\xa5u2\x07\x92\xfc\xb6\xf4\xe3vHr\xeb\xaa2/" +
	"\x1c,\x14)\xd4$\xc5\xaaF0,\x968/\xd3\xed" +
	"h\xce\xa0=I\xf5\xfd\xe2\xb7\x97\x8a\xbfP\xc0\x8c\xcd" +
	"\x7f\x81\xc9\xe0\x1c@L\x02\x8ez4\xf0\xb9Jt\x08" +
	"3\xd4\x97\x86\xe4b\xad\xc0<\xcaQ\x0d\xcaB\x88e" +
	"\x01\x95\xe7\xcc\x02A~@\xc2`*aY\xd3\xe5i" +
	"\xf4.!%\xa9\x01\xdb0I\x8d6\xbeF\x91\x0aR" +
	"\x90\xc52\xd3\xc9gY\x89\x168A/<\xeaU\xbb" +
	"+\xc3\x06\x06\x99\xedT\xf4\xc3\xac;\xf4C\xd2\xd0\x0a" +
	"\xcc\xef\x86Rs\xc8\xe5-\x80\xa80\x11\x0c\x19/@" +
	"\xf2\x95\xe2@\x947\x01,\xef\x81\xf2@\x0b\x082#" +
	" \xcac\x13\xcb\xbb\x9d\xfc-\"\xea\xe3\x12\x0a\xfe\xfc" +
	"\xc5\xf2\xf2\"/X\x0d\x82\xfc\x88\x84\xa2?\xb5\xb1\xbc" +
	"\xa3\xc9MK9\xb8i.\xe0m8Z\x12\x8a6t" +
	"\xcbH\x81\x98\xa5\xdbr7\x01\xc0x\x84}\xd4\xba\xe9" +
	"\x88.6\x90\x9c\x00\xb3\x96\x12fm!\xcc\xe6\x13\x8d" +
	"\xbe*\xa2\xda#`\xda\x09\x0bl\x89+\xa3\x1evw" +
	"A!j@N\xe6`\xa3\xc5\x96\xf4\"\xc3\xb1F\xd4" +
	"\x94\x1f\x85FEzBD\xb5?\x14\x05\xa3\xd0\xd6\x8a" +
	"\xa8\xea\xa1\x11\x92_\x0a\xa0\xf6\x8b\xa8n\x0dF\xc8\x16" +
	"\x1a!\xdf\x15Q\xfdA\x95\xe4\xa5{\xa9\xd90\x15\xcc" +
	"]@L\x01\xba%4m\x00\x084\xdb\x8f\xb1\xa4\xd9" +
	"N\xbe\xc0lG+\x00\x16}\xb9\xaf\x92\x982%K" +
	"*\x03<\xdb`#\xba\xeb\xd6\x9d\x15\xd2m\xbfu\xfd" +
	"s\xbc\xd6MR\xadQ\x0eV\xe5\x0a\xce\xfaAQ\xe7" +
	"\xf2\x8e\x9a\xb9B\xb3\xb4\x82\x0d\xa0F|\xf7\x09r\x1f" +
	"\x13Qm\x10\xee<\x13*%;\xe3\x1dXm\xb8\xd0" +
	"\xef\xdf.\xde\x9c\xc8\x99V\xe7{]D^\xdbDT" +
	";C5\xee\xa0r.\x11Q\xcd\x85j\x1cfC\x86" +
	"\x92`9?>\x8beY~\x88\xe5Bu+\x91\xcf" +
	"\xb3\xac\xce\xc4\xef\x82e^\xc9'\xd1\x8e;\xce\xd2\xf1" +
	"@eH\xd7\xfe\x1fn\x15\xb4h.\xb1\x80\xbczY" +
	"\xce\xf5\x86O\x937|\x1a\xb7\xf3\xe1\xd3\xb8\x9a\x0f\x9f" +
	"\x07\xc2S\x87D\xab\xc74;A2\x8du\xfcn\x81" +
	"n1\xd4r#=\xda\x06f\x00p\x05\xe8\xb0\x1f\x1b" +
	"\x04Q\xd7\xdd\x11m\xa1i\xb0\xc5&$\xcd>\x96\x9b" +
	"\xb8N\x13\x8d\xc4\x14\x8f\xaa\xd1\x8bj\xc6v\x1e\xd5\x8c" +
	"M<\xaa\xe9\xdbC\xe1x8,3\xd1\xe90J-" +
	"D\xa3o\x91a\x0e\xae\x83L?\x7f\xe4\x164\xbd\xcf" +
	"\xb4\x0a\x0cs\x1e\x19\x9c\xd06\x10\x88\x03\xed\x1b\xbc]" +
	"\x82OX\xecu\xcb\x0b\x08Hy\xd3\x08\x8b\xc4\xac\x80" +
	"\x16\xbeF\xd0N\x98\x13Q-\x12\x7f\xd0C\xb6po" +
	"I#\x1c\x01eQ\xf0Db\x80\x1e\xea\"\xaa\xc3\x02" +
	"&\x1d6\xec\xf8\xa53\xc3\xfe0\x19\x04\xe3\xf5\x1c\x0e" +
	"c\x0c\x04\x8c\x01\xe2H\xf9\xaa\x1aUo\x8bn\xf6\xf7" +
	"\xe4\xbba\x17\xc9\xd6C\"\xaaK\xaad\xcb\xed\xb3\x08" +
	"6#\x0b8R\xa6y\xa4Bc\x85\xca\x1d'X\x1c" +
	"\x82\x0f\xf4@}\x1a\xfc`6S0\xc3\x9e\x84&\xd0" +
	"u\xbdh\xb6\x90\xda>%\xa2\xfa\x9c\x80\x09\xe1\xb6\xeb" +
	"A\xf9,\xe1\xbbUD\xf5\xd7\x02&\xc4[\xae\x87\xe5" +
	"^\x0b@\xdd#\xa2zD\xc0D\xe4\xbf\xae\xb7\xb4\x1f" +
	"\xa6\xef\x80C\"\xaag\xc3\x0a\xe6\xc7\xe2\xa1\x99\xe6K" +
	"\x10\xdc3~\x97\xaa\x9c#\x15\xabT\xe5\xeb\xf1\x9b\xd5" +
	"dCh\xbc\x18R\xbfN(\x86-\x81\x18N:-" +
	"\xaa0/\x95\xbb;I\xdfB\x93\xd7\xda/u\xcb\xa4" +
	"Ky\xda\xa6#\xaa\x88\xe57\x89\xb7\\z\x9bpU" +
	"\xcc3\x05Lo4\xad\\\xe8\xbb#\xf4\x0f\xa0IF" +
	"2\x09\xc0\xddH3\xb5\xdbB\x11\xd5\xb5\x02bI\x99" +
	"\xd7t\x85\x94y|\x1e|\xc3\xab,\xc5\x04\x03\xf4\x7f" +
	"\x01\x00\x00\xff\xff\xf6_AX"

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
		0xe2e3f234b6ca5d2e,
		0xe94d36c9c6cc16d1,
		0xedb6aad718bd116a,
		0xf439995c3c2921c0,
		0xfa3b78ca11f5cdeb,
		0xfec3c6a7865ad9d1)
}
