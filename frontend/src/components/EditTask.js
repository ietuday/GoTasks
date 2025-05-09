import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import './EditTask.css'; // Add this for custom styles

const EditTask = () => {
  const { id } = useParams();
  const [task, setTask] = useState({ title: '', description: '', completed: false });
  const navigate = useNavigate();

  useEffect(() => {
    fetch(`http://localhost:8080/tasks/${id}`)
      .then((response) => response.json())
      .then((data) => setTask(data))
      .catch((error) => console.error('Error fetching task:', error));
  }, [id]);

  const handleSubmit = (event) => {
    event.preventDefault();

    fetch(`http://localhost:8080/tasks/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log('Task updated:', data);
        navigate('/');
      })
      .catch((error) => console.error('Error updating task:', error));
  };

  return (
    <div className="edit-task-container">
      <h2 className="edit-task-title">Edit Task</h2>
      <form onSubmit={handleSubmit} className="edit-task-form">
        <div className="form-group">
          <label htmlFor="title">Task Title</label>
          <input
            type="text"
            id="title"
            value={task.title}
            onChange={(e) => setTask({ ...task, title: e.target.value })}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="description">Task Description</label>
          <input
            type="text"
            id="description"
            value={task.description}
            onChange={(e) => setTask({ ...task, description: e.target.value })}
            required
          />
        </div>

        <div className="form-group checkbox-group">
          <label>
            <input
              type="checkbox"
              checked={task.completed}
              onChange={(e) => setTask({ ...task, completed: e.target.checked })}
            />
            {' '}Mark as Completed
          </label>
        </div>

        <button type="submit" className="update-btn">
          ✅ Update Task
        </button>
      </form>
    </div>
  );
};

export default EditTask;
