import React from 'react';
//import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import ListTaskComponent from './components/ListTaskComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import CreateTaskComponent from './components/CreateTaskComponent';
import ViewTaskComponent from './components/ViewTaskComponent';

function App() {
  return (
    <div>
        <Router>
              <HeaderComponent />
                <div className="container">
                    <Routes> 
                          <Route path = "/" exact component =
                              {ListTaskComponent}></Route>
                          <Route path = "/Tasks" component = 
                              {ListTaskComponent}></Route>
                          <Route path = "/add-Task/:id" component = 
                              {CreateTaskComponent}></Route>
                         <Route path = "/view-Task/:id" component = 
                              {ViewTaskComponent}></Route>
                         </Routes>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;