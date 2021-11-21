//go:build integration

package ping

import (
	"context"
	"fmt"
	"github.com/ycd/dstp/pkg/common"
	"testing"
)

func TestPingFallback(t *testing.T) {
	out, err := runPingFallback(context.Background(), common.Address("8.8.8.8"), 3, 10)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(out.String())
}
