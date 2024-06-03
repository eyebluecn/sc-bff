#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName="smart.classroom.bff"
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}