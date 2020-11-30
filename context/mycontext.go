package context

import (
	"context"
	"fmt"
)

func MyContext() {
	ctx := context.Background()
	process(ctx)

	ctx = context.WithValue(ctx, "traceID", "shane-Mac")
	process(ctx)

	// 若不设置traceID，则使用process使用ctx的traceID
	ctx1 := context.WithValue(ctx, "traceID", "baidu-Mac")
	process(ctx1)
}

func process(ctx context.Context) {
	traceID, ok := ctx.Value("traceID").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceID)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}
