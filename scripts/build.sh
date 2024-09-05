#!/bin/sh
# ===========================================================================
# File: build.sh
# Description: usage: ./build.sh [outdir]
# ===========================================================================
# exit when any command fails
set -e
cd "$(dirname "$0")/../"
. ./scripts/init.sh

OUTPUT_DIR=$(mkdir_output "$1")

if [ "$3" != "" ];
then
  VERSION=$3
fi

./scripts/build_plugin.sh ${OUTPUT_DIR} $2 ${VERSION}

PLUGIN_NAME="user-open.apinto.com"
OUTPUT_BINARY=$OUTPUT_DIR/${PLUGIN_NAME}

mkdir -p user_${VERSION}
mv ${OUTPUT_BINARY} ./user_${VERSION}
#cp ./plugin/plugin-user.zip ./user_${VERSION}
zip -rj ./user_${VERSION}/plugin-user.zip ./plugin/*
echo "Completed building apinto user backend."

echo ""
echo "Step 3 - printing version..."

tar -czvf user_${VERSION}_linux_amd64.tar.gz user_${VERSION}

rm -rf user_${VERSION}

cp user_${VERSION}_linux_amd64.tar.gz ${OUTPUT_DIR}

rm -rf user_${VERSION}_linux_amd64.tar.gz

echo "user_${VERSION}_linux_amd64.tar.gz 完成"

echo ""
echo "${GREEN}Completed building apinto user monolithic ${VERSION} at ${OUTPUT_BINARY} ${NC}"

