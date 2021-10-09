rm source.zip
rm main
GOOS=linux go build main.go
zip -r source.zip main
