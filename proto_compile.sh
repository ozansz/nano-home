#!/bin/bash
set -e

rm -f rpc/rts-server/iot_realtime/*.pb.go

rm -f nodemcu-client/*.pb.*
rm -f nodemcu-client/pb_*.{h,c}

protoc -I/usr/local/include -Iproto/ --go_out=plugins=grpc:rpc/rts-grpc-server/iot_realtime proto/time_service.proto
protoc -I/usr/local/include -Iproto/ --go_out=plugins=grpc:rpc/rts-mux-server/iot_realtime proto/time_service.proto
#python3 -m grpc_tools.protoc -Iproto/ --python_out=nodemcu-client --grpc_python_out=nodemcu-client proto/*.proto
python3 nanopb/generator/nanopb_generator.py -Iproto/ -Dnodemcu-client/ proto/*.proto

cp nanopb/pb_decode.* nodemcu-client/
cp nanopb/pb_encode.* nodemcu-client/
cp nanopb/pb_common.* nodemcu-client/