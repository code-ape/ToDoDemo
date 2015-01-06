from locust import HttpLocust, TaskSet, task

class WebsiteTasks(TaskSet):

  @task
  def main_page(self):
    self.client.get("/")


class WebsiteUser(HttpLocust):
  task_set = WebsiteTasks
  min_wait = 5000
  wax_wait = 15000
