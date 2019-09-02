package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEngine_ProxyFor(t *testing.T) {
	host := "test"
	config := Config{
		Backends: []Backend{
			{
				Host:    host,
				RawURLs: []string{"http://localhost:81", "http://localhost:82"},
			},
		},
	}
	engine, _ := NewEngine(config)

	t.Run("should not have data races", func(t *testing.T) {
		go func() {
			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "http://localhost/test", nil)
			engine.ProxyFor(host, recorder, req)
		}()

		go func() {
			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "http://localhost/test", nil)
			engine.ProxyFor(host, recorder, req)
		}()
	})

}
