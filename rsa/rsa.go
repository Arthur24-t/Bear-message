package rsa

// Request used by the client to send the server it's an RSA public key
// https://docs.mythic-c2.net/customizing/c2-related-development/c2-profile-code/agent-side-coding/initial-checkin#eke-by-generating-client-side-rsa-keys
type Request struct {
	Action    string `json:"action"`     // staging_rsa
	PubKey    string `json:"pub_key"`    // base64 of public RSA key
	SessionID string `json:"session_id"` // 20 character string; unique session ID for this callback
	Padding   string `json:"padding,omitempty"`
}

// Response contains the derived session key encrypted with the agent's RSA key
// https://docs.mythic-c2.net/customizing/c2-related-development/c2-profile-code/agent-side-coding/initial-checkin#eke-by-generating-client-side-rsa-keys
type Response struct {
	Action     string `json:"action"`      // staging_rsa
	ID         string `json:"uuid"`        // new UUID for the next message
	SessionKey string `json:"session_key"` // Base64( RSAPub( new aes session key ) )
	SessionID  string `json:"session_id"`  // same 20 char string back
}