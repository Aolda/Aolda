filename=$1

echo $filename
output=$(sudo ipfs add ../src/$filename)

# 명령의 반환값 확인
if [ $? -ne 0 ]; then
  echo "Error: Failed to add file to IPFS"
  exit 1
fi

# 출력에서 두 번째 필드 추출
echo "filename: $filename"
hash=$(echo $output | awk '{print $2}')

ipfs get $hash
chmod 644 $hash
sudo mv $hash $filename
sudo mv $filename ../src/

echo "Success upload the file!"
echo "$filename is replace about $hash"