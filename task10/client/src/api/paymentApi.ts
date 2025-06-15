import axios from "axios";
import { CartItem } from "../context/AppContext";

export const sendPayment = async (cartItems: CartItem[]) => {
  const token = localStorage.getItem("token");

  const res = await axios.post(
    "https://task10-back.azurewebsites.net/payment",
    cartItems,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  return res.data;
};
