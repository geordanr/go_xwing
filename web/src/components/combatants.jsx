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
            ship: e.target.value,
        });
    },
    onNameChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            name: e.target.value,
        });
    },
    onSkillChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            skill: parseInt(e.target.value),
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
                <input type="number" onChange={this.onSkillChange} />
                <button onClick={this.onRemove}>Remove</button>
            </div>
        );
    },
});

const CombatantsBase = React.createClass({
    propTypes: {
        combatants: React.PropTypes.instanceOf(Immutable.Map),
        addCombatant: React.PropTypes.func,
        onCombatantUpdate: React.PropTypes.func,
        onCombatantRemove: React.PropTypes.func,
    },
    render: function () {
        return (
            <div>
                {
                    this.props.combatants.valueSeq().map(c => {
                        return (<Combatant key={c.get('id')} id={c.get('id')} onUpdate={this.props.onCombatantUpdate} onRemove={this.props.onCombatantRemove} ship={c.get('ship')} name={c.get('name')} skill={c.get('skill')} />);
                    })
                }
                <button ref="add" onClick={this.props.addCombatant}>Add</button>
                <div>{this.props.combatants.valueSeq().map(c => {
                    return (
                        <span key={c.get('id')}>Name: {c.get('name')} Ship: {c.get('ship')} PS {c.get('skill')}<br /></span>
                    );
                })}</div>
            </div>
        );
    },
});

const Combatants = connect(
    (state) => {
        return {
            combatants: state.combatants.toMap(),
        };
    },
    (dispatch) => {
        return {
            addCombatant: () => {
                dispatch(addCombatant({name: '', id: uuid.v1(), ship: ''}));
            },
            onCombatantUpdate: (args) => {
                dispatch(updateCombatant(args));
            },
            onCombatantRemove: (args) => {
                dispatch(removeCombatant(args));
            },
        };
    }
)(CombatantsBase);

export default Combatants;
