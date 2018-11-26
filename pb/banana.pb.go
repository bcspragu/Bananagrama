// Code generated by protoc-gen-go. DO NOT EDIT.
// source: banana.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusUpdate_Status int32

const (
	StatusUpdate_UNKNOWN      StatusUpdate_Status = 0
	StatusUpdate_GAME_STARTED StatusUpdate_Status = 1
	StatusUpdate_GAME_OVER    StatusUpdate_Status = 2
)

var StatusUpdate_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "GAME_STARTED",
	2: "GAME_OVER",
}

var StatusUpdate_Status_value = map[string]int32{
	"UNKNOWN":      0,
	"GAME_STARTED": 1,
	"GAME_OVER":    2,
}

func (x StatusUpdate_Status) String() string {
	return proto.EnumName(StatusUpdate_Status_name, int32(x))
}

func (StatusUpdate_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{6, 0}
}

type Word_Orientation int32

const (
	Word_UNKNOWN    Word_Orientation = 0
	Word_HORIZONTAL Word_Orientation = 1
	Word_VERTICAL   Word_Orientation = 2
)

var Word_Orientation_name = map[int32]string{
	0: "UNKNOWN",
	1: "HORIZONTAL",
	2: "VERTICAL",
}

var Word_Orientation_value = map[string]int32{
	"UNKNOWN":    0,
	"HORIZONTAL": 1,
	"VERTICAL":   2,
}

func (x Word_Orientation) String() string {
	return proto.EnumName(Word_Orientation_name, int32(x))
}

func (Word_Orientation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{13, 0}
}

type PeelResponse_Status int32

const (
	PeelResponse_UNKNOWN         PeelResponse_Status = 0
	PeelResponse_SUCCESS         PeelResponse_Status = 1
	PeelResponse_INVALID_WORD    PeelResponse_Status = 2
	PeelResponse_DETACHED_BOARD  PeelResponse_Status = 3
	PeelResponse_NOT_ALL_LETTERS PeelResponse_Status = 4
	PeelResponse_EXTRA_LETTERS   PeelResponse_Status = 5
)

var PeelResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "INVALID_WORD",
	3: "DETACHED_BOARD",
	4: "NOT_ALL_LETTERS",
	5: "EXTRA_LETTERS",
}

var PeelResponse_Status_value = map[string]int32{
	"UNKNOWN":         0,
	"SUCCESS":         1,
	"INVALID_WORD":    2,
	"DETACHED_BOARD":  3,
	"NOT_ALL_LETTERS": 4,
	"EXTRA_LETTERS":   5,
}

func (x PeelResponse_Status) String() string {
	return proto.EnumName(PeelResponse_Status_name, int32(x))
}

func (PeelResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{14, 0}
}

