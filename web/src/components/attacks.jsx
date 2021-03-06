import React from 'react';
import { Button, Col, Panel, Row } from 'react-bootstrap';
import { connect } from 'react-redux';
import * as Immutable from 'immutable';
import uuid from 'node-uuid';

import { addAttack, updateAttack, removeAttack } from '../store/attacks';
import { CombatantSelector } from './combatants.jsx';
import Modifications from './modifications.jsx';

const Attack = React.createClass({
    propTypes: {
        id: React.PropTypes.string.isRequired,
        attackerId: React.PropTypes.string.isRequired,
        defenderId: React.PropTypes.string.isRequired,
        combatants: React.PropTypes.instanceOf(Immutable.List),
        onUpdate: React.PropTypes.func.isRequired,
        onRemove: React.PropTypes.func.isRequired,
    },
    onAttackerUpdate: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            attackerId: e.target.value.trim(),
        });
    },
    onDefenderUpdate: function (e) {
        this.props.onUpdate({
            id: this.props.id,
            defenderId: e.target.value.trim(),
        });
    },
    onRemove: function () {
        this.props.onRemove({id: this.props.id});
    },
    render: function () {
        return (
            <Panel>
                <Row>
                    <Col xs={12}>
                        <Row>
                            <Col xs={5}>
                                <CombatantSelector ref="attacker" combatantType='Attacker' combatants={this.props.combatants} onChange={this.onAttackerUpdate} />
                            </Col>
                            <Col xs={5}>
                                <CombatantSelector ref="defender" combatantType='Defender' combatants={this.props.combatants} onChange={this.onDefenderUpdate} />
                            </Col>
                            <Col xs={2}>
                                <Button bsStyle="danger" onClick={this.onRemove}><i className="fa fa-trash"></i></Button>
                            </Col>
                        </Row>
                        <Row>
                            <Col xs={12}>
                                <Modifications attackId={this.props.id} />
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </Panel>
        );
    },
});

const Attacks = React.createClass({
    propTypes: {
        attacks: React.PropTypes.instanceOf(Immutable.List).isRequired,
        combatants: React.PropTypes.instanceOf(Immutable.List).isRequired,
        addAttack: React.PropTypes.func.isRequired,
        onAttackUpdate: React.PropTypes.func.isRequired,
        onAttackRemove: React.PropTypes.func.isRequired,
    },
    render: function () {
        return (
            <Row>
                <Col xs={12}>
                    {
                        this.props.attacks.valueSeq().map(a => {
                            return (<Attack key={a.get('id')} id={a.get('id')} attackerId={a.get('attackerId')} defenderId={a.get('defenderId')} combatants={this.props.combatants} onUpdate={this.props.onAttackUpdate} onRemove={this.props.onAttackRemove} />);
                        })
                    }
                    <Button onClick={this.props.addAttack}><i className="xwing-miniatures-font xwing-miniatures-font-crit"></i> Add Attack</Button>
                </Col>
            </Row>
        );
    },
});

export default connect(
    (state) => {
        return {
            attacks: state.attacks,
            combatants: state.combatants.toList(),
        };
    },
    (dispatch) => {
        return {
            addAttack: () => {
                dispatch(addAttack({id: uuid.v1(), attackerId: '', defenderId: ''}));
            },
            onAttackUpdate: (args) => {
                dispatch(updateAttack(args));
            },
            onAttackRemove: (args) => {
                dispatch(removeAttack(args));
            },
        };
    }
)(Attacks);
