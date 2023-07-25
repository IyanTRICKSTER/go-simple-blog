package usecases

import (
	"context"
	"go-simple-blog/contracts"
	"strconv"
	"strings"
	"time"
)

func AuthGuard(ctx context.Context) contracts.IAuthenticatedRequest {
	auth := ctx.Value(contracts.AuthContex).(contracts.IAuthenticatedRequest)
	if auth != nil {
		return auth
	}
	return nil
}

func Slugify(text string) string {
	slug := strings.ToLower(text)
	return strings.Replace(slug, " ", "-", -1)
}

func UnixMilliToStr() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}
