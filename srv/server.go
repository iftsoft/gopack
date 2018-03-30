package srv

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"syscall"
	"os"
	"os/signal"
	"github.com/iftsoft/gopack/lla"
)

// Database Configuration
type ServConfig struct {
	LocalIp			string
	NetPort 		string
	ReadTimeout		int
	WriteTimeout	int
	UserTimeout		int
	SessionTime		int
	CookieName		string
	TokenKey		string
	ExtraKey		string
	StaticDir		string
}


// Print config data to console
func (cfg *ServConfig) PrintData() {
	fmt.Println("LocalIp ", cfg.LocalIp)
	fmt.Println("NetPort ", cfg.NetPort)
	fmt.Println("ReadTimeout ", cfg.ReadTimeout)
	fmt.Println("WriteTimeout ", cfg.WriteTimeout)
	fmt.Println("UserTimeout ", cfg.UserTimeout)
	fmt.Println("SessionTime ", cfg.SessionTime)
	fmt.Println("CookieName ", cfg.CookieName)
	fmt.Println("TokenKey ", len(cfg.TokenKey))
	fmt.Println("ExtraKey ", len(cfg.ExtraKey))
	fmt.Println("StaticDir ", cfg.StaticDir)
}
// Get formatted string with config data
func (cfg *ServConfig) String() string {
	str := fmt.Sprintf("Service config: " +
		"LocalIp = %s, NetPort = %s, ReadTimeout = %d sec, WriteTimeout = %d sec, " +
		"UserTimeout = %d sec, SessionTime = %d sec, CookieName = %s, " +
		"TokenKey = %d Byte, ExtraKey = %d Byte, StaticDir = \"%s\".",
		cfg.LocalIp, cfg.NetPort,  cfg.ReadTimeout,  cfg.WriteTimeout,
		cfg.UserTimeout,  cfg.SessionTime,  cfg.CookieName,
		len(cfg.TokenKey), len(cfg.ExtraKey), cfg.StaticDir)
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
	if srvCfg.ReadTimeout == 0	{	srvCfg.ReadTimeout = 30	}
	if srvCfg.WriteTimeout == 0	{	srvCfg.WriteTimeout = 20	}

	server := &http.Server{
		ReadTimeout:    time.Duration(srvCfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(srvCfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 16,
		}

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

