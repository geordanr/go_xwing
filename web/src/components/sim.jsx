import React from 'react';
import { Button, Col, ControlLabel, FormControl, FormGroup, Panel, PageHeader, Row } from 'react-bootstrap';
import { connect } from 'react-redux';
import Highcharts from 'highcharts';
import Immutable from 'immutable';

import { requestSim } from '../store/sim';

var seriesSortHelper = function (a, b) {
    if (a[0] < b[0]) { return -1 }
    else if (a[0] > b[0]) { return 1 }
    else { return 0 }
};

const ParamsBase = React.createClass({
    propTypes: {
        params: React.PropTypes.object.isRequired,
        results: React.PropTypes.object.isRequired,
        onSimulate: React.PropTypes.func.isRequired,
    },
    contextTypes: {
        postUrl: React.PropTypes.string.isRequired,
    },
    handleClick: function () {
        this.props.onSimulate(this.context.postUrl, this.props.params);
    },
    render: function () {
        return (
            <Row>
                <Col xs={12}>
                    <Row>
                        <Col xs={12}>
                            <Button bsStyle="primary" bsSize="large" onClick={this.handleClick}>Simulate</Button>
                        </Col>
                    </Row>
                    <Row>
                        <Col xs={5}>
                            {/* <pre>{JSON.stringify(this.props.params.combatants)}</pre> */}
                            {/* <pre>{JSON.stringify(this.props.params.attack_queue)}</pre> */}
                            <FormGroup controlId="simParams">
                                <ControlLabel>Parameter JSON</ControlLabel>
                                <FormControl componentClass="textarea" value={JSON.stringify(this.props.params)} />
                            </FormGroup>
                        </Col>
                        <Col xs={5}>
                            <FormGroup controlId="simResults">
                                <ControlLabel>Simulation Response JSON</ControlLabel>
                                <FormControl componentClass="textarea" value={JSON.stringify(this.props.results)} />
                            </FormGroup>
                        </Col>
                    </Row>
                    <Row>
                        <Col xs={12}>
                            <SimResults results={this.props.results} />
                        </Col>
                    </Row>
                </Col>
            </Row>

        );
    },
});

export const Params = connect(
    (state) => {
        let attack_queue = state.attacks.map(atk => {
            let mods = Immutable.Map();
            atk.mods.map(mod => {
                let step = mod.get('step');
                if (!mods.has(step)) {
                    mods = mods.set(step, Immutable.List());
                }
                mods = mods.set(step, mods.get(step).push(Immutable.List([mod.get('actor'), mod.get('mod')])));
            });
            return {
                attacker: state.combatants.get(atk.get('attackerId')).get('name'),
                defender: state.combatants.get(atk.get('defenderId')).get('name'),
                mods: mods,
            };
        });
        let combatants = state.combatants.valueSeq().map(cbt => {
            let tokens = Immutable.Map();
            ['focus', 'evade', 'targetlock'].map(k => {
                tokens = tokens.set(k, cbt.get(k));
                cbt = cbt.delete(k);
            });
            return cbt.set('tokens', tokens);
        });
        return {
            params: {
                iterations: 10000,
                combatants,
                attack_queue,
            },
            results: state.sim,
        };
    },
    (dispatch) => {
        return {
            onSimulate: (postUrl, body) => {
                dispatch(requestSim(postUrl, body));
            },
        };
    }
)(ParamsBase);

export const ShipResult = React.createClass({
    propTypes: {
        name: React.PropTypes.string.isRequired,
        hull: React.PropTypes.array.isRequired,
        shields: React.PropTypes.array.isRequired,
    },
    render: function () {
        return (
            <Panel header={this.props.name}>
                <Row>
                    <Col xs={6}>
                        <Histogram title="Hull" series={this.props.hull.sort(seriesSortHelper)} />
                    </Col>
                    <Col xs={6}>
                        <Histogram title="Shields" series={this.props.shields.sort(seriesSortHelper)} />
                    </Col>
                </Row>
            </Panel>
        );
    }
});

export const Histogram = React.createClass({
    propTypes: {
        title: React.PropTypes.string.isRequired,
        series: React.PropTypes.array.isRequired,
    },
    componentDidMount: function () {
        this.chart = new Highcharts.Chart(this.refs.highcharts, {
            chart: { type: 'column' },
            title: { text: this.props.title },
            series: [ { data: this.props.series } ],
        });
    },
    componentDidUpdate: function () {
        this.chart.series[0].update({
            data: this.props.series
        });
    },
    componentWillUnmount: function () {
        this.chart.destroy();
    },
    render: function () {
        return (
            <div ref="highcharts" style={{width: '500px', height: '300px'}}></div>
        );
    }
});

export const SimResults = React.createClass({
    propTypes: {
        results: React.PropTypes.object.isRequired,
    },
    render: function () {
        let ships = [];
        for (let name in this.props.results) {
            let shipdata = this.props.results[name];
            // console.log(`render ${name}`);
            // console.dir(shipdata);
            ships.push(<ShipResult key={name} name={name} hull={shipdata.hull.histogram} shields={shipdata.shields.histogram} />);
        }
        ships.sort(function (a, b) {
            if (a.props.name < b.props.name) { return -1 }
            else if (a.props.name > b.props.name) { return 1 }
            else { return 0 }
        });
        if (ships.length === 0) {
            ships = 'Specify simulation parameters and run the simulation.';
        }
        return (
            <Row>
                <Col xs={12}>
                    <PageHeader>Simulation Results</PageHeader>
                    {ships}
                </Col>
            </Row>
        );
    }
});
