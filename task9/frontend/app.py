from flask import Flask, render_template, request
import requests

app = Flask(__name__)

FASTAPI_URL = "http://127.0.0.1:8000/main"

@app.route("/", methods=["GET", "POST"])
def index():
    reply = None
    sentiment = None
    error = None

    if request.method == "POST":
        user_input = request.form.get("message")
        if user_input:
            try:
                response = requests.post(FASTAPI_URL, json={"text": user_input})
                data = response.json()
                if "error" in data:
                    error = data["error"]
                else:
                    reply = data["reply"]
                    sentiment = data["sentiment"]
            except Exception as e:
                error = f"No connection with model: {e}"

    return render_template("main.html", reply=reply, sentiment=sentiment, error=error)

if __name__ == "__main__":
    app.run(debug=True)