package internal

var UsersAgeMocked = []int32{2, 3, 4, 7, 9, 13, 17, 28, 44, 185, 72, 200}

type PartitionsCount struct {
	Count int32
}

var (
	ThreePartitions   = PartitionsCount{Count: 3}
	TenPartitions     = PartitionsCount{Count: 10}
	HundredPartitions = PartitionsCount{Count: 100}
)
