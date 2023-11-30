import axios from 'axios'

type TodoType = { id: number, todo: string; done: boolean; desc: string};

const listUrl = "/api/todolist_long/gdhong"
const todoUrlPrefix = "/api/todolist_long/gdhong/"

const requestAPI = () => {
  let todoList: Array<TodoType> = [];

  axios .get(listUrl)
    .then((response) => {
      todoList = response.data;
      console.log("# TodoList :", todoList);
      return todoList[0].id
    })
    .then((id) => {
      return axios.get(todoUrlPrefix + id);
    })
    .then((response) => {
      console.log("## 1번째 Todo : ", response.data);
      return todoList[1].id;
    })
    .then((id) => {
      axios.get(todoUrlPrefix + id).then((response) => {
        console.log("## 2번째 Todo : ", response.data);
      })
    })
};

requestAPI();

type Props = {};

const App = (props: Props) => {
  return <h2>Console.log를 확인하세요</h2>;
}

export default App;