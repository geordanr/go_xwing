/* global require */
require('./style.sass');

import store from './store';
import renderUI from './react.jsx';

// Fetch necessary static data before rendering
const API_URL = '/api/v1/ships';
fetch(API_URL)
    .then(resp => resp.json())
    .then(json => {
        renderUI({
            store,
            ships: [''].concat(json.data),
        });
    });
