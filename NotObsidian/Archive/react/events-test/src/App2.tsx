import {ChangeEvent, useState} from 'react'


const App2 = () => {
    const [x, setX] = useState<number>(0);
    const [y, setY] = useState<number>(0);

    // 이벤트 추가
    const changeValue = (e: ChangeEvent<HTMLInputElement>) => {
        let newValue: number = parseInt(e.target.value);
        if (isNaN(newValue)) newValue = 0;
        if (e.target.id === 'x') setX(newValue);
        else setY(newValue);
    }

  return (
    <div>
        <h3>제어 컴포넌트</h3>
        X : <input type="text" value={x} onChange = {changeValue}/>
        <br />
        Y : <input type="text" value={y} onChange = {changeValue}/>
        결과 : <span>{x + y}</span>
    </div>
  );
};

export default App2;