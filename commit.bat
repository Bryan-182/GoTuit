git add .
git commit -m "Ultimo commit"
git push
$Env:GOOS = "linux"
go build main.go
del main.zip
tar -a -cf main.zip main