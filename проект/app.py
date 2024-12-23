from fastapi import FastAPI, UploadFile
from io import BytesIO
from PIL import Image
from transformers import AutoModelForImageClassification, AutoProcessor
import torch

# Инициализация FastAPI
app = FastAPI()

# Название модели на Hugging Face
model_name = "microsoft/resnet-50"

# Загрузка модели и процессора
try:
    model = AutoModelForImageClassification.from_pretrained(model_name)
    processor = AutoProcessor.from_pretrained(model_name)
    print("Модель успешно загружена.")
except Exception as e:
    print(f"Ошибка загрузки модели: {e}")
    exit(1)

# Функция для анализа изображения
def predict(image: Image.Image) -> str:
    # Преобразование изображения в формат модели
    inputs = processor(images=image, return_tensors="pt")
    with torch.no_grad():
        outputs = model(**inputs)
    # Определение класса
    predicted_class = outputs.logits.argmax(-1).item()
    classes = model.config.id2label  # Маппинг ID -> Лейбл
    return classes[predicted_class]

# API для анализа изображения
@app.post("/analyze")
async def analyze_image(file: UploadFile):
    try:
        # Открываем изображение
        image = Image.open(BytesIO(await file.read())).convert("RGB")
        # Предсказываем результат
        result = predict(image)
        return {"result": result}
    except Exception as e:
        return {"error": str(e)}

# Запуск API
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)