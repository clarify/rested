package rest

import (
	"context"
	"testing"

	"github.com/clarify/rested/resource"
	"github.com/clarify/rested/schema"
	"github.com/stretchr/testify/assert"
)

func TestRestResourceGetSubResource(t *testing.T) {
	index := resource.NewIndex()
	index.Bind("foo", schema.Schema{}, nil, resource.DefaultConf)
	ctx := context.Background()
	rsc := restResource{}
	_, err := rsc.SubResource(ctx, "bar")
	assert.EqualError(t, err, "router not available in context")
	ctx = contextWithIndex(ctx, index)
	_, err = rsc.SubResource(ctx, "bar")
	assert.EqualError(t, err, "invalid resource reference: bar")
}
