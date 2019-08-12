package taskrunner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/client/allocdir"
	"github.com/hashicorp/nomad/client/allocrunner/interfaces"
	agentconsul "github.com/hashicorp/nomad/command/agent/consul"
	"github.com/hashicorp/nomad/nomad/structs"
)

var _ interfaces.TaskPrestartHook = &connectHook{}

// connectHook writes the bootstrap config for the Connect proxy sidecar
type connectHook struct {
	alloc          *structs.Allocation
	consulHTTPAddr string

	logger log.Logger
}

func newConnectHook(alloc *structs.Allocation, consulHTTPAddr string, logger log.Logger) *connectHook {
	h := &connectHook{
		alloc:          alloc,
		consulHTTPAddr: consulHTTPAddr,
	}
	h.logger = logger.Named(h.Name())
	return h
}

func (connectHook) Name() string {
	return "connect"
}

func (h *connectHook) Prestart(ctx context.Context, req *interfaces.TaskPrestartRequest, resp *interfaces.TaskPrestartResponse) error {
	//FIXME(schmichael) use code from #6097
	if !strings.HasPrefix(req.Task.Kind, "connect-proxy:") {
		resp.Done = true
		return nil
	}

	//FIXME(schmichael) use code & error msg from #6097
	parts := strings.SplitN(req.Task.Kind, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("Connect proxy sidecar must specify service name")
	}
	serviceName := parts[1]

	tg := h.alloc.Job.LookupTaskGroup(h.alloc.TaskGroup)

	var service *structs.Service
	for _, s := range tg.Services {
		if s.Name == serviceName {
			service = s
		}
	}

	if service == nil {
		return fmt.Errorf("Connect proxy sidecar task exists but no services configured with a sidecar")
	}

	canary := false
	if h.alloc.DeploymentStatus != nil {
		canary = h.alloc.DeploymentStatus.Canary
	}

	//FIXME(schmichael) relies on GRPCSocket being created
	grpcAddr := "unix://" + filepath.Join(allocdir.SharedAllocName, allocdir.AllocGRPCSocket)

	fn := filepath.Join(req.TaskDir.LocalDir, "bootstrap.json")
	id := agentconsul.MakeTaskServiceID(h.alloc.ID, "group-"+tg.Name, service, canary)
	h.logger.Debug("bootstrapping envoy", "sidecar_for", service.Name, "boostrap_file", fn, "sidecar_for_id", id, "grpc_addr", grpcAddr)

	// Since Consul services are registered asynchronously with this task
	// hook running, retry a small number of times with backoff.
	for tries := 3; ; tries-- {
		cmd := exec.CommandContext(ctx, "consul", "connect", "envoy",
			"-grpc-addr", grpcAddr,
			"-http-addr", h.consulHTTPAddr,
			"-bootstrap",
			"-sidecar-for", id, // must use the id not the name!
		)

		// Redirect output to local/bootstrap.json
		fd, err := os.Create(fn)
		if err != nil {
			return fmt.Errorf("error creating local/bootstrap.json for envoy: %v", err)
		}
		cmd.Stdout = fd

		buf := bytes.NewBuffer(nil)
		cmd.Stderr = buf

		// Generate bootstrap
		err = cmd.Run()

		// Close bootstrap.json
		fd.Close()

		if err == nil {
			// Happy path! Bootstrap was created, exit.
			break
		}

		// Check for error from command
		if tries == 0 {
			h.logger.Error("error creating bootstrap configuration for Connect proxy sidecar", "error", err, "stderr", buf.String())
			return fmt.Errorf("error creating bootstrap configuration for Connect proxy sidecar: %v", err)
		}

		// Sleep before retrying to give Consul services time to register
		time.Sleep(3 * time.Second)
	}

	// Bootstrap written. Mark as done and move on.
	resp.Done = true
	return nil
}
