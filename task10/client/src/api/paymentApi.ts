import axios from "axios";
import { CartItem } from "../context/AppContext";

export const sendPayment = async (cartItems: CartItem[]) => {
  const token = localStorage.getItem("token");

  const res = await axios.post(
    "task10-back-grfchafkdxbxd8bb.polandcentral-01.azurewebsites.net/payment",
    cartItems,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  return res.data;
};
