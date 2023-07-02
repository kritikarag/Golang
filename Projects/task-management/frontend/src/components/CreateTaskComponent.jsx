import React, { Component } from 'react'
import TaskService from '../services/TaskService';

class CreateTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
            task_name: '',
            task_detail: '',
            //date: ''
        }
        this.changetask_nameHandler =
            this.changetask_nameHandler.bind(this);
        this.changetask_detailHandler =
            this.changetask_detailHandler.bind(this);
        this.saveOrUpdateTask =
            this.saveOrUpdateTask.bind(this);
    }

    // step 3
    componentDidMount() {

        // step 4
        if (this.state.id === '_add') {
            return
        } else {
            TaskService.getTaskById(this.state.id)
            .then((res) => {
                let Task = res.data;
                this.setState({
                    task_name: Task.task_name,
                    task_detail: Task.task_detail,
                    //date: Task.date
                });
            });
        }
    }
    saveOrUpdateTask = (e) => {
        e.preventDefault();
        let Task = { task_name: this.state.task_name, task_detail:
             this.state.task_detail, date: this.state.date};
        console.log('Task => ' + JSON.stringify(Task));

        // step 5
        if (this.state.id === '_add') {
            TaskService.createTask(Task).then(res => {
                this.props.history.push('/Tasks');
            });
        } else {
            TaskService
            .updateTask(Task, this.state.id).then(res => {
                this.props.history.push('/Tasks');
            });
        }
    }

    changetask_nameHandler = (event) => {
        this.setState({ task_name: event.target.value });
    }

    changetask_detailHandler = (event) => {
        this.setState({ task_detail: event.target.value });
    }

    changedateHandler = (event) => {
        this.setState({ date: event.target.value });
    }

    cancel() {
        this.props.history.push('/Tasks');
    }

    getTitle() {
        if (this.state.id === '_add') {
            return <h3 className="text-center">Add Task</h3>
        } else {
            return <h3 className="text-center">Update Task</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
        <div className="container">
           <div className="row">
             <div className="card col-md-6 offset-md-3 offset-md-3">
                            {
                                this.getTitle()
                            }
                            <div className="card-body">
                                <form>
            <div className="form-group">
              <label> Task Name: </label>
                <input placeholder="First Name" 
                  name="task_name" className="form-control"
                    value={this.state.task_name} 
                      onChange={this.changetask_nameHandler} />
                          </div>
            <div className="form-group">
              <label> Task Detail: </label>
                <input placeholder="Task Detail" 
                   name="task_detail" className="form-control"
                     value={this.state.task_detail} 
                      onChange={this.changetask_detailHandler} />
                                    </div>
            <div className="form-group">
                <label> date : </label>
                    <input placeholder="date" 
                       name="date" className="form-control"
                        value={this.state.date} 
                         onChange={this.changedateHandler} />
                                    </div>

             <button className="btn btn-success" 
                  onClick={this.saveOrUpdateTask}>Save
                    </button>
             <button className="btn btn-danger" 
                  onClick={this.cancel.bind(this)} 
                     style={{ marginLeft: "10px" }}>Cancel
                        </button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default CreateTaskComponent
