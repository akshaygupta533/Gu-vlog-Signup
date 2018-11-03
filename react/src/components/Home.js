import React, { Component } from 'react';
import './Home.css'

class Home extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
        {localStorage.getItem("userlogged")!='null' && <h1 className="App-title">Welcome, {localStorage.getItem("userlogged")}
          </h1> }
        </header>
      </div>
    );
  }
}

export default Home;