// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/configure/resource/configure_resource_value.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on ListResourceValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourceValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourceValueRequestMultiError, or nil if none found.
func (m *ListResourceValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetResourceId() <= 0 {
		err := ListResourceValueRequestValidationError{
			field:  "ResourceId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListResourceValueRequestMultiError(errors)
	}

	return nil
}

// ListResourceValueRequestMultiError is an error wrapping multiple validation
// errors returned by ListResourceValueRequest.ValidateAll() if the designated
// constraints aren't met.
type ListResourceValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceValueRequestMultiError) AllErrors() []error { return m }

// ListResourceValueRequestValidationError is the validation error returned by
// ListResourceValueRequest.Validate if the designated constraints aren't met.
type ListResourceValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceValueRequestValidationError) ErrorName() string {
	return "ListResourceValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceValueRequestValidationError{}

// Validate checks the field values on ListResourceValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourceValueReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourceValueReplyMultiError, or nil if none found.
func (m *ListResourceValueReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceValueReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResourceValueReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResourceValueReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResourceValueReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListResourceValueReplyMultiError(errors)
	}

	return nil
}

// ListResourceValueReplyMultiError is an error wrapping multiple validation
// errors returned by ListResourceValueReply.ValidateAll() if the designated
// constraints aren't met.
type ListResourceValueReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceValueReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceValueReplyMultiError) AllErrors() []error { return m }

// ListResourceValueReplyValidationError is the validation error returned by
// ListResourceValueReply.Validate if the designated constraints aren't met.
type ListResourceValueReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceValueReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceValueReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceValueReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceValueReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceValueReplyValidationError) ErrorName() string {
	return "ListResourceValueReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceValueReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceValueReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceValueReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceValueReplyValidationError{}

// Validate checks the field values on UpdateResourceValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceValueRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceValueRequestMultiError, or nil if none found.
func (m *UpdateResourceValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetList()) < 1 {
		err := UpdateResourceValueRequestValidationError{
			field:  "List",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateResourceValueRequestValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateResourceValueRequestValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateResourceValueRequestValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetResourceId() <= 0 {
		err := UpdateResourceValueRequestValidationError{
			field:  "ResourceId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateResourceValueRequestMultiError(errors)
	}

	return nil
}

// UpdateResourceValueRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateResourceValueRequest.ValidateAll() if
// the designated constraints aren't met.
type UpdateResourceValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceValueRequestMultiError) AllErrors() []error { return m }

// UpdateResourceValueRequestValidationError is the validation error returned
// by UpdateResourceValueRequest.Validate if the designated constraints aren't met.
type UpdateResourceValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceValueRequestValidationError) ErrorName() string {
	return "UpdateResourceValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceValueRequestValidationError{}

// Validate checks the field values on UpdateResourceValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceValueReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceValueReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceValueReplyMultiError, or nil if none found.
func (m *UpdateResourceValueReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceValueReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateResourceValueReplyMultiError(errors)
	}

	return nil
}

// UpdateResourceValueReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateResourceValueReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateResourceValueReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceValueReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceValueReplyMultiError) AllErrors() []error { return m }

// UpdateResourceValueReplyValidationError is the validation error returned by
// UpdateResourceValueReply.Validate if the designated constraints aren't met.
type UpdateResourceValueReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceValueReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceValueReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceValueReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceValueReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceValueReplyValidationError) ErrorName() string {
	return "UpdateResourceValueReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceValueReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceValueReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceValueReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceValueReplyValidationError{}

// Validate checks the field values on ListResourceValueReply_ResourceValue
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *ListResourceValueReply_ResourceValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceValueReply_ResourceValue
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ListResourceValueReply_ResourceValueMultiError, or nil if none found.
func (m *ListResourceValueReply_ResourceValue) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceValueReply_ResourceValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for EnvId

	// no validation rules for ResourceId

	// no validation rules for Value

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return ListResourceValueReply_ResourceValueMultiError(errors)
	}

	return nil
}

// ListResourceValueReply_ResourceValueMultiError is an error wrapping multiple
// validation errors returned by
// ListResourceValueReply_ResourceValue.ValidateAll() if the designated
// constraints aren't met.
type ListResourceValueReply_ResourceValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceValueReply_ResourceValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceValueReply_ResourceValueMultiError) AllErrors() []error { return m }

// ListResourceValueReply_ResourceValueValidationError is the validation error
// returned by ListResourceValueReply_ResourceValue.Validate if the designated
// constraints aren't met.
type ListResourceValueReply_ResourceValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceValueReply_ResourceValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceValueReply_ResourceValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceValueReply_ResourceValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceValueReply_ResourceValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceValueReply_ResourceValueValidationError) ErrorName() string {
	return "ListResourceValueReply_ResourceValueValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceValueReply_ResourceValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceValueReply_ResourceValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceValueReply_ResourceValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceValueReply_ResourceValueValidationError{}

// Validate checks the field values on UpdateResourceValueRequest_Value with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *UpdateResourceValueRequest_Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceValueRequest_Value with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateResourceValueRequest_ValueMultiError, or nil if none found.
func (m *UpdateResourceValueRequest_Value) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceValueRequest_Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEnvId() <= 0 {
		err := UpdateResourceValueRequest_ValueValidationError{
			field:  "EnvId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := UpdateResourceValueRequest_ValueValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateResourceValueRequest_ValueMultiError(errors)
	}

	return nil
}

// UpdateResourceValueRequest_ValueMultiError is an error wrapping multiple
// validation errors returned by
// UpdateResourceValueRequest_Value.ValidateAll() if the designated
// constraints aren't met.
type UpdateResourceValueRequest_ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceValueRequest_ValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceValueRequest_ValueMultiError) AllErrors() []error { return m }

// UpdateResourceValueRequest_ValueValidationError is the validation error
// returned by UpdateResourceValueRequest_Value.Validate if the designated
// constraints aren't met.
type UpdateResourceValueRequest_ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceValueRequest_ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceValueRequest_ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceValueRequest_ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceValueRequest_ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceValueRequest_ValueValidationError) ErrorName() string {
	return "UpdateResourceValueRequest_ValueValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceValueRequest_ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceValueRequest_Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceValueRequest_ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceValueRequest_ValueValidationError{}