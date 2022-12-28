#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/modules/template_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/modules/template_protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done