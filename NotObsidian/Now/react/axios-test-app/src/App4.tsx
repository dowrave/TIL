import React from 'react'
import axios from 'axios'


const requestAPI = async () => {
    const url = "/api/todolist/gdhong";
    const response = await axios.get(url);
    console.log("# 응답 객체 :", response);
}

requestAPI();
type Props = {};

const App4 = (props: Props) => {
    return <h2>console Log를 확인합니다.</h2>
}


export default App4