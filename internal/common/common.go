package common

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"go.uber.org/zap"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Layout - value of Layout
const Layout = "01/02/2006"

// Active - value of status
const Active = "active"

// RPCReqTimeout - value of request timeout
const RPCReqTimeout = 10 * time.Second

// Inactive - value of status
const Inactive = "Inactive"

// Error - used for
type Error struct {
	ErrorCode      string `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`
	HTTPStatusCode int    `json:"status"`
	RequestID      string `json:"request_id"`
}

// ParseURL - parses a url into a slice (GetPathParts) and
// the query string (GetPathQueryString)
func ParseURL(urlString string) ([]string, url.Values, error) {
	pathString, queryString, err := GetPathQueryString(urlString)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return []string{}, nil, err
	}

	return GetPathParts(pathString), queryString, nil
}

// GetPathQueryString -- given url string, returns the path, and the
// query string
// Eg. "/v1/users?limit=5&cursor=s4R0Z6ecFTzTC4j=" will return
// "/v1/users", ["limit"]="5", ["cursor"]="s4R0Z6ecFTzTC4j="
func GetPathQueryString(s string) (string, url.Values, error) {
	u, err := url.Parse(s)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return "", nil, err
	}

	return u.Path, u.Query(), nil
}

// GetPathParts - given a url, returns a slice of the parts of the url
func GetPathParts(url string) []string {
	var pathParts []string

	sliceOfSubstrings := strings.Split(url, "/")

	for _, p := range sliceOfSubstrings {
		if p != "" {
			pathParts = append(pathParts, p)
		}
	}

	return pathParts
}

// RenderJSON - send JSON response
func RenderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error("Error", zap.Error(err))
		http.Error(w, err.Error(), 400)
		return
	}
}

// RenderErrorJSON - send error JSON response
func RenderErrorJSON(w http.ResponseWriter, errorCode string, errorMsg string, httpStatusCode int, requestID string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	e := Error{ErrorCode: errorCode, ErrorMsg: errorMsg, HTTPStatusCode: httpStatusCode, RequestID: requestID}
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Error("Error", zap.Error(err))
		http.Error(w, err.Error(), 400)
		return
	}
}

// GetRequestID - used for RequestID generation
func GetRequestID() string {
	return xid.New().String()
}

// GetUUID - used for UUID generation
func GetUUID() uuid.UUID {
	return uuid.New()
}

// GetUUIDBytes - used for UUID generation, to save in the db
func GetUUIDBytes() ([]byte, error) {
	return uuid.New().MarshalBinary()
}

// UUIDBytesToStr - convert a UUID retrieved from the DB as str,
// to string for sending to the client
func UUIDBytesToStr(b []byte) (string, error) {
	u, err := uuid.FromBytes(b)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return "", err
	}
	return u.String(), nil
}

// UUIDStrToUUID - convert a UUID str into UUID
func UUIDStrToUUID(s string) (uuid.UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return u, err
	}
	return u, nil
}

// UUIDStrToBytes - convert a UUID str into bytes
func UUIDStrToBytes(s string) ([]byte, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}
	return u.MarshalBinary()
}

// ParseTemplate - used for parsing template (for emails)
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Error("Error", zap.Error(err))
		return "", err
	}
	body := buf.String()
	return body, nil
}

// EncodeCursor - encode cursor
func EncodeCursor(cursor uint32) string {
	cursorStr := strconv.FormatUint(uint64(cursor), 10)
	return base64.StdEncoding.EncodeToString([]byte(cursorStr))
}

// DecodeCursor - decode cursor
func DecodeCursor(cursor string) string {
	cursorBytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return ""
	}
	return string(cursorBytes)
}

// GetTimeDetails - used to populate created_by and updated_by fields
// when inserting/updating records in the database
func GetTimeDetails() time.Time {
	tn := time.Now().UTC().Truncate(time.Second)
	return tn
}

// TimeToTimestamp - used to convert time.Time to timestamp
func TimeToTimestamp(dateAt time.Time) *timestamppb.Timestamp {
	createdAt := timestamppb.New(dateAt)
	return createdAt
}

// TimestampToTime - used to convert timestamp to time.Time
func TimestampToTime(dateAt *timestamppb.Timestamp) time.Time {
	createdAt := dateAt.AsTime()

	return createdAt
}

// ConvertTimeToTimestamp - used to convert time.Time to timestamp
func ConvertTimeToTimestamp(layout string, date string) (*timestamppb.Timestamp, error) {
	date1, err := time.Parse(layout, date)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}

	fdate := TimeToTimestamp(date1)

	return fdate, nil
}