type NewGameRequest struct {
	// What to name the game.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// How to scale the number of tile in the game.
	ScaleFactor          int32    `protobuf:"varint,2,opt,name=scale_factor,json=scaleFactor,proto3" json:"scale_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewGameRequest) Reset()         { *m = NewGameRequest{} }
func (m *NewGameRequest) String() string { return proto.CompactTextString(m) }
func (*NewGameRequest) ProtoMessage()    {}
func (*NewGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{0}
}

func (m *NewGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewGameRequest.Unmarshal(m, b)
}
func (m *NewGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewGameRequest.Marshal(b, m, deterministic)
}
func (m *NewGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewGameRequest.Merge(m, src)
}
func (m *NewGameRequest) XXX_Size() int {
	return xxx_messageInfo_NewGameRequest.Size(m)
}
func (m *NewGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewGameRequest proto.InternalMessageInfo

func (m *NewGameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewGameRequest) GetScaleFactor() int32 {
	if m != nil {
		return m.ScaleFactor
	}
	return 0
}

type NewGameResponse struct {
	// The unique identifier for the game.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewGameResponse) Reset()         { *m = NewGameResponse{} }
func (m *NewGameResponse) String() string { return proto.CompactTextString(m) }
func (*NewGameResponse) ProtoMessage()    {}
func (*NewGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{1}
}

func (m *NewGameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewGameResponse.Unmarshal(m, b)
}
func (m *NewGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewGameResponse.Marshal(b, m, deterministic)
}
func (m *NewGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewGameResponse.Merge(m, src)
}
func (m *NewGameResponse) XXX_Size() int {
	return xxx_messageInfo_NewGameResponse.Size(m)
}
func (m *NewGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewGameResponse proto.InternalMessageInfo

func (m *NewGameResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JoinGameRequest struct {
	// The unique identifier for the game to join.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the player joining the game.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGameRequest) Reset()         { *m = JoinGameRequest{} }
func (m *JoinGameRequest) String() string { return proto.CompactTextString(m) }
func (*JoinGameRequest) ProtoMessage()    {}
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{2}
}

func (m *JoinGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGameRequest.Unmarshal(m, b)
}
func (m *JoinGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGameRequest.Marshal(b, m, deterministic)
}
func (m *JoinGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGameRequest.Merge(m, src)
}
func (m *JoinGameRequest) XXX_Size() int {
	return xxx_messageInfo_JoinGameRequest.Size(m)
}
func (m *JoinGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGameRequest proto.InternalMessageInfo

func (m *JoinGameRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JoinGameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GameUpdate struct {
	// Types that are valid to be assigned to Update:
	//	*GameUpdate_PlayerUpdate
	//	*GameUpdate_StatusUpdate
	//	*GameUpdate_TileUpdate
	Update               isGameUpdate_Update `protobuf_oneof:"update"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GameUpdate) Reset()         { *m = GameUpdate{} }
func (m *GameUpdate) String() string { return proto.CompactTextString(m) }
func (*GameUpdate) ProtoMessage()    {}
func (*GameUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{3}
}

func (m *GameUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameUpdate.Unmarshal(m, b)
}
func (m *GameUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameUpdate.Marshal(b, m, deterministic)
}
func (m *GameUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameUpdate.Merge(m, src)
}
func (m *GameUpdate) XXX_Size() int {
	return xxx_messageInfo_GameUpdate.Size(m)
}
func (m *GameUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_GameUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_GameUpdate proto.InternalMessageInfo

type isGameUpdate_Update interface {
	isGameUpdate_Update()
}

type GameUpdate_PlayerUpdate struct {
	PlayerUpdate *PlayerUpdate `protobuf:"bytes,1,opt,name=player_update,json=playerUpdate,proto3,oneof"`
}

type GameUpdate_StatusUpdate struct {
	StatusUpdate *StatusUpdate `protobuf:"bytes,2,opt,name=status_update,json=statusUpdate,proto3,oneof"`
}

type GameUpdate_TileUpdate struct {
	TileUpdate *TileUpdate `protobuf:"bytes,3,opt,name=tile_update,json=tileUpdate,proto3,oneof"`
}

func (*GameUpdate_PlayerUpdate) isGameUpdate_Update() {}

func (*GameUpdate_StatusUpdate) isGameUpdate_Update() {}

func (*GameUpdate_TileUpdate) isGameUpdate_Update() {}

func (m *GameUpdate) GetUpdate() isGameUpdate_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *GameUpdate) GetPlayerUpdate() *PlayerUpdate {
	if x, ok := m.GetUpdate().(*GameUpdate_PlayerUpdate); ok {
		return x.PlayerUpdate
	}
	return nil
}

func (m *GameUpdate) GetStatusUpdate() *StatusUpdate {
	if x, ok := m.GetUpdate().(*GameUpdate_StatusUpdate); ok {
		return x.StatusUpdate
	}
	return nil
}

