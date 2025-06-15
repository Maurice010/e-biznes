import requests

BASE_URL = "https://task10-back-grfchafkdxbxd8bb.polandcentral-01.azurewebsites.net"

def test_get_products():
    res = requests.get(f"{BASE_URL}/products")
    assert res.status_code == 200
    data = res.json()
    assert isinstance(data, list)
    assert all("id" in p and "name" in p and "price" in p for p in data)

def test_post_cart_valid():
    payload = [{"productId": 1, "quantity": 2}]
    res = requests.post(f"{BASE_URL}/cart/save", json=payload)
    assert res.status_code == 200
    assert "cart_id" in res.json()

def test_post_cart_invalid_empty():
    res = requests.post(f"{BASE_URL}/cart/save", json=[])
    assert res.status_code == 400

def test_post_payment_valid():
    payload = [{"productId": 1, "quantity": 1}]
    res = requests.post(f"{BASE_URL}/payment", json=payload)
    assert res.status_code == 200
    assert "total" in res.json()

def test_post_payment_invalid_empty():
    res = requests.post(f"{BASE_URL}/payment", json=[])
    assert res.status_code in (400, 422)
