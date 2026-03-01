package race

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Race struct {
	res   *result
	ctx   context.Context
	links []string
}

func NewRace(ctx context.Context, urls []string) *Race {
	return &Race{
		res:   &result{},
		ctx:   ctx,
		links: urls,
	}
}

func (racer *Race) Start() {
	results := make(chan *resultItem, len(racer.links))

	workersContext, workersContextCancel := context.WithCancel(racer.ctx)
	var wg sync.WaitGroup

	for _, link := range racer.links {
		wg.Go(func() {
			racer.worker(workersContext, link, results)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for {
		select {
		case <-racer.ctx.Done():
			workersContextCancel()

			for r := range results {
				racer.res.Items = append(racer.res.Items, r)
			}

			racer.res.print()
			return
		case res := <-results:
			if !racer.res.HasWinner {
				res.IsWinner = true
				racer.res.HasWinner = true
			}

			racer.res.Items = append(racer.res.Items, res)
			if len(racer.res.Items) == len(racer.links) {
				racer.res.print()

				workersContextCancel()
				return
			}
		}
	}
}

var defaultHTTPClient = &http.Client{
	Timeout: 30 * time.Second,
}

func (racer *Race) worker(ctx context.Context, link string, results chan<- *resultItem) {
	if ctx.Err() != nil {
		results <- &resultItem{
			IsWinner: false,
			URL:      link,
			Err:      ctx.Err(),
		}

		return
	}

	httpResult := make(chan *resultItem)

	go func() {

		start := time.Now()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
		if err != nil {
			results <- &resultItem{
				IsWinner: false,
				URL:      link,
				Err:      err,
			}
			return
		}

		_, err = defaultHTTPClient.Do(req)

		end := time.Now()

		httpResult <- &resultItem{
			IsWinner: false,
			URL:      link,
			Err:      err,
			Time:     end.Sub(start),
		}

	}()

	for {
		select {
		case <-ctx.Done():
			results <- &resultItem{
				IsWinner: false,
				URL:      link,
				Err:      ctx.Err(),
			}
			return
		case httpR := <-httpResult:
			results <- httpR
			return
		}
	}
}
