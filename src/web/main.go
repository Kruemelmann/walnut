package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	//web server settings
	host = "localhost"
	port = "8000"
	// cookie settings
	cookieSessionID     = "walnut_session-id"
	cookieSessionMaxAge = 60 * 60 * 10 // 10 hours
)

type ctxSessionID struct{}

//TODO add the grpc bindings
type webServer struct {
	authAddr string
	authCon  *grpc.ClientConn

	farmerAddr string
	farmerCon  *grpc.ClientConn

	storeAddr string
	storeCon  *grpc.ClientConn

	dataAddr string
	dataCon  *grpc.ClientConn
}

func main() {
	ctx := context.Background()

	ws := new(webServer)
	//read envs
	extractEnv(&ws.authAddr, "AUTH_ADDRESS")
	extractEnv(&ws.farmerAddr, "FARM_ADDRESS")
	extractEnv(&ws.storeAddr, "STORE_ADDRESS")
	extractEnv(&ws.dataAddr, "DATA_ADDRESS")

	//connect grpc services
	connectGRPCService(ctx, &ws.authCon, ws.authAddr)
	connectGRPCService(ctx, &ws.farmerCon, ws.farmerAddr)
	connectGRPCService(ctx, &ws.storeCon, ws.storeAddr)
	connectGRPCService(ctx, &ws.dataCon, ws.dataAddr)

	//configure webserver
	r := mux.NewRouter()
	r.HandleFunc("/", ws.testHandler).Methods(http.MethodGet, http.MethodHead)

	var handler http.Handler = r
	handler = checkSessionID(handler)

	adress := host + ":" + port
	log.Println("Server started on " + adress)
	log.Fatal(http.ListenAndServe(adress, handler))
}

func connectGRPCService(ctx context.Context, con **grpc.ClientConn, addr string) {
	fmt.Println(addr)
}

func extractEnv(dest *string, key string) {
	viper.BindEnv(key)
	*dest = fmt.Sprintf("%s", viper.Get(key))
}
