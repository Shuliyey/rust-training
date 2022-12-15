package main

import (
	"fmt"
	"io/ioutil"
	"path"
	pb "image-resize-client/proto"
	"google.golang.org/grpc"
	"context"
	"os"
)

func main(){
	ch := make(chan error)
	go sendResizeRequest("images/walle.jpg", 100, 200, ch)
 	go sendResizeRequest("images/walle.jpg", 200, 300, ch)
	go sendResizeRequest("images/walle.jpg", 600, 400, ch)
	go sendResizeRequest("images/walle.jpg", 900, 950, ch)
	go sendResizeRequest("images/walle.jpg", 2000, 2000, ch)
	go sendResizeRequest("images/walle.jpg", 3000, 3000, ch)
	go sendResizeRequest("images/eva.jpg", 200, 300, ch)
	go sendResizeRequest("images/eva.jpg", 400, 700, ch)
	go sendResizeRequest("images/eva.jpg", 500, 750, ch)
	go sendResizeRequest("images/eva.jpg", 700, 500, ch)
	go sendResizeRequest("images/eva.jpg", 2000, 2000, ch)
	go sendResizeRequest("images/eva.jpg", 3000, 3000, ch)

	for range make([]bool, 12) {
		<- ch
	}
}

func sendResizeRequest(image_path string, width uint32, height uint32, ch chan error) error {
	conn, err := grpc.Dial("[::1]:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect to grpc endpoint [::1]:50051 (err: " + err.Error() + ")")
		ch <- err
		return err
	}

	defer conn.Close()

	client := pb.NewResizeClient(conn)

	image, err := ioutil.ReadFile(image_path)
	if err != nil {
		fmt.Println("invalid path " + image_path + "(err: " + err.Error() + ")")
		ch <- err
		return err
	}

	req := pb.ResizeRequest{
		Name: path.Base(image_path),
		Chunk: image,
		Width: width,
		Height: height,
	}

	resp, err := client.Send(context.Background(), &req)

	if (err != nil) {
		fmt.Println("grpc call resize.Resize/Send failed (err: " + err.Error() + ")")
		ch <- err
		return err
	}

	p := fmt.Sprintf("images/out/%s", resp.Name)
	file, err := os.Create(p)

	if (err != nil) {
		fmt.Println("failed to open file " + p + " (err: " + err.Error() + ")")
		ch <- err
		return err
	}

	_, err = file.Write(resp.Chunk)

	if (err != nil) {
		fmt.Println("failed to write file " + p + " (err: " + err.Error() + ")")
		ch <- err
		return err
	}

	defer file.Close()

	ch <- nil
	return nil
}