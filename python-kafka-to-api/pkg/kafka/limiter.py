import time
import threading

class DLQWriteLimiter:
    def __init__(self, rate: float, capacity: int):
        self.capacity = capacity
        self.tokens = capacity
        self.rate = rate
        self.lock = threading.Lock()
        self.last_time = time.monotonic()

    def allow(self) -> bool:
        with self.lock:
            now = time.monotonic()
            elapsed = now - self.last_time
            self.last_time = now
            self.tokens = min(self.capacity, self.tokens + elapsed * self.rate)
            if self.tokens >= 1:
                self.tokens -= 1
                return True
            return False
