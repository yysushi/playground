#!/bin/bash -e

HOGE=FUGA
export HOGE
printenv | grep FUGA > /dev/null 2>&1
