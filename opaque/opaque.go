
package opaque

import "encoding/gob"

// init registers message types with gob that are an interface for Base.Payload
func init() {
	gob.Register(Opaque{})
}

// Type represents the type of OPAQUE message
type Type int

const (
	// Undefined is the default value when a Type was not set
	Undefined Type = iota
	// RegInit is used to denote that the embedded payload contains data for the OPAQUE protocol Registration Initialization step
	RegInit
	// RegComplete is used to denote that the embedded payload contains data for the OPAQUE protocol Registration Complete step
	RegComplete
	// AuthInit is used to denote that the embedded payload contains data for the OPAQUE protocol Authorization Initialization step
	AuthInit
	// AuthComplete is used to denote that the embedded payload contains data for the OPAQUE protocol Authorization Complete step
	AuthComplete
	// ReRegister is used to instruct the Agent it needs to execute the OPAQUE Registration process with the server
	ReRegister
	// ReAuthenticate is used to instruct the Agent it needs to execute the OPAQUE Authentication process with the server
	ReAuthenticate
)

// Opaque is a structure embedded into Merlin Base messages as a payload used to complete OPAQUE registration and authentication
type Opaque struct {
	Type    Type   // The type of OPAQUE message from the constants
	Payload []byte // OPAQUE payload data
}

// String returns a string representation of the Opaque.Type
func (t Type) String() string {
	switch t {
	case Undefined:
		return "Undefined"
	case RegInit:
		return "RegInit"
	case RegComplete:
		return "RegComplete"
	case AuthInit:
		return "AuthInit"
	case AuthComplete:
		return "AuthComplete"
	case ReRegister:
		return "ReRegister"
	case ReAuthenticate:
		return "ReAuthenticate"
	default:
		return "UNKNOWN OPAQUE TYPE"
	}
}