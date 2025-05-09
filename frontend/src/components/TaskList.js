import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './TaskList.css'; // Add this to include custom styles

const TaskList = () => {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then((response) => response.json())
      .then((data) => {
        setTasks(data);
        setLoading(false);
      })
      .catch((error) => {
        setError('Failed to load tasks.');
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Loading tasks...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div className="task-list-container">
      <h2 className="task-list-title">Task List</h2>
      <button
        onClick={() => navigate('/add')}
        className="add-task-btn"
      >
        Add New Task
      </button>

      {Array.isArray(tasks) && tasks.length > 0 ? (
        <table className="task-table">
          <thead>
            <tr>
              <th>Task Name</th>
              <th>Description</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {tasks.map((task) => (
              <tr key={task.id}>
                <td>{task.title}</td>
                <td>{task.description}</td>
                <td className="task-actions">
                  <button onClick={() => handleEdit(task.id)} className="edit-btn">Edit</button>
                  <button onClick={() => handleDelete(task.id)} className="delete-btn">Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No tasks available</p>
      )}
    </div>
  );
};

// Example Edit and Delete handlers
const handleEdit = (taskId) => {
  console.log('Edit task with id:', taskId);
};

const handleDelete = (taskId) => {
  console.log('Delete task with id:', taskId);
};

export default TaskList;
