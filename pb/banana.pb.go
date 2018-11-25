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
	return fileDescriptor_0022f69d17e210ba, []int{7, 0}
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
	return fileDescriptor_0022f69d17e210ba, []int{0}
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
	return fileDescriptor_0022f69d17e210ba, []int{1}
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
	return fileDescriptor_0022f69d17e210ba, []int{2}
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
	return fileDescriptor_0022f69d17e210ba, []int{3}
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
	return fileDescriptor_0022f69d17e210ba, []int{4}
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
	return fileDescriptor_0022f69d17e210ba, []int{5}
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
	return fileDescriptor_0022f69d17e210ba, []int{6}
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
	return fileDescriptor_0022f69d17e210ba, []int{7}
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
	return fileDescriptor_0022f69d17e210ba, []int{8}
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

func init() {
	proto.RegisterEnum("StatusUpdate_Status", StatusUpdate_Status_name, StatusUpdate_Status_value)
	proto.RegisterType((*Tiles)(nil), "Tiles")
	proto.RegisterType((*NewGameRequest)(nil), "NewGameRequest")
	proto.RegisterType((*NewGameResponse)(nil), "NewGameResponse")
	proto.RegisterType((*JoinGameRequest)(nil), "JoinGameRequest")
	proto.RegisterType((*GameUpdate)(nil), "GameUpdate")
	proto.RegisterType((*PlayerUpdate)(nil), "PlayerUpdate")
	proto.RegisterType((*Player)(nil), "Player")
	proto.RegisterType((*StatusUpdate)(nil), "StatusUpdate")
	proto.RegisterType((*TileUpdate)(nil), "TileUpdate")
	proto.RegisterType((*DumpRequest)(nil), "DumpRequest")
	proto.RegisterType((*DumpResponse)(nil), "DumpResponse")
}

func init() { proto.RegisterFile("banana.proto", fileDescriptor_0022f69d17e210ba) }

