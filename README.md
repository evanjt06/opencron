# OpenCron
OpenCron is a lightweight, in-memory cron-style job scheduler written in Go. It supports both one-time delayed jobs and recurring interval-based jobs, all executed asynchronously via goroutines.

## Features
- Schedule tasks to run **once after a delay** or **on a recurring interval**
- In-memory job store (no external dependencies)
- Asynchronous execution using goroutines
- Thread-safe with mutex locks for concurrent access
- Simple time-based polling (checks every second)

## API
- `ScheduleOnce(delay time.Duration, callback func())`
- `ScheduleRepeat(interval time.Duration, callback func())`
- `Start()` — starts the scheduler

# Motivations
I made OpenCron as a basic exercise to understand how cron-style scheduling systems work. OpenScheduler is an in-memory job scheduler that lets you schedule tasks either once (after a delay) or on a repeating interval. It’s inspired by how systems like cron, Celery Beat, or distributed job queues manage delayed and recurring job execution.

Tasks are stored in-memory and run asynchronously (in the background) using goroutines. For recurring jobs, OpenScheduler will requeue the task at the specified interval. For one-time jobs, it deletes the task after execution. Thread-safety is ensured with mutexes, so the scheduler is safe for concurrent usage. All jobs are checked every single second for execution eligibility based on the current time.
