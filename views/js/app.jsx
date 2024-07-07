import Home from './homec.jsx';

class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}


//note component
const Note = ({ note }) => (
    <div className="note">
        <h3>{note.title}</h3>
        <p>{note.description}</p>
        <small>{note.created_at}</small>
    </div>
);

//user's page
class LoggedIn extends React.Component {

}

ReactDOM.render(<App />, document.getElementById('app'));
