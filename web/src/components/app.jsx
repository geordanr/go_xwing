import React from 'react';

import Combatants from './combatants.jsx';

const App = React.createClass({
    propTypes: {
        ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    childContextTypes: {
        ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    getChildContext: function () {
        return {
            ships: this.props.ships,
        };
    },
    render: function () {
        return (
            <div>
                <Combatants />
            </div>
        );
    },
});

export default App;
