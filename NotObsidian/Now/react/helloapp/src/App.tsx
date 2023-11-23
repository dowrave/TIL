import React from 'react'
import CountryList from './CountryList'
import { useState } from 'react'
import styles from './styles'
import AppCssModule from './App.module.css'
import Footer from './Footer'
import {
  BasicButton, ItalicButton, UnderLineButton, WhiteUnderlineButton
} from './Buttons'

export type CountryType = {
  no : number;
  country: string;
  visited: boolean;
};



const App = () => {
  const [msg, setMsg] = useState<string>("World");
  const [theme, setTheme] = useState<string>("basic");
  const [list, setList] = useState<Array<CountryType>>([
    { no: 1, country: "이집트", visited: false},
    { no: 2, country: "일본", visited: true},
    { no: 3, country: "피지", visited: false},
    { no: 4, country: "콜롬비아", visited: false},
  ])
  // let list: Array<CountryType> = [
    // { no: 1, country: "이집트", visited: false},
    // { no: 2, country: "일본", visited: true},
    // { no: 3, country: "피지", visited: false},
    // { no: 4, country: "콜롬비아", visited: false},

  // ]

  // let msg = "<i>World</i>";
  // let msg = (<i>World</i>);
  const addResult = (x: number, y: number) => {
    return (
      <div className="card card-body bg-light mb-3">
        {x} + {y} = {x + y} 
      </div>
    );
  };

  return (
    <div className="container">
      {/* <h2>Hello <span dangerouslySetInnerHTML = {{__html:msg}} />!</h2> */}
      <h2 className = {AppCssModule.test}>Hello {msg}!</h2>
      <hr style={styles.dashStyle}/>
      {addResult(4, 3)}

      {/* countries는 CountryList.tsx에서 정의되었음! */}
      <CountryList countries={list}/>
      <BasicButton>기본</BasicButton>
      <ItalicButton>이탤릭</ItalicButton>
      <UnderLineButton>언더라인</UnderLineButton>
      <WhiteUnderlineButton>화이트 언더라인</WhiteUnderlineButton>
      <Footer theme = {theme} />
    </div>
  )
}

export default App