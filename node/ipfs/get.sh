hash=$1

echo $hash

ipfs get $hash
ipfs pin add $hash

mv $hash ./src/$hash.js