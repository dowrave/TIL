import InputTodo from './InputTodo'
import TodoList from './TodoList'

// 속성은 전부 위에서 정의되었던 것들임
// Context API 도입하면서 주석 처리
// type AppProps = {
//     todoList: Array<TodoListItemType>;
//     addTodo: (todo: string) => void;
//     toggleDone: (no: number) => void;
//     deleteTodo: (no: number) => void;
// }

const App = () => { // Context API 주석 처리 -> props도 필요 없음
  return (
    <div className="container">
        <div className="card card-body bg-light">
            <div className="title">:: Todolist App</div>
        </div>
        <div className="card card-default card-borderless">
            <div className="card-body">
                {/* <InputTodo addTodo={props.addTodo} />
                <TodoList todoList = {props.todoList}
                    toggleDone={props.toggleDone} deleteTodo={props.deleteTodo}/> */}
                {/* Context API 도입 후 이렇게 변경 */}
                <InputTodo />
                <TodoList />
            </div>
        </div>
    </div>
  )
}

export default App;