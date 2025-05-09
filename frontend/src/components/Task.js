import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import './TaskDetails.css';

const TaskDetails = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [task, setTask] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(`http://localhost:8080/tasks/${id}`)
      .then((res) => res.json())
      .then((data) => setTask(data))
      .catch((err) => {
        setError('Failed to load task details.');
        console.error(err);
      });
  }, [id]);

  if (error) return <p>{error}</p>;
  if (!task) return <p>Loading task details...</p>;

  return (
    <div className="task-details-container">
      <h2>Task Details</h2>
      <div className="task-card">
        <p><strong>Title:</strong> {task.title}</p>
        <p><strong>Description:</strong> {task.description}</p>
        <p><strong>Status:</strong> {task.completed ? '✅ Completed' : '⏳ Incomplete'}</p>
      </div>
      <button onClick={() => navigate('/')} className="back-btn">⬅️ Go Back</button>
    </div>
  );
};

export default TaskDetails;
