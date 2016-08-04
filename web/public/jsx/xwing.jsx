/* global Highcharts, uuid */
const API_URL='http://localhost:8080/api/v1';
var SHIPS = null;

var seriesSortHelper = function (a, b) {
    if (a[0] < b[0]) { return -1 }
    else if (a[0] > b[0]) { return 1 }
    else { return 0 }
};

// var ParameterJSONInputTextbox = React.createClass({
//     getInitialState: function () {
//         return {
//             data: null
//         };
//     },
//     handleChange: function () {
//         try {
//             this.setState({data: JSON.parse(this.refs.textarea.value)});
//         } catch (e) {
//             // console.warn(`Cannot parse JSON: ${this.refs.textarea.value}, error: ${e}`);
//             this.setState({data: null});
//         }
//     },
//     componentDidMount: function () {
//         this.refs.textarea.value = `{
//     "iterations": 10000,
//     "combatants": [
//         {
//             "name": "Miranda Doni",
//             "ship": "K-Wing",
//             "skill": 8,
//             "initiative": true,
//             "tokens": {
//                 "targetlock": "Howlrunner"
//             }
//         },
//         {
//             "name": "Howlrunner",
//             "ship": "TIE Fighter",
//             "skill": 8,
//             "initiative": false,
//             "tokens": {
//                 "focus": 1,
//                 "evade": 1
//             }
//         }
//     ],
//     "attack_queue": [
//         {
//             "attacker": "Miranda Doni",
//             "defender": "Howlrunner",
//             "mods": {
//                 "Declare Target": [
//                     ["attacker", "Twin Laser Turret"]
//                 ],
//                 "Modify Attack Dice": [
//                     ["attacker", "Spend Target Lock"],
//                     ["attacker", "Spend Focus Token"]
//                 ],
//                 "Modify Defense Dice": [
//                     ["defender", "Spend Focus Token"],
//                     ["defender", "Spend Evade Token"]
//                 ],
//                 "Perform Additional Attack": [
//                     ["attacker", "Gunner"]
//                 ]
//             }
//         },
//         {
//             "attacker": "Howlrunner",
//             "defender": "Miranda Doni",
//             "mods": {
//                 "Modify Attack Dice": [
//                     ["attacker", "Spend Focus Token"]
//                 ],
//                 "Roll Defense Dice": [
//                     ["defender", "Roll Defense Dice"]
//                 ],
//                 "Modify Defense Dice": [
//                     ["defender", "Spend Focus Token"]
//                 ],
//                 "Compare Results": [
//                     ["attacker", "Crack Shot"],
//                     ["attacker", "Compare Results" ]
//                 ]
//             }
//         }
//     ]
// }`;
//     },
//     render: function () {
//         return (
//             <textarea ref="textarea" rows="15" cols="200" className="json-input" onChange={this.handleChange}></textarea>
//         );
//     }
// });

// var ParameterJSONSubmitButton = React.createClass({
//     propTypes: {
//         onClick: React.PropTypes.func.isRequired
//     },
//     handleClick: function (e) {
//         this.props.onClick(e);
//     },
//     render: function () {
//         return (
//             <button onClick={this.handleClick}>Submit</button>
//         );
//     }
// });

// var ParameterJSON = React.createClass({
//     propTypes: {
//         url: React.PropTypes.string.isRequired
//     },
//     handleSubmit: function () {
//         $.post({
//             url: this.props.url,
//             data: JSON.stringify(this.refs.textbox.state.data),
//             contentType: 'application/json',
//         })
//         .done(function (data) {
//             $(window).trigger('xwing-sim:resultsReturned', data);
//         });
//     },
//     render: function () {
//         return (
//             <div>
//                 <span>Simulation Parameters (JSON)</span>
//                 <ParameterJSONInputTextbox ref="textbox"/>
//                 <ParameterJSONSubmitButton ref="button" onClick={this.handleSubmit} />
//             </div>
//         );
//     }
// });

var SimResults = React.createClass({
    getInitialState: function () {
        return {
            results: null
        };
    },
    componentDidMount: function () {
        $(window).on('xwing-sim:resultsReturned', (e, data) => {
            // console.log('set state'); console.dir(data);
            this.setState({results: data});
        });
    },
    componentWillUnmount: function () {
        $(window).off('xwing-sim:resultsReturned');
    },
    render: function () {
        let ships = [];
        for (let name in this.state.results) {
            let shipdata = this.state.results[name];
            // console.log(`render ${name}`);
            // console.dir(shipdata);
            ships.push(<ShipResult key={name} name={name} hull={shipdata.hull.histogram} shields={shipdata.shields.histogram} />);
        }
        ships.sort(function (a, b) {
            if (a.props.name < b.props.name) { return -1 }
            else if (a.props.name > b.props.name) { return 1 }
            else { return 0 }
        });
        return (
            <div>
                <span>Simulation Results</span>
                {ships}
            </div>
        );
    }
});

var ShipResult = React.createClass({
    propTypes: {
        name: React.PropTypes.string.isRequired,
        hull: React.PropTypes.array.isRequired,
        shields: React.PropTypes.array.isRequired,
    },
    render: function () {
        return (
            <div>
                <span>{this.props.name}</span>
                <Histogram title="Hull" series={this.props.hull.sort(seriesSortHelper)} />
                <Histogram title="Shields" series={this.props.shields.sort(seriesSortHelper)} />
            </div>
        );
    }
});

var Histogram = React.createClass({
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
        // console.log(`unmount destroy ${this.chart}`);
        this.chart.destroy();
    },
    render: function () {
        return (
            <div ref="highcharts" style={{width: '500px', height: '300px'}}></div>
        );
    }
});

// var params = React.createElement(ParameterJSON, {
//     url: '/api/v1/sim',
// });
var results = React.createElement(SimResults, null);

var Combatant = React.createClass({
    getInitialState: function () {
        return {
            name: '',
            ship: '',
            skill: 1,
            tokens: {},
        };
    },
    handleNameChange: function () {
        this.setState({name: this.refs.name.value});
    },
    handleShipChange: function () {
        this.setState({ship: this.refs.ship.value});
    },
    handlePilotSkillChange: function () {
        this.setState({skill: parseInt(this.refs.pilotskill.value)});
    },
    render: function () {
        let ships = SHIPS.map((item) => { return <option key={item}>{item}</option> });
        return (
            <div>
                <input type="text" ref="name" onChange={this.handleNameChange} value={this.state.name}/>
                <select ref="ship" onChange={this.handleShipChange}>{ships} value={this.state.ship}</select>
                <input type="number" ref="pilotskill" onChange={this.handlePilotSkillChange} value={this.state.skill}/>
            </div>
        );
    }
});

var Combatants = React.createClass({
    getInitialState: function () {
        return {
            combatants: []
        };
    },
    handleAddCombatant: function () {
        this.setState({combatants: this.state.combatants.concat([<Combatant key={uuid.v1()}/>])});
        $(window).trigger('xwing-sim:combatantsChanged', this.state.combatants);
    },
    render: function () {
        return (
            <div>
                {this.state.combatants}
                <button ref="add" onClick={this.handleAddCombatant}>Add Combatant</button>
            </div>
        );
    }
});

var Attack = React.createClass({
    getInitialState: function () {
        return {
            combatants: null,
            attacker: '',
            defender: '',
        };
    },
    componentDidMount: function () {
        $(window).on('xwing-sim:combatantsChanged', (e, data) => {
            this.setState({combatants: data});
        });
    },
    componentWillUnmount: function () {
        $(window).off('xwing-sim:combatantsChanged');
    },
    render: function () {
        return (
            <div>
                <label>
                    Attacker
                    <select ref="attacker">{this.state.combatants}</select>
                </label>
            </div>
        );
    }
});

var Attacks = React.createClass({
    getInitialState: function () {
        return {
            attacks: []
        };
    },
    handleAddAttack: function () {
        this.setState({attacks: this.state.attacks.concat([<Attack key={uuid.v1()}/>])});
    },
    render: function () {
        return (
            <div>
                {this.state.attacks}
                <button ref="add" onClick={this.handleAddAttack}>Add Attack</button>
            </div>
        );
    }
});

// fetch ship data, then render things
$.getJSON(API_URL + '/ships', (data) => {
    SHIPS = data.data;
})
.then(() => {
    // ReactDOM.render(
    //     params,
    //     document.getElementById('params')
    // );

    ReactDOM.render(
        results,
        document.getElementById('results')
    );

    ReactDOM.render(
        React.createElement(Combatants),
        document.getElementById('combatants')
    );

    ReactDOM.render(
        React.createElement(Attacks),
        document.getElementById('attacks')
    );
});
