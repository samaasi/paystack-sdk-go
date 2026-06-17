package paystackapi

import "context"

// FetchNextPageFunc is a function signature for fetching a specific page of results.
type FetchNextPageFunc[T any] func(ctx context.Context, page, perPage int) (Response[[]T], error)

// Iterator provides a convenient way to iterate over paginated API responses.
type Iterator[T any] struct {
	ctx         context.Context
	fetchNext   FetchNextPageFunc[T]
	CurrentPage []T
	Meta        *Meta
	Index       int
	Err         error
	initialized bool
}

// NewIterator creates a new Iterator. The fetchNext function should capture any additional
// query parameters (like filters) and use the provided page and perPage to fetch results.
func NewIterator[T any](ctx context.Context, fetchNext FetchNextPageFunc[T]) *Iterator[T] {
	return &Iterator[T]{
		ctx:       ctx,
		fetchNext: fetchNext,
		Index:     -1,
	}
}

// Next advances the iterator to the next item. It fetches the next page automatically if needed.
// It returns true if there is a next item, or false if there are no more items or an error occurred.
func (it *Iterator[T]) Next() bool {
	if it.Err != nil {
		return false
	}

	if !it.initialized {
		it.initialized = true
		if !it.fetchPage(1, 50) {
			return false
		}
	}

	if it.Index+1 < len(it.CurrentPage) {
		it.Index++
		return true
	}

	if it.Meta != nil {
		if it.Meta.Page >= it.Meta.PageCount {
			return false
		}
		if !it.fetchPage(it.Meta.Page+1, it.Meta.PerPage) {
			return false
		}

		it.Index = 0
		return len(it.CurrentPage) > 0
	}

	return false
}

// fetchPage is a helper to fetch a specific page and update iterator state.
func (it *Iterator[T]) fetchPage(page, perPage int) bool {
	resp, err := it.fetchNext(it.ctx, page, perPage)
	if err != nil {
		it.Err = err
		return false
	}

	it.CurrentPage = resp.Data
	it.Meta = resp.Meta
	return true
}

// Value returns the current item. It panics if called before Next() returns true.
func (it *Iterator[T]) Value() T {
	return it.CurrentPage[it.Index]
}

// Error returns any error encountered during iteration.
func (it *Iterator[T]) Error() error {
	return it.Err
}
