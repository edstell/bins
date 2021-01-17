package validation

import (
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
		v := m.Get(fd)
		required := isRequired(fd)
		present := m.Has(fd)
		if required && !present {
			return errors.MissingParam(name(fd))
		}
		if !required && !present {
			return nil
		}
		if fd.IsMap() {
			if err := validateMap(fd.MapValue().Kind(), v); err != nil {
				return err
			}
		} else if fd.IsList() {
			if err := validateList(fd.Kind(), v); err != nil {
				return err
			}
		} else {
			if err := validate(fd.Kind(), v); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateMap(k protoreflect.Kind, v protoreflect.Value) error {
	var err error
	v.Map().Range(func(_ protoreflect.MapKey, v protoreflect.Value) bool {
		if err = validate(k, v); err != nil {
			return false
		}
		return true
	})
	return err
}

func validateList(k protoreflect.Kind, v protoreflect.Value) error {
	list := v.List()
	for i := 0; i < list.Len(); i++ {
		if err := validate(k, list.Get(i)); err != nil {
			return err
		}
	}
	return nil
}

func validate(k protoreflect.Kind, v protoreflect.Value) error {
	if k == protoreflect.MessageKind {
		return validateMessage(v.Message())
	}
	// TODO once more validation options are added, they'll be checked here.
	return nil
}

func isRequired(fd protoreflect.FieldDescriptor) bool {
	if oneof := fd.ContainingOneof(); oneof != nil {
		opts := oneof.Options().(*descriptorpb.OneofOptions)
		isRequired, _ := proto.GetExtension(opts, validationproto.E_AtLeastOne).(bool)
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
