// Package main implements a server for movieinfo service.
package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/glatham1/CloudNativeCourse/Lab5/movieapi"
	"google.golang.org/grpc"

	"errors"
	"fmt"
)

const (
	port = ":50051"
)

// server is used to implement movieapi.MovieInfoServer
type server struct {
	movieapi.UnimplementedMovieInfoServer
}

// Map representing a database
var moviedb = map[string][]string{"Pulp fiction": []string{"1994", "Quentin Tarantino", "John Travolta,Samuel Jackson,Uma Thurman,Bruce Willis"}}

// GetMovieInfo implements movieapi.MovieInfoServer
func (s *server) GetMovieInfo(ctx context.Context, in *movieapi.MovieRequest) (*movieapi.MovieReply, error) {
	title := in.GetTitle()
	log.Printf("Received: %v", title)
	reply := &movieapi.MovieReply{}
	if val, ok := moviedb[title]; !ok { // Title not present in database
		return reply, nil
	} else {
		if year, err := strconv.Atoi(val[0]); err != nil {
			reply.Year = -1
		} else {
			reply.Year = int32(year)
		}
		reply.Director = val[1]
		cast := strings.Split(val[2], ",")
		reply.Cast = append(reply.Cast, cast...)

	}

	return reply, nil

}

//Set MovieInfo function and implents movieapi.MovieData
func (s *server) SetMovieInfo(ctx context.Context, in *movieapi.MovieData) (*movieapi.Status, error) {
	//gather input from movieapi
	title := in.GetTitle()
	yearInput := in.GetYear()
	director := in.GetDirector()
	castInput := in.GetCast()
	
	//adjusts movieapi "status" and inputs
	status := &movieapi.Status{}
	year := fmt.Sprint(yearInput)
	cast := strings.Join(castInput, ",")
	
	//creates a map and adds new movie info to map
	movieMap := make([]string, 0)
	movieMap = append(movieMap, year, director, cast)
	
	//movie checks to see if it already exists in the database
	if _, ok := moviedb[title]; !ok {
		//if the movie doesnt exist it adds it to the database and sends a success
		//message for the status
		moviedb[title] = movieMap
		status.Code = "success"
		
		return status, nil
	}
	//if the movie is in the database it sets the status message to a failure and returns an error message
	status.Code = "failure"
	
	return status, errors.New("Movie already in database")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	movieapi.RegisterMovieInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
