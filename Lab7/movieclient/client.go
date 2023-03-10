// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/glatham1/CloudNativeCourse/Lab7/movieapi"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
)

//change these
var newTitle = "The Shawshank Redemption"
var newYear int32 = 1994
var newDirector = "Frank Darabont"
var newCast = []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton"}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := movieapi.NewMovieInfoClient(conn)

	// Contact the server and print out its response.
	title := defaultTitle
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s \nYear: %d \nDirector: %s \nCast: %v", title, r.GetYear(), r.GetDirector(), r.GetCast())
	
	//creates a temp variable and sends the set new info of the movie to the api and then prints the status
	temp, err := c.SetMovieInfo(ctx, &movieapi.MovieData{Title: newTitle, Year: newYear, Director: newDirector, Cast: newCast})
	log.Print("SetMovieInfo Status:" + temp.GetCode())

	//calls the get movie info to display the new movie that was set
	display, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: newTitle})
	log.Printf("Movie Info for %s \nYear: %d \nDirector: %s \nCast: %v", newTitle, display.GetYear(), display.GetDirector(), display.GetCast())
}
