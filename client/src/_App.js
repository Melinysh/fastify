import React, { Component } from 'react';
import './App.css';

import WebTorrent from 'webtorrent';

import _ from 'lodash';

const sleep = t => new Promise(res => setTimeout(res, t));

const client = new WebTorrent();

var saveData = (function () {
  var a = document.createElement("a");
  document.body.appendChild(a);
  a.style = "display: none";
  return function (data, fileName) {
    var url = window.URL.createObjectURL(data);
    a.href = url;
    a.download = fileName;
    a.click();
    window.URL.revokeObjectURL(url);
  };
}());

async function fetchthings(that) {
  console.log("test")
  const response = await fetch(
    "http://ec2-54-83-190-222.compute-1.amazonaws.com:8080/", 
    {
      method: "POST",
      body: `http://i.imgur.com/Dag9eaxh.jpg`
    }
  )

  console.log(response)

  const torrentLink = await response.text();
  const torrent = await new Promise(res => client.add(torrentLink, res));
  // const torrent = await new Promise(res => client.add("magnet:?xt=urn:btih:6a9759bffd5c0af65319979fb7832189f4f3c35d&dn=sintel.mp4&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fsintel-1024-surround.mp4", res));

  torrent.on('download', function (bytes) {
    that.setState({
      downloads:{
        ...that.state.downloads,
        "test": {
          progress: torrent.progress
        }
      }
    })

    console.log('just downloaded: ' + bytes)
    console.log('total downloaded: ' + torrent.downloaded);
    console.log('download speed: ' + torrent.downloadSpeed)
    console.log('progress: ' + torrent.progress)
  })
  
  console.log("heres da torrent:", torrent)

  torrent.files[0].appendTo('body')
}

class App extends Component {
  state = {
    downloads: _.reduce(
      _.range(0,12), 
      (a,x) => { return { ...a, [x]: {progress: 0} }; }, 
      {}
    )
    // {
    //   "testfile.jayhess": { progress: 0, },
    //   "testf123": { progress: 0, },
    //   "poosad": { progress: 0, },
    // }
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
        {
          _.map(this.state.downloads, (torrent, name) => {
            return (
              <div key={name} className="file">
                <div 
                  className="file_loadingBar" 
                  style={{
                    width: `${torrent.progress * 100}%`,
                    "background-color": (torrent.progress === 1) ? "#51c34a" : "#81D4FA"
                  }}
                ></div>
                <div className="file_name">
                  {name}
                </div>

              </div>
            )
          })
        }
        <button onClick={() => fetchthings(this)}>poke steevo</button>
      </div>
    );
  }
}

export default App;

