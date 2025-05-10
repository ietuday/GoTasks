**Q: What are Goroutines and how are they different from threads?**  
A: Goroutines are lightweight threads managed by the Go runtime. They're cheaper in memory (≈2KB stack) and multiplexed onto OS threads, allowing high concurrency with low overhead.

**Q: What is the difference between buffered and unbuffered channels in Go?**  
A: Unbuffered channels block until both sender and receiver are ready. Buffered channels have a fixed size and can store values without immediate receiver readiness.

**Q: What’s the purpose of the select statement in Go?**  
A: `select` waits on multiple channel operations and executes whichever is ready first. It's used for concurrent channel management.

**Q: Explain the init() function in Go.**  
A: The `init()` function is automatically called before `main()`. It’s used for initialization logic and runs once per file/package.

**Q: How does Go handle errors?**  
A: Go returns errors explicitly using the `error` type. It encourages handling errors at each step rather than try-catch-style exceptions.

**Q: What's the difference between useEffect and useLayoutEffect in React?**  
A: `useEffect` runs after the DOM is painted. `useLayoutEffect` runs synchronously before paint, ideal for layout read/write operations.

**Q: What is a key prop in React and why is it important?**  
A: Keys help React identify changed, added, or removed elements in a list. They improve performance and consistency during re-renders.

**Q: What is the Virtual DOM in React?**  
A: The Virtual DOM is a lightweight, in-memory representation of the real DOM. React compares it with the previous version to apply efficient updates (diffing).

**Q: What’s the difference between controlled and uncontrolled components in React?**  
A: Controlled components are driven by React state. Uncontrolled components rely on the DOM and are accessed via refs.

**Q: How does React Router v6 handle routing?**  
A: React Router v6 uses `Routes` and `Route` with nested routes and element-based syntax. It supports lazy loading and declarative navigation.

**Q: What’s the difference between a Docker image and a container?**  
A: An image is a snapshot of the environment. A container is a running instance of that image.

**Q: What is a Dockerfile and why is it used?**  
A: A Dockerfile defines how to build a Docker image using base images, commands, and file copies.

**Q: How do you persist data in Docker?**  
A: By using Docker volumes, which map container data to host locations and survive restarts and rebuilds.

**Q: What does docker-compose up --build do?**  
A: It builds images (if needed), creates containers, and starts all services defined in the docker-compose file.

**Q: What are multi-stage builds in Docker?**  
A: Multi-stage builds allow you to separate the build environment from the final runtime image, reducing image size and improving security.

**Q: How can you share data between containers?**  
A: You can use Docker volumes or shared networks. Volumes allow multiple containers to read/write to the same data path.

**Q: How does Go achieve concurrency?**  
A: Using goroutines and channels. Goroutines run concurrently and communicate using channels to avoid shared memory pitfalls.

**Q: What is the difference between == and === in React (JavaScript)?**  
A: `==` does type coercion. `===` checks for both value and type equality. Always prefer `===` for reliable comparisons.

**Q: What is a Docker volume vs bind mount?**  
A: Volumes are managed by Docker and are ideal for persistent data. Bind mounts map specific host paths into containers and are more flexible during dev.

**Q: How does React’s reconciliation work?**  
A: React uses a diffing algorithm to compare the new Virtual DOM with the previous one and updates only what's changed in the actual DOM.

