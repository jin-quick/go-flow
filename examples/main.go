/**
 * Name: main.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/18
 * Description:
 */

package main

import (
	"context"
	"github.com/jin-quick/go-flow/flows"
	"github.com/jin-quick/go-flow/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger(&logger.Config{
		Level:  "debug",
		Format: "text",
	})
	flow := flows.NewInquirySupportFlow()
	ctx := context.Background()
	logger.Info(ctx, "execute 4")
	if err := flow.Execute(context.Background(), "4"); err != nil {
		logger.Error(ctx, "failed to call flow.Execute", zap.Error(err))
	}
}

// End
