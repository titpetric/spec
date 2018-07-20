#!/bin/bash
set -e
PROJECT=$(basename $(dirname $(readlink -f $0)))

NAMES=$(ls cmd/* -d | xargs -n1 basename)
for NAME in $NAMES; do
	OSES=${OSS:-"linux"}
	ARCHS=${ARCHS:-"amd64"}
	for ARCH in $ARCHS; do
		for OS in $OSES; do
			echo $OS $ARCH $NAME
			GOOS=${OS} GOARCH=${ARCH} CGO_ENABLED=0 GOARM=7 go build -o build/${NAME}-${OS}-${ARCH} cmd/${NAME}/*.go
			if [ $? -eq 0 ]; then
				echo " " OK
			fi
			if [ "$OS" == "windows" ]; then
				mv build/${NAME}-${OS}-${ARCH} build/${NAME}-${OS}-${ARCH}.exe
			fi
		done
	done
done

echo "Resulting files:"
find build -type f | xargs -n1 echo " "

if [ -d "/usr/local/bin" ]; then
	cp build/$PROJECT-linux-amd64 /usr/local/bin/$PROJECT
fi
if [ -d "/src/misc/bin" ]; then
	cp build/$PROJECT-linux-amd64 /src/misc/bin/$PROJECT
fi