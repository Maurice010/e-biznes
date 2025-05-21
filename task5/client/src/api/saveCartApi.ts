import axios from "axios";
import { CartItem } from "../context/AppContext";

export const saveCart = async (items: CartItem[]) => {
    return axios.post("http://localhost:8080/cart/save", items).then(res => res.data);
};
