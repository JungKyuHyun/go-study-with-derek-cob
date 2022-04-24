package context

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Initialize : 함수 3개
func Initialize() {
	log.SetHandler(text.New(os.Stdout))
	// empty Contex 생성.
	ctx := context.Background()

	// 로거 생성후 context에 연결.
	ctx, e := FromContext(ctx, log.Log)

	// 필드 생성
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go Cookbook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
}