# 동재의 P2P 통신 구현 일기

webSocket, channel 사용해보기
race condition(동기화 안되는 문제), mutexes 해결
마스터가 없는 탈중앙 네트워크 가능..?

```
go run -race main.go -mode=rest -port=3000
go run -race main.go -mode=rest -port=4000
```

## Connect

host 입장에서 3000 port 노드와 연결
[POST] {host}/peers

body

```
{
    "address": "127.0.0.1",
    "port": "3000"
}
```

## Connect Check

[GET] {host}/peers
host 노드와 연결된 IP들 확인
