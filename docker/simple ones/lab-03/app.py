from pathlib import Path

from flask import Flask, jsonify

app = Flask(__name__)


@app.get("/health")
def health():
    return jsonify(status="ok")


@app.get("/message")
def message():
    message_path = Path(__file__).resolve().parent / "assets" / "message.txt"
    return jsonify(message=message_path.read_text(encoding="utf-8").strip())
