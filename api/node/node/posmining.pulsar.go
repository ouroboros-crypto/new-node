// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package node

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Posmining                 protoreflect.MessageDescriptor
	fd_Posmining_owner           protoreflect.FieldDescriptor
	fd_Posmining_coin            protoreflect.FieldDescriptor
	fd_Posmining_lastTransaction protoreflect.FieldDescriptor
)

func init() {
	file_node_node_posmining_proto_init()
	md_Posmining = File_node_node_posmining_proto.Messages().ByName("Posmining")
	fd_Posmining_owner = md_Posmining.Fields().ByName("owner")
	fd_Posmining_coin = md_Posmining.Fields().ByName("coin")
	fd_Posmining_lastTransaction = md_Posmining.Fields().ByName("lastTransaction")
}

var _ protoreflect.Message = (*fastReflection_Posmining)(nil)

type fastReflection_Posmining Posmining

func (x *Posmining) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Posmining)(x)
}

func (x *Posmining) slowProtoReflect() protoreflect.Message {
	mi := &file_node_node_posmining_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Posmining_messageType fastReflection_Posmining_messageType
var _ protoreflect.MessageType = fastReflection_Posmining_messageType{}

type fastReflection_Posmining_messageType struct{}

func (x fastReflection_Posmining_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Posmining)(nil)
}
func (x fastReflection_Posmining_messageType) New() protoreflect.Message {
	return new(fastReflection_Posmining)
}
func (x fastReflection_Posmining_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Posmining
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Posmining) Descriptor() protoreflect.MessageDescriptor {
	return md_Posmining
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Posmining) Type() protoreflect.MessageType {
	return _fastReflection_Posmining_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Posmining) New() protoreflect.Message {
	return new(fastReflection_Posmining)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Posmining) Interface() protoreflect.ProtoMessage {
	return (*Posmining)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Posmining) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_Posmining_owner, value) {
			return
		}
	}
	if x.Coin != nil {
		value := protoreflect.ValueOfMessage(x.Coin.ProtoReflect())
		if !f(fd_Posmining_coin, value) {
			return
		}
	}
	if x.LastTransaction != "" {
		value := protoreflect.ValueOfString(x.LastTransaction)
		if !f(fd_Posmining_lastTransaction, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Posmining) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "node.node.Posmining.owner":
		return x.Owner != ""
	case "node.node.Posmining.coin":
		return x.Coin != nil
	case "node.node.Posmining.lastTransaction":
		return x.LastTransaction != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: node.node.Posmining"))
		}
		panic(fmt.Errorf("message node.node.Posmining does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Posmining) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "node.node.Posmining.owner":
		x.Owner = ""
	case "node.node.Posmining.coin":
		x.Coin = nil
	case "node.node.Posmining.lastTransaction":
		x.LastTransaction = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: node.node.Posmining"))
		}
		panic(fmt.Errorf("message node.node.Posmining does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Posmining) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "node.node.Posmining.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "node.node.Posmining.coin":
		value := x.Coin
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "node.node.Posmining.lastTransaction":
		value := x.LastTransaction
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: node.node.Posmining"))
		}
		panic(fmt.Errorf("message node.node.Posmining does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Posmining) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "node.node.Posmining.owner":
		x.Owner = value.Interface().(string)
	case "node.node.Posmining.coin":
		x.Coin = value.Message().Interface().(*v1beta1.Coin)
	case "node.node.Posmining.lastTransaction":
		x.LastTransaction = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: node.node.Posmining"))
		}
		panic(fmt.Errorf("message node.node.Posmining does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Posmining) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "node.node.Posmining.coin":
		if x.Coin == nil {
			x.Coin = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Coin.ProtoReflect())
	case "node.node.Posmining.owner":
		panic(fmt.Errorf("field owner of message node.node.Posmining is not mutable"))
	case "node.node.Posmining.lastTransaction":
		panic(fmt.Errorf("field lastTransaction of message node.node.Posmining is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: node.node.Posmining"))
		}
		panic(fmt.Errorf("message node.node.Posmining does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Posmining) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "node.node.Posmining.owner":
		return protoreflect.ValueOfString("")
	case "node.node.Posmining.coin":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "node.node.Posmining.lastTransaction":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: node.node.Posmining"))
		}
		panic(fmt.Errorf("message node.node.Posmining does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Posmining) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in node.node.Posmining", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Posmining) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Posmining) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Posmining) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Posmining) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Posmining)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Coin != nil {
			l = options.Size(x.Coin)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.LastTransaction)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Posmining)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.LastTransaction) > 0 {
			i -= len(x.LastTransaction)
			copy(dAtA[i:], x.LastTransaction)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LastTransaction)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Coin != nil {
			encoded, err := options.Marshal(x.Coin)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Posmining)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Posmining: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Posmining: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Coin == nil {
					x.Coin = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Coin); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastTransaction", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LastTransaction = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: node/node/posmining.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Posmining struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner           string        `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Coin            *v1beta1.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
	LastTransaction string        `protobuf:"bytes,3,opt,name=lastTransaction,proto3" json:"lastTransaction,omitempty"`
}

func (x *Posmining) Reset() {
	*x = Posmining{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_node_posmining_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Posmining) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posmining) ProtoMessage() {}

// Deprecated: Use Posmining.ProtoReflect.Descriptor instead.
func (*Posmining) Descriptor() ([]byte, []int) {
	return file_node_node_posmining_proto_rawDescGZIP(), []int{0}
}

func (x *Posmining) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Posmining) GetCoin() *v1beta1.Coin {
	if x != nil {
		return x.Coin
	}
	return nil
}

func (x *Posmining) GetLastTransaction() string {
	if x != nil {
		return x.LastTransaction
	}
	return ""
}

var File_node_node_posmining_proto protoreflect.FileDescriptor

var file_node_node_posmining_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a,
	0x09, 0x50, 0x6f, 0x73, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x04, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x94, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x42, 0x0e, 0x50, 0x6f, 0x73, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x75, 0x72, 0x6f, 0x62, 0x6f, 0x72, 0x6f, 0x73, 0x2d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0xa2, 0x02, 0x03, 0x4e, 0x4e, 0x58, 0xaa, 0x02, 0x09, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0xca, 0x02, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x4e, 0x6f, 0x64,
	0x65, 0xe2, 0x02, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x3a, 0x3a, 0x4e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_node_posmining_proto_rawDescOnce sync.Once
	file_node_node_posmining_proto_rawDescData = file_node_node_posmining_proto_rawDesc
)

func file_node_node_posmining_proto_rawDescGZIP() []byte {
	file_node_node_posmining_proto_rawDescOnce.Do(func() {
		file_node_node_posmining_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_node_posmining_proto_rawDescData)
	})
	return file_node_node_posmining_proto_rawDescData
}

var file_node_node_posmining_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_node_node_posmining_proto_goTypes = []interface{}{
	(*Posmining)(nil),    // 0: node.node.Posmining
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_node_node_posmining_proto_depIdxs = []int32{
	1, // 0: node.node.Posmining.coin:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_node_node_posmining_proto_init() }
func file_node_node_posmining_proto_init() {
	if File_node_node_posmining_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_node_posmining_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Posmining); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_node_posmining_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_node_posmining_proto_goTypes,
		DependencyIndexes: file_node_node_posmining_proto_depIdxs,
		MessageInfos:      file_node_node_posmining_proto_msgTypes,
	}.Build()
	File_node_node_posmining_proto = out.File
	file_node_node_posmining_proto_rawDesc = nil
	file_node_node_posmining_proto_goTypes = nil
	file_node_node_posmining_proto_depIdxs = nil
}
