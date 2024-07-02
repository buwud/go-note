class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            notes: []
        }
    }
    serverRequest() {
        $.get("http://localhost:8000/notes", res => {
            this.setState({
                notes: res
            });
        });
    }

    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>GoNote</h1>
                    <p>A load of notes :3</p>
                    <p>Sign in to take notes </p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
                <div className="container">
                    <div className="col-lg-12">
                        <br />
                        <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                        <h2>GoNote</h2>
                        <p>Let's feed you with some notes</p>
                        <div className="row">
                            {this.state.notes.map(function (note, i) {
                                return (<Note key={i} note={note} />);
                            })}
                        </div>
                    </div>
                </div>
            </div>


        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            notes: []
        }
    }
    serverRequest() {
        $.get("http://localhost:3000/api/jokes", res => {
            this.setState({
                notes: res
            });
        });
    }

    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br />
                    <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                    <h2>GoNote</h2>
                    <p>Let's feed you with some notes</p>
                    <div className="row">
                        {this.state.notes.map(function (note, i) {
                            return (<Note key={i} note={note} />);
                        })}
                    </div>
                </div>
            </div>
        )
    }
}

class Note extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            liked: ""
        }
        this.like = this.like.bind(this);
    }

    like() {
        // ... we'll add this block later
    }
    serverRequest() {
        $.get("http://localhost:8000/notes/", res => {
            this.setState({
                notes: res
            });
        });
    }

    render() {
        return (
            <div className="col">
                <div className="panel panel-default">
                    <div className="panel-heading">#{this.props.joke.id} <span className="pull-right">{this.state.liked}</span></div>
                    <div className="panel-body">
                        {this.props.joke.joke}
                    </div>
                    <div className="panel-footer">
                        {this.props.joke.likes} Likes &nbsp;
                        <a onClick={this.like} className="btn btn-default">
                            <span className="glyphicon glyphicon-thumbs-up"></span>
                        </a>
                    </div>
                </div>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));
