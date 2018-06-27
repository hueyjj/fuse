#!/usr/bin/env bash

TEST_DATA='{
    "app_name": "yt_download_music",
    "data": {
        "format": {
            "value": "m4a"
        }
    }
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