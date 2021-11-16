package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {}

func TestHandler(t *testing.T) {
	data := "hello, world"
	// svr := Server(&StubStore{data})

	// request := httptest.NewRequest(http.MethodGet, "/", nil)
	// response := httptest.NewRecorder()

	// svr.ServeHTTP(response, request)

	// if response.Body.String() != data {
	// 	t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	// }

	t.Run("tells store to canel work if request is canceled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was nto told to cancel")
		}
	})
}
