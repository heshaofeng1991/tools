package main

import (
	"combo/domain/service"
	ent "combo/ent"
	admin "combo/gen/admin"
	healthy "combo/gen/healthy"
	web "combo/gen/web"
	"combo/internal/dao"

	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	_ "combo/ent/runtime"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		versionF  = flag.String("version", "v1", "API version")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger

		entClient *ent.Client
	)
	{
		// Initialize the logger.
		logger = log.New(os.Stderr, "[combobox] ", log.Ltime)

		// Initialize the db.
		entClient = dao.Open()
	}

	// Initialize the services.
	var (
		healthySvc healthy.Service
		adminSvc   admin.Service
		webSvc     web.Service
	)
	{
		healthySvc = service.NewHealthy(logger)
		adminSvc = service.NewAdmin(entClient, logger)
		webSvc = service.NewWeb(entClient, logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		healthyEndpoints *healthy.Endpoints
		adminEndpoints   *admin.Endpoints
		webEndpoints     *web.Endpoints
	)
	{
		healthyEndpoints = healthy.NewEndpoints(healthySvc)
		adminEndpoints = admin.NewEndpoints(adminSvc)
		webEndpoints = web.NewEndpoints(webSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://0.0.0.0:80"
			addr = strings.Replace(addr, "{version}", *versionF, -1)
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, healthyEndpoints, adminEndpoints, webEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
