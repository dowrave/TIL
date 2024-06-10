import { all, fork, takeEvery, call, put} from 'redux-saga/effects'
import TimeActionCreator, { TIME_ACTION } from '../redux/TimeActionCreator'

// 비동기 처리 함수 : 2초 후 현재 시각을 응답함
const changeTimeApi = () => {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve({ currentTime: new Date()});
        }, 2000)
    })
}

// 비동기 처리 진행하는 작업자 사가
function* changeTime() {
    try{
        // Promise 완료까지 changeTimeApi를 블록되도록 fork한다.
        // 완료 시 response로 리턴
        const response: {currentTime: Date} = yield call(changeTimeApi);

        // 리턴받은 response를 put() 해서 액션을 디스패치한다(전달한다)
        yield put(TimeActionCreator.changeTimeCompleted(response.currentTime));

    } catch (e) {
        console.error(e);
        // 실패 : 실패 시의 액션을 전달함
        yield put(TimeActionCreator.changeTimeFailed());
    }
}

// CHANGE_TIME_REQUEST 액션을 감시한다. 액션 포착 시 changeTime 작업자 사가를 시작한다.
function* watcher_changeTime() {
    yield takeEvery(TIME_ACTION.CHANGE_TIME_REQUEST, changeTime);
}

// watcher.changeTime() 을 이용하는 새로운 사가 작업을 시작한다
export default function* timeSaga() {
    yield all([fork(watcher_changeTime)]);
}