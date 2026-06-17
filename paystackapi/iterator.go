package paystackapi

import "context"

// FetchNextPageFunc is a function signature for fetching a specific page of results.
type FetchNextPageFunc[T any] func(ctx context.Context, page, perPage int) (Response[[]T], error)

// Iterator provides a convenient way to iterate over paginated API responses.
type Iterator[T any] struct {
	ctx         context.Context
	fetchNext   FetchNextPageFunc[T]
	currentPage []T
	meta        *Meta
	index       int
	err         error
	initialized bool
}

// NewIterator creates a new Iterator.
func NewIterator[T any](ctx context.Context, fetchNext FetchNextPageFunc[T]) *Iterator[T] {
	return &Iterator[T]{
		ctx:       ctx,
		fetchNext: fetchNext,
		index:     -1,
	}
}

// Next advances the iterator to the next item. Returns false when exhausted or on error.
func (it *Iterator[T]) Next() bool {
	if it.err != nil {
		return false
	}

	if !it.initialized {
		it.initialized = true
		if !it.fetchPage(1, 50) {
			return false
		}
	}

	if it.index+1 < len(it.currentPage) {
		it.index++
		return true
	}

	if it.meta != nil {
		if it.meta.Page >= it.meta.PageCount {
			return false
		}
		if !it.fetchPage(it.meta.Page+1, it.meta.PerPage) {
			return false
		}

		it.index = 0
		return len(it.currentPage) > 0
	}

	return false
}

func (it *Iterator[T]) fetchPage(page, perPage int) bool {
	resp, err := it.fetchNext(it.ctx, page, perPage)
	if err != nil {
		it.err = err
		return false
	}

	it.currentPage = resp.Data
	it.meta = resp.Meta
	return true
}

// Value returns the current item. Must only be called after Next() returns true.
func (it *Iterator[T]) Value() T {
	return it.currentPage[it.index]
}

// Err returns any error encountered during iteration.
func (it *Iterator[T]) Err() error {
	return it.err
}
