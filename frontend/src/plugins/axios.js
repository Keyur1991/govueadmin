import axios from "redaxios"

export const $axios = axios.create({
    baseURL: import.meta.env.VITE_BACKEND_API,
    withCredentials: true
})