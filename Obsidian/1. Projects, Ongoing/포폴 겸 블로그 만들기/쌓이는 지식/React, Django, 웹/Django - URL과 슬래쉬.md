1. URL 설정 시 끝에 `/`을 붙이는 건 자동으로 설정되어 있다.
	- `settings.py`의 `APPEND_SLASH = True`라는 듯.

2. 프론트엔드에서 요청을 보낼 떄, 끝에 `/`이 누락되었음에도 요청이 수락되는 경우가 있고, 오류가 발생하는 경우가 있다. 후자는 특히 Post 요청에서 발생하는데, **POST 데이터를 유지하며 리디렉션하는 것이 기술적으로 복잡**하기 때문이다.

3. 결국 기본적으로 알아야 하는 것은 아래 2가지이다.
	- **Django의 URL 끝 부분 마무리는 `/`으로 끝나며, Django에서 URL 설정을 할 때는 빼도 무방하다.**
	- 프론트엔드에서 백엔드의 URL로 접근할 때는, 반드시 `/`으로 마무리해야 한다
		- get 요청처럼 ?뒤에 파라미터 넣을 게 아닌 이상