from typing import Optional, Tuple
from locust import HttpUser, LoadTestShape, between, task

CLIENT_ID = "locust"
CLIENT_SECRET = "locust"

USER_CREDENTIAL = {
  "email": "string@mail.com",
  "password": "string"
}

class CreateCli(HttpUser):
    wait_time = between(15, 30)
    token = None

    def on_start(self):
        res = self.client.post("/auth/login", json=USER_CREDENTIAL, timeout=15)
        self.token = res.json()["token"]
        print(self.token)

    @task()
    def get_product_detail(self):
        self.client.post("/notification/list", headers={"Authorization": f"Bearer {self.token}"}, timeout=15)

class StagesShape(LoadTestShape):
    stages = [
        {"time": 3600, "users": 100, "spawn_rate": 50}, # first 3 hours
    ]

    def tick(self):
        run_time = self.get_run_time()
        for stage in self.stages:
            if run_time < stage["time"]:
                return stage["users"], stage["spawn_rate"]
        return None