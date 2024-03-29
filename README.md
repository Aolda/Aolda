# Aolda-FlexiContract

## Aolda Start

```
cd node
go run main.go
```
만약 bind 문제가 생긴다면 main 함수에서 아래에 해당하는 api와 관련된 모든 부분을 주석처리하고 실행해주세요.

```
go func() {
		defer wg.Done()
		api.Listening()
	}() 
```
## API
https://aolda.kro.kr/
에 들어가시면 해당 서비스에 대해서 사용하실 수 있습니다.

예시로,

Call Aolda 탭에서

```
fileHash: add.js
functionName: add
args: 1,2
```
를 넣어주시고 실행하시면 됩니다.

## Add file
실행하기를 원하는 Js code를 넣고 file upload를 하시면 fileHash 값이 나옵니다.

해당 fileHash 값을 통해 아올다를 실행할 수 있습니다.


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

