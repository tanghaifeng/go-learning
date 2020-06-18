package tests

import (
	"fmt"
	sync_learn_01 "go-learning/sync-learn-01"
	"math"
	"os"
	"sync"
	"testing"
	"text/tabwriter"
)

func TestRwmutx(t *testing.T) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	var m sync.RWMutex
	for i := 0; i < 1; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, sync_learn_01.Test(count, &m, m.RLocker()), sync_learn_01.Test(count, &m, &m))
	}

}
