#!/usr/bin/env bash

pushd $GOPATH/src/github.com/hueyjj/fuse
    make buildmodule
    echo
    bin/sundermodule get
    echo
popd