// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: profile.proto

package pb

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _profile_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Profile) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Introduction

	// no validation rules for Sex

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProfileValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProfileValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProfileValidationError is the validation error returned by Profile.Validate
// if the designated constraints aren't met.
type ProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileValidationError) ErrorName() string { return "ProfileValidationError" }

// Error satisfies the builtin error interface
func (e ProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileValidationError{}

// Validate checks the field values on GetProfileReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetProfileReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() < 1 {
		return GetProfileReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetProfileReqValidationError is the validation error returned by
// GetProfileReq.Validate if the designated constraints aren't met.
type GetProfileReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileReqValidationError) ErrorName() string { return "GetProfileReqValidationError" }

// Error satisfies the builtin error interface
func (e GetProfileReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileReqValidationError{}

// Validate checks the field values on BatchGetProfilesReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BatchGetProfilesReq) Validate() error {
	if m == nil {
		return nil
	}

	_BatchGetProfilesReq_UserIds_Unique := make(map[int64]struct{}, len(m.GetUserIds()))

	for idx, item := range m.GetUserIds() {
		_, _ = idx, item

		if _, exists := _BatchGetProfilesReq_UserIds_Unique[item]; exists {
			return BatchGetProfilesReqValidationError{
				field:  fmt.Sprintf("UserIds[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_BatchGetProfilesReq_UserIds_Unique[item] = struct{}{}
		}

		// no validation rules for UserIds[idx]
	}

	return nil
}

// BatchGetProfilesReqValidationError is the validation error returned by
// BatchGetProfilesReq.Validate if the designated constraints aren't met.
type BatchGetProfilesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetProfilesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetProfilesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetProfilesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetProfilesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetProfilesReqValidationError) ErrorName() string {
	return "BatchGetProfilesReqValidationError"
}

// Error satisfies the builtin error interface
func (e BatchGetProfilesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetProfilesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetProfilesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetProfilesReqValidationError{}

// Validate checks the field values on BatchGetProfilesRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BatchGetProfilesRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProfiles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BatchGetProfilesResValidationError{
					field:  fmt.Sprintf("Profiles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// BatchGetProfilesResValidationError is the validation error returned by
// BatchGetProfilesRes.Validate if the designated constraints aren't met.
type BatchGetProfilesResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetProfilesResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetProfilesResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetProfilesResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetProfilesResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetProfilesResValidationError) ErrorName() string {
	return "BatchGetProfilesResValidationError"
}

// Error satisfies the builtin error interface
func (e BatchGetProfilesResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetProfilesRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetProfilesResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetProfilesResValidationError{}

// Validate checks the field values on CreateProfileReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateProfileReq) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		return CreateProfileReqValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetIntroduction()); l < 1 || l > 1000 {
		return CreateProfileReqValidationError{
			field:  "Introduction",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
	}

	if _, ok := Sex_name[int32(m.GetSex())]; !ok {
		return CreateProfileReqValidationError{
			field:  "Sex",
			reason: "value must be one of the defined enum values",
		}
	}

	if m.GetUserId() < 1 {
		return CreateProfileReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateProfileReqValidationError is the validation error returned by
// CreateProfileReq.Validate if the designated constraints aren't met.
type CreateProfileReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProfileReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProfileReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProfileReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProfileReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProfileReqValidationError) ErrorName() string { return "CreateProfileReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateProfileReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProfileReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProfileReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProfileReqValidationError{}

// Validate checks the field values on UpdateProfileReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateProfileReq) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 10 {
		return UpdateProfileReqValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetIntroduction()); l < 1 || l > 1000 {
		return UpdateProfileReqValidationError{
			field:  "Introduction",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
	}

	if m.GetUserId() < 1 {
		return UpdateProfileReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// UpdateProfileReqValidationError is the validation error returned by
// UpdateProfileReq.Validate if the designated constraints aren't met.
type UpdateProfileReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProfileReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProfileReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProfileReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProfileReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProfileReqValidationError) ErrorName() string { return "UpdateProfileReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateProfileReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProfileReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProfileReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProfileReqValidationError{}

// Validate checks the field values on DeleteProfileReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteProfileReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() < 1 {
		return DeleteProfileReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DeleteProfileReqValidationError is the validation error returned by
// DeleteProfileReq.Validate if the designated constraints aren't met.
type DeleteProfileReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProfileReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProfileReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProfileReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProfileReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProfileReqValidationError) ErrorName() string { return "DeleteProfileReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteProfileReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProfileReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProfileReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProfileReqValidationError{}
