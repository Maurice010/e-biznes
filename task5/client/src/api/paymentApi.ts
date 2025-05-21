import axios from "axios";
import { CartItem } from "../context/AppContext";

export const sendPayment = async (cartItems: CartItem[]) => {
    const res = await axios.post("http://localhost:8080/payment", cartItems);
    return res.data;
};
