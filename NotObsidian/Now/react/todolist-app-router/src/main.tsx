import React from 'react'
import ReactDOM from 'react-dom/client'
import 'bootstrap/dist/css/bootstrap.css'
import './index.css'

import App from './App.tsx'
// import AppContainer from './AppContainer'

import Appstore from './redux/Appstore.ts'
import { Provider } from 'react-redux'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Provider store={Appstore}>
      <App />
    </Provider>
  </React.StrictMode>,
)
