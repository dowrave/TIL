import { useEffect, useState } from 'react'
import App from './App'
import { produce } from 'immer'
import axios from 'axios'

export type TodoItemType = { id: number; todo: string; desc: string; done: boolean};
export type StatesType = { todoList: Array<TodoItemType>, isLoading: boolean }
export type CallbacksType = {
    fetchTodoList: () => void;
    addTodo: (todo: string, desc: string) => void;
    deleteTodo: (id: number, desc: string) => void;
    toggleDone: (id: number) => void;
    updateTodo: (id: number, todo: string, desc: string, done: boolean) => void;
}

// 사용자 데이터
const USER = 'gdhong';
// const BASEURI = '/api/todolist/' + USER;

const BASEURI = "api/todolist_long/" + USER; // 지연 시간 추가


const AppContainer = (props: Props) => {

    let [todoList, setTodoList] = useState<Array<TodoItemType>>([]);
    let [isLoading, setIsLoading] = useState<boolean>(false);

    useEffect(() => {
        fetchTodoList();
    }, [])
    
    // 할 일 목록 조회 기능 제공
    const fetchTodoList = async () => {
        setTodoList([]);
        setIsLoading(true);
        try {
            const response = await axios.get(BASEURI);
            setTodoList(response.data);
        } catch (e) {
            if (e instanceof Error) alert("조회 실패 : ", e.message);
            else alert("조회 실패 : " + e);
        }
        setIsLoading(false);
    }

    // 할 일 추가 기능 제공
    // 성공시 마지막 인자로 전달한 callback을 호출한다.
    const addTodo = async (todo: string, desc: string, callback: () => void) => {
        setIsLoading(true);
        try {
            const response = await axios.post(BASEURI, { todo, desc });
            if (response.data.status == "success") {
                // 1건의 할 일 추가 성공시 전체 할 일을 조회하지 않고, 추가된 1건만 state에 추가한다.
                let newTodoList = produce(todoList, (draft) => {
                    draft.push({ ...response.data.item, done: false })
                });
                setTodoList(newTodoList);
                callback();
            } else {
                alert("할 일 추가 실패 :" + response.data.message);
            }
        } catch (e) {
            if (e instanceof Error) alert("할 일 추가 실패:" + e.message);
            else alert("할 일 추가 실패:" + e);
        }
        setIsLoading(false);

    }

    // 할 일 1건을 삭제하는 기능 제공
    const deleteTodo = async (id: number) => {
        setIsLoading(true);

        try {
            const response = await axios.delete(`${BASEURI}/${id}`);
            if (response.data.status === "success") {
                let index = todoList.findIndex((todo) => todo.id === id);
                let newTodoList = produce(todoList, (draft) => {
                    draft.splice(index, 1);
                });
                setTodoList(newTodoList);
            } else {
                alert("할 일 삭제 실패 : " + response.data.message);
            }
        } catch (e) {
            if (e instanceof Error) alert("할 일 삭제 실패" + response.data.message)
            else alert("할 일 삭제 실패:" + e);
        }
        setIsLoading(false);

    }

    // 완료 여부를 토글하는 기능 제공
    const toggleDone = async (id: number) => {
        setIsLoading(true);
        
        try {
            let todoItem = todoList.find((todo) => todo.id === id);
            const response = await axios.put(`${BASEURI}/${id}`, { ...todoItem, done: !todoItem?.done });
            if (response.data.status === "success") {
                let index = todoList.findIndex((todo) => todo.id === id);
                let newTodoList = produce(todoList, (draft) => {
                    draft[index].done = !draft[index].done;
                });
                setTodoList(newTodoList);
            } else {
                alert("완료 토글 실패 : " + response.data.message);
            }
        } catch (e) {
            if (e instanceof Error) alert("완료 토글 실패 : " + e.message);
            else alert("완료 토글 실패 : " + e);
        }
        setIsLoading(false);

    }

    // 할 일 수정 기능을 제공한다.
    // 성공 시 마지막 인자로 전달된 callback을 호출한다.
    const updateTodo = async (id: number, todo: string, desc: string, done: boolean, callback: () => void) => {
        setIsLoading(true);

        try {
            const response = await axios.put(`${BASEURI}/${id}`, { todo, desc, done });
            if (response.data.status == "success") {
                let index = todoList.findIndex((todo) => todo.id === id);
                let newTodoList = produce(todoList, (draft) => {
                    draft[index] = {...draft[index], todo, desc, done };
            });
            setTodoList(newTodoList);
            callback();
        } else {
            alert("할 일 수정 실패 : " + response.data.message);
        }
    } catch (e) {
        if (e instanceof Error) alert("할 일 수정 실패 : " + e.message);
        else alert("할 일 수정 실패 : " + e);
    }
    setIsLoading(false);

}
    // 상태, 액션을 states, callbacks 객체로 묶어서 한꺼번에 속성을 전달한다.
    const callbacks: CallbacksType = { fetchTodoList, addTodo, updateTodo, deleteTodo, toggleDone }
    const states: StatesType = { todoList, isLoading };
    return <App callbacks={callbacks} states={states} />;
}

export default AppContainer;