import React, { useState, useEffect, ChangeEvent, useTransition } from 'react'
import logo from './logo.svg'

// 대량으로 만들 아이템 타입 정의
type ItemType = {
  id: number;
  keyword: string;
}

const App = () => {
  const [keyword, setKeyword] = useState<string>("");
  const [results, setResults] = useState<Array<ItemType>>([]);
  const [isPending, setIsPending] = useTransition();

  const handleChange = (e : ChangeEvent<HTMLInputElement>) => {
    setKeyword(e.target.value);
  };

  // 컴포넌트 마운트 시 results 상태에 빈 배열 생성
  // keyword 상태 변경 시 keyword에 입력이 되었다면 5000개의 아이템 생성
  // 대량 데이터 업데이트는 긴급하지 않다
  useEffect(() => {
    if (keyword.trim() === "") {
      setResults([]);
    } else {
      const items = Array.from(Array(5000), (item, index) => {
        return { id: index, keyword: keyword };
      });
      // setResults(items);
      startTransition(() => {
        setResults(items);
      })
    }
  }, [keyword]);

  // results 상태를 이용해 div 대량 생성
  const divRows = results.map((item) => (
    <div key={item.id}>
      id: {item.id}
      <br />
      keyword: {item.keyword}
      <br />
      <img src={logo} style={{ width: 100, height: 100}} />
    </div>
  ));

  // 사용자가 입력 필드에 타이핑하면 handleChange 함수를 실행해 상태 변경
  // onChange 이벤트에 의해 바뀐 값 렌더링은 긴급한 업데이트가 요구된다.
  return (
    <div style={{ margin: 10}}>
      <div className="SearchInput">
        Keyword : <input type="text" value={keyword} onChange={handleChange} />
      </div>
      <hr />
      {/* <div>{divRows}</div> */}
      <div>{
        // isPending 동안 true면 fallback UI를 렌더링
        isPending ? <h2>로딩 중입니다~~</h2> : divRows
}   </div>
    </div>
  )

}

export default App;