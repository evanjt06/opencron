package internal

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       string
	Interval time.Duration
	RunAt    time.Time
	Callback func()
	Repeat   bool
}

type Scheduler struct {
	Jobs []*Job
	Mu   sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Jobs: []*Job{},
	}
}

func (s *Scheduler) ScheduleOnce(delay time.Duration, callback func()) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	job := &Job{
		ID:       fmt.Sprintf("job-%d", time.Now().UnixNano()),
		RunAt:    time.Now().Add(delay),
		Callback: callback,
		Repeat:   false,
	}
	s.Jobs = append(s.Jobs, job)
}

func (s *Scheduler) ScheduleRepeat(interval time.Duration, callback func()) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	job := &Job{
		ID:       fmt.Sprintf("job-%d", time.Now().UnixNano()),
		RunAt:    time.Now().Add(interval),
		Interval: interval,
		Callback: callback,
		Repeat:   true,
	}
	s.Jobs = append(s.Jobs, job)
}

func (s *Scheduler) Start() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			now := time.Now()

			s.Mu.Lock()
			for i := 0; i < len(s.Jobs); i++ {
				job := s.Jobs[i]
				if now.After(job.RunAt) || now.Equal(job.RunAt) {
					go job.Callback()

					if job.Repeat {
						job.RunAt = job.RunAt.Add(job.Interval)
					} else {
						s.Jobs = append(s.Jobs[:i], s.Jobs[i+1:]...)
						i-- // adjust index removal of an element in a slice shifts everything leftwards by 1
					}
				}
			}
			s.Mu.Unlock()
		}
	}()
}
