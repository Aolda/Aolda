# Aolda_dev-p2p

## progress

- [X] 1:1 ws network
- [X] 1:1 msg read / write
- [X] 1:n multinode p2p network
- [ ] 1:1 db synchronization
- [ ] 1:n db synchronization
- [ ] 1:n msg read / write
- [ ] Kademlia implement

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
