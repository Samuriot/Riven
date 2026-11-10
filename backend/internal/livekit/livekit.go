package livekit

import (
	// Blank imports keep the LiveKit SDKs in go.mod until real code uses them.
	"github.com/livekit/protocol/auth"
	_ "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/server-sdk-go/v2"

	"time"
	_ "context"
)

// Will create a token for any user, with a bool param for admin tokens
func CreateToken(apiKey, apiSecret, room, identity, name string, isAdmin bool) (string, error) {
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room: room,
		RoomAdmin: isAdmin,
	}
	grant.SetCanPublish(true)
	grant.SetCanSubscribe(true)
	grant.SetCanUpdateOwnMetadata(isAdmin)

	return auth.NewAccessToken(apiKey, apiSecret).
		SetVideoGrant(grant).
		SetIdentity(identity).
		SetName(name).
		SetValidFor(time.Hour).
		ToJWT()
}