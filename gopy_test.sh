cd $HOME/Codes/playground/go/gopy

playground_folder="gocloud-hy9be"
readonly playground_folder

rm -rf playground_folder

mkdir playground_folder
cd playground_folder

go mod init hy9be/gocloud
go get github.com/hy9be/go-cloud
gopy pkg -output=out -vm=python3 github.com/hy9be/go-cloud > $HOME/Downloads/gopy-gocloud-$(date +%s%3N).log