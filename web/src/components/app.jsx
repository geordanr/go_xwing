import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import Combatants from './combatants.jsx';
import Attacks from './attacks.jsx';

class App extends React.Component {
    getChildContext() {
        return {
            ships: this.props.ships,
        };
    }
    render() {
        return (
            <div>
                <Combatants />
                <Attacks />
            </div>
        );
    }
}
App.propTypes = {
    ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
};
App.childContextTypes = {
    ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
};

function renderUI(args) {
    ReactDOM.render(
        <Provider store={args.store}>
            <App ships={args.ships} />
        </Provider>,
        document.getElementById('app')
    );
}

export default renderUI;
