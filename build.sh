#构建脚本
#!/bin/sh

function check_code() {
	EXCODE=$?
	if [ "$EXCODE" != "0" ]; then
		echo "build fail."
		exit $EXCODE
	fi
}

out="bin"
echo "build file to ./$out"

mkdir -p "$out/config"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./$out/todo_app  -tags v1 -v main.go

sources=`find ./conf -name "*.json"`
check_code
for source in $sources;do
	yes | echo $source|sed "s/.*\/\(.*\.json\).*/cp -f & .\/$out\/config\/\1/"|bash
	check_code
done

echo "build success."