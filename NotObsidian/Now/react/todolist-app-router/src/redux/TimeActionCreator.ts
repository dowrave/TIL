// import { AnyAction } from "redux";
// import { ThunkDispatch } from "redux-thunk"
// import { createAction, createAsyncThunk } from '@reduxjs/toolkit'

export const TIME_ACTION = {
    CHANGE_TIME_REQUEST: "changeTimeRequest" as const,
    CHANGE_TIME_COMPLETED: "changeTimeCompleted" as const,
    CHANGE_TIME_FAILED: "changeTimeFailed" as const,
}

// export type TimeActionType = ReturnType<typeof TimeActionCreator.changeTime>;

// const timeout = (delaytime: number) => new Promise((resolve) => setTimeout(resolve, delayTime));

const TimeActionCreator = {
    changeTimeRequest() {
        return {type: TIME_ACTION.CHANGE_TIME_REQUEST};
    },
    changeTimeCompleted(currentTime: Date) {
        return {
            type: TIME_ACTION.CHANGE_TIME_COMPLETED,
            payload: { currentTime: currentTime }
        };
    },
    changeTimeFailed() {
        return {type: TIME_ACTION.CHANGE_TIME_FAILED }
    }
    // asyncChangeTime() {
        // 의도적 지연 시간 1초
        // return (dispatch: ThunkDispatch<{}, {}, AnyAction>) => {
        //     dispatch(this.changeTimeRequest());
        //     setTimeout(() => {
        //         dispatch(this.changeTimeCompleted({ currentTime: new Date() }))
        //     }, 1000)
        // }
    // asyncChangeTime: createAsyncThunk(
    //     "changeTime",
    //     async (undefined, thunkAPI) => {
    //         await timeout(2000);
    //         return { currentTime: new Date() };
    //     }
    // )
}


export type TimeActionType =
    | ReturnType<typeof TimeActionCreator.changeTimeCompleted>
    | ReturnType<typeof TimeActionCreator.changeTimeRequest>
    | ReturnType<typeof TimeActionCreator.changeTimeFailed>

export default TimeActionCreator;