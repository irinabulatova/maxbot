package maxbot

import (
	"context"
	"testing"
	"time"

	maxbot "github.com/max-messenger/max-bot-api-client-go/v2"
	"github.com/max-messenger/max-bot-api-client-go/v2/model"
)

func TestContext(t *testing.T) {
	ctx := NewContext(context.Background(), &maxbot.Api{}, model.Update{})

	timeNow := time.Now().Unix()
	ctx.WithValue(ctx.Context(), "timestamp", timeNow)

	result, ok := ctx.Context().Value("timestamp").(int64)
	if !ok {
		t.Error("expected timestamp")
	}

	if result != timeNow {
		t.Error("expected timestamp")
	}
}
