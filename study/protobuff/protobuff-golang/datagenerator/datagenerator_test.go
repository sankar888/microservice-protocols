package datagenerator

import (
	"fmt"
	"log"
	"os"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestDatagenerator(t *testing.T) {
	var gen DataGenerator = NewDeviceRecordGenerator()
	rec := gen.Get()
	bytes, err := proto.Marshal(rec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes)
}

func TestDatageneratorJson(t *testing.T) {
	var gen DataGenerator = NewDeviceRecordGenerator()
	rec := gen.Get()
	json := protojson.Format(rec)
	fmt.Println(json)
}

func BenchmarkDatagenerator(b *testing.B) {
	var gen DataGenerator = NewDeviceRecordGenerator()
	file, err := os.Create("/temp/devices.proto")
	if err != nil {
		log.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rec := gen.Get()
		bytes, err := proto.Marshal(rec)
		if err != nil {
			log.Fatal(err)
		}
		file.Write(bytes)
	}
	file.Sync()
	file.Close()
	fmt.Printf("No of Records in Proto file %d\n", gen.RecordCount())
	b.StopTimer()
}

func BenchmarkDatageneratorJson(b *testing.B) {
	var gen DataGenerator = NewDeviceRecordGenerator()
	file, err := os.Create("/temp/devices.proto.json")
	if err != nil {
		log.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rec := gen.Get()
		bytes, err := protojson.Marshal(rec)
		if err != nil {
			log.Fatal(err)
		}
		file.Write(bytes)
	}
	file.Sync()
	file.Close()
	fmt.Printf("No of Records in Json file %d\n", gen.RecordCount())
	b.StopTimer()
}
