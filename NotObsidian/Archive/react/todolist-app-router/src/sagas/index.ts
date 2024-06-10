import { all, fork } from 'redux-saga/effects'
import timeSaga from './timeSaga'

// timeSaga 외에 다른 사가가 있다면 모두 모아 병렬로 fork 시킨다.
export default function* rootSaga() {
    yield all([fork(timeSaga)]);
}