git add .
git commit -m "Ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build main.go
del main.zip
tar -a -cf main.zip main