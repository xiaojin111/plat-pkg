package tcphealthcheck

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	"bufio"
	"errors"
	"io"
	"net"
	"strings"

	mlog "github.com/jinmukeji/go-pkg/log"
)

var slog = mlog.StandardLogger()

type tcpHealthCheck struct {
	addr    string
	enabled bool
}

func (p *tcpHealthCheck) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "tcp_healthcheck_addr",
			Usage:       "TCP healthcheck listen address e.g :9901",
			EnvVar:      "TCP_HEALTHCHECK_ADDR",
			Value:       ":9901", // default address
			Destination: &(p.addr),
		},
		cli.BoolFlag{
			Name:        "enable_tcp_healthcheck",
			Usage:       "Enable TCP healthcheck",
			EnvVar:      "ENABLE_TCP_HEALTHCHECK",
			Destination: &(p.enabled),
		},
	}
}

func (p *tcpHealthCheck) Commands() []cli.Command {
	return nil
}

func (p *tcpHealthCheck) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		// 什么都不包装，透传
		return h
	}
}

func (p *tcpHealthCheck) Init(ctx *cli.Context) error {
	if p.enabled {
		go serveTCP(p.addr)
	}
	return nil
}

func (p *tcpHealthCheck) String() string {
	return "TCPHealthCheck"
}

func NewPlugin() plugin.Plugin {
	return &tcpHealthCheck{}
}

func serveTCP(addr string) {
	slog.Infof("[TCP Health] serving at %s", addr)

	l, err := net.Listen("tcp4", addr)
	if err != nil {
		slog.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			slog.Fatalf("[TCP Health] accept connection error: %v", err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	slog.Debugf("[TCP Health] check from %s", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			// 有的健康检查不发送数据过来，类似 HAProxy 的 TCP 健康检查方式
			// 这种情况视为正常
			if errors.Is(err, io.EOF) {
				break
			}

			slog.Warnf("[TCP Health] read data error: %v", err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		_, err = c.Write([]byte("healthy"))
		if err != nil {
			slog.Warnf("[TCP Health] write reply error: %v", err)
		}
	}

	//nolint:errcheck
	c.Close()
}
