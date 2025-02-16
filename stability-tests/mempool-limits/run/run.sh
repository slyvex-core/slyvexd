#!/bin/bash

APPDIR=/tmp/slyvexd-temp
SLYVEXD_RPC_PORT=29587

rm -rf "${APPDIR}"

slyvexd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${SLYVEXD_RPC_PORT}" --profile=6061 &
SLYVEXD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${SLYVEXD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $SLYVEXD_PID

wait $SLYVEXD_PID
SLYVEXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "slyvexd exit code: $SLYVEXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SLYVEXD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
