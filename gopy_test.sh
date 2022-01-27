cd $HOME/Codes/playground/go/gopy

playground_folder="gocloud"
readonly playground_folder

rm -rf $playground_folder

mkdir $playground_folder
cd $playground_folder

go mod init hy9be/gocloud
rm -rf $HOME/go/pkg/mod/github.com/hy9be/gocloud@v0.24.0/
go get github.com/hy9be/gocloud
gopy pkg -output=out -vm=python3 -symbols=false github.com/hy9be/gocloud #> $HOME/Downloads/gopy-gocloud-$(date +%s%3N).log