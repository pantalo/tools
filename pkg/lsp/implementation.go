// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"

	"golang.org/x/tools/pkg/lsp/protocol"
	"golang.org/x/tools/pkg/lsp/source"
)

func (s *Server) implementation(ctx context.Context, params *protocol.ImplementationParams) ([]protocol.Location, error) {
	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, source.Go)
	defer release()
	if !ok {
		return nil, err
	}
	return source.Implementation(ctx, snapshot, fh, params.Position)
}
