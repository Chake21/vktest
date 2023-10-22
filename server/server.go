package server

import (
	"context"
	"log"
	"sync"
	"vktest/internal"

	"vktest/vktest/api"
)

type Server struct {
}

func (s *Server) CountOfUsers(ctx context.Context, req *api.CountOfUsersRequest) (*api.CountOfUsersResponse, error) {
	res := s.calculateCountOfUsers(req.Array, req.AgeFrom, req.AgeTo)
	return &api.CountOfUsersResponse{Count: res}, nil
}

func (s *Server) StreamCountOfUsers(req *api.CountOfUsersRequest, stream api.VKTest_StreamCountOfUsersServer) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	arrays := arrayPartitioner(req.Array, internal.TenPartitions)
	for _, v := range arrays {
		wg.Add(1)
		go func(array []int32) {
			defer wg.Done()
			err := stream.Send(&api.CountOfUsersResponse{Count: s.calculateCountOfUsers(req.Array, req.AgeTo, req.AgeFrom)})
			if err != nil {
				errChan <- err
			}
		}(v)
	}
	wg.Wait()
	if len(errChan) != 0 {
		log.Fatal("Errors in concurrency calculation!")
	}
	return nil
}

func (s *Server) calculateCountOfUsers(array []int32, ageFrom, ageTo int32) uint64 {
	var counter uint64
	if ageTo > ageFrom {
		for _, v := range array {
			if v <= ageTo && v >= ageFrom {
				counter++
			}
		}
		return counter
	}
	for _, v := range array {
		if v >= ageTo && v <= ageFrom {
			counter++
		}
	}
	return counter
}

func arrayPartitioner(array []int32, countOfPartitions internal.PartitionsCount) [][]int32 {
	var res [][]int32
	length := int32(len(array))
	size := length / countOfPartitions.Count
	var start int32
	for start < length-size {
		res = append(res, array[start:start+size])
		start += size
	}
	res = append(res, array[start:])
	return res
}
