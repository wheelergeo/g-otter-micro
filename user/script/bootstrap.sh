#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=user
echo "$CURDIR/bin/$BinaryName"
$CURDIR/bin/$BinaryName $BinaryName
