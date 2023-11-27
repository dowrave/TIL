import { ChangeEvent, useState } from 'react' // useState 훅을 가져온다


const App = () => {
  const [msg, setMsg] = useState<string>(""); // useState 훅

  return (
    <div>
      <input type="text" value = {msg}
      onChange = {(e: ChangeEvent<HTMLInputElement>) => setMsg(e.target.value)} />
      <br />
      <span>입력 메시지 : {msg}</span>
    </div>
  )
}

export default App;