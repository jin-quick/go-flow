/**
 * Name: flow_test.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/25
 * Description:
 */

package flows

import (
	"context"
	"testing"

	"github.com/jin-quick/go-flow/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestFlow_Execute(t *testing.T) {
	logger.InitLogger(&logger.Config{
		Level:  "debug",
		Format: "text",
	})
	f := NewInquirySupportFlow()

	tests := []struct {
		name string
		args string
	}{
		{
			name: "Normal: 1",
			args: "1",
		},
		{
			name: "Normal: 4",
			args: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := f.Execute(context.Background(), tt.args)
			assert.Empty(t, err)
		})
	}
}

// End
