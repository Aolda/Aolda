# Aolda

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
solc --optimize --abi ./contracts/MySmartContract.sol -o build
solc --optimize --bin ./contracts/MySmartContract.sol -o build
```
- ./contracts/MySmartContract.sol 를 로컬 환경에 맞게 고치셈
  
```
abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go
```
- 위와 동일 해당 bin파일과 abi를 기반으로 .go 파일 생성
- https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9 <-----여기 참고>