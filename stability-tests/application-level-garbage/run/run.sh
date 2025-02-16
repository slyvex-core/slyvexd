#!/bin/bash
rm -rf /tmp/slyvexd-temp

slyvexd --devnet --appdir=/tmp/slyvexd-temp --profile=6061 --loglevel=debug &
SLYVEXD_PID=$!
SLYVEXD_KILLED=0
function killSlyvexdIfNotKilled() {
    if [ $SLYVEXD_KILLED -eq 0 ]; then
      kill $SLYVEXD_PID
    fi
}
trap "killSlyvexdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:22611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SLYVEXD_PID

wait $SLYVEXD_PID
SLYVEXD_KILLED=1
SLYVEXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "slyvexd exit code: $SLYVEXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SLYVEXD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
