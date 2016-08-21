import React from 'react';

const ShipSelector = React.createClass({
    propTypes: {
        onChange: React.PropTypes.func.isRequired,
    },
    contextTypes: {
        ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    render: function () {
        return (
            <select ref="sel" onChange={this.props.onChange}>
                {this.context.ships.map(item => { return <option key={item}>{item}</option> })}
            </select>
        );
    },
});

export default ShipSelector;
