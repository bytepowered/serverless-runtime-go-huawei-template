GOOS=linux
GOARCH=amd64
# Binary name
BINARY=handler

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s"
BDFLAGS=CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH}

# Release
BUILD_DIR=./build
# Binary
OUTPUT="${BUILD_DIR}/${BINARY}"

# Builds the project
build:
		rm -rf ${BUILD_DIR}
		mkdir -p ${BUILD_DIR}

		# Build for linux
		${BDFLAGS} go build ${LDFLAGS} -a -installsuffix cgo -o ${OUTPUT} ./main.go

		# Write version
		(cd ${BUILD_DIR} && zip fss-cloud-openai-${GOOS}-${GOARCH}.zip handler)

install:
		go install

clean:
		go clean

.PHONY:  clean build