// This file is part of MinIO Direct CSI
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package metrics

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/golang/glog"
)

const (
	port        = "80"
	metricsPath = "direct-csi/metrics"
)

func ServeMetrics(ctx context.Context, nodeId string) {

	server := &http.Server{
		Handler: metricsHandler(nodeId),
	}

	lc := net.ListenConfig{}
	listener, lErr := lc.Listen(ctx, "tcp", fmt.Sprintf(":%v", port))
	if lErr != nil {
		panic(lErr)
	}

	glog.Infof("Starting metrics exporter in port: %s", port)
	if err := server.Serve(listener); err != nil {
		glog.Errorf("Failed to listen and serve metrics server: %v", err)
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
