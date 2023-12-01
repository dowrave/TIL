import { configureStore } from '@reduxjs/toolkit'
import { combineReducers, Middleware } from 'redux';
import TodoReducer, { TodoStatesType } from './TodoReducer'
import TimeReducer, { HomeStatesType } from './TimeReducer'
import createSagaMiddleware from 'redux-saga'
import rootSaga from '../sagas'

export type RootStatesType = {
    home: HomeStatesType;
    todos: TodoStatesType;
}

const sagaMiddleware = createSagaMiddleware();

const RootReducer = combineReducers({
    home: TimeReducer,
    todos: TodoReducer
})

// const mw1: Middleware = (store) => (next) =>(action) => {
//     console.log("### mw1 전")
//     next(action);
//     console.log("### mw1 후")
// }

// const mw2: Middleware = (store) => (next) =>(action) => {
//     console.log("### mw2 전")
//     next(action);
//     console.log("### mw2 후")
//     console.log(store.getState());
// }

const loggerMW: Middleware = (store) => (next) => (action) => {
    console.log("### action 실행 : ", action);
    // console.log("### action 변경 전 상태 : ", store.getState());
    next(action);
    // console.log("### action 변경 후 상태 : ", store.getState());
};

// const Appstore = configureStore({ reducer: RootReducer });
const AppStore = configureStore({
    reducer: RootReducer,
    // middleware: (getDefaultMiddleware) => {
    //     return getDefaultMiddleware({ serializableCheck: false}).concat(loggerMW);
    middleware: (getDefaultMiddleware) => {
        return getDefaultMiddleware({ serializableCheck : false })
            .concat(loggerMW).concat(sagaMiddleware)
        },
    devTools: process.env.NODE_ENV !== "production",
    }
)

// 사가 미들웨어를 등록한 뒤, rootSaga를 이용해 사가를 구동한다.
sagaMiddleware.run(rootSaga);

export default AppStore;