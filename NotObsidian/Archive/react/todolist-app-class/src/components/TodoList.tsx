import React, { Component, PureComponent } from "react";
import { TodoListItemType } from "../AppContainer";
import TodoListItem from "./TodoListItem";

type Props = {
  todoList: Array<TodoListItemType>;
  toggleDone: Function;
  deleteTodo: Function;
};

export default class TodoList extends PureComponent<Props> {
  // shouldComponentUpdate(nextProps:Props, nextState: {}) {
  //   // todoList의 속성이 변경되었을 때 true를 반환한다.
  //   if (nextProps.todoList !== this.props.todoList) return true;
  //   return false;
  // }

  render() {
    console.log("## TodoList 렌더")
    let items = this.props.todoList.map((item: TodoListItemType) => {
      return <TodoListItem key={item.no} todoItem={item} deleteTodo={this.props.deleteTodo} toggleDone={this.props.toggleDone} />;
    });

    return (
      <div className="row">
        {" "}
        <div className="col">
          <ul className="list-group">{items}</ul>
        </div>
      </div>
    );
  }
}
