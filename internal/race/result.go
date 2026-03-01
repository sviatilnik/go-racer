package race

import (
	"fmt"
	"time"
)

type resultItem struct {
	IsWinner bool
	URL      string
	Time     time.Duration
	Err      error
}

type result struct {
	HasWinner bool
	Items     []*resultItem
}

func (r *result) print() {
	if r.HasWinner {
		r.printWithWinner()
	} else {
		r.printWithoutWinner()
	}
}

func (r *result) printWithWinner() {
	for _, item := range r.Items {
		if item.IsWinner {
			fmt.Printf("🚀 Самый быстрый: %s (%d ms) \n", item.URL, item.Time.Milliseconds())
		}
	}
	fmt.Println("Остальные результаты:: ")
	for _, item := range r.Items {
		if !item.IsWinner {
			fmt.Printf("\t✓ %s (%d ms) \n", item.URL, item.Time.Milliseconds())
		}
	}
}

func (r *result) printWithoutWinner() {
	fmt.Println("❌ Ни один URL не ответил успешно")
	fmt.Println("Ошибки: ")

	for _, item := range r.Items {
		fmt.Printf("\t ✗ %s : %s\n", item.URL, item.Err.Error())
	}
}
