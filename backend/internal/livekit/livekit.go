// Package livekit will hold riven-api's LiveKit integration
// (token signing, room API, webhooks).
package livekit

import (
	// Blank imports keep the LiveKit SDKs in go.mod until real code uses them.
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	_ "github.com/livekit/server-sdk-go/v2"

	"time"
	"context"
)

func CreateJoinToken(apiKey, apiSecret, room, identity, name string) (string, error) {
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room: room,
	}
	grant.SetCanPublish(true)
	grant.SetCanSubscribe(true)

	return auth.NewAccessToken(apiKey, apiSecret).
		SetVideoGrant(grant).
		SetIdentity(identity).
		SetName(name).
		SetValidFor(time.Hour).
		ToJWT()
}

func CreateHostToken(room string) {
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room: room,
		RoomAdmin: true,
	}
	grant.SetCanPublish(true)
	grant.SetCanSubscribe(true)
	grant.SetCanUpdateOwnMetadata(true)
}