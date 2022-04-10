# go에서 환경변수 조작.

os 패키지를 사용해서 쉽게 환경변수를 작업할 수 있다.

## 예제 분석

- LoadConfig 함수 : path 상에 있는 데이터를 유연히 가져올 수 있는 포멧을 보여줌.
- 🙋‍♂️ 왜 유연한가?? 이과정에서 오버헤드가 많이 발생하거나 추장 종속성을 요구하지 않기 떄문.