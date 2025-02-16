#!/bin/bash
slyvexdsanity --command-list-file ./commands-list --profile=7000
TEST_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ]; then
  echo "slyvexdsanity test: PASSED"
  exit 0
fi
echo "slyvexdsanity test: FAILED"
exit 1
