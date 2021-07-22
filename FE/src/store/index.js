import { createStore } from 'vuex';
import products from './modules/products.module.js';
import cart from './modules/cart.module.js'
const debug = process.env.NODE_ENV !== 'production';

const store = createStore({
  strict: debug,
  modules: {
    products,
    cart
  },
});
export default store;
