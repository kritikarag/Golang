import React, { Component } from 'react'
import taskService from '../services/TaskService'

class ListTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                tasks: []
        }
        this.addTask = this.addTask.bind(this);
        this.edittask = this.edittask.bind(this);
        this.deletetask = this.deletetask.bind(this);
    }

    deletetask(id){
        taskService.deletetask(id).then( res => {
            this.setState({tasks: 
                this.state.tasks
                .filter(task => task.id !== id)});
        });
    }
    viewtask(id){
        this.props.history.push(`/view-task/${id}`);
    }
    edittask(id){
        this.props.history.push(`/add-task/${id}`);
    }

    componentDidMount(){
        taskService.gettasks().then((res) => {
            if(res.data==null)
            {
                this.props.history.push('/add-task/_add');
            }
            this.setState({ tasks: res.data});
        });
    }

    addTask(){
        this.props.history.push('/add-task/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">
                     tasks List</h2>
                 <div className = "row">
                    <button className="btn btn-primary"
                     onClick={this.addTask}> Add task</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className 
                        = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Task Name</th>
                                    <th> Task Details</th>
                                    <th> Task Date </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.tasks.map(
                                        task => 
                                        <tr key = {task.id}>
                                   <td> {task.task_name} </td>   
                                   <td> {task.task_detail}</td>
                                   <td> {task.date}</td>
                                             <td>
                      <button onClick={ () => 
                          this.edittask(task.id)} 
                               className="btn btn-info">Update 
                                 </button>
                       <button style={{marginLeft: "10px"}}
                          onClick={ () => this.deletetask(task.id)} 
                             className="btn btn-danger">Delete 
                                 </button>
                       <button style={{marginLeft: "10px"}} 
                           onClick={ () => this.viewtask(task.id)}
                              className="btn btn-info">View 
                                  </button>
                                    </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>
                 </div>
            </div>
        )
    }
}

export default ListTaskComponent