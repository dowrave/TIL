import React from 'react'
import TodoList from './TodoList'
import UserInfo from './UserInfo'

const App = () => {
  return (
    <div style={{padding : 10}}>
      <React.Suspense fallback={<h2> UserInfo 로딩 중</h2>}>
          <UserInfo />
      </React.Suspense> 
      <hr />
      <React.Suspense fallback={<h2> TodoList 로딩 중</h2>}>
          <TodoList />
      </React.Suspense> 
    </div>
  )
}

export default App;