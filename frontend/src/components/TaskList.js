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
        // Assuming data is an array
        setTasks(data);
        setLoading(false);
      })
      .catch((error) => {
        setError('Failed to load tasks.');
        setLoading(false);
        console.error('Error loading tasks:', error); // Log actual error
      });
  }, []);

  const handleEdit = (taskId) => {
    console.log('Navigate to edit task with id:', taskId);
    navigate(`/edit/${taskId}`); // Navigate to the edit page for the selected task
  };

  const handleDelete = (taskId) => {
    if (window.confirm('Are you sure you want to delete this task?')) {
      fetch(`http://localhost:8080/tasks/${taskId}`, {
        method: 'DELETE',
      })
        .then((response) => {
          if (response.ok) {
            setTasks(tasks.filter(task => task.id !== taskId));
            alert('Task deleted successfully!');
          } else {
            alert('Failed to delete task.');
          }
        })
        .catch((error) => {
          console.error('Error deleting task:', error);
          alert('Failed to delete task.');
        });
    }
  };

  const handleComplete = (taskId) => {
    const updatedTasks = tasks.map((task) =>
      task.id === taskId
        ? { ...task, completed: !task.completed }
        : task
    );
    setTasks(updatedTasks);

    // Optionally, you can update the task completion status in the backend as well
    fetch(`http://localhost:8080/tasks/${taskId}`, {
      method: 'PUT',
      headers: {
      'Content-Type': 'application/json',
      },
      body: JSON.stringify(updatedTasks.find((task) => task.id === taskId)),
    })
      .then((response) => {
      if (!response.ok) {
        alert('Failed to update task completion status.');
      }
      })
      .catch((error) => {
      console.error('Error updating task completion status:', error);
      alert('Failed to update task completion status.');
      })
      .finally(() => {
      // call task list again to refresh the data
      fetch('http://localhost:8080/tasks')
        .then((response) => response.json())
        .then((data) => {
          setTasks(data);
        })
        .catch((error) => {
          console.error('Error loading tasks:', error);
        });
      } 
    );
  };

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
              <th>Completed</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {tasks.map((task) => (
              <tr key={task.id}>
                <td>{task.title}</td>
                <td>{task.description}</td>
                <td>
                  <input
                    type="checkbox"
                    checked={task.completed}
                    onChange={() => handleComplete(task.id)}
                  />
                </td>
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

export default TaskList;
