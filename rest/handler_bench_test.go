package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clarify/rested/resource"
	"github.com/clarify/rested/resource/testing/mem"
	"github.com/clarify/rested/schema"
)

func BenchmarkServeHTTP(b *testing.B) {
	i := resource.NewIndex()
	i.Bind("foo", schema.Schema{}, mem.NewHandler(), resource.DefaultConf)
	h, _ := NewHandler(i)
	r, _ := http.NewRequest("GET", "/foo", nil)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
	}
}
