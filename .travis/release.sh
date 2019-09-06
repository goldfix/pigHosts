#!/usr/bin/env bash

set -euxo pipefail

if [ ! "$TRAVIS_PULL_REQUEST" = "false"  ]
then
	echo "TRAVIS_PULL_REQUEST - $TRAVIS_PULL_REQUEST"
	exit 0
fi

if [ -n "$TRAVIS_TAG" ]
then
	echo "TRAVIS_TAG - $TRAVIS_TAG"
	task package
	exit 0
fi

if [ -z "$TRAVIS_TAG" ]
then
	echo "TRAVIS_TAG - EMPTY"
	task package-snap
	exit 0
fi

