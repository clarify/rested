package rest_test

import (
	"context"
	"net/http"

	"github.com/clarify/rested/resource"
	"github.com/clarify/rested/rest"
)

type myResponseFormatter struct {
	// Extending default response sender
	rest.DefaultResponseFormatter
}

// Add a wrapper around the list with pagination info
func (r myResponseFormatter) FormatList(ctx context.Context, headers http.Header, l *resource.ItemList, skipBody bool) (context.Context, interface{}) {
	ctx, data := r.DefaultResponseFormatter.FormatList(ctx, headers, l, skipBody)
	return ctx, map[string]interface{}{
		"meta": map[string]int{
			"total":  l.Total,
			"offset": l.Offset,
		},
		"list": data,
	}
}

func ExampleResponseSender() {
	index := resource.NewIndex()
	api, _ := rest.NewHandler(index)
	api.ResponseFormatter = myResponseFormatter{}
}
