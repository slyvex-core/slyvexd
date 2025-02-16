#!/bin/bash
rm -rf /tmp/slyvexd-temp

NUM_CLIENTS=128
slyvexd --devnet --appdir=/tmp/slyvexd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SLYVEXD_PID=$!
SLYVEXD_KILLED=0
function killSlyvexdIfNotKilled() {
  if [ $SLYVEXD_KILLED -eq 0 ]; then
    kill $SLYVEXD_PID
  fi
}
trap "killSlyvexdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $SLYVEXD_PID

wait $SLYVEXD_PID
SLYVEXD_EXIT_CODE=$?
SLYVEXD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "slyvexd exit code: $SLYVEXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SLYVEXD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
