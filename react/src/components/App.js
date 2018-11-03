import React, { Component } from 'react';
import Home from './Home';
import SignIn from './SignIn';
import SignUp from './SignUp';
import DeleteUser from './DeleteUser';

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

class App extends Component {

  constructor(){
    super();
    let user = localStorage.getItem("userlogged")
    let LoggedIn = localStorage.getItem("LoggedIn")
    this.state={
      anylogged: false,
      userlogged: null,
    }
    if(LoggedIn==="true"){
      this.state.userlogged=user
      this.state.anylogged=true
    }
    this.handleLogout=this.handleLogout.bind(this)
  }

  handleLogout(){
    localStorage.setItem("userlogged",null)
    localStorage.setItem("LoggedIn",JSON.stringify(false))
    this.setState({anylogged:false,userlogged:null})
  }


  render() {
    return (
      <div>
        <Router>
          <div>
            <nav className="navbar navbar-default">
              <div className="container-fluid">
                <div className="navbar-header">
                  <Link className="navbar-brand" to={'/'}>React App</Link>
                </div>
                {console.log(this.state.anylogged)}
                <ul className="nav navbar-nav">
                  <li><Link to={'/'}>Home</Link></li>
                { (localStorage.LoggedIn==="true" && localStorage.userlogged==="admin") && 
                  <li><Link to={'/DeleteUser'}>Delete User</Link></li>
                }

                { !(localStorage.LoggedIn==="true") &&
                    <li><Link to={'/SignIn'}>Sign In</Link></li>}
                {    !(localStorage.LoggedIn==="true") &&
                    <li><Link to={'/SignUp'}>Sign Up</Link></li>}
                </ul>
              </div>
            </nav>
            {    (localStorage.LoggedIn==="true") &&
                  <li><button type="button" onClick={this.handleLogout}>Logout</button></li>}
            <Switch>
                 <Route exact path='/' component={Home} />
                 <Route exact path='/DeleteUser' component={DeleteUser} />
                 <Route exact path='/SignUp' component={SignUp} />
                 <Route exact path='/SignIn' component={SignIn} />
            </Switch>
          </div>
        </Router>
      </div>
    );
  }
}

export default App;
