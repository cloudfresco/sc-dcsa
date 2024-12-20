// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ebl/v1/issuerequestresponse.proto

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

// Validate checks the field values on CreateIssuanceRequestResponseRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CreateIssuanceRequestResponseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateIssuanceRequestResponseRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreateIssuanceRequestResponseRequestMultiError, or nil if none found.
func (m *CreateIssuanceRequestResponseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateIssuanceRequestResponseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransportDocumentReference

	// no validation rules for IssuanceResponseCode

	// no validation rules for Reason

	// no validation rules for CreatedDateTime

	// no validation rules for IssuanceRequestId

	// no validation rules for UserId

	// no validation rules for UserEmail

	// no validation rules for RequestId

	if len(errors) > 0 {
		return CreateIssuanceRequestResponseRequestMultiError(errors)
	}

	return nil
}

// CreateIssuanceRequestResponseRequestMultiError is an error wrapping multiple
// validation errors returned by
// CreateIssuanceRequestResponseRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateIssuanceRequestResponseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateIssuanceRequestResponseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateIssuanceRequestResponseRequestMultiError) AllErrors() []error { return m }

// CreateIssuanceRequestResponseRequestValidationError is the validation error
// returned by CreateIssuanceRequestResponseRequest.Validate if the designated
// constraints aren't met.
type CreateIssuanceRequestResponseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIssuanceRequestResponseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIssuanceRequestResponseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIssuanceRequestResponseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIssuanceRequestResponseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIssuanceRequestResponseRequestValidationError) ErrorName() string {
	return "CreateIssuanceRequestResponseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIssuanceRequestResponseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIssuanceRequestResponseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIssuanceRequestResponseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIssuanceRequestResponseRequestValidationError{}

// Validate checks the field values on CreateIssuanceRequestResponseResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CreateIssuanceRequestResponseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateIssuanceRequestResponseResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreateIssuanceRequestResponseResponseMultiError, or nil if none found.
func (m *CreateIssuanceRequestResponseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateIssuanceRequestResponseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIssuanceRequestResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateIssuanceRequestResponseResponseValidationError{
					field:  "IssuanceRequestResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateIssuanceRequestResponseResponseValidationError{
					field:  "IssuanceRequestResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIssuanceRequestResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateIssuanceRequestResponseResponseValidationError{
				field:  "IssuanceRequestResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateIssuanceRequestResponseResponseMultiError(errors)
	}

	return nil
}

// CreateIssuanceRequestResponseResponseMultiError is an error wrapping
// multiple validation errors returned by
// CreateIssuanceRequestResponseResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateIssuanceRequestResponseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateIssuanceRequestResponseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateIssuanceRequestResponseResponseMultiError) AllErrors() []error { return m }

// CreateIssuanceRequestResponseResponseValidationError is the validation error
// returned by CreateIssuanceRequestResponseResponse.Validate if the
// designated constraints aren't met.
type CreateIssuanceRequestResponseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateIssuanceRequestResponseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateIssuanceRequestResponseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateIssuanceRequestResponseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateIssuanceRequestResponseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateIssuanceRequestResponseResponseValidationError) ErrorName() string {
	return "CreateIssuanceRequestResponseResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateIssuanceRequestResponseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateIssuanceRequestResponseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateIssuanceRequestResponseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateIssuanceRequestResponseResponseValidationError{}

// Validate checks the field values on UpdateIssuanceRequestResponseRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateIssuanceRequestResponseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateIssuanceRequestResponseRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateIssuanceRequestResponseRequestMultiError, or nil if none found.
func (m *UpdateIssuanceRequestResponseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateIssuanceRequestResponseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransportDocumentReference

	// no validation rules for IssuanceResponseCode

	// no validation rules for Reason

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for UserEmail

	// no validation rules for RequestId

	if len(errors) > 0 {
		return UpdateIssuanceRequestResponseRequestMultiError(errors)
	}

	return nil
}

// UpdateIssuanceRequestResponseRequestMultiError is an error wrapping multiple
// validation errors returned by
// UpdateIssuanceRequestResponseRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateIssuanceRequestResponseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateIssuanceRequestResponseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateIssuanceRequestResponseRequestMultiError) AllErrors() []error { return m }

// UpdateIssuanceRequestResponseRequestValidationError is the validation error
// returned by UpdateIssuanceRequestResponseRequest.Validate if the designated
// constraints aren't met.
type UpdateIssuanceRequestResponseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateIssuanceRequestResponseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateIssuanceRequestResponseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateIssuanceRequestResponseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateIssuanceRequestResponseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateIssuanceRequestResponseRequestValidationError) ErrorName() string {
	return "UpdateIssuanceRequestResponseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateIssuanceRequestResponseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateIssuanceRequestResponseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateIssuanceRequestResponseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateIssuanceRequestResponseRequestValidationError{}

// Validate checks the field values on UpdateIssuanceRequestResponseResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateIssuanceRequestResponseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateIssuanceRequestResponseResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateIssuanceRequestResponseResponseMultiError, or nil if none found.
func (m *UpdateIssuanceRequestResponseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateIssuanceRequestResponseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateIssuanceRequestResponseResponseMultiError(errors)
	}

	return nil
}

// UpdateIssuanceRequestResponseResponseMultiError is an error wrapping
// multiple validation errors returned by
// UpdateIssuanceRequestResponseResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateIssuanceRequestResponseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateIssuanceRequestResponseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateIssuanceRequestResponseResponseMultiError) AllErrors() []error { return m }

