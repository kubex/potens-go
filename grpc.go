package potens

import (
	"crypto/tls"
	"errors"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/cubex/portcullis-go/keys"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// GetGrpcContext context to use when communicating with other services
func (app *Application) GetGrpcContext() context.Context {
	md := metadata.Pairs(
		keys.GetAppIDKey(), app.Definition().AppID,
		keys.GetAppVendorKey(), app.Definition().VendorID,
	)
	return metadata.NewContext(context.Background(), md)
}

// CreateServer creates a gRPC server with your tls certificates
func (app *Application) CreateServer() error {

	app.server = grpc.NewServer()

	//Do not secure with imperium for initial development
	if true {
		return nil
	}

	if app.imperiumKey == nil || app.imperiumCertificate == nil || app.hostname == "" {
		return errors.New("CreateServer called before GetCertificate, or GetCertificate call failed")
	}

	cert, err := tls.X509KeyPair(app.imperiumCertificate, app.imperiumKey)
	if err != nil {
		return err
	}

	app.server = grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	return nil
}

//GetServer returns the grpc server
func (app *Application) GetServer() *grpc.Server {
	return app.server
}

func (app *Application) Serve() error {

	lis, err := net.Listen("tcp", app.hostname+":"+app.PortString())
	if err != nil {
		return err
	}

	return app.server.Serve(lis)
}

func (app *Application) GetServiceConnection(service string) (*grpc.ClientConn, error) {
	location := os.Getenv(strings.ToUpper(service) + EnvServiceLocationSuffix)

	kubexServiceDomain := os.Getenv(EnvKubexServiceDomain)
	if kubexServiceDomain == "" {
		kubexServiceDomain = KubexProductionServicesDomain
	}

	if location == "" {
		location = strings.ToLower(service) + "." + kubexServiceDomain
		location += ":" + strconv.FormatInt(int64(KubexDefaultGRPCPort), 10)
	}

	return grpc.Dial(location, grpc.WithInsecure())
	//return grpc.Dial(location, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
}
