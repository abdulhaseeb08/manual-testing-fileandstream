package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/abdulhaseeb08/protocol/livekit"
	lksdk "github.com/abdulhaseeb08/server-sdk-go"
)

func main() {
	egressclient := lksdk.NewEgressClient("ws://localhost:7880", "API7KyKknWj9RVW", "6AKnxnqWWKDXyE8iMM2zffgZYv6Qfm09DuPRVFcHZybB")

	request := &pb.TrackCompositeEgressRequest{
		RoomName:     "nov",
		AudioTrackId: "TR_AMiJwt5QM43rCr",
		VideoTrackId: "TR_VCDZJkdoQffgMY",
		Output: &pb.TrackCompositeEgressRequest_FileAndStream{
			FileAndStream: &pb.FileAndStreamOutput{
				FileType: pb.EncodedFileType_MP4,
				Filepath: "again.mp4",
				Urls:     []string{"rtmp://a.rtmp.youtube.com/live2/52ah-1jja-k4mb-7716-8hhp"},
			},
		},
	}

	ctx := context.Background()

	fmt.Println("Starting Egress")
	info, err := egressclient.StartTrackCompositeEgress(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	egressID := info.EgressId

	time.Sleep(50 * time.Second)

	fmt.Println("Stopping Egress")
	info, err = egressclient.StopEgress(ctx, &pb.StopEgressRequest{
		EgressId: egressID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Info: ", info)
}
