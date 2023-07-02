import axios from 'axios';

const Task_API_BASE_URL = "http://localhost:9080/Tasks";

class TaskService {

    getTasks(){
        return axios.get(Task_API_BASE_URL);
    }

    createTask(Task){
        return axios.post(Task_API_BASE_URL, Task);
    }

    getTaskById(TaskId){
        return axios.get(Task_API_BASE_URL + '/' + TaskId);
    }

    updateTask(Task, TaskId){
        return axios.put(Task_API_BASE_URL + '/' + TaskId, Task);
    }

    deleteTask(TaskId){
        return axios.delete(Task_API_BASE_URL + '/' + TaskId);
    }
}

export default new TaskService()