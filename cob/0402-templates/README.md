# text/template과 html/template을 활용한 작업

- Go 템플릿은 템플릿 중첩, 함수 가져오기, 변수 표현, 데이터의 반복문 처리 등과 같은 작업을 간단히 처리할 수 있다.
- CSV writer보다 더 정교한 기능이 필요하다면 템플릿이 훌륭한 솔루션이 될 수 있다.
- 서버측 데이터를 클라이언트에 렌더링해야 할 때는 템플릿이 적합하다.

## 예제 분석

- Go는 text/template과 html/template이라는 두 개의 템플릿 패키지를 제공한다.
- 일반적으로 웹사이트를 렌더링할 때는 html/template를 사용하고(자바스크립트 주입과 같은 보안 침해를 방지) 그 외에는 text/template를 사용하면 된다.
- 템플릿 패키지는 최신 템플릿 라이브러리에서 기대할 수 있는 기능을 제공한다. HTML과 자바스크립트로 결과를 내보낼 때 템플릿 결합과 애플리케이션 로직의 추가가 쉬우며 안정성을 보장한다.