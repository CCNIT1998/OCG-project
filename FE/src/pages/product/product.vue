<template>
  <main class="main">
    <div class="row">
      
      <Collection />
      <!-- category -->
      <div class="col-sm-0 col-md-3 category">
        <div class="btn-group dropend category-btn">
          <button
            type="button"
            class="btn btn-light"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            <p>Category</p>
            <i class="bi bi-caret-down"></i>
          </button>
        </div>

        <div class="btn-group dropend category-btn category-lv1">
          <button
            type="button"
            class="btn btn-light"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            <p>Loa</p>
            <i class="bi bi-caret-right"></i>
          </button>
        </div>

        <div class="btn-group dropend category-btn category-lv1">
          <button
            type="button"
            class="btn btn-light"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            <p>Tai Nghe</p>
            <i class="bi bi-caret-right"></i>
          </button>
        </div>

        <select
          class="form-select form-select-lg mb-3"
          aria-label=".form-select-lg example"
          @change="sortPrice($event.target.value)"
        >
          <option selected>Price</option>
          <option value="asc">Low to High</option>
          <option value="desc">High to Low</option>
        </select>

        <!-- advertisement -->
        <div class="advertisement">
          <img src="https://via.placeholder.com/250x600" alt="advertisement" />
        </div>
      </div>
      <!-- end category -->

      <!-- products -->
      <div class="col-sm-12 col-md-9">
        
        <div class="filter"></div>
        <div class="row">
          <div
            class="col-sm-12 col-md-6 col-lg-4"
            v-for="product in products"
            :key="product.ID"
          >
            <router-link :to="'/products/' + product.id">
              <div class="card">
                <img
                  v-if="product.image"
                  :src="`http://localhost:3000/images/` + product.image[0].url"
                  class="card-img-top"
                />

                <img
                  v-else
                  src="https://via.placeholder.com/250x250"
                  class="card-img-top"
                />
                <div class="card-body">
                  <h5 class="card-title">{{ product.name }}</h5>
                  <div class="d-flex justify-content-between">
                    <p class="card-text">{{ formatMoney(product.price) }}<sup>vnđ</sup></p>
                    <a
                      @click="addProductToCart(product)"
                      href="#"
                      class="btn btn-outline-primary"
                      >Add to Cart</a
                    >
                  </div>
                </div>
              </div>
            </router-link>
          </div>

          <!-- pagination -->
          <div class="pagiantio">
            <Pagination
              :length="lastPage"
              :pageSize="limit"
              :pageIndex="pageIndex"
              @change="changePage"
            >
            </Pagination>
          </div>
          <!-- end pagination -->
        </div>
        <!-- end products -->
      </div>
    </div>
  </main>
</template>
<script>
import Collection from "@/components/Collection.vue";
import { mapState } from "vuex";
import Pagination from "@/components/Pagination.vue";
import parseQueryProduct from "@/utils/parseQueryProduct.js";
import formatMoney from "@/utils/formatMoney.js";
export default {
  name: "Product",
  components: {
    Pagination,
    Collection,
  },

  computed: {
    ...mapState("products", [
      "isLoading",
      "products",
      "totalItems",
      "pageIndex",
      "limit",
      "lastPage",
      "sort",
      "search",
      "order",
    ]),
    
  },
  methods: {
    formatMoney,
    changePage(pageIndex) {
      this.$router.push({
        name: "Product",
        query: { page: pageIndex, sort: this.sort, order: this.order },
      });
    },
    sortPrice(sort) {
      this.$router.push({
        name: "Product",
        query: { page: this.pageIndex, sort: sort, order: "price" },
      });
    },

    addProductToCart(item) {
      console.log(item);
      this.$store.dispatch("cart/addProductToCart", item);
    },
  },

  watch: {
    async "$route.query"() {
      let queryUrl = parseQueryProduct(this.$route.query);
      await this.$store.dispatch("products/getProducts", queryUrl);
    },
  },

  async created() {
    let queryUrl = parseQueryProduct(this.$route.query);
    queryUrl.limit = 6;
    await this.$store.dispatch("products/getProducts", queryUrl);
  },
};
</script>
<style scoped>
main {
  /* height: 100vh; */
  padding: 10px 40px;
}

.category-lv1 button {
  padding-left: 50px;
}
a {
  text-decoration: none !important;
  color: black;
}
.card img {
  width: 100%;
  height: 250px;
  object-fit: cover;
  margin: 0 auto;
}
.pagiantio {
  margin-top: 15px;
}
.category-btn {
  display: block;
  border-top: 1px solid rgba(0, 0, 0, 0.205);
}
.category-btn-price {
  transform: translate(10px, 0px);
}
.product-navigation {
  margin-top: 15px;
}
.category-btn button {
  width: 100%;
  display: flex;
  justify-content: space-between;
}

.advertisement img {
  max-height: 556px;
  width: 100%;
}

.card-title {
  font-size: 15px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
}
.card-text {
  font-size: 13px;
  transform: translate(0, 50%);
}

.card-img-top {
  width: 100%;
  transform: scale(0.95);
}
.card-img-top:hover {
  transform: scale(1);
  transition: cubic-bezier(0.075, 0.82, 0.165, 1), 0.6s;
}

@media (max-width: 768px) {
  .category {
    display: none;
  }
}
</style>
