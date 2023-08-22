package serialize_test

import (
	pb "github.com/DitoAdriel99/grpc/computer/proto"
	"github.com/DitoAdriel99/grpc/computer/sample"
	"github.com/DitoAdriel99/grpc/computer/serialize"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../../tmp/laptop.bin"
	jsonFile := "../../tmp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serialize.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serialize.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serialize.WriteProtobufToJsonFile(laptop1, jsonFile)
	require.NoError(t, err)

}
