import React from 'react';
import { connect } from 'react-redux';
import * as Immutable from 'immutable';
import uuid from 'node-uuid';

import ShipSelector from './ships.jsx';
import { addCombatant, updateCombatant, removeCombatant } from '../store/combatants';

const Combatant = React.createClass({
    propTypes: {
        id: React.PropTypes.string.isRequired,
        onUpdate: React.PropTypes.func.isRequired,
        onRemove: React.PropTypes.func.isRequired,
        ship: React.PropTypes.string.isRequired,
        name: React.PropTypes.string.isRequired,
        skill: React.PropTypes.number.isRequired,
    },
    onShipChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            ship: e.target.value.trim(),
        });
    },
    onNameChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            name: e.target.value.trim(),
        });
    },
    onSkillChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            skill: parseInt(e.target.value.trim()),
        });
    },
    onRemove: function () {
        this.props.onRemove({id: this.props.id});
    },
    render: function () {
        return (
            <div>
                <input type="text" placeholder="Pilot name" onChange={this.onNameChange} />
                <ShipSelector onChange={this.onShipChange} />
                <input type="number" placeholder="Pilot skill" onChange={this.onSkillChange} />
                <button onClick={this.onRemove}>Remove</button>
            </div>
        );
    },
});

const Combatants = React.createClass({
    propTypes: {
        combatants: React.PropTypes.instanceOf(Immutable.Map).isRequired,
        addCombatant: React.PropTypes.func.isRequired,
        onCombatantUpdate: React.PropTypes.func.isRequired,
        onCombatantRemove: React.PropTypes.func.isRequired,
    },
    render: function () {
        return (
            <div>
                {
                    this.props.combatants.valueSeq().map(c => {
                        return (<Combatant key={c.get('id')} id={c.get('id')} onUpdate={this.props.onCombatantUpdate} onRemove={this.props.onCombatantRemove} ship={c.get('ship')} name={c.get('name')} skill={c.get('skill')} />);
                    })
                }
                <button onClick={this.props.addCombatant}>Add Combatant</button>
            </div>
        );
    },
});

export default connect(
    (state) => {
        return {
            combatants: state.combatants,
        };
    },
    (dispatch) => {
        return {
            addCombatant: () => {
                dispatch(addCombatant({id: uuid.v1(), name: '', ship: ''}));
            },
            onCombatantUpdate: (args) => {
                dispatch(updateCombatant(args));
            },
            onCombatantRemove: (args) => {
                dispatch(removeCombatant(args));
            },
        };
    }
)(Combatants);

export const CombatantSelector = React.createClass({
    propTypes: {
        combatants: React.PropTypes.instanceOf(Immutable.List).isRequired,
    },
    render: function () {
        return (
            <select>
                {
                    this.props.combatants.valueSeq().map(c => {
                        return (
                            <option key={c.get('id')} value={c.get('id')}>{c.get('name')}</option>
                        );
                    })
                }
            </select>
        );
    },
});
