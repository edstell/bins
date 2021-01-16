package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/edstell/lambda/libraries/errors"
	validationproto "github.com/edstell/lambda/tools/protoc-gen-service/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Check implementations actually perform the validation.
type check interface {
	// Perform will run the expression implemented by this check, returning
	// an error if the expression failed (i.e. returned false).
	Perform() error
}

// Validate is the only function exported from this package. It derives any
// checks required for the message passed, then performs each, returning the
// first error encountered.
func Validate(pb protoreflect.ProtoMessage) error {
	return validateMessage(pb.ProtoReflect())
}

func validateMessage(m protoreflect.Message) error {
	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if isRequired(fd) && !m.Has(fd) {
			return errors.BadRequest(fmt.Sprintf("missing param: %s", fd.TextName()))
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
	fmt.Println(fd.TextName(), "kind", fd.Kind().String())
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
		opts := fd.Options().(*descriptorpb.FieldOptions)
		what, ok := proto.GetExtension(opts, validationproto.E_Validation).(string)
		if !ok {
			return nil
		}
		checks := createChecks(what, fd, v)
		for _, check := range checks {
			if err := check.Perform(); err != nil {
				return err
			}
		}
	}
	return err
}

func isRequired(fd protoreflect.FieldDescriptor) bool {
	opts := fd.Options().(*descriptorpb.FieldOptions)
	what, ok := proto.GetExtension(opts, validationproto.E_Validation).(string)
	if !ok {
		return false
	}
	return strings.Contains(what, "required")
}

func createChecks(what string, fd protoreflect.FieldDescriptor, v protoreflect.Value) []check {
	checks := []check{}
	for _, what := range strings.Split(what, ",") {
		switch what {
		case "required":
			checks = append(checks, &required{fd: fd, v: v})
		}
	}
	return checks
}

type required struct {
	fd protoreflect.FieldDescriptor
	v  protoreflect.Value
}

// Perform the 'required' check by checking whether the field value is its zero
// value (in which case the check fails).
func (r *required) Perform() error {
	val := r.v.Interface()
	typ := reflect.TypeOf(val)
	if !typ.Comparable() {
		panic("complex type encountered")
	}
	if reflect.Zero(typ) == reflect.ValueOf(val) {
		return errors.BadRequest(fmt.Sprintf("missing param: %s", r.fd.TextName()))
	}
	return nil
}
