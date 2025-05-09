import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom'; // Updated import

const EditTask = () => {
  const { id } = useParams(); // Get the task ID from the URL
  const [task, setTask] = useState({ name: '', description: '' });
  const navigate = useNavigate(); // Updated to useNavigate()

  useEffect(() => {
    // Fetch the task to edit
    fetch(`http://localhost:8080/tasks/${id}`)
      .then((response) => response.json())
      .then((data) => setTask(data))
      .catch((error) => console.error('Error fetching task:', error));
  }, [id]);

  const handleSubmit = (event) => {
    event.preventDefault();

    const updatedTask = { name: task.name, description: task.description };

    fetch(`http://localhost:8080/tasks/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedTask),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log('Task updated:', data);
        navigate('/'); // Redirect to the TaskList page after update
      })
      .catch((error) => console.error('Error updating task:', error));
  };

  return (
    <div>
      <h2>Edit Task</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name</label>
          <input
            type="text"
            value={task.name}
            onChange={(e) => setTask({ ...task, name: e.target.value })}
            required
          />
        </div>
        <div>
          <label>Description</label>
          <input
            type="text"
            value={task.description}
            onChange={(e) => setTask({ ...task, description: e.target.value })}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Update Task
        </button>
      </form>
    </div>
  );
};

export default EditTask;
