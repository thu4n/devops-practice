import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';

function App() {

  // Why is react semantic so ugly?
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('/api/message')
    .then(response => response.json())
    .then(data => setMessage(data.text))
    .catch(error => console.error('Error:', error));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <h1>{message}</h1>
      </header>
    </div>
  );
}

export default App;
