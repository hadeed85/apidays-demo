// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/example/com/aircraft/v1/aircraft_seats_service.proto

package aircraftv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request and response messages
type SeatStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The row number of the seat to query.
	RowNumber int32 `protobuf:"varint,1,opt,name=rowNumber,proto3" json:"rowNumber,omitempty"`
	// The letter of the seat to query.
	SeatLetter    string `protobuf:"bytes,2,opt,name=seatLetter,proto3" json:"seatLetter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeatStatusRequest) Reset() {
	*x = SeatStatusRequest{}
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeatStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatStatusRequest) ProtoMessage() {}

func (x *SeatStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatStatusRequest.ProtoReflect.Descriptor instead.
func (*SeatStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP(), []int{0}
}

func (x *SeatStatusRequest) GetRowNumber() int32 {
	if x != nil {
		return x.RowNumber
	}
	return 0
}

func (x *SeatStatusRequest) GetSeatLetter() string {
	if x != nil {
		return x.SeatLetter
	}
	return ""
}

type SeatStatusResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The row number of the seat.
	RowNumber int32 `protobuf:"varint,1,opt,name=rowNumber,proto3" json:"rowNumber,omitempty"`
	// The letter of the seat.
	SeatLetter string `protobuf:"bytes,2,opt,name=seatLetter,proto3" json:"seatLetter,omitempty"`
	// Whether the seat is occupied.
	Occupied      bool `protobuf:"varint,3,opt,name=occupied,proto3" json:"occupied,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeatStatusResponse) Reset() {
	*x = SeatStatusResponse{}
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeatStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatStatusResponse) ProtoMessage() {}

func (x *SeatStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatStatusResponse.ProtoReflect.Descriptor instead.
func (*SeatStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP(), []int{1}
}

func (x *SeatStatusResponse) GetRowNumber() int32 {
	if x != nil {
		return x.RowNumber
	}
	return 0
}

func (x *SeatStatusResponse) GetSeatLetter() string {
	if x != nil {
		return x.SeatLetter
	}
	return ""
}

func (x *SeatStatusResponse) GetOccupied() bool {
	if x != nil {
		return x.Occupied
	}
	return false
}

type UpdateSeatStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The row number of the seat to update.
	RowNumber int32 `protobuf:"varint,1,opt,name=rowNumber,proto3" json:"rowNumber,omitempty"`
	// The letter of the seat to update.
	SeatLetter string `protobuf:"bytes,2,opt,name=seatLetter,proto3" json:"seatLetter,omitempty"`
	// The new occupied status of the seat.
	Occupied      bool `protobuf:"varint,3,opt,name=occupied,proto3" json:"occupied,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSeatStatusRequest) Reset() {
	*x = UpdateSeatStatusRequest{}
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSeatStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeatStatusRequest) ProtoMessage() {}

func (x *UpdateSeatStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeatStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateSeatStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSeatStatusRequest) GetRowNumber() int32 {
	if x != nil {
		return x.RowNumber
	}
	return 0
}

func (x *UpdateSeatStatusRequest) GetSeatLetter() string {
	if x != nil {
		return x.SeatLetter
	}
	return ""
}

func (x *UpdateSeatStatusRequest) GetOccupied() bool {
	if x != nil {
		return x.Occupied
	}
	return false
}

type UpdateSeatStatusResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Indicates whether the update was successful.
	Success       bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSeatStatusResponse) Reset() {
	*x = UpdateSeatStatusResponse{}
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSeatStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeatStatusResponse) ProtoMessage() {}

func (x *UpdateSeatStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeatStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateSeatStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSeatStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SeatStatusSubscriptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeatStatusSubscriptionRequest) Reset() {
	*x = SeatStatusSubscriptionRequest{}
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeatStatusSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatStatusSubscriptionRequest) ProtoMessage() {}

func (x *SeatStatusSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatStatusSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SeatStatusSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP(), []int{4}
}

type SeatStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The row number of the seat.
	RowNumber int32 `protobuf:"varint,1,opt,name=rowNumber,proto3" json:"rowNumber,omitempty"`
	// The letter of the seat.
	SeatLetter string `protobuf:"bytes,2,opt,name=seatLetter,proto3" json:"seatLetter,omitempty"`
	// Whether the seat is occupied.
	Occupied      bool `protobuf:"varint,3,opt,name=occupied,proto3" json:"occupied,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SeatStatus) Reset() {
	*x = SeatStatus{}
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeatStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatStatus) ProtoMessage() {}

