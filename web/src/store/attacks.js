import * as Immutable from 'immutable';
import { createAction, handleActions } from 'redux-actions';

// Actions
export const addAttack = createAction('ADD_ATTACK');
export const removeAttack = createAction('REMOVE_ATTACK');

// Reducers
export const attacks = handleActions({
    ADD_ATTACK: (state, action) => {
        return state.withMutations(list => {
            list.push(Immutable.Map(action.attack));
        });
    },
    REMOVE_ATTACK: (state, action) => {
        var item = state.find(atk => atk.get('id') === action.id);
        var idx = state.indexOf(item);
        if (idx != -1) {
            return state.delete(idx);
        }
        return state;
    },
}, Immutable.List());
