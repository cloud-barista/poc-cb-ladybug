# cb-fw-template
Template for a new Cloud-Barista framework

### 개요
`cb-fw-template` 는 개발자가 Cloud-Barista 프레임워크를 새로 만들 때 활용할 수 있는 소스코드 템플릿입니다.

### 사용법
[Repo 메인 페이지](https://github.com/cloud-barista/cb-fw-template)에 있는 "Use this template" 버튼을 누르고
본인의 계정 (예: `jihoon-seo`) 아래에 
새로운 repo (예: `jihoon-seo/cb-cat`) 를 만듭니다.

이후, 필요에 따라
`pkg/apiserver/apiserver.go` 에 REST API 엔드포인트를 추가하고
`pkg/common/resource.go` 에 REST API 처리 함수를 작성하고
적절한 위치에 코어 로직 코드를 작성하면 됩니다.

### 소스 트리 설명
- `cmd/cb-myfw`: 
  이 repo에서는 소스 컴파일 결과 `cb-myfw` 라는 바이너리 파일이 생성된다고 가정합니다. 
  이 repo에서 `cb-myfw` 라고 되어 있는 부분들을 당신이 원하는 이름 (예: `cb-cat`) 으로 바꾸시면 됩니다.
- `pkg/`: REST API 서버, 코어 로직 등 프레임워크의 주요 코드가 위치하는 디렉토리입니다.
- `test/`: REST API 호출을 통해 프레임워크를 테스트 해 볼 수 있는 셸 스크립트가 있습니다.
