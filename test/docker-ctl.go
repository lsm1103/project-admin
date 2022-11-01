package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	//helper, err := connhelper.GetCommandConnectionHelper("ssh://pujian:pujian666@172.16.10.183:22")
	////helper, err := connhelper.GetConnectionHelper("ssh://pujian:pujian666@172.16.10.183:22")
	//if err != nil {
	//	panic(err)
	//}
	//httpClient := &http.Client{
	//	Transport: &http.Transport{
	//		DialContext: helper.Dialer,
	//	},
	//}
	//cl, err := client.NewClientWithOpts(
	//	client.WithHTTPClient(httpClient),
	//	client.WithHost(helper.Host),
	//	client.WithDialContext(helper.Dialer),
	//)

	cl, err := client.NewClient("tcp://172.16.10.183:2375", "", nil, nil)
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	fmt.Println(cl.ImageList(context.Background(), types.ImageListOptions{}))
}