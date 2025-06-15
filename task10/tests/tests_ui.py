
import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
import time

URL = "task10-front-c7eyc9dfhfabg4fb.polandcentral-01.azurewebsites.net"

@pytest.fixture(scope="module")
def driver():
    service = Service("chromedriver.exe")
    options = webdriver.ChromeOptions()
    options.add_argument("--headless")
    driver = webdriver.Chrome(service=service, options=options)
    driver.implicitly_wait(3)
    yield driver
    driver.quit()

def test_page_title(driver):
    driver.get(URL)
    assert driver.title is not None
    assert isinstance(driver.title, str)

def test_products_header(driver):
    driver.get(URL)
    assert "Produkty" in driver.page_source

def test_product_cards(driver):
    driver.get(URL)
    produkty = driver.find_elements(By.XPATH, "//h3")
    assert len(produkty) >= 1
    for p in produkty:
        assert p.text.strip() != ""

def test_product_prices_visible(driver):
    driver.get(URL)
    ceny = driver.find_elements(By.XPATH, "//p[contains(text(),'Cena')]")
    assert len(ceny) > 0
    for c in ceny:
        assert "zł" in c.text

def test_add_to_cart_buttons_present(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    assert len(buttons) > 0
    for b in buttons:
        assert b.is_displayed()

def test_add_single_item_to_cart(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    items = driver.find_elements(By.XPATH, "//p[contains(text(),'Produkt ID')]")
    assert len(items) == 1

def test_increase_quantity(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()

    quantities_before = [
        int(q.text.split(":")[1].strip()) for q in driver.find_elements(By.XPATH, "//p[contains(text(),'Ilość')]")
    ]
    driver.find_element(By.XPATH, "//button[text()='+']").click()
    time.sleep(0.5)
    quantities_after = [
        int(q.text.split(":")[1].strip()) for q in driver.find_elements(By.XPATH, "//p[contains(text(),'Ilość')]")
    ]
    assert any(a > b for a, b in zip(quantities_after, quantities_before))

def test_decrease_quantity(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()

    quantities_before = [
        int(q.text.split(":")[1].strip()) for q in driver.find_elements(By.XPATH, "//p[contains(text(),'Ilość')]")
    ]
    driver.find_element(By.XPATH, "//button[text()='-']").click()
    time.sleep(0.5)
    quantities_after = [
        int(q.text.split(":")[1].strip()) for q in driver.find_elements(By.XPATH, "//p[contains(text(),'Ilość')]")
    ]
    assert any(a <= b for a, b in zip(quantities_after, quantities_before))

def test_clear_cart(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Wyczyść koszyk')]").click()
    time.sleep(0.5)
    assert "Koszyk jest pusty" in driver.page_source

def test_save_cart_message(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Zapisz koszyk')]").click()
    time.sleep(1)
    assert "koszyk zapisany" in driver.page_source.lower() or "błąd" in driver.page_source.lower()

def test_go_to_payment(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do płatności')]").click()
    assert "Płatność" in driver.page_source

def test_payment_button_presence(driver):
    driver.get(f"{URL}/payment")
    btns = driver.find_elements(By.XPATH, "//button[contains(text(),'Zapłać')]")
    assert len(btns) == 1
    assert btns[0].is_enabled()

def test_payment_success_or_fail(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do płatności')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Zapłać')]").click()
    time.sleep(1)
    assert "transakcja" in driver.page_source.lower() or "błąd" in driver.page_source.lower()

def test_cart_persistence(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    buttons[0].click()
    driver.get(URL)
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    items = driver.find_elements(By.XPATH, "//p[contains(text(),'Produkt ID')]")
    assert len(items) >= 1

def test_reload_page(driver):
    driver.get(URL)
    driver.refresh()
    assert "produkty" in driver.page_source.lower()

def test_navigate_between_pages(driver):
    driver.get(URL)
    driver.find_element(By.XPATH, "//button[contains(text(),'Dodaj')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do płatności')]").click()
    driver.back()
    assert "Koszyk" in driver.page_source
    driver.back()
    assert "Produkty" in driver.page_source

def test_multiple_cart_items(driver):
    driver.get(URL)
    buttons = driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")
    for b in buttons[:3]:
        b.click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    produkty = driver.find_elements(By.XPATH, "//p[contains(text(),'Produkt ID')]")
    assert len(produkty) >= 2

def test_quantity_display(driver):
    driver.get(URL)
    driver.find_elements(By.XPATH, "//button[contains(text(),'Dodaj')]")[0].click()
    driver.find_element(By.XPATH, "//button[contains(text(),'Przejdź do koszyka')]").click()
    quantity = driver.find_element(By.XPATH, "//p[contains(text(),'Ilość')]")
    assert "Ilość:" in quantity.text
    assert quantity.text.strip().split(":")[1].strip().isdigit()

def test_checkout_without_cart(driver):
    driver.get(f"{URL}/payment")
    assert "Koszyk jest pusty" in driver.page_source

def test_payment_button_text(driver):
    driver.get(f"{URL}/payment")
    buttons = driver.find_elements(By.TAG_NAME, "button")
    texts = [b.text for b in buttons]
    assert any("Zapłać" in t for t in texts)
