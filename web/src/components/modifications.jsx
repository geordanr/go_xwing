import React from 'react';
import { Button, Col, ControlLabel, FormControl, FormGroup, Row } from 'react-bootstrap';
import { connect } from 'react-redux';
import * as Immutable from 'immutable';

import { addAttackMod, updateAttackMod, removeAttackMod } from '../store/attacks';

const Modification = React.createClass({
    propTypes: {
        id: React.PropTypes.string.isRequired,
        attackId: React.PropTypes.string.isRequired,
        onRemove: React.PropTypes.func.isRequired,
        onUpdate: React.PropTypes.func.isRequired,
    },
    onActorChanged: function (e) {
        this.props.onUpdate(this.props.attackId, this.props.id, {
            actor: e.target.value.trim(),
        });
    },
    onModChanged: function (e) {
        this.props.onUpdate(this.props.attackId, this.props.id, {
            mod: e.target.value.trim(),
        });
    },
    onStepChanged: function (e) {
        this.props.onUpdate(this.props.attackId, this.props.id, {
            step: e.target.value.trim(),
        });
    },
    onRemove: function () {
        this.props.onRemove(this.props.attackId, this.props.id);
    },
    render: function () {
        return (
            <Row>
                <Col xs={12}>
                    <StepSelector onChange={this.onStepChanged} />
                    <ActorSelector onChange={this.onActorChanged} />
                    <ModSelector onChange={this.onModChanged} />
                    <Button onClick={this.onRemove}>Remove Mod</Button>
                </Col>
            </Row>
        );
    },
});

const ActorSelector = React.createClass({
    propTypes: {
        onChange: React.PropTypes.func.isRequired,
    },
    render: function () {
        return (
            <FormGroup controlId="actorSelector">
                <ControlLabel>Actor</ControlLabel>
                <FormControl componentClass="select" onChange={this.props.onChange}>
                    <option value="attacker">Attacker</option>
                    <option value="defender">Defender</option>
                </FormControl>
            </FormGroup>
        );
    },
});

const StepSelector = React.createClass({
    contextTypes: {
        steps: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    propTypes: {
        onChange: React.PropTypes.func.isRequired,
    },
    render: function () {
        return (
            <FormGroup controlId="stepSelector">
                <ControlLabel>Step</ControlLabel>
                <FormControl componentClass="select" onChange={this.props.onChange}>
                    {this.context.steps.map(step => {
                        return (
                            <option key={step}>{step}</option>
                        );
                    })}
                </FormControl>
            </FormGroup>
        );
    },
});

const ModSelector = React.createClass({
    contextTypes: {
        modifications: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    propTypes: {
        onChange: React.PropTypes.func.isRequired,
    },
    render: function () {
        return (
            <FormGroup controlId="modSelector">
                <FormControl componentClass="select" onChange={this.props.onChange}>
                    {this.context.modifications.map(mod => {
                        return (
                            <option key={mod}>{mod}</option>
                        );
                    })}
                </FormControl>
            </FormGroup>
        );
    },
});

const Modifications = React.createClass({
    propTypes: {
        attackId: React.PropTypes.string.isRequired,
        attacks: React.PropTypes.instanceOf(Immutable.List).isRequired,
        addMod: React.PropTypes.func.isRequired,
        updateMod: React.PropTypes.func.isRequired,
        removeMod: React.PropTypes.func.isRequired,
    },
    contextTypes: {
        steps: React.PropTypes.arrayOf(React.PropTypes.string).isRequired,
    },
    onAddMod: function () {
        this.props.addMod(this.props.attackId);
    },
    render: function () {
        let atk = this.props.attacks.find(a => a.get('id') === this.props.attackId);
        return (
            <Row>
                <Col xs={12}>
                    {
                        atk.get('mods').map(mod => {
                            return (
                                <Modification id={mod.get('id')} key={mod.get('id')} attackId={this.props.attackId} onUpdate={this.props.updateMod} onRemove={this.props.removeMod} />
                            );
                        })
                    }
                    <Button onClick={this.onAddMod}>Add Mod</Button>
                </Col>
            </Row>
        );
    },
});

export default connect(
    (state) => {
        return {
            attacks: state.attacks,
        };
    },
    (dispatch) => {
        return {
            addMod: (attackId) => {
                dispatch(addAttackMod(attackId));
            },
            updateMod: (attackId, modId, args) => {
                dispatch(updateAttackMod({
                    attackId,
                    modId,
                    args,
                }));
            },
            removeMod: (attackId, modId) => {
                dispatch(removeAttackMod({
                    attackId,
                    modId,
                }));
            },
        };
    }
)(Modifications);
