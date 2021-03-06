package isogrids

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	key string
)

func init() {
	h := md5.New()
	io.WriteString(h, "hello")
	key = fmt.Sprintf("%x", h.Sum(nil)[:])
}

func TestHexa(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	Hexa(rec, key, colorTheme, 10, 10)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

	rec = httptest.NewRecorder()
	Hexa(rec, key, colorTheme, 600, 60)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}

	rec = httptest.NewRecorder()
	Hexa(rec, key, colorTheme, 400, 40)
	if rec.Code != http.StatusOK {
		t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
	}
}
