package closer

import (
	"context"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/logger"
	"strings"
	"sync"
)

type Closer struct {
	mu    sync.Mutex
	funcs []Func
	logger *logger.Logger
}

func New(logger *logger.Logger) (Closer, error){
	return Closer{
		logger: logger,
	}, nil
}

func (c *Closer) Add(f Func) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		msgs     = make([]string, 0, len(c.funcs))
		complete = make(chan struct{}, 1)
	)

	go func() {
		for _, f := range c.funcs {
			if err := f(ctx); err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	case <-complete:
		break
	}

	if len(msgs) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(msgs, "\n"),
		)
	}

	return nil
}

type Func func(ctx context.Context) error
