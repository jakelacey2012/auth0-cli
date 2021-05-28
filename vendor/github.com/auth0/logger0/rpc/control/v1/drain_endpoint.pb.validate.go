// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc/control/v1/drain_endpoint.proto

package controlv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// define the regex for a UUID once up-front
var _drain_endpoint_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListDrainsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListDrainsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Tenant

	// no validation rules for PageSize

	// no validation rules for PageToken

	return nil
}

// ListDrainsRequestValidationError is the validation error returned by
// ListDrainsRequest.Validate if the designated constraints aren't met.
type ListDrainsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDrainsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDrainsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDrainsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDrainsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDrainsRequestValidationError) ErrorName() string {
	return "ListDrainsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListDrainsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDrainsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDrainsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDrainsRequestValidationError{}

// Validate checks the field values on ListDrainsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListDrainsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetDrains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDrainsResponseValidationError{
					field:  fmt.Sprintf("Drains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	return nil
}

// ListDrainsResponseValidationError is the validation error returned by
// ListDrainsResponse.Validate if the designated constraints aren't met.
type ListDrainsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDrainsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDrainsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDrainsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDrainsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDrainsResponseValidationError) ErrorName() string {
	return "ListDrainsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListDrainsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDrainsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDrainsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDrainsResponseValidationError{}

// Validate checks the field values on CreateDrainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateDrainRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTenant()) < 1 {
		return CreateDrainRequestValidationError{
			field:  "Tenant",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetSink() == nil {
		return CreateDrainRequestValidationError{
			field:  "Sink",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSink()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateDrainRequestValidationError{
				field:  "Sink",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateDrainRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateDrainRequestValidationError is the validation error returned by
// CreateDrainRequest.Validate if the designated constraints aren't met.
type CreateDrainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDrainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDrainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDrainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDrainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDrainRequestValidationError) ErrorName() string {
	return "CreateDrainRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDrainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDrainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDrainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDrainRequestValidationError{}

// Validate checks the field values on CreateDrainResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateDrainResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDrain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateDrainResponseValidationError{
				field:  "Drain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateDrainResponseValidationError is the validation error returned by
// CreateDrainResponse.Validate if the designated constraints aren't met.
type CreateDrainResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDrainResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDrainResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDrainResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDrainResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDrainResponseValidationError) ErrorName() string {
	return "CreateDrainResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDrainResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDrainResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDrainResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDrainResponseValidationError{}

// Validate checks the field values on UpdateDrainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDrainRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return UpdateDrainRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if m.GetSink() == nil {
		return UpdateDrainRequestValidationError{
			field:  "Sink",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSink()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateDrainRequestValidationError{
				field:  "Sink",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateDrainRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

func (m *UpdateDrainRequest) _validateUuid(uuid string) error {
	if matched := _drain_endpoint_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdateDrainRequestValidationError is the validation error returned by
// UpdateDrainRequest.Validate if the designated constraints aren't met.
type UpdateDrainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDrainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDrainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDrainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDrainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDrainRequestValidationError) ErrorName() string {
	return "UpdateDrainRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDrainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDrainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDrainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDrainRequestValidationError{}

// Validate checks the field values on UpdateDrainResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDrainResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDrain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateDrainResponseValidationError{
				field:  "Drain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateDrainResponseValidationError is the validation error returned by
// UpdateDrainResponse.Validate if the designated constraints aren't met.
type UpdateDrainResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDrainResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDrainResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDrainResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDrainResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDrainResponseValidationError) ErrorName() string {
	return "UpdateDrainResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDrainResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDrainResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDrainResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDrainResponseValidationError{}

// Validate checks the field values on DeleteDrainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteDrainRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return DeleteDrainRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *DeleteDrainRequest) _validateUuid(uuid string) error {
	if matched := _drain_endpoint_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteDrainRequestValidationError is the validation error returned by
// DeleteDrainRequest.Validate if the designated constraints aren't met.
type DeleteDrainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDrainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDrainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDrainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDrainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDrainRequestValidationError) ErrorName() string {
	return "DeleteDrainRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDrainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDrainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDrainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDrainRequestValidationError{}