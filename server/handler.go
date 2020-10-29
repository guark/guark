// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package server

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/guark/guark/app"
	"github.com/guark/guark/log"
)

type Handler struct {
	log   log.Logger
	embed *app.Embed
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}

	if mt := mime.TypeByExtension(filepath.Ext(r.URL.Path)); mt != "" {
		w.Header().Set("Content-Type", mt)
	}

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {

		w.Header().Set("Content-Encoding", "gzip")

		gz, e := h.embed.Data(r.URL.Path)
		if e != nil {
			w.Write([]byte(e.Error()))
			return
		}

		w.Write(*gz)
		return
	}

	data, e := h.embed.UngzipData(r.URL.Path)
	if e != nil {
		w.Write([]byte(e.Error()))
		return
	}

	w.Write(data)
}
