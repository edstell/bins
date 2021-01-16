package validation

import (
	"reflect"

	"github.com/edstell/lambda/libraries/errors"
	validationproto "github.com/edstell/lambda/tools/protoc-gen-service/proto/validation"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Validate is the only function exported from this package. It derives any
// validation checks required for the message passed, then performs each,
// returning the first error encountered.
func Validate(pb protoreflect.ProtoMessage) error {
	return validateMessage(pb.ProtoReflect())
}

func validateMessage(m protoreflect.Message) error {
	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if isRequired(fd) && !m.Has(fd) {
			return errors.MissingParam(name(fd))
		}
		if err := validate(fd, m.Get(fd)); err != nil {
			return err
		}
	}
	return nil
}

func validateMap(fd protoreflect.FieldDescriptor, v protoreflect.Value) error {
	var err error
	v.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		if err = validate(fd.MapValue(), v); err != nil {
			return false
		}
		return true
	})
	return err
}

func validateList(fd protoreflect.FieldDescriptor, v protoreflect.Value) error {
	list := v.List()
	for i := 0; i < list.Len(); i++ {
		if err := validate(nil, list.Get(i)); err != nil {
			return err
		}
	}
	return nil
}

func validate(fd protoreflect.FieldDescriptor, v protoreflect.Value) error {
	var err error
	switch fd.Kind() {
	case protoreflect.MessageKind:
		if fd.IsMap() {
			err = validateMap(fd, v)
		} else if fd.IsList() {
			err = validateList(fd, v)
		} else {
			err = validateMessage(v.Message())
		}
	default:
		if isRequired(fd) && !isPresent(fd, v) {
			return errors.MissingParam(name(fd))
		}
	}
	return err
}

func isRequired(fd protoreflect.FieldDescriptor) bool {
	if oneof := fd.ContainingOneof(); oneof != nil {
		opts := oneof.Options().(*descriptorpb.OneofOptions)
		isRequired, _ := proto.GetExtension(opts, validationproto.E_MustExist).(bool)
		return isRequired
	}
	opts := fd.Options().(*descriptorpb.FieldOptions)
	isRequired, _ := proto.GetExtension(opts, validationproto.E_Required).(bool)
	if isRequired {
		return true
	}
	return false
}

func name(fd protoreflect.FieldDescriptor) string {
	if oneof := fd.ContainingOneof(); oneof != nil {
		return string(oneof.Name())
	}
	return fd.TextName()
}

func isPresent(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
	val := v.Interface()
	typ := reflect.TypeOf(val)
	if !typ.Comparable() {
		panic("complex type encountered")
	}
	if reflect.Zero(typ) == reflect.ValueOf(val) {
		return false
	}
	return true
}
