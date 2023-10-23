package server

import (
	"context"
	"sync"

	"vktest/internal"
	"vktest/vktest/api"
)

type Server struct {
}

func (s *Server) CountOfUsers(ctx context.Context, req *api.CountOfUsersRequest) (*api.CountOfUsersResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		found, count := s.calculateCountOfUsers(req.Array, req.AgeFrom, req.AgeTo)
		return &api.CountOfUsersResponse{
			Found: found,
			Count: count,
		}, nil
	}
}

func (s *Server) StreamCountOfUsers(req *api.CountOfUsersRequest, stream api.VKTest_StreamCountOfUsersServer) error {
	countArray := internal.NewIntArray()
	boolArray := internal.NewBoolArray()
	var wg sync.WaitGroup
	arrays := arrayPartitioner(req.Array, internal.TenPartitions)
	for _, v := range arrays {
		wg.Add(1)
		go func(array []int32) {
			defer wg.Done()
			found, count := s.calculateCountOfUsers(array, req.AgeFrom, req.AgeTo)
			countArray.Append(int32(count))
			boolArray.Append(found)
		}(v)
	}
	wg.Wait()
	err := stream.Send(&api.CountOfUsersResponse{
		Found: boolArray.AnyTrue(),
		Count: uint64(countArray.Sum()),
	})
	return err
}

func (s *Server) calculateCountOfUsers(array []int32, ageFrom, ageTo int32) (bool, uint64) {
	var (
		counter uint64
		found   bool
	)
	if ageTo > ageFrom {
		for _, v := range array {
			if v <= ageTo && v >= ageFrom {
				counter++
			}
		}
		if counter > 0 {
			found = true
		}
		return found, counter
	}
	for _, v := range array {
		if v >= ageTo && v <= ageFrom {
			counter++
		}
	}
	if counter > 0 {
		found = true
	}
	return found, counter
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
