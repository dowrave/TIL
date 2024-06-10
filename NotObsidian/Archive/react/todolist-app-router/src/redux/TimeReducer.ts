import { TimeActionType, TIME_ACTION } from "./TimeActionCreator";
// import { createReducer } from "@reduxjs/toolkit"
// import TimeActionCreator from './TimeActionCreator'

export type HomeStatesType = { currentTime : Date; isChanging: boolean; };

const initialState = {
    currentTime: new Date(),
    isChanging: false
}

const TimeReducer = (state = initialState, action: TimeActionType) => {
    switch(action.type) {
        case TIME_ACTION.CHANGE_TIME_REQUEST:
            return { ...state, isChanging: true };
        case TIME_ACTION.CHANGE_TIME_COMPLETED:
            return { ...state, currentTime: action.payload.currentTime, isChanging: false };
        case TIME_ACTION.CHANGE_TIME_FAILED:
            return { ...state, isChaning: false };
        default:
            return state; 
    }
}
// const TimeReducer = createReducer(initialState, (builder) => {
//     builder
//         .addCase(TimeActionCreator.asyncChangeTime.pending, (state, action) => {
//             state.isChanging = true;
//         })
//         .addCase(TimeActionCreator.asyncChangeTime.fulfilled, (state, action) => {
//             state.isChanging = false;
//         })
// })

export default TimeReducer;