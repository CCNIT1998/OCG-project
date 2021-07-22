<template >
  <div class="container">
    <h3>related products</h3>
    <div class="row">
      <div
        class="col-md-6 col-lg-3"
        v-for="item in products"
        :key="item.ID"
      >
        <router-link :to="'/products/' + item.id">
          <div class="card">
            <div>
              <img
                  v-if="item.image"
                  :src="`http://localhost:3000/images/` + item.image[0].url"
                  class="card-img-top"
                />

                <img
                  v-else
                  src="https://via.placeholder.com/250x250"
                  class="card-img-top"
                />
             
            </div>
            <div class="card-body">
              <h5 class="card-title">{{ item.name }}</h5>
              <div class="d-flex justify-content-between">
                <p class="card-text">{{ formatMoney(item.price) }}<sup>vnđ</sup> </p>
                <button
                  @click="addProductToCart(item)"
                  class="btn btn-outline-primary"
                >
                  Add to Cart
                </button>
              </div>
            </div>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
import { mapState } from "vuex";
import rootImageUrl from "@/utils/rootImageUrl";
import formatMoney from "@/utils/formatMoney.js";

export default {
  name: "ProductSale",
  computed: {
    ...mapState("products", ["products", "product", "isLoading"]),
  },
  async created() {
    await this.$store.dispatch("products/getProducts", {
      category: this.product.category_id,
      limit: 4,
    });
  },
  methods: {
    formatMoney,
    rootImageUrl,
    addProductToCart(item) {
      console.log(item);
      this.$store.dispatch("cart/addProductToCart", item);
    },
  },
};
</script>
<style scoped>
a {
  text-decoration: none;
  color: black;
}

h5.card-title {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
  font-size: 15px;
  overflow: hidden;
}
h3 {
  margin: 15px 0px;
  border-top: 1px solid #00071326;
  font-size: 18px;
  padding-top: 10px;
}
.card {
  max-width: 250px;
  min-width: 250px;
  margin: 0 auto;
}
.container {
  margin-bottom: 20px;
}
</style>