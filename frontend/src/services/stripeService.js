import axios from "axios";

const baseClient = () =>
    axios.create({
        baseURL: window.configs.vite_app_endpoint,
    });

export default {
    async getProducts() {
        return await baseClient().get("/products");
    },

    async checkout(cart, success_url, failure_url) {
        return await baseClient().post("/checkout", {
            cart,
            success_url,
            failure_url,
        });
    },
};