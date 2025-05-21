import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Products from "./components/Products";
import Payment from "./components/Payment";
import Cart from "./components/Cart";
import { AppProvider } from "./context/AppContext";

export default function App() {
  return (
    <AppProvider>
      <Router>
        <Routes>
          <Route path="/" element={<Products />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/payment" element={<Payment />} />
        </Routes>
      </Router>
    </AppProvider>
  );
}
