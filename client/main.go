package main

import (
	pb "Cw_TileService/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type Authentication struct {
	Login string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string , error){
	return map[string]string{
		"login": a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}


func makingConnection() (*grpc.ClientConn, error){
	// creating creds for grpc client connection
	//creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	//if err != nil {
	//	log.Fatalf("could not load tls cert: %s", err)
	//}
	//
	//auth := Authentication{
	//	Login:    "nayan",
	//	Password: "makasare",
	//}

	//conn , err := grpc.Dial("cloudwalker.services.tv:7771", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	conn , err := grpc.Dial("192.168.1.143:7773", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, err
}

func main(){
	// stress testing run the rpc for how many times
	run_test_index := 1;

	var wg sync.WaitGroup
	conn, err := makingConnection()
	if err != nil {
		log.Fatal(err)
	}
	c :=  pb.NewTileServiceClient(conn)


	wg.Add(run_test_index)
	for i:=0 ; i < run_test_index ; i++ {
		go CloudwalkerPrimePages(c, &wg)
		//go GetPage(c, &wg)
		//go GetCarousel(c, &wg)
		//go GetRow(c, &wg)
	}
	wg.Wait()
}


func CloudwalkerPrimePages(c pb.TileServiceClient, wg *sync.WaitGroup){
	resp, err := c.CloudwalkerPrimePages(context.Background(), &pb.PrimePagesRequest{
		Vendor:               "cvte",
		Brand:                "shinko",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Pages)
	wg.Done()
}

func GetPage(c pb.TileServiceClient, wg *sync.WaitGroup){
	resp, err := c.GetPage(context.Background(), &pb.PageRequest{
		Vendor:               "cvte",
		Brand:                "shinko",
		PageName:             "movies",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	wg.Done()
}

func GetCarousel(c pb.TileServiceClient, wg *sync.WaitGroup){
	resp, err := c.GetCarousel(context.Background(), &pb.CarouselRequest{
		Vendor:               "cvte",
		Brand:                "shinko",
		PageName:             "movies",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	wg.Done()
}

func GetRow(c pb.TileServiceClient, wg *sync.WaitGroup){
	resp, err := c.GetRow(context.Background(), &pb.RowRequest{
		Vendor:               "cvte",
		Brand:                "shinko",
		PageName:             "movies",
		RowName:              "hindi_dub_movies",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	wg.Done()
}