func (m *GameUpdate) GetTileUpdate() *TileUpdate {
	if x, ok := m.GetUpdate().(*GameUpdate_TileUpdate); ok {
		return x.TileUpdate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GameUpdate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GameUpdate_OneofMarshaler, _GameUpdate_OneofUnmarshaler, _GameUpdate_OneofSizer, []interface{}{
		(*GameUpdate_PlayerUpdate)(nil),
		(*GameUpdate_StatusUpdate)(nil),
		(*GameUpdate_TileUpdate)(nil),
	}
}

func _GameUpdate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GameUpdate)
	// update
	switch x := m.Update.(type) {
	case *GameUpdate_PlayerUpdate:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlayerUpdate); err != nil {
			return err
		}
	case *GameUpdate_StatusUpdate:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StatusUpdate); err != nil {
			return err
		}
	case *GameUpdate_TileUpdate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TileUpdate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GameUpdate.Update has unexpected type %T", x)
	}
	return nil
}

func _GameUpdate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GameUpdate)
	switch tag {
	case 1: // update.player_update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlayerUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &GameUpdate_PlayerUpdate{msg}
		return true, err
	case 2: // update.status_update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StatusUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &GameUpdate_StatusUpdate{msg}
		return true, err
	case 3: // update.tile_update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TileUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &GameUpdate_TileUpdate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GameUpdate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GameUpdate)
	// update
	switch x := m.Update.(type) {
	case *GameUpdate_PlayerUpdate:
		s := proto.Size(x.PlayerUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GameUpdate_StatusUpdate:
		s := proto.Size(x.StatusUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GameUpdate_TileUpdate:
		s := proto.Size(x.TileUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PlayerUpdate struct {
	Players              []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlayerUpdate) Reset()         { *m = PlayerUpdate{} }
func (m *PlayerUpdate) String() string { return proto.CompactTextString(m) }
func (*PlayerUpdate) ProtoMessage()    {}
func (*PlayerUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{4}
}

func (m *PlayerUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerUpdate.Unmarshal(m, b)
}
func (m *PlayerUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerUpdate.Marshal(b, m, deterministic)
}
func (m *PlayerUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerUpdate.Merge(m, src)
}
func (m *PlayerUpdate) XXX_Size() int {
	return xxx_messageInfo_PlayerUpdate.Size(m)
}
func (m *PlayerUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerUpdate proto.InternalMessageInfo

func (m *PlayerUpdate) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

type Player struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{5}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StatusUpdate struct {
	Status               StatusUpdate_Status `protobuf:"varint,1,opt,name=status,proto3,enum=StatusUpdate_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StatusUpdate) Reset()         { *m = StatusUpdate{} }
func (m *StatusUpdate) String() string { return proto.CompactTextString(m) }
func (*StatusUpdate) ProtoMessage()    {}
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{6}
}

func (m *StatusUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusUpdate.Unmarshal(m, b)
}
func (m *StatusUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusUpdate.Marshal(b, m, deterministic)
}
func (m *StatusUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusUpdate.Merge(m, src)
}
func (m *StatusUpdate) XXX_Size() int {
	return xxx_messageInfo_StatusUpdate.Size(m)
}
func (m *StatusUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_StatusUpdate proto.InternalMessageInfo

func (m *StatusUpdate) GetStatus() StatusUpdate_Status {
	if m != nil {
		return m.Status
	}
	return StatusUpdate_UNKNOWN
}

type TileUpdate struct {
	AllTiles             *Tiles   `protobuf:"bytes,1,opt,name=all_tiles,json=allTiles,proto3" json:"all_tiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TileUpdate) Reset()         { *m = TileUpdate{} }
func (m *TileUpdate) String() string { return proto.CompactTextString(m) }
func (*TileUpdate) ProtoMessage()    {}
func (*TileUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{7}
}

func (m *TileUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TileUpdate.Unmarshal(m, b)
}
func (m *TileUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TileUpdate.Marshal(b, m, deterministic)
}
func (m *TileUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TileUpdate.Merge(m, src)
}
func (m *TileUpdate) XXX_Size() int {
	return xxx_messageInfo_TileUpdate.Size(m)
}
func (m *TileUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_TileUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_TileUpdate proto.InternalMessageInfo

func (m *TileUpdate) GetAllTiles() *Tiles {
	if m != nil {
		return m.AllTiles
	}
	return nil
}

type Tiles struct {
	// Tiles are just lists of letters
	Letters              []string `protobuf:"bytes,1,rep,name=letters,proto3" json:"letters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tiles) Reset()         { *m = Tiles{} }
func (m *Tiles) String() string { return proto.CompactTextString(m) }
func (*Tiles) ProtoMessage()    {}
func (*Tiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{8}
}

func (m *Tiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tiles.Unmarshal(m, b)
}
func (m *Tiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tiles.Marshal(b, m, deterministic)
}
func (m *Tiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tiles.Merge(m, src)
}
func (m *Tiles) XXX_Size() int {
	return xxx_messageInfo_Tiles.Size(m)
}
func (m *Tiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Tiles.DiscardUnknown(m)
}

var xxx_messageInfo_Tiles proto.InternalMessageInfo

func (m *Tiles) GetLetters() []string {
	if m != nil {
		return m.Letters
	}
	return nil
}

type DumpRequest struct {
	// The unique identifier for the game you're
	// dumping in.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpRequest) Reset()         { *m = DumpRequest{} }
func (m *DumpRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRequest) ProtoMessage()    {}
func (*DumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{9}
}

func (m *DumpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpRequest.Unmarshal(m, b)
}
func (m *DumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpRequest.Marshal(b, m, deterministic)
}
func (m *DumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRequest.Merge(m, src)
}
func (m *DumpRequest) XXX_Size() int {
	return xxx_messageInfo_DumpRequest.Size(m)
}
func (m *DumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRequest proto.InternalMessageInfo

func (m *DumpRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DumpResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpResponse) Reset()         { *m = DumpResponse{} }
func (m *DumpResponse) String() string { return proto.CompactTextString(m) }
func (*DumpResponse) ProtoMessage()    {}
func (*DumpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{10}
}

func (m *DumpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpResponse.Unmarshal(m, b)
}
func (m *DumpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpResponse.Marshal(b, m, deterministic)
}
func (m *DumpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpResponse.Merge(m, src)
}
func (m *DumpResponse) XXX_Size() int {
	return xxx_messageInfo_DumpResponse.Size(m)
}
func (m *DumpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DumpResponse proto.InternalMessageInfo

type PeelRequest struct {
	// The unique identifier for the game you're placing in.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Board                *Board   `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeelRequest) Reset()         { *m = PeelRequest{} }
func (m *PeelRequest) String() string { return proto.CompactTextString(m) }
func (*PeelRequest) ProtoMessage()    {}
func (*PeelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{11}
}

func (m *PeelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeelRequest.Unmarshal(m, b)
}
func (m *PeelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeelRequest.Marshal(b, m, deterministic)
}
func (m *PeelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeelRequest.Merge(m, src)
}
func (m *PeelRequest) XXX_Size() int {
	return xxx_messageInfo_PeelRequest.Size(m)
}
func (m *PeelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeelRequest proto.InternalMessageInfo

func (m *PeelRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeelRequest) GetBoard() *Board {
	if m != nil {
		return m.Board
	}
	return nil
}

type Board struct {
	Words                []*Word  `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Board) Reset()         { *m = Board{} }
func (m *Board) String() string { return proto.CompactTextString(m) }
func (*Board) ProtoMessage()    {}
func (*Board) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{12}
}

func (m *Board) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Board.Unmarshal(m, b)
}
func (m *Board) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Board.Marshal(b, m, deterministic)
}
func (m *Board) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Board.Merge(m, src)
}
func (m *Board) XXX_Size() int {
	return xxx_messageInfo_Board.Size(m)
}
func (m *Board) XXX_DiscardUnknown() {
	xxx_messageInfo_Board.DiscardUnknown(m)
}

var xxx_messageInfo_Board proto.InternalMessageInfo

func (m *Board) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

type Word struct {
	Letters     string           `protobuf:"bytes,1,opt,name=letters,proto3" json:"letters,omitempty"`
	Orientation Word_Orientation `protobuf:"varint,2,opt,name=orientation,proto3,enum=Word_Orientation" json:"orientation,omitempty"`
	// The x-position of the top-left letter on the board.
	X int32 `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	// The y-position of the top-left letter on the board.
	Y                    int32    `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{13}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetLetters() string {
	if m != nil {
		return m.Letters
	}
	return ""
}

func (m *Word) GetOrientation() Word_Orientation {
	if m != nil {
		return m.Orientation
	}
	return Word_UNKNOWN
}

func (m *Word) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Word) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type PeelResponse struct {
	Status               PeelResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=PeelResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PeelResponse) Reset()         { *m = PeelResponse{} }
func (m *PeelResponse) String() string { return proto.CompactTextString(m) }
func (*PeelResponse) ProtoMessage()    {}
func (*PeelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0022f69d17e210ba, []int{14}
}

func (m *PeelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeelResponse.Unmarshal(m, b)
}
func (m *PeelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeelResponse.Marshal(b, m, deterministic)
}
func (m *PeelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeelResponse.Merge(m, src)
}
func (m *PeelResponse) XXX_Size() int {
	return xxx_messageInfo_PeelResponse.Size(m)
}
func (m *PeelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeelResponse proto.InternalMessageInfo

func (m *PeelResponse) GetStatus() PeelResponse_Status {
	if m != nil {
		return m.Status
	}
	return PeelResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("StatusUpdate_Status", StatusUpdate_Status_name, StatusUpdate_Status_value)
	proto.RegisterEnum("Word_Orientation", Word_Orientation_name, Word_Orientation_value)
	proto.RegisterEnum("PeelResponse_Status", PeelResponse_Status_name, PeelResponse_Status_value)
	proto.RegisterType((*NewGameRequest)(nil), "NewGameRequest")
	proto.RegisterType((*NewGameResponse)(nil), "NewGameResponse")
	proto.RegisterType((*JoinGameRequest)(nil), "JoinGameRequest")
	proto.RegisterType((*GameUpdate)(nil), "GameUpdate")
	proto.RegisterType((*PlayerUpdate)(nil), "PlayerUpdate")
	proto.RegisterType((*Player)(nil), "Player")
	proto.RegisterType((*StatusUpdate)(nil), "StatusUpdate")
	proto.RegisterType((*TileUpdate)(nil), "TileUpdate")
	proto.RegisterType((*Tiles)(nil), "Tiles")
	proto.RegisterType((*DumpRequest)(nil), "DumpRequest")
	proto.RegisterType((*DumpResponse)(nil), "DumpResponse")
	proto.RegisterType((*PeelRequest)(nil), "PeelRequest")
	proto.RegisterType((*Board)(nil), "Board")
	proto.RegisterType((*Word)(nil), "Word")
	proto.RegisterType((*PeelResponse)(nil), "PeelResponse")
}

func init() { proto.RegisterFile("banana.proto", fileDescriptor_0022f69d17e210ba) }

var fileDescriptor_0022f69d17e210ba = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0xb8, 0x71, 0x7e, 0xae, 0x9d, 0xc4, 0x9d, 0xef, 0x5b, 0x44, 0xa1, 0x48, 0xcd, 0x94,
	0x45, 0x25, 0xd0, 0x88, 0xa6, 0x80, 0x90, 0x58, 0x39, 0x89, 0x69, 0x03, 0x21, 0xae, 0x26, 0x6e,
	0x8b, 0xba, 0xb1, 0x26, 0xcd, 0x20, 0x59, 0x72, 0x63, 0x13, 0x3b, 0xb4, 0x7d, 0x26, 0x24, 0x16,
	0x88, 0x07, 0x44, 0x1e, 0xdb, 0x89, 0x13, 0xda, 0x9d, 0xcf, 0x99, 0x7b, 0xe6, 0xfe, 0xcc, 0xf1,
	0x05, 0x7d, 0xca, 0xe7, 0x7c, 0xce, 0x69, 0xb8, 0x08, 0xe2, 0x80, 0x9c, 0x42, 0x63, 0x2c, 0xee,
	0x4e, 0xf9, 0xad, 0x60, 0xe2, 0xfb, 0x52, 0x44, 0x31, 0xc6, 0x50, 0x9a, 0xf3, 0x5b, 0xd1, 0x42,
	0x07, 0xe8, 0xa8, 0xc6, 0xe4, 0x37, 0xee, 0x80, 0x1e, 0xdd, 0x70, 0x5f, 0xb8, 0xdf, 0xf8, 0x4d,
	0x1c, 0x2c, 0x5a, 0xca, 0x01, 0x3a, 0x52, 0x99, 0x26, 0xb9, 0x8f, 0x92, 0x22, 0x1d, 0x68, 0xae,
	0x2e, 0x8a, 0xc2, 0x60, 0x1e, 0x09, 0xdc, 0x00, 0xc5, 0x9b, 0x65, 0xf7, 0x28, 0xde, 0x8c, 0xbc,
	0x85, 0xe6, 0xa7, 0xc0, 0x9b, 0x17, 0x93, 0x6d, 0x85, 0xac, 0x92, 0x2b, 0xeb, 0xe4, 0xe4, 0x37,
	0x02, 0x48, 0x34, 0x17, 0xe1, 0x8c, 0xc7, 0x02, 0xbf, 0x81, 0x7a, 0xe8, 0xf3, 0x07, 0xb1, 0x70,
	0x97, 0x92, 0x90, 0x6a, 0xad, 0x5b, 0xa7, 0xe7, 0x92, 0x4d, 0xa3, 0xce, 0x76, 0x98, 0x1e, 0x16,
	0x70, 0xa2, 0x8a, 0x62, 0x1e, 0x2f, 0xa3, 0x5c, 0xa5, 0x64, 0xaa, 0x89, 0x64, 0xd7, 0xaa, 0xa8,
	0x80, 0x31, 0x05, 0x2d, 0xf6, 0x7c, 0x91, 0x6b, 0x76, 0xa5, 0x46, 0xa3, 0x8e, 0xe7, 0x8b, 0x95,
	0x02, 0xe2, 0x15, 0xea, 0x55, 0xa1, 0x9c, 0x86, 0x92, 0x63, 0xd0, 0x8b, 0xf5, 0xe0, 0x0e, 0x54,
	0xd2, 0x7a, 0xa2, 0x16, 0x3a, 0xd8, 0x3d, 0xd2, 0xba, 0x95, 0xac, 0x5e, 0x96, 0xf3, 0x64, 0x1f,
	0xca, 0x29, 0xf5, 0xd8, 0x13, 0x90, 0x18, 0xf4, 0x62, 0xa9, 0xf8, 0x15, 0x94, 0xd3, 0x52, 0x65,
	0x54, 0xa3, 0xfb, 0xff, 0x46, 0x27, 0x19, 0x60, 0x59, 0x0c, 0x79, 0x07, 0xe5, 0x94, 0xc1, 0x1a,
	0x54, 0x2e, 0xc6, 0x9f, 0xc7, 0xf6, 0xd5, 0xd8, 0xd8, 0xc1, 0x06, 0xe8, 0xa7, 0xe6, 0x17, 0xcb,
	0x9d, 0x38, 0x26, 0x73, 0xac, 0x81, 0x81, 0x70, 0x1d, 0x6a, 0x92, 0xb1, 0x2f, 0x2d, 0x66, 0x28,
	0xe4, 0x18, 0x60, 0xdd, 0x2c, 0x3e, 0x84, 0x1a, 0xf7, 0x7d, 0x37, 0x69, 0x38, 0xca, 0xc6, 0x5e,
	0x96, 0xc3, 0x88, 0x58, 0x95, 0xfb, 0xbe, 0xfc, 0x22, 0x1d, 0x50, 0xe5, 0x07, 0x6e, 0x41, 0xc5,
	0x17, 0x71, 0x9c, 0xb7, 0x5c, 0x63, 0x39, 0x24, 0xcf, 0x41, 0x1b, 0x2c, 0x6f, 0xc3, 0x27, 0x4c,
	0x40, 0x1a, 0xa0, 0xa7, 0xc7, 0xa9, 0x8f, 0xc8, 0x07, 0xd0, 0xce, 0x85, 0xf0, 0x9f, 0xf2, 0xcc,
	0x3e, 0xa8, 0xd3, 0x80, 0x2f, 0x66, 0xd9, 0x93, 0x96, 0x69, 0x2f, 0x41, 0x2c, 0x25, 0xc9, 0x0b,
	0x50, 0x25, 0xc6, 0xcf, 0x40, 0xbd, 0x0b, 0x16, 0xb3, 0x7c, 0xfe, 0x2a, 0xbd, 0x0a, 0x92, 0x28,
	0xc9, 0x91, 0x9f, 0x08, 0x4a, 0x09, 0xde, 0x2c, 0x1a, 0x15, 0x8a, 0xc6, 0x27, 0xa0, 0x05, 0x0b,
	0x4f, 0xcc, 0x63, 0x1e, 0x7b, 0xc1, 0x5c, 0x26, 0x6b, 0x74, 0xf7, 0xe4, 0x2d, 0xd4, 0x5e, 0x1f,
	0xb0, 0x62, 0x14, 0xd6, 0x01, 0xdd, 0x4b, 0xdb, 0xa8, 0x0c, 0xdd, 0x27, 0xe8, 0xa1, 0x55, 0x4a,
	0xd1, 0x03, 0x79, 0x0f, 0x5a, 0x41, 0xb7, 0xf9, 0x30, 0x0d, 0x80, 0x33, 0x9b, 0x0d, 0xaf, 0xed,
	0xb1, 0x63, 0x8e, 0x0c, 0x84, 0x75, 0xa8, 0x5e, 0x5a, 0xcc, 0x19, 0xf6, 0xcd, 0x91, 0xa1, 0x90,
	0x5f, 0x08, 0xf4, 0x74, 0x22, 0xd9, 0x9f, 0xf6, 0xaf, 0x19, 0x8a, 0xc7, 0xdb, 0x66, 0x08, 0x1f,
	0x37, 0x83, 0x06, 0x95, 0xc9, 0x45, 0xbf, 0x6f, 0x4d, 0x26, 0x06, 0x4a, 0x9c, 0x31, 0x1c, 0x5f,
	0x9a, 0xa3, 0xe1, 0xc0, 0xbd, 0xb2, 0xd9, 0xc0, 0x50, 0x30, 0x86, 0xc6, 0xc0, 0x72, 0xcc, 0xfe,
	0x99, 0x35, 0x70, 0x7b, 0xb6, 0xc9, 0x06, 0xc6, 0x2e, 0xfe, 0x0f, 0x9a, 0x63, 0xdb, 0x71, 0xcd,
	0xd1, 0xc8, 0x1d, 0x59, 0x8e, 0x63, 0xb1, 0x89, 0x51, 0xc2, 0x7b, 0x50, 0xb7, 0xbe, 0x3a, 0xcc,
	0x5c, 0x51, 0x6a, 0xf7, 0x0f, 0x82, 0x7a, 0x4f, 0xae, 0x9d, 0x89, 0x58, 0xfc, 0xf0, 0x6e, 0x92,
	0x8a, 0x2b, 0xd9, 0xba, 0xc0, 0x4d, 0xba, 0xb9, 0x81, 0xda, 0x06, 0xdd, 0xde, 0x24, 0x2f, 0xa1,
	0x9a, 0x6f, 0x0e, 0x6c, 0xd0, 0xad, 0x25, 0xd2, 0xd6, 0xe8, 0x7a, 0x3d, 0xbc, 0x46, 0xf8, 0x10,
	0x4a, 0x49, 0xf7, 0x58, 0xa7, 0x05, 0xd7, 0xb4, 0xeb, 0x1b, 0x23, 0x49, 0x82, 0x12, 0x8f, 0x61,
	0x9d, 0x16, 0x9c, 0xd8, 0xae, 0xd3, 0xa2, 0xf1, 0x7a, 0xa5, 0x6b, 0x25, 0x9c, 0x4e, 0xcb, 0x72,
	0x53, 0x9e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x7a, 0x95, 0x58, 0x39, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BananaServiceClient is the client API for BananaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BananaServiceClient interface {
	NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error)
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (BananaService_JoinGameClient, error)
	Peel(ctx context.Context, in *PeelRequest, opts ...grpc.CallOption) (*PeelResponse, error)
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*DumpResponse, error)
}

type bananaServiceClient struct {
	cc *grpc.ClientConn
}

func NewBananaServiceClient(cc *grpc.ClientConn) BananaServiceClient {
	return &bananaServiceClient{cc}
}

func (c *bananaServiceClient) NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error) {
	out := new(NewGameResponse)
	err := c.cc.Invoke(ctx, "/BananaService/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bananaServiceClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (BananaService_JoinGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BananaService_serviceDesc.Streams[0], "/BananaService/JoinGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &bananaServiceJoinGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BananaService_JoinGameClient interface {
	Recv() (*GameUpdate, error)
	grpc.ClientStream
}

type bananaServiceJoinGameClient struct {
	grpc.ClientStream
}

func (x *bananaServiceJoinGameClient) Recv() (*GameUpdate, error) {
	m := new(GameUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bananaServiceClient) Peel(ctx context.Context, in *PeelRequest, opts ...grpc.CallOption) (*PeelResponse, error) {
	out := new(PeelResponse)
	err := c.cc.Invoke(ctx, "/BananaService/Peel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bananaServiceClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*DumpResponse, error) {
	out := new(DumpResponse)
	err := c.cc.Invoke(ctx, "/BananaService/Dump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BananaServiceServer is the server API for BananaService service.
type BananaServiceServer interface {
	NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error)
	JoinGame(*JoinGameRequest, BananaService_JoinGameServer) error
	Peel(context.Context, *PeelRequest) (*PeelResponse, error)
	Dump(context.Context, *DumpRequest) (*DumpResponse, error)
}

func RegisterBananaServiceServer(s *grpc.Server, srv BananaServiceServer) {
	s.RegisterService(&_BananaService_serviceDesc, srv)
}

func _BananaService_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BananaServiceServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BananaService/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BananaServiceServer).NewGame(ctx, req.(*NewGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BananaService_JoinGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinGameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BananaServiceServer).JoinGame(m, &bananaServiceJoinGameServer{stream})
}

type BananaService_JoinGameServer interface {
	Send(*GameUpdate) error
	grpc.ServerStream
}

type bananaServiceJoinGameServer struct {
	grpc.ServerStream
}

func (x *bananaServiceJoinGameServer) Send(m *GameUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _BananaService_Peel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BananaServiceServer).Peel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BananaService/Peel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BananaServiceServer).Peel(ctx, req.(*PeelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BananaService_Dump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BananaServiceServer).Dump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BananaService/Dump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BananaServiceServer).Dump(ctx, req.(*DumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BananaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BananaService",
	HandlerType: (*BananaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _BananaService_NewGame_Handler,
		},
		{
			MethodName: "Peel",
			Handler:    _BananaService_Peel_Handler,
		},
		{
			MethodName: "Dump",
			Handler:    _BananaService_Dump_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinGame",
			Handler:       _BananaService_JoinGame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "banana.proto",
}
