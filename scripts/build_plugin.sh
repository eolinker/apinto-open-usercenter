#!/bin/sh
# ===========================================================================
# File: build.sh
# Description: usage: ./build.sh [outdir]
# ===========================================================================
# exit when any command fails

set -e
PLUGIN_NAME="user-open.apinto.com"
cd "$(dirname "$0")/../"
. ./scripts/init.sh

OUTPUT_DIR=$1
BUILD_MODE=$2
VERSION=$3
ARCH=$4
if [ "$ARCH" == "" ];
then
  ARCH="amd64"
fi
echo "build user plugin for "${ARCH}

OUTPUT_BINARY=$OUTPUT_DIR/${PLUGIN_NAME}

GO_VERSION=`go version | { read _ _ v _; echo ${v#go}; }`

if [ "$(version ${GO_VERSION})" -lt "$(version 1.18)" ];
then
   echo "${RED}Precheck failed.${NC} Require go version >= 1.19. Current version ${GO_VERSION}."; exit 1;
fi


# Step 1 - Build the frontend release version into the backend/server/dist folder
# Step 2 - Build the monolithic app by building backend release version together with the backend/server/dist (leveraing embed introduced in Golang 1.19).
echo "Start building apinto user monolithic ${VERSION}..."

echo ""
echo "Step 1 - building apinto user frontend..."

if [ "$BUILD_MODE" = "all" ] || [ ! -d "frontend/dist" ];then
  echo "begin frontend building..."
  if ! command -v yarn > /dev/null
  then 
    npm install yarn -g
  
  echo "cd frontend && yarn install --registry https://registry.npmmirror.com --legacy-peer-deps "
  cd frontend && yarn install --registry https://registry.npmmirror.com --legacy-peer-deps
  echo "yarn build"
  yarn build
  cd ../
  else
      npm --prefix ./frontend run build
  fi
else
  echo "skip frontend building..."
fi

echo "Completed building apinto user frontend."

echo "${VERSION}"
echo "Step 2 - building apinto user backend..."
# -ldflags="-w -s" means omit DWARF symbol table and the symbol table and debug information
go mod tidy
echo "${PWD}" "GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build --tags "mysql"   -o ${OUTPUT_BINARY} ./app/user-center"
echo `GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build --tags "mysql"   -o ${OUTPUT_BINARY} ./app/user-center`
