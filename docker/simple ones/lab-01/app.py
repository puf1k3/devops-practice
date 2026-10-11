import os

print("Hello from Docker!")
print("Environment:", os.getenv("APP_ENV", "development"))
print("Working directory:", os.getcwd())