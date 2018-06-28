#!/usr/bin/env bash

TEST_DATA='{
    "command_name": "yt_download_music",
    "options": {
        "format": {
            "value": "m4a"
        }
    },
    "args": [
        "https://www.youtube.com/watch?v=7dGwk5-QMpc"
    ]
}'

echo
echo "Test data:"
echo $TEST_DATA
echo

pushd $GOPATH/src/github.com/hueyjj/fuse
    make buildmodule
    echo
    bin/sundermodule "$TEST_DATA"
    echo
popd