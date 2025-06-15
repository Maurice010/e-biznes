import axios from "axios";
import { CartItem } from "../context/AppContext";

export const saveCart = async (items: CartItem[]) => {
  const token = localStorage.getItem("token");

  const res = await axios.post(
    "https://task10-back-grfchafkdxbxd8bb.polandcentral-01.azurewebsites.net/cart/save",
    items,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  return res.data;
};
