hash=$1
name=$2

echo $hash
echo $name

ipfs get $hash
ipfs pin add $hash

mv $hash $name
mv $name ./src/ #실행하는 주체인 node/main.go 위주로

