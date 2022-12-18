package aoc

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestFetcher(t *testing.T) {
	dir, err := os.MkdirTemp("", "aoc")
	if err != nil {
		t.Fatalf("could not create temp directory: %v", err)
	}
	t.Cleanup(func() {
		_ = os.RemoveAll(dir)
	})

	// write token file

	count := 0

	handler := func(w http.ResponseWriter, r *http.Request) {
		count++
		if count > 1 {
			t.Errorf("unexpected request")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// check URL
		if want, got := "https://adventofcode.com/2022/day/12/input", r.URL.String(); got != want {
			t.Errorf("want: %#v, got %#v", want, got)
		}

		// check session
		c, err := r.Cookie("session")
		if err != nil {
			t.Errorf("could not extract session cookie: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if want, got := "secret", c.Value; got != want {
			t.Errorf("want: %#v, got: %#v", want, got)
		}

		w.Header().Add("Content-Type", "text/plain")
		_, _ = w.Write([]byte("server input"))
	}

	f := fetcher{
		dir: dir,
		client: &http.Client{
			Transport: &testRoundTripper{handler},
		},
	}

	// request without token, should panic
	tokenPath := filepath.Join(dir, "token")
	assertPanic(t, "missing token file; you must store your AoC session token in file "+tokenPath, func() {
		f.fetchInput(2022, 12)
	})

	// write token file
	if err := os.WriteFile(tokenPath, []byte("secret"), 0o644); err != nil {
		t.Fatalf("failed to write token file: %v", err)
	}

	// first proper request, should contact server
	data := f.fetchInput(2022, 12)
	input, err := io.ReadAll(data)
	if err != nil {
		t.Fatalf("failed to read input data: %v", err)
	}

	if count != 1 {
		t.Errorf("server was not contacted")
	}

	if want, got := "server input", string(input); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}

	// second request, should hit cache
	data = f.fetchInput(2022, 12)
	input, err = io.ReadAll(data)
	if err != nil {
		t.Fatalf("failed to read input data: %v", err)
	}

	if count > 1 {
		t.Errorf("server was contacted twice")
	}

	if want, got := "server input", string(input); got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}

type testRoundTripper struct {
	handler http.HandlerFunc
}

func (t *testRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	t.handler.ServeHTTP(rec, r)
	return rec.Result(), nil
}
