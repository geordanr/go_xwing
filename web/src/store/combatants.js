import * as Immutable from 'immutable';
import { createAction, handleActions } from 'redux-actions';

var CombatantRecord = Immutable.Record({
    id: '',
    ship: '',
    name: '',
    skill: 0,
    // nested makes merging suck
    focus: 0,
    evade: 0,
    targetlock: '',
});

// Actions
export const addCombatant = createAction('ADD_COMBATANT');
export const updateCombatant = createAction('UPDATE_COMBATANT');
export const removeCombatant = createAction('REMOVE_COMBATANT');

// Reducers
export const combatants = handleActions({
    ADD_COMBATANT: (state, action) => {
        return state.withMutations(map => {
            map.set(action.payload.id, new CombatantRecord(action.payload));
        });
    },
    UPDATE_COMBATANT: (state, action) => {
        return state.withMutations(map => {
            let c = map.get(action.payload.id);
            // TODO ensure it's in there
            map.set(action.payload.id, c.merge(action.payload));
        });
    },
    REMOVE_COMBATANT: (state, action) => {
        return state.withMutations(map => {
            map.delete(action.payload.id);
        });
    },
}, Immutable.Map());
