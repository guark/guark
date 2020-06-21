// +build !windows

package window

import (
	// "io"
	// "mime"
	"net"
	"net/http"

	// "path/filepath"

	"github.com/guark/guark/app"
	"github.com/sirupsen/logrus"
)

func (s *Server) serve() {

	var err error

	// TODO: get new port here.

	s.ln, err = net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		s.Log.Panic(err)
	}

	go http.Serve(s.ln, &srvHandler{mbd: s.App.Embed, log: s.Log})
}

type srvHandler struct {
	mbd *app.Embed
	log *logrus.Entry
}

func (h srvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))

	// if r.URL.Path == "/" {
	// 	r.URL.Path = "/index.html"
	// }

	// if ctype := mime.TypeByExtension(filepath.Ext(r.URL.Path)); ctype != "" {
	// 	w.Header().Set("Content-Type", ctype)
	// }

	// gz, e := h.assets.ReadAll(r.URL.Path)

	// if e != nil {
	// 	w.Write([]byte(e.Error()))
	// 	return
	// }

	// if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {

	// 	w.Header().Set("Content-Encoding", "gzip")
	// 	w.Write(gz)
	// 	return
	// }

	// reader, e := gzip.NewReader(bytes.NewReader(gz))

	// if e != nil {
	// 	w.Write([]byte(e.Error()))
	// 	return
	// }

	// if _, err := io.Copy(w, reader); err != nil {
	// 	h.log.Error(err)
	// }
}
