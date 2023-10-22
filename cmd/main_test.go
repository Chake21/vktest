package main

import (
	"context"
	"log"
	"testing"

	"vktest/internal"
	"vktest/vktest/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connection(target string) *grpc.ClientConn {
	connect, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot connect to service")
	}
	return connect
}

var (
	conn   = connection("localhost:1488")
	client = api.NewVKTestClient(conn)
	ctx    = context.Background()
)

func TestPositive(t *testing.T) {
	t.Parallel()
	var (
		res *api.CountOfUsersResponse
		err error
	)
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 2,
		AgeTo:   8,
	}
	res, err = client.CountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if !(res.Count == 4) {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestBelowZero(t *testing.T) {
	t.Parallel()
	var (
		res *api.CountOfUsersResponse
		err error
	)
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: -1,
		AgeTo:   8,
	}
	res, err = client.CountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cannot calculate count of users BelowZero Case: ", err)
	}
	if !(res.Count == 4) {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestInversive(t *testing.T) {
	t.Parallel()
	var (
		res *api.CountOfUsersResponse
		err error
	)
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 8,
		AgeTo:   2,
	}
	res, err = client.CountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cannot calculate count of users Inversion Case: ", err)
	}
	if !(res.Count == 4) {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestZeroLength(t *testing.T) {
	t.Parallel()
	var (
		res *api.CountOfUsersResponse
		err error
	)
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 3,
		AgeTo:   3,
	}
	res, err = client.CountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cannot calculate count of users ZeroLength Case: ", err)
	}
	if !(res.Count == 1) {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestNegative(t *testing.T) {
	t.Parallel()
	var (
		res *api.CountOfUsersResponse
		err error
	)
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 10,
		AgeTo:   12,
	}
	res, err = client.CountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cannot calculate count of users Negative Case: ", err)
	}
	if res.Found {
		t.Fatal("Found any results, expected no one!")
	}
	if !(res.Count == 0) {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestStreamPositive(t *testing.T) {
	t.Parallel()
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 2,
		AgeTo:   8,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cant read from stream", err)
	}
	if !res2.Found {
		t.Fatal("Not found!")
	}
	if !(res2.Count == 4) {
		t.Fatal("Result of calculation isnt correct!")
	}
}
