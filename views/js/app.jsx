class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

const Note = ({ note }) => (
    <div className="note">
        <h3>{note.title}</h3>
        <p>{note.description}</p>
        <small>{note.created_at}</small>
    </div>
);

class Home extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            notes: []
        };
    }

    componentDidMount() {
        this.serverRequest();
    }

    serverRequest() {
        fetch("http://localhost:8000/notes")
            .then(response => response.json())
            .then(res => {
                if (res.success) {
                    this.setState({
                        notes: res.data.rows
                    });
                }
            })
            .catch(error => console.error('Error fetching notes:', error));
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
                            {this.state.notes.map((note, i) => (
                                <Note key={i} note={note} />
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        );
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
        $.get("http://localhost:8000/notes/", res => {
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

ReactDOM.render(<App />, document.getElementById('app'));