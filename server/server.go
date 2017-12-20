// Vision Screening Upload Simulator
// Copyright (C) 2017 Andrew Allen
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

var path = ""

func init() {
	path = fmt.Sprintf(
		"%s/src/github.com/achew22/acceptance-testing-vision-upload-server",
		os.Getenv("GOPATH"))
}

type Server struct {
	l    *log.Logger
	port int
}

func New(l *log.Logger, port int) *Server {
	return &Server{
		l:    l,
		port: port,
	}
}

func (s *Server) Run() {
	s.l.Printf("Starting AGPL'd server.")
	s.l.Printf("All server assets from %s are available at https://localhost:%d/assets/",
		path, s.port)

	mux := http.NewServeMux()
	assetsPrefix := "/assets/"
	mux.Handle(assetsPrefix, http.StripPrefix(assetsPrefix, http.FileServer(http.Dir(path))))
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
	})

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			// Order doesn't matter since it is attacker chosen.
			// List derived from Mozilla modern cipher list (Dec. 20, 2017).
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	err := srv.ListenAndServeTLS("certs/good_certificate.crt", "certs/good_key.pem")

	if err != nil {
		s.l.Fatalf("Unable to start listening on port %d with error:\n%v", s.port, err)
	}
}
