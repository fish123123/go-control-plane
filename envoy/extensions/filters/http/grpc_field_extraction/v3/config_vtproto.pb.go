//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_field_extraction/v3/config.proto

package grpc_field_extractionv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *GrpcFieldExtractionConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcFieldExtractionConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *GrpcFieldExtractionConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.ExtractionsByMethod) > 0 {
		for k := range m.ExtractionsByMethod {
			v := m.ExtractionsByMethod[k]
			baseI := i
			size, err := v.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.DescriptorSet != nil {
		if vtmsg, ok := interface{}(m.DescriptorSet).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.DescriptorSet)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FieldExtractions) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldExtractions) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FieldExtractions) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.RequestFieldExtractions) > 0 {
		for k := range m.RequestFieldExtractions {
			v := m.RequestFieldExtractions[k]
			baseI := i
			size, err := v.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestFieldValueDisposition) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestFieldValueDisposition) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RequestFieldValueDisposition) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.Disposition.(*RequestFieldValueDisposition_DynamicMetadata); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RequestFieldValueDisposition_DynamicMetadata) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RequestFieldValueDisposition_DynamicMetadata) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.DynamicMetadata)
	copy(dAtA[i:], m.DynamicMetadata)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DynamicMetadata)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *GrpcFieldExtractionConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DescriptorSet != nil {
		if size, ok := interface{}(m.DescriptorSet).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.DescriptorSet)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ExtractionsByMethod) > 0 {
		for k, v := range m.ExtractionsByMethod {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.SizeVT()
			}
			l += 1 + protohelpers.SizeOfVarint(uint64(l))
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + l
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *FieldExtractions) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RequestFieldExtractions) > 0 {
		for k, v := range m.RequestFieldExtractions {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.SizeVT()
			}
			l += 1 + protohelpers.SizeOfVarint(uint64(l))
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + l
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *RequestFieldValueDisposition) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Disposition.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RequestFieldValueDisposition_DynamicMetadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DynamicMetadata)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
