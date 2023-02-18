package device

import (
	"fmt"
	"log"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestCreateWifiStats(t *testing.T) {
	var stats *SSIDStats = &SSIDStats{
		BytesSent:       100,
		BytesReceived:   1000,
		PacketsSent:     11000,
		PacketsReceived: 18000,
	}
	fmt.Println(stats)
	bytes, err := proto.Marshal(stats)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes)

	var statsFromWire *SSIDStats = &SSIDStats{}
	err = proto.Unmarshal(bytes, statsFromWire)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(statsFromWire)
}
