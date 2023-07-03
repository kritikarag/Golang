import React from "react";
//import logo from './logo.svg';
import './App.css';
import { Container } from 'semantic-ui-react';
import TaskManagement from './task_management';
function App() {
  return (
    <div>
         <Container>
          <TaskManagement/>
         </Container>
    </div>
    
  );
}

export default App;