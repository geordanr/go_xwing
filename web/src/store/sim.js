import { createAction, handleActions } from 'redux-actions';

// Actions
export function requestSim(postUrl, body) {
    return function (dispatch) {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        return fetch(postUrl, {
            method: 'POST',
            headers,
            body: JSON.stringify(body),
        })
        .then(resp => resp.json())
        .then(json => {
            dispatch(receiveSimResults(json));
        });
    };
}

export const receiveSimResults = createAction('RECEIVE_SIM_RESULTS');

// Reducers
export const sim = handleActions({
    RECEIVE_SIM_RESULTS: (state, action) => {
        return action.payload;
    },
}, {});
