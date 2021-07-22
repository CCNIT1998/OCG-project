const state = () => ({
  products: [],
  discount: 0,
  isLoading: false,
  addToCartResult: "",
  totalItems: 0,
  isShowCartDropdown: false,
});

const getters = {
  totalItems(state) {
    return state.products.reduce(
      (total, product) => total + product.quantity,
      0
    );
  },

  subTotal(state) {
    let sum = state.products.reduce(
      (totalPrice, product) =>
        totalPrice +
        product.quantity * product.price,
      0
    );
    return sum 
  },
  total(state) {
    let sum = state.products.reduce(
      (totalPrice, product) =>
        totalPrice +
        product.quantity * product.price,
      0
    );
    return (sum - sum * state.discount / 100)
  },
  
};

const actions = {
  addProductToCart({ state, commit }, product) {
    const isExists = state.products.find((p) => p.id === product.id);
    if (isExists) {
      commit("setAddToCartResult", "This item was already in your cart.");
      alert("product has been added to cart")
    } else {
      // product.quantity = 1;
      // product.totalPrice = product.quantity * product.price;
      const newProduct = { ...product };
      newProduct.quantity = 1;
      newProduct.totalPrice =newProduct.quantity * newProduct.price;
      console.log("hello");
      if (newProduct) {
        commit("setAddToCartResult", "Item successfully added to your cart.");
        commit("addProductToCart", newProduct);
      }
    }

    commit("setShowCartDropdown", true);
  },

  async getProductsInCart({ commit }) {
    commit("setLoading", true);

    const products = JSON.parse(localStorage.getItem("cart"));

    commit("setProducts", products);
    commit("setLoading", false);
  },

  updateProduct({ commit }, { value, index }) {
    if (!value) {
      value = 1;
    }
    commit("updateProductQuantity", { value, index });
  },

  deleteProduct({ commit }, index) {
    console.log(index);
    commit("deleteProduct", index);
  },

  addCoupon({ commit }, value) {
    console.log(value)
    commit("addCoupon", value);
  },
};

const mutations = {
  setShowCartDropdown(state, status) {
    state.isShowCartDropdown = status;
  },

  setLoading(state, status) {
    state.isLoading = status;
  },

  setProducts(state, products) {
    state.products = products.map((product) => {
      product.totalPrice = product.quantity * product.price;

      return product;
    });
  },

  updateProductQuantity(state, { value, index }) {
    value = parseInt(value, 10);
    index = parseInt(index, 10);
    console.log(value, index);
    const product = state.products[index];
    value = Number(value);

    if (value > 1) {
      product.quantity = value;
    } else {
      product.quantity = 1;
    }

    product.totalPrice = product.price * product.quantity;
  },

  addProductToCart(state, product) {
    state.products.push(product);
    localStorage.setItem("cart", JSON.stringify(state.products));
  },

  setAddToCartResult(state, message) {

    state.addToCartResult = message;

  },

  deleteProduct(state, index) {
    state.products.splice(index, 1);
    localStorage.setItem("cart", JSON.stringify(state.products));
  },
  addCoupon(state, value) {
    console.log(value);
    switch (value) {
      case "coupon1": {
        state.discount = 50;
        alert("Discount 50%")
        break;
      }
      case "coupon2": {
        state.discount = 10;
        alert("Discount 10%")

        break;
      }
      default:
        console.log(state.products)
        alert("Coupon does not exist  ")

        state.discount = 0;
        break;
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};