import * as Immutable from 'immutable';
import { createAction, handleActions } from 'redux-actions';

var AttackRecord = Immutable.Record({
    id: '',
    attackerId: '',
    defenderId: '',
});

// Actions
export const addAttack = createAction('ADD_ATTACK');
export const updateAttack = createAction('UPDATE_ATTACK');
export const removeAttack = createAction('REMOVE_ATTACK');

// Reducers
export const attacks = handleActions({
    ADD_ATTACK: (state, action) => {
        return state.withMutations(list => {
            list.push(new AttackRecord(action.payload));
        });
    },
    UPDATE_ATTACK: (state, action) => {
        let item = state.find(atk => atk.get('id') === action.payload.id);
        let idx = state.indexOf(item);
        if (idx != -1) {
            return state.set(idx, item.merge(action.payload));
        }
        return state;
    },
    REMOVE_ATTACK: (state, action) => {
        let item = state.find(atk => atk.get('id') === action.payload.id);
        let idx = state.indexOf(item);
        if (idx != -1) {
            return state.delete(idx);
        }
        return state;
    },
}, Immutable.List());
