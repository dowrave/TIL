- Django에서 [[Django - CORS 관련 설정]]을 했음에도, 프론트엔드에서 요청을 보내도 계속 에러가 발생함
```
- Django에서 CORS 설정 했을 때
Request header field access-control-allow-origin is not allowed by Access-Control-Allow-Headers in preflight response

- Django에서 설정을 껐을 때
Response to preflight request doesn't pass access control check: Redirect is not allowed for a preflight request.
```

- 어떻게 해결했는지 모르겠는데 해결됐다. 프론트엔드에서 axios.get에 넣는 링크에 엔터를 넣었는데, 그게 화근이 됐나?

