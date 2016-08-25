import React from 'react';
import { ControlLabel, FormControl, FormGroup } from 'react-bootstrap';

const ShipSelector = React.createClass({
    propTypes: {
        onChange: React.PropTypes.func.isRequired,
    },
    contextTypes: {
        ships: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    render: function () {
        return (
            <FormGroup controlId="shipSelector">
                <ControlLabel>Ship</ControlLabel>
                <FormControl componentClass="select" placeholder="Select ship" onChange={this.props.onChange}>
                    {this.context.ships.map(item => { return <option key={item}>{item}</option> })}
                </FormControl>
            </FormGroup>
        );
    },
});

export default ShipSelector;
