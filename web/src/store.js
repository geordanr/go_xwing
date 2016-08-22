import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunkMiddleware from 'redux-thunk';

import { ships } from './store/ships';
import { attacks } from './store/attacks';
import { combatants } from './store/combatants';
import { sim } from './store/sim';

const store = createStore(combineReducers({
    attacks,
    combatants,
    ships,
    sim,
}), applyMiddleware(thunkMiddleware));

export default store;
