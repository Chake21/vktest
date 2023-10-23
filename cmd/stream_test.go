package main

import (
	"testing"

	"vktest/internal"
	"vktest/vktest/api"
)

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
	if *res2.Count != 4 {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestStreamSearch(t *testing.T) {
	t.Parallel()
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 0,
		AgeTo:   1000,
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
		t.Fatal("Not found! Expected at least one!")
	}
}

func TestStreamBelowZero(t *testing.T) {
	t.Parallel()
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: -1,
		AgeTo:   8,
	}

	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if *res2.Count != 4 {
		t.Fatal("Incorrect calculation!")
	}
}

func TestStreamInversive(t *testing.T) {
	t.Parallel()
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 8,
		AgeTo:   2,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if *res2.Count != 4 {
		t.Fatal("Result of calculation is incorrect!")
	}
}

func TestStreamEmptyArray(t *testing.T) {
	req := api.CountOfUsersRequest{
		Array:   []int32{},
		AgeFrom: 8,
		AgeTo:   2,
	}
	res, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("cannot read from stream!")
	}
	_, err2 := res.Recv()
	if err2 == nil {
		t.Fatal("Error is empty!")
	}
}

func TestStreamZeroLength(t *testing.T) {
	t.Parallel()

	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 8,
		AgeTo:   8,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if res2.Found {
		t.Fatal("Founded result, expected no one!")
	}
	if *res2.Count != 0 {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestStreamZeroLengthEqualToNumberInArray(t *testing.T) {
	t.Parallel()

	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 3,
		AgeTo:   3,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if !res2.Found {
		t.Fatal("Found nothing! Expected at least something!")
	}
	if *res2.Count != 1 {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestStreamNegative(t *testing.T) {
	req := api.CountOfUsersRequest{
		Array:   internal.UsersAgeMocked,
		AgeFrom: 10,
		AgeTo:   12,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if res2.Found {
		t.Fatal("Found any results, expected no one!")
	}
	if *res2.Count != 0 {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestStreamSearchLefterThenArraysBorder(t *testing.T) {
	t.Parallel()
	req := api.CountOfUsersRequest{
		Array:   []int32{55, 66, 77, 88},
		AgeFrom: 10,
		AgeTo:   12,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if res2.Found {
		t.Fatal("Found any results, expected no one!")
	}
	if *res2.Count != 0 {
		t.Fatal("Result of calculation isnt correct!")
	}
}

func TestStreamSearchRighterThenArraysBorder(t *testing.T) {
	t.Parallel()
	req := api.CountOfUsersRequest{
		Array:   []int32{55, 66, 77, 88},
		AgeFrom: 90,
		AgeTo:   100,
	}
	cli, err := client.StreamCountOfUsers(ctx, &req)
	if err != nil {
		t.Fatal("Cant read from stream! ", err)
	}
	res2, err2 := cli.Recv()
	if err2 != nil {
		t.Fatal("Cannot calculate count of users Positive Case: ", err)
	}
	if res2.Found {
		t.Fatal("Found any results, expected no one!")
	}
	if *res2.Count != 0 {
		t.Fatal("Result of calculation isnt correct!")
	}
}
