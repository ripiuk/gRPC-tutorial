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
	if err != nil {log.Fatalf("Unexpected error: %v\n", err)}
	fmt.Printf("Blog has been created: %v\n", res)
	blogID := res.GetBlog().GetId()


	// Read blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "123j1h21"})
	if err2 != nil {
		// Should return InvalidArgument
		fmt.Printf("Error while reading: %v\n", err2)
	}

	_, err3 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5d932d4d6fd02778036254d2"})
	if err3 != nil {
		// Should return NotFound
		fmt.Printf("Error while reading: %v\n", err3)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	resp, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error while reading: %v\n", readBlogErr)
	}
	fmt.Printf("Blog read results: %v\n", resp)

	// Update blog
	fmt.Println("Updating the blog")

	updatedBlog := &blogpb.Blog{
		Id: blogID,
		AuthorId: "Changed Author",
		Title: "Some first Blog (edited)",
		Content: "Here is some content, with additional info",
	}
	updateResp, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: updatedBlog})
	if updateErr != nil {
		fmt.Printf("Error while updating: %v\n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateResp)
}
