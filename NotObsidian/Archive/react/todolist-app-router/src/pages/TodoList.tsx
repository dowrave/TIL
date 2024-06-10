import { Link } from 'react-router-dom'
import TodoItem from './TodoItem'
import TodoActionCreator from '../redux/TodoActionCreator';
import { AnyAction, Dispatch } from 'redux'
import { useDispatch, useSelector } from 'react-redux';
import { TodoStatesType, TodoItemType } from '../redux/TodoReducer'
import RootStatesType from '../redux/AppStore'
// import { CallbacksType, StatesType } from '../AppContainer'

// type PropsType = { states: StatesType, callbacks: CallbacksType };
type PropsType = { todoList: Array<TodoItemType>;
                deleteTodo: (id: number) => void;
                toggleDone: (id: number) => void;
}

const TodoList = ({ todoList, deleteTodo, toggleDone }: PropsType) => {
    let todoItems = todoList.map((item) => {
        // return <TodoItem key={item.id} todoItem={item} callbacks = {callbacks} />;
        return <TodoItem key={item.id} todoItem={item} deleteTodo={deleteTodo} toggleDone={toggleDone} />
    });

    return (
        <>
        <div className='row'>
            <div className="col p-3">
                <Link className="btn btn-primary" to="/todos/add">
                    할 일 추가
                </Link>
                <button className="btn btn-primary ms-1" onClick={() => callbacks.fetchTodoList()}>
                    할 일 목록 새로고침
                </button>
            </div>
        </div>
        <div className="row">
            <div className="col">
                <ul className="list-group">{todoItems}</ul>
            </div>
        </div>
        </>
    )
}

// const mapStateToProps = (state : RootStatesType) => ( {
//     todoList: state.todos.todoList,
// })

// const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => ({
//     deleteTodo: (id: number) => dispatch(TodoActionCreator.deleteTodo({ id })),
//     toggleDone: (id: number) => dispatch(TodoActionCreator.toggleDone({ id })),
// })

// 훅을 이용한 컨테이너 컴포넌트 생성
const TodoListContainer = () => {
    const dispatch = useDispatch();
    const todoList = useSelector((state: RootStatesType) => state.todos.todoList);
    const deleteTodo = (id: number) => dispatch(TodoActionCreator.deleteTodo({ id }));
    const toggleDone = (id: number) => dispatch(TodoActionCreator.toggleDone({ id }));

    return <TodoList todoList={todoList} deleteTodo ={deleteTodo} toggleDone = {toggleDone} />;
};

// export default TodoList;
// export default connect(mapStateToProps, mapDispatchToProps)(TodoList);
export default TodoListContainer;