// UpdateIssuanceRequestResponseResponseValidationError is the validation error
// returned by UpdateIssuanceRequestResponseResponse.Validate if the
// designated constraints aren't met.
type UpdateIssuanceRequestResponseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateIssuanceRequestResponseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateIssuanceRequestResponseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateIssuanceRequestResponseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateIssuanceRequestResponseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateIssuanceRequestResponseResponseValidationError) ErrorName() string {
	return "UpdateIssuanceRequestResponseResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateIssuanceRequestResponseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateIssuanceRequestResponseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateIssuanceRequestResponseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateIssuanceRequestResponseResponseValidationError{}

// Validate checks the field values on IssuanceRequestResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IssuanceRequestResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IssuanceRequestResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IssuanceRequestResponseMultiError, or nil if none found.
func (m *IssuanceRequestResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IssuanceRequestResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIssuanceRequestResponseD()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "IssuanceRequestResponseD",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "IssuanceRequestResponseD",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIssuanceRequestResponseD()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssuanceRequestResponseValidationError{
				field:  "IssuanceRequestResponseD",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetIssuanceRequestResponseT()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "IssuanceRequestResponseT",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "IssuanceRequestResponseT",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIssuanceRequestResponseT()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssuanceRequestResponseValidationError{
				field:  "IssuanceRequestResponseT",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCrUpdUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "CrUpdUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "CrUpdUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCrUpdUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssuanceRequestResponseValidationError{
				field:  "CrUpdUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCrUpdTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "CrUpdTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssuanceRequestResponseValidationError{
					field:  "CrUpdTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCrUpdTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssuanceRequestResponseValidationError{
				field:  "CrUpdTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IssuanceRequestResponseMultiError(errors)
	}

	return nil
}

// IssuanceRequestResponseMultiError is an error wrapping multiple validation
// errors returned by IssuanceRequestResponse.ValidateAll() if the designated
// constraints aren't met.
type IssuanceRequestResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IssuanceRequestResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IssuanceRequestResponseMultiError) AllErrors() []error { return m }

// IssuanceRequestResponseValidationError is the validation error returned by
// IssuanceRequestResponse.Validate if the designated constraints aren't met.
type IssuanceRequestResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssuanceRequestResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssuanceRequestResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssuanceRequestResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssuanceRequestResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssuanceRequestResponseValidationError) ErrorName() string {
	return "IssuanceRequestResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IssuanceRequestResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssuanceRequestResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssuanceRequestResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssuanceRequestResponseValidationError{}

// Validate checks the field values on IssuanceRequestResponseD with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IssuanceRequestResponseD) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IssuanceRequestResponseD with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IssuanceRequestResponseDMultiError, or nil if none found.
func (m *IssuanceRequestResponseD) ValidateAll() error {
	return m.validate(true)
}

func (m *IssuanceRequestResponseD) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Uuid4

	// no validation rules for IdS

	// no validation rules for TransportDocumentReference

	// no validation rules for IssuanceResponseCode

	// no validation rules for Reason

	// no validation rules for IssuanceRequestId

	if len(errors) > 0 {
		return IssuanceRequestResponseDMultiError(errors)
	}

	return nil
}

// IssuanceRequestResponseDMultiError is an error wrapping multiple validation
// errors returned by IssuanceRequestResponseD.ValidateAll() if the designated
// constraints aren't met.
type IssuanceRequestResponseDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IssuanceRequestResponseDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IssuanceRequestResponseDMultiError) AllErrors() []error { return m }

// IssuanceRequestResponseDValidationError is the validation error returned by
// IssuanceRequestResponseD.Validate if the designated constraints aren't met.
type IssuanceRequestResponseDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssuanceRequestResponseDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssuanceRequestResponseDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssuanceRequestResponseDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssuanceRequestResponseDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssuanceRequestResponseDValidationError) ErrorName() string {
	return "IssuanceRequestResponseDValidationError"
}

// Error satisfies the builtin error interface
func (e IssuanceRequestResponseDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssuanceRequestResponseD.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssuanceRequestResponseDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssuanceRequestResponseDValidationError{}

// Validate checks the field values on IssuanceRequestResponseT with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IssuanceRequestResponseT) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IssuanceRequestResponseT with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IssuanceRequestResponseTMultiError, or nil if none found.
func (m *IssuanceRequestResponseT) ValidateAll() error {
	return m.validate(true)
}

func (m *IssuanceRequestResponseT) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCreatedDateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IssuanceRequestResponseTValidationError{
					field:  "CreatedDateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IssuanceRequestResponseTValidationError{
					field:  "CreatedDateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedDateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IssuanceRequestResponseTValidationError{
				field:  "CreatedDateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IssuanceRequestResponseTMultiError(errors)
	}

	return nil
}

// IssuanceRequestResponseTMultiError is an error wrapping multiple validation
// errors returned by IssuanceRequestResponseT.ValidateAll() if the designated
// constraints aren't met.
type IssuanceRequestResponseTMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IssuanceRequestResponseTMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IssuanceRequestResponseTMultiError) AllErrors() []error { return m }

// IssuanceRequestResponseTValidationError is the validation error returned by
// IssuanceRequestResponseT.Validate if the designated constraints aren't met.
type IssuanceRequestResponseTValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssuanceRequestResponseTValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssuanceRequestResponseTValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssuanceRequestResponseTValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssuanceRequestResponseTValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssuanceRequestResponseTValidationError) ErrorName() string {
	return "IssuanceRequestResponseTValidationError"
}

// Error satisfies the builtin error interface
func (e IssuanceRequestResponseTValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssuanceRequestResponseT.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssuanceRequestResponseTValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssuanceRequestResponseTValidationError{}
