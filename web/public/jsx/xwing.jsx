var seriesSortHelper = function (a, b) {
    if (a[0] < b[0]) {
        return -1
    } else if (a[0] > b[0]) {
        return 1
    } else {
        return 0
    }
}

var ParameterJSONInputTextbox = React.createClass({
    getInitialState: function () {
        return {
            data: null
        }
    },
    handleChange: function (e) {
        try {
            this.state.data = JSON.parse(this.refs.textarea.value);
        } catch (e) {
            // console.warn(`Cannot parse JSON: ${this.refs.textarea.value}, error: ${e}`);
            this.state.data = null;
        }
    },
    componentDidMount: function () {
        this.refs.textarea.value = `{
    "iterations": 10000,
    "combatants": [
        {
            "name": "Miranda Doni",
            "ship": "K-Wing",
            "skill": 8,
            "initiative": true,
            "tokens": {
                "targetlock": "Howlrunner"
            }
        },
        {
            "name": "Howlrunner",
            "ship": "TIE Fighter",
            "skill": 8,
            "initiative": false,
            "tokens": {
                "focus": 1,
                "evade": 1
            }
        }
    ],
    "attack_queue": [
        {
            "attacker": "Miranda Doni",
            "defender": "Howlrunner",
            "mods": {
                "Declare Target": [
                    ["attacker", "Twin Laser Turret"]
                ],
                "Modify Attack Dice": [
                    ["attacker", "Spend Target Lock"],
                    ["attacker", "Spend Focus Token"]
                ],
                "Modify Defense Dice": [
                    ["defender", "Spend Focus Token"],
                    ["defender", "Spend Evade Token"]
                ],
                "Perform Additional Attack": [
                    ["attacker", "Gunner"]
                ]
            }
        },
        {
            "attacker": "Howlrunner",
            "defender": "Miranda Doni",
            "mods": {
                "Modify Attack Dice": [
                    ["attacker", "Spend Focus Token"]
                ],
                "Roll Defense Dice": [
                    ["defender", "Roll Defense Dice"]
                ],
                "Modify Defense Dice": [
                    ["defender", "Spend Focus Token"]
                ],
                "Compare Results": [
                    ["attacker", "Crack Shot"],
                    ["attacker", "Compare Results" ]
                ]
            }
        }
    ]
}`;
    },
    render: function () {
        return (
            <textarea ref="textarea" rows="15" cols="200" className="json-input" onChange={this.handleChange}></textarea>
        )
    }
});

var ParameterJSONSubmitButton = React.createClass({
    handleClick: function (e) {
        this.props.onClick(e);
    },
    render: function () {
        return (
            <button onClick={this.handleClick}>Submit</button>
        )
    }
});

var ParameterJSON = React.createClass({
    handleSubmit: function (e) {
        $.post({
            url: this.props.url,
            data: JSON.stringify(this.refs.textbox.state.data),
            contentType: 'application/json',
        })
        .done(function (data) {
            $(window).trigger('xwing-sim:resultsReturned', data);
        });
    },
    render: function () {
        return (
            <div>
                <span>Simulation Parameters (JSON)</span>
                <ParameterJSONInputTextbox ref="textbox"/>
                <ParameterJSONSubmitButton ref="button" onClick={this.handleSubmit} />
            </div>
        )
    }
});

var SimResults = React.createClass({
    getInitialState: function () {
        return {
            results: null
        }
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
        var ships = [];
        for (var name in this.state.results) {
            var shipdata = this.state.results[name];
            console.log(`render ${name}`);
            // console.dir(shipdata);
            ships.push(<ShipResult key={name} name={name} hull={shipdata.hull.histogram} shields={shipdata.shields.histogram} />);
        }
        ships.sort(function (a, b) {
            if (a.props.name < b.props.name) {
                return -1
            } else if (a.props.name > b.props.name) {
                return 1
            } else {
                return 0
            }
        });
        return (
            <div>
                <span>Simulation Results</span>
                {ships}
            </div>
        )
    }
});

var ShipResult = React.createClass({
    render: function () {
        return (
            <div>
                <span>{this.props.name}</span>
                <Histogram title="Hull" series={this.props.hull.sort(seriesSortHelper)} />
                <Histogram title="Shields" series={this.props.shields.sort(seriesSortHelper)} />
            </div>
        )
    }
});

var Histogram = React.createClass({
    componentDidMount: function () {
        this.chart = new Highcharts.Chart(this.refs.highcharts, {
            chart: { type: 'column' },
            title: { text: this.props.title },
            series: [ { data: this.props.series } ],
        });
    },
    componentDidUpdate: function (prevProps, prevState) {
        this.chart.series[0].update({
            data: this.props.series
        });
    },
    componentWillUnmount: function () {
        console.log(`unmount destroy ${this.chart}`);
        this.chart.destroy();
    },
    render: function () {
        return (
            <div ref="highcharts" style={{width: '500px', height: '300px'}}></div>
        )
    }
});

var params = React.createElement(ParameterJSON, {
    url: "/api/v1/sim",
});
var results = React.createElement(SimResults, null);

ReactDOM.render(
    params,
    document.getElementById('params')
);

ReactDOM.render(
    results,
    document.getElementById('results')
);
