import * as Immutable from 'immutable';
import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunkMiddleware from 'redux-thunk';

import { ships } from './store/ships';
import { attacks } from './store/attacks';
import { combatants } from './store/combatants';

const store = createStore(combineReducers({
    attacks,
    combatants,
    ships,
}), applyMiddleware(thunkMiddleware));

export default store;
