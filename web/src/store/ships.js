import { createAction, handleActions } from 'redux-actions';

// Actions
export function fetchShips(url) {
    return function (dispatch) {
        console.log(`fetch(${url})`);
        return fetch(url)
        .then(response => response.json())
        .then(json => dispatch(receiveShips(json.data)));
    };
}
export const receiveShips = createAction('RECEIVE_SHIPS');

// Reducers
export const ships = handleActions({
    RECEIVE_SHIPS: (state, action) => action.payload,
}, []);
