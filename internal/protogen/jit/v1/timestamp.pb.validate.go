// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: jit/v1/timestamp.proto

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

// Validate checks the field values on CreateTimestampResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTimestampResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTimestampResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTimestampResponseMultiError, or nil if none found.
func (m *CreateTimestampResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTimestampResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp1()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTimestampResponseValidationError{
					field:  "Timestamp1",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTimestampResponseValidationError{
					field:  "Timestamp1",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp1()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTimestampResponseValidationError{
				field:  "Timestamp1",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTimestampResponseMultiError(errors)
	}

	return nil
}

// CreateTimestampResponseMultiError is an error wrapping multiple validation
// errors returned by CreateTimestampResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateTimestampResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTimestampResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTimestampResponseMultiError) AllErrors() []error { return m }

// CreateTimestampResponseValidationError is the validation error returned by
// CreateTimestampResponse.Validate if the designated constraints aren't met.
type CreateTimestampResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTimestampResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTimestampResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTimestampResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTimestampResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTimestampResponseValidationError) ErrorName() string {
	return "CreateTimestampResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTimestampResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTimestampResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTimestampResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTimestampResponseValidationError{}

// Validate checks the field values on Timestamp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Timestamp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Timestamp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimestampMultiError, or nil
// if none found.
func (m *Timestamp) ValidateAll() error {
	return m.validate(true)
}

func (m *Timestamp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestampD()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TimestampValidationError{
					field:  "TimestampD",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TimestampValidationError{
					field:  "TimestampD",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampD()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimestampValidationError{
				field:  "TimestampD",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampT()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TimestampValidationError{
					field:  "TimestampT",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TimestampValidationError{
					field:  "TimestampT",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampT()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimestampValidationError{
				field:  "TimestampT",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TimestampMultiError(errors)
	}

	return nil
}

// TimestampMultiError is an error wrapping multiple validation errors returned
// by Timestamp.ValidateAll() if the designated constraints aren't met.
type TimestampMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimestampMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimestampMultiError) AllErrors() []error { return m }

// TimestampValidationError is the validation error returned by
// Timestamp.Validate if the designated constraints aren't met.
type TimestampValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampValidationError) ErrorName() string { return "TimestampValidationError" }

// Error satisfies the builtin error interface
func (e TimestampValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestamp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampValidationError{}

// Validate checks the field values on TimestampD with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TimestampD) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TimestampD with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimestampDMultiError, or
// nil if none found.
func (m *TimestampD) ValidateAll() error {
	return m.validate(true)
}

func (m *TimestampD) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Uuid4

	// no validation rules for IdS

	// no validation rules for EventTypeCode

	// no validation rules for EventClassifierCode

	// no validation rules for DelayReasonCode

	// no validation rules for ChangeRemark

	if len(errors) > 0 {
		return TimestampDMultiError(errors)
	}

	return nil
}

// TimestampDMultiError is an error wrapping multiple validation errors
// returned by TimestampD.ValidateAll() if the designated constraints aren't met.
type TimestampDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimestampDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimestampDMultiError) AllErrors() []error { return m }

// TimestampDValidationError is the validation error returned by
// TimestampD.Validate if the designated constraints aren't met.
type TimestampDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampDValidationError) ErrorName() string { return "TimestampDValidationError" }

// Error satisfies the builtin error interface
func (e TimestampDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestampD.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampDValidationError{}

// Validate checks the field values on TimestampT with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TimestampT) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TimestampT with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimestampTMultiError, or
// nil if none found.
func (m *TimestampT) ValidateAll() error {
	return m.validate(true)
}

func (m *TimestampT) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEventDateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TimestampTValidationError{
					field:  "EventDateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TimestampTValidationError{
					field:  "EventDateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEventDateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimestampTValidationError{
				field:  "EventDateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TimestampTMultiError(errors)
	}

	return nil
}

// TimestampTMultiError is an error wrapping multiple validation errors
// returned by TimestampT.ValidateAll() if the designated constraints aren't met.
type TimestampTMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimestampTMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimestampTMultiError) AllErrors() []error { return m }

// TimestampTValidationError is the validation error returned by
// TimestampT.Validate if the designated constraints aren't met.
type TimestampTValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimestampTValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimestampTValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimestampTValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimestampTValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimestampTValidationError) ErrorName() string { return "TimestampTValidationError" }

// Error satisfies the builtin error interface
func (e TimestampTValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimestampT.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimestampTValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimestampTValidationError{}

// Validate checks the field values on CreateTimestampRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTimestampRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTimestampRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTimestampRequestMultiError, or nil if none found.
func (m *CreateTimestampRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTimestampRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EventTypeCode

	// no validation rules for EventClassifierCode

	// no validation rules for EventDateTime

	// no validation rules for DelayReasonCode

	// no validation rules for ChangeRemark

	// no validation rules for UserId

	// no validation rules for UserEmail

	// no validation rules for RequestId

	if len(errors) > 0 {
		return CreateTimestampRequestMultiError(errors)
	}

	return nil
}

// CreateTimestampRequestMultiError is an error wrapping multiple validation
// errors returned by CreateTimestampRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateTimestampRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTimestampRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTimestampRequestMultiError) AllErrors() []error { return m }

// CreateTimestampRequestValidationError is the validation error returned by
// CreateTimestampRequest.Validate if the designated constraints aren't met.
type CreateTimestampRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTimestampRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTimestampRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTimestampRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTimestampRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTimestampRequestValidationError) ErrorName() string {
	return "CreateTimestampRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTimestampRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTimestampRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTimestampRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTimestampRequestValidationError{}

// Validate checks the field values on GetTimestampsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTimestampsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTimestampsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTimestampsResponseMultiError, or nil if none found.
func (m *GetTimestampsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTimestampsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTimestamps() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTimestampsResponseValidationError{
						field:  fmt.Sprintf("Timestamps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTimestampsResponseValidationError{
						field:  fmt.Sprintf("Timestamps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTimestampsResponseValidationError{
					field:  fmt.Sprintf("Timestamps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextCursor

	if len(errors) > 0 {
		return GetTimestampsResponseMultiError(errors)
	}

	return nil
}

// GetTimestampsResponseMultiError is an error wrapping multiple validation
// errors returned by GetTimestampsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTimestampsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTimestampsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTimestampsResponseMultiError) AllErrors() []error { return m }

// GetTimestampsResponseValidationError is the validation error returned by
// GetTimestampsResponse.Validate if the designated constraints aren't met.
type GetTimestampsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTimestampsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTimestampsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTimestampsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTimestampsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTimestampsResponseValidationError) ErrorName() string {
	return "GetTimestampsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTimestampsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTimestampsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTimestampsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTimestampsResponseValidationError{}

// Validate checks the field values on GetTimestampsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTimestampsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTimestampsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTimestampsRequestMultiError, or nil if none found.
func (m *GetTimestampsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTimestampsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Limit

	// no validation rules for NextCursor

	// no validation rules for UserEmail

	// no validation rules for RequestId

	if len(errors) > 0 {
		return GetTimestampsRequestMultiError(errors)
	}

	return nil
}

// GetTimestampsRequestMultiError is an error wrapping multiple validation
// errors returned by GetTimestampsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTimestampsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTimestampsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTimestampsRequestMultiError) AllErrors() []error { return m }

// GetTimestampsRequestValidationError is the validation error returned by
// GetTimestampsRequest.Validate if the designated constraints aren't met.
type GetTimestampsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTimestampsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTimestampsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTimestampsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTimestampsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTimestampsRequestValidationError) ErrorName() string {
	return "GetTimestampsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTimestampsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTimestampsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTimestampsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTimestampsRequestValidationError{}