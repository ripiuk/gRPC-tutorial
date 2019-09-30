package main

import (
	"context"
	"fmt"
	"log"

	"gRPC_course/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Running Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8000", opts)
	if err != nil {
		log.Fatalf("Could not conect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// Create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Sashko",
		Title: "Some first Blog",
		Content: "Here is some content",
	}

	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {log.Fatalf("Unexpected error: %v", err)}
	fmt.Printf("Blog has been created: %v", res)
}
