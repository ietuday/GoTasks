import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Import useNavigate
import './AddTask.css'; // Add this for styling

const AddTask = () => {
  const [taskName, setTaskName] = useState('');
  const [taskDescription, setTaskDescription] = useState('');
  const navigate = useNavigate(); // Initialize navigate

  const handleSubmit = (event) => {
    event.preventDefault();

    const newTask = { title: taskName, description: taskDescription }; // Use state variables

    fetch('http://localhost:8080/tasks', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newTask),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log('Task added:', data);
        navigate('/task'); // Redirect to the Task List page after adding a task
      })
      .catch((error) => console.error('Error adding task:', error));
  };

  return (
    <div className="add-task-container">
      <h2 className="form-title">Add New Task</h2>
      <form onSubmit={handleSubmit} className="task-form">
        <div className="form-group">
          <label>Task Name</label>
          <input
            type="text"
            className="form-input"
            value={taskName}
            onChange={(e) => setTaskName(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Description</label>
          <textarea
            className="form-input"
            value={taskDescription}
            onChange={(e) => setTaskDescription(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="submit-btn">Add Task</button>
      </form>
    </div>
  );
};

export default AddTask;
