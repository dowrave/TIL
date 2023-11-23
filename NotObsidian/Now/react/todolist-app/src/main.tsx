import React from 'react'
import ReactDOM from 'react-dom/client'
import AppContainer from './AppContainer'
import 'bootstrap/dist/css/bootstrap.css'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <AppContainer />    {/* 요기 변경 */}
  </React.StrictMode>
)
