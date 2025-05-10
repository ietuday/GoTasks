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

### How do you handle dependency injection in Go?

Go doesn’t have a built-in DI framework like Java or C#. Instead, you pass dependencies explicitly via constructors (manual injection), often using interfaces for flexibility. Libraries like `wire` (Google) or `fx` (Uber) can help automate it, but many Go devs prefer the simplicity of keeping DI manual and transparent.

### What’s the difference between `map` and `slice` in Go?

A `slice` is an ordered, resizable sequence of elements—think dynamic array. A `map` is an unordered collection of key-value pairs. Slices preserve order and are index-accessible, maps allow fast key lookups. Slices are useful for lists; maps are best when you need associations (like a dictionary).

### Explain the difference between `nil` interface and interface with nil concrete value in Go.

A `nil` interface is literally zero—it holds no type and no value. But an interface that holds a concrete type with a nil value is **not** nil. This subtle difference can lead to bugs in error handling or logging. Always check both type and value when comparing interfaces to nil.

### How does `React.memo` improve performance?

`React.memo` is a higher-order component that memoizes the result. It skips re-rendering a component if props haven’t changed. It’s useful for pure functional components that don’t need to re-render on every parent change. Add a custom comparison function if shallow comparison isn’t enough.

### What's the role of `key` in React lists?

The `key` helps React identify which items in a list have changed, been added, or removed. It must be unique among siblings. Using `index` as a key is discouraged unless list items are static, because it breaks identity and can cause bugs in component state or animations.

### How does Docker handle networking between containers?

Docker sets up a default bridge network. Containers on the same network can reach each other via container names. You can also define custom user-defined bridge networks in `docker-compose`, which offer DNS-based discovery and better isolation. For external access, ports are published to the host.

### What’s the difference between `CMD` and `ENTRYPOINT` in Docker?

Both define the default command for a container. `CMD` provides default arguments that can be overridden. `ENTRYPOINT` defines the executable. Used together, they form a command: `ENTRYPOINT` + `CMD`. Use `ENTRYPOINT` for fixed behavior and `CMD` for default args (like a config file or mode).

### What are Go's zero values?

Go gives default zero values to variables that aren't explicitly initialized. For example, `int` is `0`, `string` is `""`, `bool` is `false`, and pointers/interfaces are `nil`. This makes it safe to use uninitialized variables but requires awareness to avoid bugs from assuming something was set.

### Explain reconciliation in React.

Reconciliation is React’s diffing algorithm for comparing the previous virtual DOM with the new one. It tries to minimize DOM mutations. If a node has the same type and key, it updates props. If not, it removes and re-adds. Keys play a vital role in minimizing unnecessary DOM changes.

### How do you reduce image size in Docker?

1. Use smaller base images (e.g., `alpine`).
2. Use multistage builds.
3. Clean up temporary files and package caches.
4. Avoid unnecessary dependencies.
5. Combine `RUN` steps to reduce layers.

Example:
```dockerfile
RUN apt-get update && apt-get install -y curl \
 && rm -rf /var/lib/apt/lists/*
