#!/usr/bin/bash

VERSION=2.7.1
ARCH=loong64
DIR="influxdb2_linux_${ARCH}"
mkdir "${DIR}"
cp -v README.md "${DIR}/"
cp -v LICENSE "${DIR}/"
cp -v bin/linux/influxd "${DIR}"
tar czvf "influxdb2-${VERSION}-linux-${ARCH}.tar.gz" "$DIR"
