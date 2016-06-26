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
    render: function () {
        return (
            <textarea ref="textarea" className="json-input" onChange={this.handleChange}></textarea>
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
        console.log("Posting:"); console.dir(this.refs.textbox.state.data);
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
            this.setState({results: data});
        });
    },
    componentWillUnmount: function () {
        $(window).off('xwing-sim:resultsReturned');
    },
    render: function () {
        return (
            <div>
                <span>Simulation Results</span>
                <pre>{JSON.stringify(this.state.results)}</pre>
            </div>
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
