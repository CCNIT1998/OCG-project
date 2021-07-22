<template>
  <div>
    <div class="row container product-detail">
      <!-- show images -->
      <div class="col-lg-5 col text-center">
        <div class="show-picture">


          <img
            :src="rootImageUrl(product.image?product.image[indexActive].url:'')"
            alt="img"
            class="img-show"
          />
        </div>
        <div class="image">

          <div class="img-scroll d-flex">
            <div
              v-for="(image, index) in product.image"
              :key="index"
              class="picture"
            >
              <img
                :src="rootImageUrl(image.url)"
                alt="img"
                :class="{ active: index == indexActive ? true : false }"
                @mouseover="changeActiveImg(index)"
              />
            </div>
          </div>
        </div>
      </div>
      <!-- end show images -->

      <!-- detail-product -->
      <div class="col-lg-7 col">
        <div class="product-content">
          <h5 class="title">{{ product.name }}</h5>
          <div class="row">
            <div class="col-3">Description</div>
            <div class="col-9">
              <p class="description">
                Some quick example text to build on the card title and make up
                the bulk of the card's content. Some quick example text to build
                on the card title and make up the bulk of the card's content.
                Some quick example text to build on the card title and make up
                the bulk of the card's content. Some quick example text to build
                on the card title and make up the bulk of the card's content.
              </p>
            </div>
          </div>
          <div class="row">
            <div class="col-3">Color</div>
            <div class="col-9">
              <div
                class="btn-group"
                role="group"
                aria-label="Basic outlined example"
              >
                <button type="button" class="btn btn-outline">White</button>
                <button type="button" class="btn btn-outline">Black</button>
                <button type="button" class="btn btn-outline">Gree</button>
              </div>
            </div>
          </div>
          <br /><br />
          <div class="row">
            <div class="col-3">Price</div>
            <div class="col-9">
              {{ formatMoney(product.price) }}<sup>vnđ</sup>
            </div>
          </div>
          <router-link to="/Cart">
            <div class="product-cart">
              <!-- <button type="button" class="btn btn-outline-primary">
                <i class="bi bi-cart-plus"></i>Add To Cart
              </button> -->
              <a href="" class="btn btn-outline-primary">Buy Now</a>
            </div>
          </router-link>
        </div>
      </div>
    </div>

    <ProductSale />
  </div>
</template>
<script>
import ProductSale from "@/pages/product/productSale.vue";
import { mapState } from "vuex";
import rootImageUrl from "@/utils/rootImageUrl";
import formatMoney from "@/utils/formatMoney.js";
export default {
  name: "ProductDetail",
  components: {
    ProductSale,
  },
  data() {
    return {
      imgActive: "https://via.placeholder.com/250x252",
      indexActive: 0,
   
    };
  },

  computed: {
    ...mapState("products", ["product", "isLoading"]),
  },

  async created() {
    // Get product detail
    await this.$store.dispatch(
      "products/getProductById",
      this.$route.params.id
    );
    console.log(this.$route.params.id);


  },

  async beforeRouteUpdate(to) {
    this.$store.dispatch("products/getProductById", to.params.id);
  },

  methods: {
    changeActiveImg(index) {
      this.imgActive = this.product.image[index];
      this.indexActive = index;
    },
    rootImageUrl,
    formatMoney,
  },
};
</script>
<style scoped>
.product-cart {
  margin-top: 20px;
}
.product-cart i {
  margin-right: 10px;
}

.product-cart button:hover {
  opacity: 1;
}
.product-cart a:hover {
  opacity: 1;
}
.description {
  text-align: justify;
}

.product-cart button {
  margin-right: 5px;
  background-color: rgba(208, 1, 27, 0.12);
  color: #d0011b;
  opacity: 0.6;
  border: none;
  border-radius: 0;
}

.product-cart a {
  /* margin-left: 14px; */
  margin-right: 5px;
  background-color: #d0011b;
  color: white;
  opacity: 0.9;
  border: none;
  border-radius: 0;
}

.product-detail {
  margin: 20px 0px;
}
.product-content h5 {
  margin-bottom: 16px;
}

.image {
  display: inline-block;
}

.img-scroll {
  margin-top: 5px;
}

.img-show {
  /* max-width: 350px;
  min-width: 350px; */
  width: 350px !important;
  height: 350px !important;
}

.picture img {
  max-width: 60px;
  min-width: 60px;
  margin: 0px 2px;
  cursor: pointer;
}

.active {
  border: 1px solid #ee4d2d;
}
.btn-group button {
  margin-right: 15px;
  border: 1px solid rgb(201, 165, 144);
}

.btn-group button:hover {
  border: 1px solid #ee4d2d;
}
</style>