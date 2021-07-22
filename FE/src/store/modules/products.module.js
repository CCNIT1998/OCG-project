import { getProducts, getProductById } from '@/service/products.service.js';

const state = {
  products: [],
  total: 0,
  product: {},
  isLoading: false,
  pageIndex: 1,
  lastPage: 0,
  limit: 6,
  sort: '',
  order: '',
  search: '',
  category: 0,
};
// const getters = {
//   listImage(state) {
//     let images = [...state.product.image]
//     console.log(images)
//       images = images.map((item) => {
//         console.log(item)
//         return "http://localhost:3000/images/" + item.url
//       })
//       return images
//   }

// }



const actions = {
  async getProducts(
    { state, commit },
    { pageIndex, limit, sort, order, search, category }
  ) {
    commit('setLoading', true);
    if (pageIndex) commit('setPageIndex', pageIndex);
    if (limit) commit('setLimit', limit);
    if (sort) commit('setSort', sort);
    if (order) commit('setOrder', order);
    if (search !== undefined) commit('setSearch', search);
    if (category) commit('setCategory', category);

    const response = await getProducts({
      page: state.pageIndex,
      limit: state.limit,
      sort: state.sort,
      order: state.order,
      search: state.search,
      categoryId: state.category,
    });
    commit('setProducts', response.data);
    commit('setLastPage', response.meta.last_page);
    commit('setTotal', response.meta.total);
    commit('setLoading', false);
  },

  async getProductById({ commit }, productId) {
    commit("setLoading", true);

    const product = await getProductById(productId);

    commit("setProduct", product);
    commit("setLoading", false);
  },
};
const mutations = {
  setLoading(state, status) {
    state.isLoading = status;
  },
  setProducts(state, response) {
    state.products = response;
  },
  setProduct(state, product) {
    state.product = product;
  },
  setPageIndex(state, index) {
    index = parseInt(index, 10);
    state.pageIndex = index;
  },
  setLimit(state, limit) {
    state.limit = limit;
  },
  setOrder(state, order) {
    state.order = order;
  },
  setSearch(state, search) {
    state.search = search;
  },
  setCategory(state, category) {
    state.category = category;
  },
  setTotal(state, total) {
    state.total = total;
  },
  setLastPage(state, total) {
    state.lastPage = total;
  },
  setSort(state, sort) {
    state.sort = sort;
  },
};
export default {
  namespaced: true,
  state,
  actions,
  mutations,

};