var fileDescriptor_0022f69d17e210ba = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdd, 0x8e, 0xda, 0x3c,
	0x10, 0xdd, 0x64, 0xd9, 0x04, 0x26, 0x09, 0x44, 0xa3, 0xef, 0x02, 0xa1, 0xaf, 0x12, 0x78, 0x6f,
	0x90, 0x5a, 0x59, 0x5d, 0xfa, 0x73, 0xbf, 0x68, 0x29, 0x55, 0xab, 0xb2, 0x95, 0x61, 0x5b, 0xa9,
	0x37, 0xc8, 0x80, 0x2b, 0x45, 0x32, 0x49, 0x1a, 0x9b, 0x56, 0x7d, 0x8a, 0xbe, 0x4b, 0x9f, 0xb0,
	0x8a, 0x9d, 0x40, 0x16, 0xb5, 0x77, 0x9e, 0x93, 0x73, 0x3c, 0x67, 0x3c, 0x27, 0x10, 0x6e, 0x78,
	0xca, 0x53, 0x4e, 0xf3, 0x22, 0xd3, 0x19, 0x19, 0xc1, 0xd5, 0x2a, 0x91, 0x42, 0x61, 0x1f, 0x7c,
	0x29, 0xb4, 0x16, 0x85, 0xea, 0x3b, 0xc3, 0xcb, 0x71, 0x87, 0xd5, 0x25, 0x99, 0x43, 0x77, 0x21,
	0x7e, 0xcc, 0xf9, 0x5e, 0x30, 0xf1, 0xed, 0x20, 0x94, 0x46, 0x84, 0x56, 0xca, 0xf7, 0xa2, 0xef,
	0x0c, 0x9d, 0x71, 0x87, 0x99, 0x33, 0x8e, 0x20, 0x54, 0x5b, 0x2e, 0xc5, 0xfa, 0x2b, 0xdf, 0xea,
	0xac, 0xe8, 0xbb, 0x43, 0x67, 0x7c, 0xc5, 0x02, 0x83, 0xbd, 0x31, 0x10, 0x19, 0x41, 0xef, 0x78,
	0x91, 0xca, 0xb3, 0x54, 0x09, 0xec, 0x82, 0x9b, 0xec, 0xaa, 0x7b, 0xdc, 0x64, 0x47, 0x5e, 0x41,
	0xef, 0x5d, 0x96, 0xa4, 0xcd, 0x66, 0x67, 0x94, 0x63, 0x73, 0xf7, 0xd4, 0x9c, 0xfc, 0x76, 0x00,
	0x4a, 0xcd, 0x43, 0xbe, 0xe3, 0x5a, 0xe0, 0x4b, 0x88, 0x72, 0xc9, 0x7f, 0x8a, 0x62, 0x7d, 0x30,
	0x80, 0x51, 0x07, 0x93, 0x88, 0x7e, 0x34, 0xa8, 0x65, 0xbd, 0xbd, 0x60, 0x61, 0xde, 0xa8, 0x4b,
	0x95, 0xd2, 0x5c, 0x1f, 0x54, 0xad, 0x72, 0x2b, 0xd5, 0xd2, 0xa0, 0x27, 0x95, 0x6a, 0xd4, 0x48,
	0x21, 0xd0, 0x89, 0x14, 0xb5, 0xe6, 0xd2, 0x68, 0x02, 0x5a, 0x3e, 0xea, 0x51, 0x01, 0xfa, 0x58,
	0x4d, 0xdb, 0xe0, 0x59, 0x2a, 0xb9, 0x81, 0xb0, 0xe9, 0x07, 0x47, 0xe0, 0x5b, 0x3f, 0x76, 0x03,
	0xc1, 0xc4, 0xaf, 0xfc, 0xb2, 0x1a, 0x27, 0xff, 0x83, 0x67, 0xa1, 0xbf, 0xad, 0x80, 0x68, 0x08,
	0x9b, 0x56, 0xf1, 0x19, 0x78, 0xd6, 0xaa, 0x61, 0x75, 0x27, 0xff, 0x3d, 0x9a, 0xa4, 0x2a, 0x58,
	0xc5, 0x21, 0xaf, 0xc1, 0xb3, 0x08, 0x06, 0xe0, 0x3f, 0x2c, 0xde, 0x2f, 0xee, 0x3f, 0x2f, 0xe2,
	0x0b, 0x8c, 0x21, 0x9c, 0xdf, 0x7e, 0x98, 0xad, 0x97, 0xab, 0x5b, 0xb6, 0x9a, 0xdd, 0xc5, 0x0e,
	0x46, 0xd0, 0x31, 0xc8, 0xfd, 0xa7, 0x19, 0x8b, 0x5d, 0x72, 0x03, 0x70, 0x1a, 0x16, 0xaf, 0xa1,
	0xc3, 0xa5, 0x5c, 0x97, 0x03, 0xab, 0xea, 0xd9, 0x3d, 0xf3, 0x18, 0x8a, 0xb5, 0xb9, 0x94, 0xe6,
	0x44, 0x9e, 0x40, 0x70, 0x77, 0xd8, 0xe7, 0xff, 0xd8, 0x30, 0xe9, 0x42, 0x68, 0x3f, 0xdb, 0x90,
	0x4c, 0x7e, 0x39, 0x10, 0x4d, 0x4d, 0x68, 0x97, 0xa2, 0xf8, 0x9e, 0x6c, 0xcb, 0xc9, 0xfc, 0x2a,
	0x49, 0xd8, 0xa3, 0x8f, 0xc3, 0x39, 0x88, 0xe9, 0x79, 0xc8, 0x9e, 0x42, 0xbb, 0x0e, 0x15, 0xc6,
	0xf4, 0x2c, 0x5f, 0x83, 0x80, 0x9e, 0x92, 0xf3, 0xdc, 0xc1, 0x6b, 0x68, 0x95, 0xcd, 0x31, 0xa4,
	0x0d, 0x8b, 0x83, 0x88, 0x36, 0x1d, 0x4d, 0x5b, 0x5f, 0xdc, 0x7c, 0xb3, 0xf1, 0xcc, 0x2f, 0xf4,
	0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x43, 0xc5, 0xcf, 0x52, 0x03, 0x00, 0x00,
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
