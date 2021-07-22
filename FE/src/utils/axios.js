import axios from "axios";
const rootPath = "http://localhost:3000";

const axiosInstance = axios.create({
  baseURL: rootPath,
});
// axiosInstance.interceptors.request.use(
//   (config) => new Promise((resolve) => setTimeout(() => resolve(config), 1000))
// );
export default axiosInstance;