package main

import (
	"fmt"
	"strconv"
)

const (
	Mod = 8192
)

func ParseShardId(oId string) (int64, int, int, error) {
	id, err := strconv.ParseInt(oId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return -1, -1, -1, err
	}
	timestamp := (id >> 23) + 1314220021721

	shardId := (id << 41) >> 51

	if shardId <= 0 {
		shardId = int64(Mod) + shardId
	}

	seqId := (id << 54) >> 54
	if seqId <= 0 {
		seqId = 1024 + seqId
	}

	return timestamp, int(shardId - 1), int(seqId - 1), nil
}

func main() {
	_, b, _, _ := ParseShardId("1958607939402952060")
	fmt.Println(b)
}
