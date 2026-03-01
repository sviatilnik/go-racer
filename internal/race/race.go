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

func NewRace(ctx context.Context, URLs []string) *Race {
	return &Race{
		res:   &result{},
		ctx:   ctx,
		links: URLs,
	}
}

func (race *Race) Start() {
	results := make(chan *resultItem, len(race.links))

	workersContext, workersContextCancel := context.WithCancel(race.ctx)
	var wg sync.WaitGroup

	for _, link := range race.links {
		wg.Go(func() {
			race.worker(workersContext, link, results)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for {
		select {
		case <-race.ctx.Done():
			workersContextCancel()

			for r := range results {
				race.res.Items = append(race.res.Items, r)
			}

			race.res.print()
			return
		case res := <-results:
			if !race.res.HasWinner {
				res.IsWinner = true
				race.res.HasWinner = true
			}

			race.res.Items = append(race.res.Items, res)
			if len(race.res.Items) == len(race.links) {
				race.res.print()

				workersContextCancel()
				return
			}
		}
	}
}

func (r *Race) worker(ctx context.Context, link string, results chan<- *resultItem) {

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

		_, err = http.DefaultClient.Do(req)

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
