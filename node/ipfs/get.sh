hash=$1

echo $hash

ipfs get $hash
ipfs pin add $hash

# 파일 이름에 이미 .js 확장자가 포함되어 있는지 확인 후 추가
if [[ $hash != *".js" ]]; then
  mv $hash ./src/$hash.js
fi

mv $hash ./src/