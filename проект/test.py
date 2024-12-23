import requests

url = "http://127.0.0.1:8000/analyze"
file_path = "C:/Users/legyx/Desktop/овца.jpg"  # Путь к вашему тестовому изображению

with open(file_path, "rb") as f:
    files = {"file": (file_path, f, "image/jpeg")}
    response = requests.post(url, files=files)

print(response.json())
