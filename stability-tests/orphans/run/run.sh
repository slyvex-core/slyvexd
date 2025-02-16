#!/bin/bash
rm -rf /tmp/slyvexd-temp

slyvexd --simnet --appdir=/tmp/slyvexd-temp --profile=6061 &
SLYVEXD_PID=$!

sleep 1

orphans --simnet -alocalhost:22511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $SLYVEXD_PID

wait $SLYVEXD_PID
SLYVEXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "slyvexd exit code: $SLYVEXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SLYVEXD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
