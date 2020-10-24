// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package server

import (
	"bytes"
	"compress/gzip"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/guark/guark/app"
	"github.com/guark/guark/log"
)

type Handler struct {
	log   log.Log
	embed *app.Embed
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}

	if mt := mime.TypeByExtension(filepath.Ext(r.URL.Path)); mt != "" {
		w.Header().Set("Content-Type", mt)
	}

	gz, e := h.embed.Data(r.URL.Path)

	if e != nil {
		w.Write([]byte(e.Error()))
		return
	}

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {

		w.Header().Set("Content-Encoding", "gzip")
		w.Write(*gz)
		return
	}

	reader, e := gzip.NewReader(bytes.NewReader(*gz))

	if e != nil {
		w.Write([]byte(e.Error()))
		return
	}

	if _, err := io.Copy(w, reader); err != nil {
		h.log.Error(err)
	}
}
