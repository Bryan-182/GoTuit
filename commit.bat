git add .
git commit -m "Ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
GOARCH=amd64 GOOS=linux go build main.go -ldflags="-s -w"
del main.zip
tar.exe -a -cf main.zip main.exe