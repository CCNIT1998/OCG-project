import axios from '../utils/axios';

const getProducts = async ({
  page,
  limit,
  sort,
  order,
  search,
  categoryId,
}) => {
  try {
    const respon = await axios.get(
      `/products?page=${page}&limit=${limit}&sort=${sort}&order=${order}&search=${search}&category=${categoryId}`
    );
    return respon.data
  } catch (e) {
    return
  }
};

const getProductById = async (productId) => {
  try {
    const respon = await axios.get(`/products/${productId}`);
    return respon.data;
  } catch (e) {
    return
  }

};


export { getProducts, getProductById }