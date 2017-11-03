package srv

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"syscall"
	"os"
	"os/signal"
	"go-ticket/lla"
)

// Database Configuration
type ServConfig struct {
	LocalIp string
	NetPort string
}


// Print config data to console
func (cfg *ServConfig) PrintData() {
	fmt.Println("LocalIp ", cfg.LocalIp)
	fmt.Println("NetPort ", cfg.NetPort)
}
// Get formatted string with config data
func (cfg *ServConfig) String() string {
	str := fmt.Sprintf("Database config: " +
		"LocalIp = %s, NetPort = %s.",
		cfg.LocalIp, cfg.NetPort)
	return str
}


// Log Agent for service layer logging
var srvLog lla.LogAgent

func InitLoggerSRV(level int){
	srvLog.Init(level, "SRV")
}

///////////////////////////////////////////////////////////////////////



func Run(srvCfg *ServConfig) error {
	netSpec := fmt.Sprintf("%s:%s", srvCfg.LocalIp, srvCfg.NetPort)
	srvLog.Info("Starting, HTTP on: %s\n", netSpec)

	listener, err := net.Listen("tcp", netSpec)
	if err != nil {
		srvLog.Error("Error creating listener: %v", err)
		return err
	}

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	go server.Serve(listener)

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	srvLog.Info("Got signal: %v, exiting.", s)
}

