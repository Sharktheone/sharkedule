import axios, {AxiosInstance} from "axios"

const MODE = "dev"
const DEV_URL = "http://localhost:8080/api"

const PROD_URL = window.location.host + "/api"

export const api: AxiosInstance = axios.create({
    baseURL: MODE === "dev" ? DEV_URL : PROD_URL,
})