#!/bin/bash
rm -rf /tmp/slyvexd-temp

slyvexd --devnet --appdir=/tmp/slyvexd-temp --profile=6061 &
SLYVEXD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:22611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SLYVEXD_PID

wait $SLYVEXD_PID
SLYVEXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "slyvexd exit code: $SLYVEXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SLYVEXD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
