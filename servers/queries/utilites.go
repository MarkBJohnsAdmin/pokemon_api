package queries

import (
	"encoding/json"
	"net/http"
)

//	===============================================================================================

const PkmnQueryErr = "failed to get Pokemon data"

const PkmnNotFound = "Pokemon not found"

const LrndMoveQueryErr = "failed to get Learned Move data"

const LrndMoveNotFound = "Learned Move not found"

const PdexQueryErr = "failed to get Pokedex data"

const PdexNotFound = "Pokedex not found"

const AbltyQueryErr = "failed to get Ability data"

const AbltyNotFound = "Ability not found"

const MoveQueryErr = "failed to get Move data"

const MoveNotFound = "Move not found"

const EncodeFail = "failed to encode data"

//	-----------------------------------------------------------------------------------------------

const ContentType = "Content-Type"

const ApplicationJSON = "application/json"

//	-----------------------------------------------------------------------------------------------

const AccessControlAllowOrigin = "Access-Control-Allow-Origin"

const AccessControlAllowMethods = "Access-Control-Allow-Methods"

const GET_OPTIONS = "GET, OPTIONS"

const AccessControlAllowHeaders = "Access-Control-Allow-Headers"

const MethodNotAllowed = "Method not allowed"

//	===============================================================================================

// func ReturnAsJSON[T any](w http.ResponseWriter, data T) error {
// 	w.Header().Set(ContentType, ApplicationJSON)

// 	return json.NewEncoder(w).Encode(data)
// }

type Responder struct {
	Writer http.ResponseWriter
	Status int
}

func (s *Responder) SetHeader(key, val string) {
	s.Writer.Header().Set(key, val)
}

func (s *Responder) SetCORSHeaders() {
	s.SetHeader(AccessControlAllowOrigin, "*")
	s.SetHeader(AccessControlAllowMethods, GET_OPTIONS)
	s.SetHeader(AccessControlAllowHeaders, ContentType)
}

func (s *Responder) ReturnAsJSON(data interface{}) error {
	s.SetHeader(ContentType, ApplicationJSON)
	s.Writer.WriteHeader(s.Status)
	return json.NewEncoder(s.Writer).Encode(data)
}

func (s *Responder) SetStatus(status int) {
	s.Status = status
}

func (s *Responder) InternalServerErr(msg string) {
	http.Error(s.Writer, msg, http.StatusInternalServerError)
}

func (s *Responder) NotFoundErr(msg string) {
	http.Error(s.Writer, msg, http.StatusNotFound)
}

func (s *Responder) MethodNotAllowedErr() {
	http.Error(s.Writer, MethodNotAllowed, http.StatusMethodNotAllowed)
}

func (s *Responder) EncodingErr() {
	http.Error(s.Writer, EncodeFail, http.StatusInternalServerError)
}
