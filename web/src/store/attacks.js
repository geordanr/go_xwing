import * as Immutable from 'immutable';
import { createAction, handleActions } from 'redux-actions';
import uuid from 'node-uuid';

var AttackRecord = Immutable.Record({
    id: '',
    attackerId: '',
    defenderId: '',
    mods: Immutable.List(), // list of {id, stepname, modname}, marshaled elsewhere
});

// Actions
export const addAttack = createAction('ADD_ATTACK');
export const updateAttack = createAction('UPDATE_ATTACK');
export const removeAttack = createAction('REMOVE_ATTACK');
export const addAttackMod = createAction('ADD_ATTACK_MOD');
export const updateAttackMod = createAction('UPDATE_ATTACK_MOD');
export const removeAttackMod = createAction('REMOVE_ATTACK_MOD');

// Reducers
export const attacks = handleActions({
    ADD_ATTACK: (state, action) => {
        return state.withMutations(list => {
            list.push(new AttackRecord(action.payload));
        });
    },
    UPDATE_ATTACK: (state, action) => {
        let atk = state.find(atk => atk.get('id') === action.payload.id);
        let idx = state.indexOf(atk);
        if (idx != -1) {
            return state.set(idx, atk.merge(action.payload));
        }
        return state;
    },
    REMOVE_ATTACK: (state, action) => {
        let atk = state.find(atk => atk.get('id') === action.payload.id);
        let idx = state.indexOf(atk);
        if (idx != -1) {
            return state.delete(idx);
        }
        return state;
    },
    ADD_ATTACK_MOD: (state, action) => {
        let atk = state.find(atk => atk.get('id') === action.payload);
        let idx = state.indexOf(atk);
        if (idx != -1) {
            return state.set(idx, atk.merge({
                mods: atk.get('mods').push(Immutable.Map({
                    id: uuid.v1(),
                    step: '',
                    actor: '',
                    mod: '',
                })),
            }));
        }
        return state;
    },
    UPDATE_ATTACK_MOD: (state, action) => {
        let atk = state.find(atk => atk.get('id') === action.payload.attackId);
        let atkIdx = state.indexOf(atk);
        if (atkIdx != -1) {
            let mods = atk.get('mods');
            let mod = mods.find(m => m.get('id') === action.payload.modId);
            let modIdx = mods.indexOf(mod);
            if (modIdx != -1) {
                let updated = atk.mods.get(modIdx).merge(action.payload.args);
                return state.set(atkIdx, atk.merge({mods: mods.set(modIdx, updated)}));
            }
        }
        return state;
    },
    REMOVE_ATTACK_MOD: (state, action) => {
        let atk = state.find(atk => atk.get('id') === action.payload.attackId);
        let atkIdx = state.indexOf(atk);
        if (atkIdx != -1) {
            let mods = atk.get('mods');
            let mod = mods.find(m => m.get('id') === action.payload.modId);
            let modIdx = mods.indexOf(mod);
            if (modIdx != -1) {
                return state.set(atkIdx, atk.merge({mods: mods.delete(modIdx)}));
            }
        }
        return state;
    },
}, Immutable.List());
