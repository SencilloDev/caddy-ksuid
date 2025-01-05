// Copyright 2025 Sencillo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ksuid

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/segmentio/ksuid"
)

func init() {
	caddy.RegisterModule(Ksuid{})
	httpcaddyfile.RegisterHandlerDirective("ksuid", parseCaddyfile)
}

type Ksuid struct{}

func (Ksuid) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.ksuid",
		New: func() caddy.Module { return new(Ksuid) },
	}
}

func (m *Ksuid) Provision(ctx caddy.Context) error {

	return nil
}

func (m Ksuid) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	id := ksuid.New().String()

	repl.Set("ksuid.id", id)

	return next.ServeHTTP(w, r)
}

func (m *Ksuid) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {

	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := new(Ksuid)
	err := m.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Ksuid)(nil)
	_ caddyhttp.MiddlewareHandler = (*Ksuid)(nil)
	_ caddyfile.Unmarshaler       = (*Ksuid)(nil)
)
