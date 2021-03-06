import store from './store';
import renderUI from './components/app.jsx';

window.store = store;

// Fetch necessary static data before rendering
const API_URL = '/api/v1';
Promise.all([
    fetch(`${API_URL}/ships`),
    fetch(`${API_URL}/modifications`),
    fetch(`${API_URL}/steps`),
])
    .then(responses => Promise.all(responses.map(response => response.json())))
    .then(jsons => {
        renderUI({
            store,
            postUrl: `${API_URL}/sim`,
            ships: [''].concat(jsons[0].data),
            modifications: [''].concat(jsons[1].data),
            steps: [''].concat(jsons[2].data),
        });
    });
