import ReactDOM from 'react-dom';
import React from 'react';
import { Provider } from 'react-redux';

import App from './components/app.jsx';
// import { results } from './components/simresults.jsx';
// import { Attacks } from './components/attacks.jsx';

function renderUI(args) {
    ReactDOM.render(
        <Provider store={args.store}>
            <App ships={args.ships} />
        </Provider>,
        document.getElementById('app')
    );
}

export default renderUI;
