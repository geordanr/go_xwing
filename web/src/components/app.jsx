import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import Combatants from './combatants.jsx';
import Attacks from './attacks.jsx';
import { Params } from './sim.jsx';

class App extends React.Component {
    getChildContext() {
        return {
            postUrl: this.props.postUrl,
            ships: this.props.ships,
            modifications: this.props.modifications,
            steps: this.props.steps,
        };
    }
    render() {
        return (
            <div>
                <Combatants />
                <Attacks />
                <Params />
            </div>
        );
    }
}
App.propTypes = {
    postUrl: React.PropTypes.string.isRequired,
    ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    modifications: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    steps: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
};
App.childContextTypes = {
    postUrl: React.PropTypes.string.isRequired,
    ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    modifications: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    steps: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
};

function renderUI(args) {
    ReactDOM.render(
        <Provider store={args.store}>
            <App postUrl={args.postUrl} ships={args.ships} modifications={args.modifications} steps={args.steps} />
        </Provider>,
        document.getElementById('app')
    );
}

export default renderUI;
