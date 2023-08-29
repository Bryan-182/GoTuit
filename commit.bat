git add .
git commit -m "Ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
del main.zip
tar.exe -a -cf main.zip main