package paystackapi

import (
	"context"
	"errors"
	"testing"
)

type item struct {
	ID int
}

func makePages(pages [][]item) FetchNextPageFunc[item] {
	return func(ctx context.Context, page, perPage int) (Response[[]item], error) {
		if page > len(pages) {
			return Response[[]item]{}, nil
		}
		data := pages[page-1]
		resp := Response[[]item]{
			Data: data,
			Meta: &Meta{
				Page:      page,
				PerPage:   perPage,
				PageCount: len(pages),
				Total:     len(pages) * perPage,
			},
		}
		return resp, nil
	}
}

func TestIterator_SinglePage(t *testing.T) {
	fetch := makePages([][]item{{{ID: 1}, {ID: 2}, {ID: 3}}})
	it := NewIterator(fetch)

	var collected []item
	for it.Next(context.Background()) {
		collected = append(collected, it.Value())
	}
	if it.Err() != nil {
		t.Fatalf("unexpected error: %v", it.Err())
	}
	if len(collected) != 3 {
		t.Errorf("expected 3 items, got %d", len(collected))
	}
	if collected[0].ID != 1 || collected[2].ID != 3 {
		t.Errorf("unexpected items: %v", collected)
	}
}

func TestIterator_MultiPage(t *testing.T) {
	fetch := makePages([][]item{
		{{ID: 1}, {ID: 2}},
		{{ID: 3}, {ID: 4}},
		{{ID: 5}},
	})
	it := NewIterator(fetch)

	var collected []item
	for it.Next(context.Background()) {
		collected = append(collected, it.Value())
	}
	if it.Err() != nil {
		t.Fatalf("unexpected error: %v", it.Err())
	}
	if len(collected) != 5 {
		t.Errorf("expected 5 items across 3 pages, got %d", len(collected))
	}
}

func TestIterator_EmptyPage(t *testing.T) {
	fetch := makePages([][]item{{}})
	it := NewIterator(fetch)

	if it.Next(context.Background()) {
		t.Error("expected false for empty page")
	}
	if it.Err() != nil {
		t.Fatalf("unexpected error: %v", it.Err())
	}
}

func TestIterator_FetchError(t *testing.T) {
	fetchErr := errors.New("network error")
	fetch := func(ctx context.Context, page, perPage int) (Response[[]item], error) {
		return Response[[]item]{}, fetchErr
	}
	it := NewIterator[item](fetch)

	if it.Next(context.Background()) {
		t.Error("expected false when fetch returns error")
	}
	if !errors.Is(it.Err(), fetchErr) {
		t.Errorf("expected fetchErr, got %v", it.Err())
	}
}

func TestIterator_ErrorMidIteration(t *testing.T) {
	fetchErr := errors.New("mid error")
	calls := 0
	fetch := func(ctx context.Context, page, perPage int) (Response[[]item], error) {
		calls++
		if calls == 1 {
			return Response[[]item]{
				Data: []item{{ID: 1}},
				Meta: &Meta{Page: 1, PerPage: 1, PageCount: 2, Total: 2},
			}, nil
		}
		return Response[[]item]{}, fetchErr
	}
	it := NewIterator[item](fetch)

	count := 0
	for it.Next(context.Background()) {
		count++
	}
	if count != 1 {
		t.Errorf("expected 1 item before error, got %d", count)
	}
	if !errors.Is(it.Err(), fetchErr) {
		t.Errorf("expected fetchErr after mid-iteration failure, got %v", it.Err())
	}
}

func TestIterator_NoNextAfterError(t *testing.T) {
	fetchErr := errors.New("fail")
	fetch := func(ctx context.Context, page, perPage int) (Response[[]item], error) {
		return Response[[]item]{}, fetchErr
	}
	it := NewIterator[item](fetch)

	it.Next(context.Background())
	// Calling Next again after error should return false immediately
	if it.Next(context.Background()) {
		t.Error("expected false on subsequent Next() after error")
	}
}
