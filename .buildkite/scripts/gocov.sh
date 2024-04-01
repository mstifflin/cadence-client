#!/bin/sh

set -ex

echo "ls -la /usr/bin/"
ls -la /usr/bin/
whoami

buildkite-agent artifact download ".build/*/coverage/unit_test_cover.out" . --step ":golang: unit-test" --build "$BUILDKITE_BUILD_ID"
buildkite-agent artifact download ".build/*/coverage/integ_test_sticky_off_cover.out" . --step ":golang: integration-test-sticky-off" --build "$BUILDKITE_BUILD_ID"
buildkite-agent artifact download ".build/*/coverage/integ_test_sticky_on_cover.out" . --step ":golang: integration-test-sticky-on" --build "$BUILDKITE_BUILD_ID"
buildkite-agent artifact download ".build/*/coverage/integ_test_grpc_cover.out" . --step ":golang: integration-test-grpc-adapter" --build "$BUILDKITE_BUILD_ID"

echo "download complete"

# report coverage
make cover_ci

# cleanup
rm -rf .build
