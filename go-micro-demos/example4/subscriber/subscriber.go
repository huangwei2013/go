package subscriber

import (
	"context"
	"fmt"
	"github.com/lpxxn/gomicrorpc/example4/proto/model"
)

func Handler(ctx context.Context, msg *model.CommonReq) error {
	fmt.Printf("Received message: %s \n", msg.Action)
	return nil
}