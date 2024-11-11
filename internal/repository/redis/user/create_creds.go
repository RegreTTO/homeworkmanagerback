package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"homeworktodolist/internal/entity"
	"time"
)

func (r *Repo) CreateCreds(ctx context.Context, creds entity.UserCreds) (sessionKey string, err error) {
	sessionKey = uuid.NewString()

	credsBytes, err := json.Marshal(credsData{
		UserID:  creds.UserID,
		Role:    creds.Role,
		GroupId: creds.GroupId,
	})
	if err != nil {
		return "", err
	}

	if err = r.client.Set(ctx, sessionKey, credsBytes, r.config.AuthTTL*time.Second).Err(); err != nil {
		return "", err
	}

	return
}
