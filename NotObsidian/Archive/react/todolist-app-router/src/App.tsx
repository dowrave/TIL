import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Layout from './components/Layout'
// import Loading from './components/Loading'

import Home from './pages/Home'
import About from './pages/About'
import TodoList from './pages/TodoList'
import AddTodo from './pages/AddTodo'
import EditTodo from './pages/EditTodo'
import NotFound from './pages/NotFound'
// import { CallbacksType, StatesType } from './AppContainer'

// type PropsType = {
//   states: StatesType,
//   callbacks: CallbacksType;
// }

const App = () => {
  return(
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />   
          <Route path="about" element={<About />} />
          <Route path="todos" element={<TodoList />} />
          <Route path="todos/add" element={<AddTodo />} />    
          <Route path="todos/edit/:id" element={<EditTodo />} />
          <Route path='*' element={<NotFound />}/>
        </Route>
      </Routes>
      {/* {states.isLoading ? <Loading /> : ""} */}
    </Router>
  )

}

export default App