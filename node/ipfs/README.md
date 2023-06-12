# File upload CLI
```
sudo ipfs daemon &
```
daemon를 실행해서 ipfs network 참가


**sudo를 붙이지 않으면 파일 권한 관련 문제 발생**

# usage
```
aolda add -file [filename]
```
src 폴더 내부에 있는 file를 ipfs network로 올림

```
aolda get -filehash [filehash]
```
해당 filehash를 ipfs network로 부터 다운로드

파일이름은 hash 그대로 저장