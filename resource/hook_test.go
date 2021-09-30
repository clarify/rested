package resource

import (
	"context"
	"errors"
	"testing"

	"github.com/searis/rest-layer/schema/query"
	"github.com/stretchr/testify/assert"
)

func TestHookUseFind(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(FindEventHandlerFunc(func(ctx context.Context, q *query.Query) error {
		called = true
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 1)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	err = h.onFind(context.Background(), nil)
	assert.NoError(t, err)
	assert.True(t, called)

	err = h.use(FindEventHandlerFunc(func(ctx context.Context, q *query.Query) error {
		return errors.New("error")
	}))
	assert.NoError(t, err)
	err = h.onFind(context.Background(), nil)
	assert.EqualError(t, err, "error")
}

func TestHookUseFound(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(FoundEventHandlerFunc(func(ctx context.Context, q *query.Query, list **ItemList, err *error) {
		called = true
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 1)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	h.onFound(context.Background(), nil, nil, nil)
	assert.True(t, called)

	err = h.use(FoundEventHandlerFunc(func(ctx context.Context, q *query.Query, list **ItemList, err *error) {
		*list = &ItemList{}
		*err = errors.New("error")
	}))
	assert.NoError(t, err)
	var list *ItemList
	err = nil
	h.onFound(context.Background(), nil, &list, &err)
	assert.EqualError(t, err, "error")
	assert.NotNil(t, list)
}

func TestHookUseGet(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(GetEventHandlerFunc(func(ctx context.Context, id interface{}) error {
		called = true
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 1)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	_ = h.onGet(context.Background(), nil)
	assert.True(t, called)

	err = h.use(GetEventHandlerFunc(func(ctx context.Context, id interface{}) error {
		return errors.New("error")
	}))
	assert.NoError(t, err)
	err = h.onGet(context.Background(), nil)
	assert.EqualError(t, err, "error")
}

func TestHookUseGot(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(GotEventHandlerFunc(func(ctx context.Context, item **Item, err *error) {
		called = true
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 1)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	h.onGot(context.Background(), nil, nil)
	assert.True(t, called)

	err = h.use(GotEventHandlerFunc(func(ctx context.Context, item **Item, err *error) {
		*item = &Item{}
		*err = errors.New("error")
	}))
	assert.NoError(t, err)
	var item *Item
	err = nil
	h.onGot(context.Background(), &item, &err)
	assert.EqualError(t, err, "error")
	assert.NotNil(t, item)
}

func TestHookUseInsert(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(InsertEventHandlerFunc(func(ctx context.Context, items []*Item) error {
		called = true
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 1)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	_ = h.onInsert(context.Background(), nil)
	assert.True(t, called)

	err = h.use(InsertEventHandlerFunc(func(ctx context.Context, items []*Item) error {
		return errors.New("error")
	}))
	assert.NoError(t, err)
	err = h.onInsert(context.Background(), nil)
	assert.EqualError(t, err, "error")
}

func TestHookUseInserted(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(InsertedEventHandlerFunc(func(ctx context.Context, items []*Item, err *error) {
		called = true
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 1)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	h.onInserted(context.Background(), nil, nil)
	assert.True(t, called)

	err = h.use(InsertedEventHandlerFunc(func(ctx context.Context, items []*Item, err *error) {
		*err = errors.New("error")
	}))
	assert.NoError(t, err)
	err = nil
	h.onInserted(context.Background(), nil, &err)
	assert.EqualError(t, err, "error")
}

func TestHookUseUpdate(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(UpdateEventHandlerFunc(func(ctx context.Context, item *Item, original *Item) error {
		called = true
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 1)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	_ = h.onUpdate(context.Background(), nil, nil)
	assert.True(t, called)

	err = h.use(UpdateEventHandlerFunc(func(ctx context.Context, item *Item, original *Item) error {
		return errors.New("error")
	}))
	assert.NoError(t, err)
	err = h.onUpdate(context.Background(), nil, nil)
	assert.EqualError(t, err, "error")
}

func TestHookUseUpdated(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(UpdatedEventHandlerFunc(func(ctx context.Context, item *Item, original *Item, err *error) {
		called = true
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 1)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	h.onUpdated(context.Background(), nil, nil, nil)
	assert.True(t, called)

	err = h.use(UpdatedEventHandlerFunc(func(ctx context.Context, item *Item, original *Item, err *error) {
		*err = errors.New("error")
	}))
	assert.NoError(t, err)
	err = nil
	h.onUpdated(context.Background(), nil, nil, &err)
	assert.EqualError(t, err, "error")
}

func TestHookUseDelete(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(DeleteEventHandlerFunc(func(ctx context.Context, item *Item) error {
		called = true
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 1)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	_ = h.onDelete(context.Background(), nil)
	assert.True(t, called)

	err = h.use(DeleteEventHandlerFunc(func(ctx context.Context, item *Item) error {
		return errors.New("error")
	}))
	assert.NoError(t, err)
	err = h.onDelete(context.Background(), nil)
	assert.EqualError(t, err, "error")
}

func TestHookUseDeleted(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(DeletedEventHandlerFunc(func(ctx context.Context, item *Item, err *error) {
		called = true
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 1)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
	h.onDeleted(context.Background(), nil, nil)
	assert.True(t, called)

	err = h.use(DeletedEventHandlerFunc(func(ctx context.Context, item *Item, err *error) {
		*err = errors.New("error")
	}))
	assert.NoError(t, err)
	err = nil
	h.onDeleted(context.Background(), nil, &err)
	assert.EqualError(t, err, "error")
}

func TestHookUseClear(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(ClearEventHandlerFunc(func(ctx context.Context, q *query.Query) error {
		called = true
		return nil
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 1)
	assert.Len(t, h.onClearedH, 0)
	_ = h.onClear(context.Background(), nil)
	assert.True(t, called)

	err = h.use(ClearEventHandlerFunc(func(ctx context.Context, q *query.Query) error {
		return errors.New("error")
	}))
	assert.NoError(t, err)
	err = h.onClear(context.Background(), nil)
	assert.EqualError(t, err, "error")
}

func TestHookUseCleared(t *testing.T) {
	h := eventHandler{}
	called := false
	err := h.use(ClearedEventHandlerFunc(func(ctx context.Context, q *query.Query, deleted *int, err *error) {
		called = true
	}))
	assert.NoError(t, err)
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 1)
	deleted := 0
	h.onCleared(context.Background(), nil, &deleted, nil)
	assert.True(t, called)

	err = h.use(ClearedEventHandlerFunc(func(ctx context.Context, q *query.Query, deleted *int, err *error) {
		*err = errors.New("error")
		*deleted = 2
	}))
	assert.NoError(t, err)
	err = nil
	h.onCleared(context.Background(), nil, &deleted, &err)
	assert.EqualError(t, err, "error")
	assert.Equal(t, 2, deleted)
}

func TestHookUseNonEventHandler(t *testing.T) {
	h := eventHandler{}
	err := h.use("something else")
	assert.EqualError(t, err, "does not implement any event handler interface")
	assert.Len(t, h.onFindH, 0)
	assert.Len(t, h.onFoundH, 0)
	assert.Len(t, h.onGetH, 0)
	assert.Len(t, h.onGotH, 0)
	assert.Len(t, h.onInsertH, 0)
	assert.Len(t, h.onInsertedH, 0)
	assert.Len(t, h.onUpdateH, 0)
	assert.Len(t, h.onUpdatedH, 0)
	assert.Len(t, h.onDeleteH, 0)
	assert.Len(t, h.onDeletedH, 0)
	assert.Len(t, h.onClearH, 0)
	assert.Len(t, h.onClearedH, 0)
}
