package ariaHelper

import (
	"aria-go-mirror-bot/env"
	"context"
	"fmt"
	aria "github.com/zyxar/argo/rpc"
	"os/exec"
	"time"
)

var RPC aria.Protocol

func init() {
	var err error
	RPC, err = aria.New(context.Background(), "http://localhost:6800/jsonrpc", "", time.Second, nil)
	if err != nil {
		panic(err)
	}

	if err := launchAria2cDaemon(); err != nil {
		panic(err)
	}
	version, err := RPC.GetVersion()
	if err != nil {
		panic(err)
	}

	_ = fmt.Sprintf("ARIA2C started! : %s", version.Version)
}
func launchAria2cDaemon() (err error) {
	args := env.Config.AriaArgs
	if env.Config.RpcSecret != "" {
		args = append(args, "--rpc-secret="+env.Config.RpcSecret)
	}

	cmd := exec.Command("aria2c", args...)
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("failed to open rpc from aria2c : %v", err)
	}
	return cmd.Process.Release()
}