func (x *SeatStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatStatus.ProtoReflect.Descriptor instead.
func (*SeatStatus) Descriptor() ([]byte, []int) {
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP(), []int{5}
}

func (x *SeatStatus) GetRowNumber() int32 {
	if x != nil {
		return x.RowNumber
	}
	return 0
}

func (x *SeatStatus) GetSeatLetter() string {
	if x != nil {
		return x.SeatLetter
	}
	return ""
}

func (x *SeatStatus) GetOccupied() bool {
	if x != nil {
		return x.Occupied
	}
	return false
}

var File_proto_example_com_aircraft_v1_aircraft_seats_service_proto protoreflect.FileDescriptor

const file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDesc = "" +
	"\n" +
	":proto/example/com/aircraft/v1/aircraft_seats_service.proto\x12\x17example.com.aircraft.v1\"Q\n" +
	"\x11SeatStatusRequest\x12\x1c\n" +
	"\trowNumber\x18\x01 \x01(\x05R\trowNumber\x12\x1e\n" +
	"\n" +
	"seatLetter\x18\x02 \x01(\tR\n" +
	"seatLetter\"n\n" +
	"\x12SeatStatusResponse\x12\x1c\n" +
	"\trowNumber\x18\x01 \x01(\x05R\trowNumber\x12\x1e\n" +
	"\n" +
	"seatLetter\x18\x02 \x01(\tR\n" +
	"seatLetter\x12\x1a\n" +
	"\boccupied\x18\x03 \x01(\bR\boccupied\"s\n" +
	"\x17UpdateSeatStatusRequest\x12\x1c\n" +
	"\trowNumber\x18\x01 \x01(\x05R\trowNumber\x12\x1e\n" +
	"\n" +
	"seatLetter\x18\x02 \x01(\tR\n" +
	"seatLetter\x12\x1a\n" +
	"\boccupied\x18\x03 \x01(\bR\boccupied\"4\n" +
	"\x18UpdateSeatStatusResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\x1f\n" +
	"\x1dSeatStatusSubscriptionRequest\"f\n" +
	"\n" +
	"SeatStatus\x12\x1c\n" +
	"\trowNumber\x18\x01 \x01(\x05R\trowNumber\x12\x1e\n" +
	"\n" +
	"seatLetter\x18\x02 \x01(\tR\n" +
	"seatLetter\x12\x1a\n" +
	"\boccupied\x18\x03 \x01(\bR\boccupied2\xf8\x02\n" +
	"\x14AircraftSeatsService\x12h\n" +
	"\rGetSeatStatus\x12*.example.com.aircraft.v1.SeatStatusRequest\x1a+.example.com.aircraft.v1.SeatStatusResponse\x12w\n" +
	"\x10UpdateSeatStatus\x120.example.com.aircraft.v1.UpdateSeatStatusRequest\x1a1.example.com.aircraft.v1.UpdateSeatStatusResponse\x12}\n" +
	"\x1cSubscribeToSeatStatusUpdates\x126.example.com.aircraft.v1.SeatStatusSubscriptionRequest\x1a#.example.com.aircraft.v1.SeatStatus0\x01B\x88\x02\n" +
	"\x1bcom.example.com.aircraft.v1B\x19AircraftSeatsServiceProtoP\x01ZOaircraft-seats-service-go/internal/gen/proto/example/com/aircraft/v1;aircraftv1\xa2\x02\x03ECA\xaa\x02\x17Example.Com.Aircraft.V1\xca\x02\x17Example\\Com\\Aircraft\\V1\xe2\x02#Example\\Com\\Aircraft\\V1\\GPBMetadata\xea\x02\x1aExample::Com::Aircraft::V1b\x06proto3"

var (
	file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescOnce sync.Once
	file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescData []byte
)

func file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescGZIP() []byte {
	file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescOnce.Do(func() {
		file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDesc), len(file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDesc)))
	})
	return file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDescData
}

var file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_goTypes = []any{
	(*SeatStatusRequest)(nil),             // 0: example.com.aircraft.v1.SeatStatusRequest
	(*SeatStatusResponse)(nil),            // 1: example.com.aircraft.v1.SeatStatusResponse
	(*UpdateSeatStatusRequest)(nil),       // 2: example.com.aircraft.v1.UpdateSeatStatusRequest
	(*UpdateSeatStatusResponse)(nil),      // 3: example.com.aircraft.v1.UpdateSeatStatusResponse
	(*SeatStatusSubscriptionRequest)(nil), // 4: example.com.aircraft.v1.SeatStatusSubscriptionRequest
	(*SeatStatus)(nil),                    // 5: example.com.aircraft.v1.SeatStatus
}
var file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_depIdxs = []int32{
	0, // 0: example.com.aircraft.v1.AircraftSeatsService.GetSeatStatus:input_type -> example.com.aircraft.v1.SeatStatusRequest
	2, // 1: example.com.aircraft.v1.AircraftSeatsService.UpdateSeatStatus:input_type -> example.com.aircraft.v1.UpdateSeatStatusRequest
	4, // 2: example.com.aircraft.v1.AircraftSeatsService.SubscribeToSeatStatusUpdates:input_type -> example.com.aircraft.v1.SeatStatusSubscriptionRequest
	1, // 3: example.com.aircraft.v1.AircraftSeatsService.GetSeatStatus:output_type -> example.com.aircraft.v1.SeatStatusResponse
	3, // 4: example.com.aircraft.v1.AircraftSeatsService.UpdateSeatStatus:output_type -> example.com.aircraft.v1.UpdateSeatStatusResponse
	5, // 5: example.com.aircraft.v1.AircraftSeatsService.SubscribeToSeatStatusUpdates:output_type -> example.com.aircraft.v1.SeatStatus
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_init() }
func file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_init() {
	if File_proto_example_com_aircraft_v1_aircraft_seats_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDesc), len(file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_goTypes,
		DependencyIndexes: file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_depIdxs,
		MessageInfos:      file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_msgTypes,
	}.Build()
	File_proto_example_com_aircraft_v1_aircraft_seats_service_proto = out.File
	file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_goTypes = nil
	file_proto_example_com_aircraft_v1_aircraft_seats_service_proto_depIdxs = nil
}
