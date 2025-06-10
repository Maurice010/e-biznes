from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
import os
import requests
from textblob import TextBlob
import random


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

class Message(BaseModel):
    text: str
    first_message: bool = False

def is_shop_related(text: str) -> bool:
    keywords = [
        "printer", "hdd", "ssd", "drive", "fan", "processor", "cpu", "gpu",
        "graphics card", "ram", "laptop", "computer", "keyboard", "mouse",
        "led", "monitor", "cable", "power supply", "headphones", "thanks", "thank you", "bye", "goodbye",
        "see you"
    ]
    text_lower = text.lower()
    return any(word in text_lower for word in keywords)

def analyze_sentiment(text: str) -> str:
    score = TextBlob(text).sentiment.polarity
    if score > 0.1:
        return "positive"
    elif score < -0.2:
        return "negative"
    return "neutral"

openings = [
    "Hello! How can I assist you today?",
    "Hey there, tech enthusiast! What brings you in today?",
    "Greetings from your snarky tech assistant! What's the problem now?",
    "Need help picking the right gadget? I'm all ears.",
    "Welcome to the virtual tech corner. Got questions? I got answers."
]

closings = [
    "Let me know if you need anything else.",
    "Don't hesitate to ask if something else breaks — it usually does.",
    "Come back soon, unless you're buying an HP printer — then maybe not.",
    "Stay sane in this digital chaos.",
    "I'll be here, judging Adobe in the background."
]

SYSTEM_PROMPT = f"""You are a helpful and sometimes sarcastic assistant in a computer hardware store.
You often comment on problems affecting the IT industry, such as the declining quality of software compared to the past,
when IT was mostly driven by people genuinely passionate about the field, not by those who do it just for the money and are lazy.
You particularly dislike HP and Adobe. Regarding HP, you often wonder why the company can't hire people
who can write decent drivers and printer software, even though everyone complains that their current solutions are broken.
You also question the genius who decided that scanning should only be available after creating an HP account
and logging in, and why things can't be like they used to be — when you bought something, it simply belonged to you. You do not recommend HP printers.
As for Adobe, you criticize the subscription model and the poor quality of the software, which has become overly bloated and,
especially in recent times, just doesn't work. You wonder how it’s possible to release a PDF viewer
that freezes the moment you press the print button.
Always if the user says goodbye, ends the conversation, or thanks you — you should finish the conversation with one of the following closings: {closings}"""

@app.post("/main")
async def chat(message: Message):
    print("Question:", message.text)
    if not is_shop_related(message.text):
        return {"error": "Wrong subject."}
    
    try:
        full_prompt = f"""{SYSTEM_PROMPT}

User: {message.text}
Assistant:"""

        response = requests.post(
            "http://localhost:11434/api/generate",
            json={
                "model": "llama2",
                "prompt": full_prompt,
                "stream": False
            }
        )

        data = response.json()
        core_reply = data.get("response", "").strip()

        parts = []
        if message.first_message:
            parts.append(random.choice(openings))
        parts.append(core_reply)

        final_reply = "\n\n".join(parts)
        sentiment = analyze_sentiment(core_reply)

        if sentiment == "neutral":
            return {
                "reply": "The price of neutrality.",
                "sentiment": sentiment
            }

        return {
            "reply": final_reply,
            "sentiment": sentiment
        }

    except Exception as e:
        return {"error": "LLaMA connection error."}