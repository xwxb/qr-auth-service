package model

import "time"

// Verification represents the verification information.
type Verification struct {
	UUID        string    `json:"uuid"`
	Expiration  time.Time `json:"expiration"`
	VerificationStatus bool   `json:"verificationStatus"`
}