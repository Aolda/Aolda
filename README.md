# Aolda_dev-kad

## progress

- [X] 1:1 ws network
- [X] 1:1 msg read / write
- [X] 1:n multinode p2p network
- [X] 1:1 db synchronization
- [ ] Public IP base connection
- [ ] DHT(Kademlia base)
- [ ] Distributed DB(Swarm base)
- [ ] k8s operation

## usage

```
go run . -port={portNum}
```

### [POST] {host}/peers

req.body
```
{
    "Address": "localhost",
    "Port": "3000"
}
```

해당 node로 연결 요청
- call AddPeer()

### [POST] {host}/dbsync

req.body
```
{
    "Address": "localhost",
    "Port": "3000"
    "FileName":"add.js,div.js"
}
```

- 원하는 파일을 전송
- 만약 똑같은 이름의 파일이 있다면 덮어쓰기


### [GET] {host}/peers

res.body
```
[
    "localhost:3001",
    "localhost:3002",
    "localhost:4000"
]
```

host node와 연결된 nodes 확인
- call AllPeers()

### [GET] {host}/dbsync
```
[
    {
        "modified": "2023-04-13 15:50:45.811246588 +0900 KST",
        "name": "add.js",
        "size": 101
    },
    {
        "modified": "2023-04-13 15:50:45.811620125 +0900 KST",
        "name": "div.js",
        "size": 100
    },
    {
        "modified": "2023-04-07 13:34:51.871255162 +0900 KST",
        "name": "math.js",
        "size": 421
    },
    {
        "modified": "2023-04-07 13:34:51.871394655 +0900 KST",
        "name": "mod.js",
        "size": 100
    },
    {
        "modified": "2023-04-07 13:34:51.871517316 +0900 KST",
        "name": "mul.js",
        "size": 100
    },
    {
        "modified": "2023-04-07 13:34:51.871648601 +0900 KST",
        "name": "sub.js",
        "size": 101
    }
]
```
- src에 있는 모든 파일 리스트를 return


## function description

```
func CliStart()
```
- port 번호 입력
- call RestStart()

```
func RestStart()
```
- set router
- listening /peers
- listening /ws

```
func peersAPI()
```
- called by /peers
- in POST, call AddPeer()
- in GET, call AllPeers()

```
func AddPeer()
```
- make websocket connection
- call initPeer()

```
func AllPeers()
```
- return Peers

```
func initPeer()
```
- return connected Peer
- save connected Peer in Peers
- call readListener()
- call writeListener()
- call write()

```
func writeListener()
```
- write the message for Peers data to All nodes

```
func readListener()
```
- convert msg to JSON/struct Peer
- call AddPeer() for new Peer data

```
func write()
```
- return Peers data to go channel(writeListener)



## commit message guidline

```
<type>(<scope>): <subject>
```

### **\<type>**

- feat : 새로운 기능 추가, 기존의 기능을 요구 사항에 맞추어 수정
- fix : 기능에 대한 버그 수정
- build : 빌드 관련 수정
- chore : 패키지 매니저 수정, 그 외 기타 수정 ex) .gitignore
- ci : CI 관련 설정 수정
- docs : 문서(주석) 수정
- style : 코드 스타일, 포맷팅에 대한 수정
- refactor : 기능의 변화가 아닌 코드 리팩터링 ex) 변수 이름 변경
- test : 테스트 코드 추가/수정
- release : 버전 릴리즈

### **\<scope>**

생략 가능, dir name

### **\<subject>**

- Limit to 50 characters
- Start with capital letters
- Don't put a period at the end
- Be used as an imperative and does not use the past tense
- Explain something and why rather than how

## go-ethereum

```
solc --optimize --abi ./contracts/AoldaClient.sol -o build
solc --optimize --bin ./contracts/AoldaClient.sol -o build
```

- ./contracts/MySmartContract.sol 를 로컬 환경에 맞게 고치셈

```
abigen --abi=./build/AoldaClient.abi --bin=./build/AoldaClient.bin --pkg=aoldaClient --out=./build/AoldaClient.go
```

- 위와 동일 해당 bin파일과 abi를 기반으로 .go 파일 생성
- https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9 <-----여기 참고>

# 테스트해보기

## 1. 컨트랙트 배포하기

.env.sample 참고해서 .env 작성

```
cd contract
npx hardhat run --network ganache scripts/deploy.ts
```

## 2. 노드 실행하기

```
cd node
go run main.go
```

## 3. 컨트랙트 함수호출하기

```
cd sender
node call.sample.js
```

## 4. body

```json
{
  "eventName": "",
  "payload": {}
}
```

