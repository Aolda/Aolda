# Aolda_dev-p2p

## warning

```
sudo ipfs daemon &
```
ipfs daemon를 백그라운드에 실행하고 main.go를 실행

## progress

- [X] 1:1 ws network
- [X] 1:1 msg read / write
- [X] 1:n multinode p2p network
- [X] 1:1 db synchronization
- [X] Public IP base connection
- [X] DHT(Kademlia base)
- [X] Distributed DB(IPFS base)
- [ ] k8s operation

## usage
node/main.go 파일 실행

```
go run main.go
```

KDH를 통한 Peer 탐색
```
Searching for peers...
Failed connecting to  12D3KooWDiAYcsKoznuM7Srvxiyiev4E7AJZPskREytrdxtyTVSi , error: no addresses
Failed connecting to  12D3KooWDejdKJUyNndjhraWYjfhbGqUaezqHNruv39UYUBLVhNB , error: no addresses
Failed connecting to  12D3KooWDnVPRzHuxfeLwC3EcY8soq4gCtYz4h3xJfHYkBLKzVWt , error: no addresses
Failed connecting to  12D3KooWDzubhS7uBFAZrgxrqodcNiPauWD7tSmDUeHrza1WGgt7 , error: no addresses
Failed connecting to  12D3KooWAAySgPfySCtg6YtTDqbFh4twHg23JtvbuvA353Z5DfoM , error: no addresses
Failed connecting to  12D3
...
Connected to: 12D3KooWC1iTVpMFfKbdN8rkZBRCvRFT9BEn1nXPkS94jDjEDYrW
Failed connecting to  12D3KooWDXU1jJr8jVsRMwSodEySxTAfPZhcoU29kxfJ5bdWZJFY , error: no addresses
...
Failed connecting to  12D3KooWSnpXUJX6wbTwp4He7SG3ZVYer1UErqUgwQ5LjjrMNqGL , error: no addresses
Failed connecting to  12D3KooWNpFvRJYPYCYPirsaRRELNkGddWCxcB6Jv5FaQwx7gf8G , error: no addresses
Failed connecting to  12D3KooWPBHw7LUzWiFzHqbASVrAfStrZMdAdBtxbEXMBMS8equh , error: no addresses
Failed connecting to  12D3KooWQXUUzd4poaoButF3tRzYXwPhzXqhYZ9fHao4h3bQFvUJ , error: no addresses
Peer discovery complete
```
Connected to: 나 Peer discovery complete의 Log가 발생하면 연결이 완료된 것


특정 message를 보내서 행동을 할 수 있음

1. exec/[file name]/[function name]/[argv]
```
exec/add.js/add/1 2
```
2. upload/[filename]
```
upload/add.js
```
**해당 add.js는 src에 있는 파일을 대상으로 함**

3. get/[hash value]/[file name]
```
get/Qmb4vr9WeYJZvTS9drzD4UTVjzFW2nZLEPZAhCKMkBxaz1/mul1.js
```
**해당 hash 값은 add를 할 때, 나오는 값**


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
