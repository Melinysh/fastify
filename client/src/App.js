import React, { Component } from 'react';
import './App.css';

import logo from '../public/favicon.png'

import _ from 'lodash';


function debounce(func, wait, immediate) {
  var timeout;
  return function() {
    var context = this, args = arguments;
    var later = function() {
      timeout = null;
      if (!immediate) func.apply(context, args);
    };
    var callNow = immediate && !timeout;
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
    if (callNow) func.apply(context, args);
  };
};

const sleep = t => new Promise(res => setTimeout(res, t));

const fetchthings = debounce(async function(that, link = "http://i.imgur.com/Dag9eaxh.jpg") {

  that.setState({...that.state, isLoading: true});
  const response = await fetch(
    "http://ec2-52-90-47-131.compute-1.amazonaws.com:8080/", 
    {
      method: "POST",
      body: link
    }
  )


  const torrentLink = await response.text();

  console.log(torrentLink)

  that.setState({...that.state, link: torrentLink});


  return torrentLink;
}, 2000)

class App extends Component {
  state = {
    downloads: _.reduce(
      _.range(0,12), 
      (a,x) => { return { ...a, [x]: {progress: 0} }; }, 
      {}
    )
  }

  componentDidMount() {
    _.map(this.state.downloads, (torrent, name) => {
      const handle = setInterval(() => {
        let newProgress = this.state.downloads[name].progress + Math.random() * 0.1;

        if (newProgress > 1) {
          clearInterval(handle);
          newProgress = 1
        }

        this.setState({
          downloads: {
            ...this.state.downloads,
            [name]: { progress: newProgress }
          }
        })
      }, 250);
    });
  }

  render() {
    return (
      <div className="App">
        <div className="imgContainer">
          <img src={logo} />
        </div>
        <p>Give us a file link, and we'll give you a <em>superfast</em> torrent back!<br/>Speed up your downloads today!</p>
        <input placeholder="Paste link here!" onChange={e => fetchthings(this, e.target.value)}></input>

        {
          this.state.link 
            ? <a href={this.state.link}>Download!</a> 
            : ( this.state.isLoading 
              ? <a>Loading...</a> 
              : null 
            )
        }
      </div>
    );
  }
}

export default App;
