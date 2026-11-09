// Package livekit will hold riven-api's LiveKit integration
// (token signing, room API, webhooks).
package livekit

import (
	// Blank imports keep the LiveKit SDKs in go.mod until real code uses them.
	_ "github.com/livekit/protocol/auth"
	_ "github.com/livekit/server-sdk-go/v2"
)
