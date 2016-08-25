import React from 'react';
import { Button, Col, ControlLabel, FormControl, FormGroup, Row } from 'react-bootstrap';
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
        targetlock: React.PropTypes.string.isRequired,
        combatants: React.PropTypes.instanceOf(Immutable.Map).isRequired,
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
    onFocusTokensChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            focus: parseInt(e.target.value.trim()),
        });
    },
    onEvadeTokensChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            evade: parseInt(e.target.value.trim()),
        });
    },
    onTargetLockChange: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            targetlock: this.props.combatants.get(e.target.value.trim()).get('name'),
        });
    },
    onRemove: function () {
        this.props.onRemove({id: this.props.id});
    },
    render: function () {
        return (
            <Row>
                <Col xs={12}>
                    <FormGroup controlId="pilotName">
                        <ControlLabel>Pilot Name</ControlLabel>
                        <FormControl type="text" placeholder="Pilot name" onChange={this.onNameChange} />
                    </FormGroup>
                    <ShipSelector onChange={this.onShipChange} />
                    <FormGroup controlId="pilotSkill">
                        <ControlLabel>Skill</ControlLabel>
                        <FormControl type="number" placeholder="Pilot skill" onChange={this.onSkillChange} />
                    </FormGroup>
                    <FormGroup controlId="focusTokens">
                        <ControlLabel>Focus Tokens</ControlLabel>
                        <FormControl type="number" placeholder="Focus tokens" onChange={this.onFocusTokensChange} />
                    </FormGroup>
                    <FormGroup controlId="evadeTokens">
                        <ControlLabel>Evade Tokens</ControlLabel>
                        <FormControl type="number" placeholder="Evade tokens" onChange={this.onEvadeTokensChange} />
                    </FormGroup>
                    <FormGroup controlId="targetLock">
                        <ControlLabel>Target Lock</ControlLabel>
                        <CombatantSelector combatantType="lock target" combatants={this.props.combatants} onChange={this.onTargetLockChange} />
                    </FormGroup>
                    <Button onClick={this.onRemove}>Remove</Button>
                </Col>
            </Row>

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
            <Row>
                <Col xs={12}>
                    {
                        this.props.combatants.valueSeq().map(c => {
                            return (<Combatant key={c.get('id')} id={c.get('id')} combatants={this.props.combatants} onUpdate={this.props.onCombatantUpdate} onRemove={this.props.onCombatantRemove} ship={c.get('ship')} name={c.get('name')} skill={c.get('skill')} targetlock={c.get('targetlock')} />);
                        })
                    }
                    <Button onClick={this.props.addCombatant}>Add Combatant</Button>
                </Col>
            </Row>
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
        combatantType: React.PropTypes.string.isRequired,
        combatants: React.PropTypes.instanceOf(Immutable.List).isRequired,
        onChange: React.PropTypes.func.isRequired,
    },
    render: function () {
        return (
            <FormGroup controlId="combatantSelector">
                <FormControl componentClass="select" onChange={this.props.onChange}>
                    <option value=''>Select {this.props.combatantType}</option>
                    {
                        this.props.combatants.valueSeq().map(c => {
                            return (
                                <option key={c.get('id')} value={c.get('id')}>{c.get('name')}</option>
                            );
                        })
                    }
                </FormControl>
            </FormGroup>
        );
    },
});
