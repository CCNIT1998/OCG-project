<template>
  <div class="row main">
    <div class="col-sm-0 col-md-3">
      <Category />
    </div>

    <div class="col-sm-12 col-md-9">
      <div class="row container product-detail">
        <!-- show images -->
        <div class="col-lg-5 col text-center">
          <div class="show-picture">
            <img
              :src="
                rootImageUrl(
                  product.image ? product.image[indexActive].url : ''
                )
              "
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
          <hr />
          <input
            type="file"
            id="file"
            ref="file"
            v-on:change="handleFileUpload()"
          />
          <button v-on:click="submitFile()">Submit</button>
        </div>
        <!-- end show images -->

        <!-- detail-product -->
        <div class="col-lg-7 col">
          <div class="product-content">
            <h5 class="title"></h5>

            <!-- name -->
            <div class="row">
              <div class="col-3">Name</div>
              <div class="col-9">
                <input
                  type="email"
                  class="form-control name"
                  id="exampleFormControlInput1"
                  :value="product.name"
                />
              </div>
            </div>
            <br />
            <!-- trade_mark -->
            <div class="row">
              <div class="col-3">Trade Mark</div>
              <div class="col-9">
                <select
                  :value="product.trade_mark"
                  class="form-select"
                  id="TradeMark"
                >
                  <option value="Việt Nam">Việt Nam</option>
                  <option value="Chine">Chine</option>
                  <option value="US">US</option>
                </select>
              </div>
            </div>

            <br />
            <!-- Category -->
            <div class="row">
              <div class="col-3">Category</div>
              <div class="col-9">
                <input
                  type="email"
                  class="form-control"
                  id="exampleFormControlInput1"
                  value="Tai Nghe"
                />
              </div>
            </div>

            <!-- Price -->
            <br />
            <div class="row">
              <div class="col-3">Price</div>
              <div class="col-9">
                <input
                  type="email"
                  class="form-control price"
                  id="exampleFormControlInput1"
                  :value="product.price"
                />
              </div>
            </div>

            <!-- Description -->
            <br />
            <div class="row">
              <div class="col-3">Description</div>
              <div class="col-9">
                <p class="description">
                  <textarea
                    class="form-control description"
                    id="exampleFormControlTextarea1"
                    rows="3"
                    :value="product.description"
                  ></textarea>
                </p>
              </div>
            </div>

            <div class="product-cart">
              <button
                @click="updateProduct()"
                type="button"
                class="btn btn-outline-primary"
              >
                Update
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Category from "@/admin/product/category.vue";
import { mapState } from "vuex";
import rootImageUrl from "@/utils/rootImageUrl";
import axios from "axios";
export default {
  name: "productUpdate",
  components: {
    Category,
  },

  data() {
    return {
      imgActive: "https://via.placeholder.com/250x252",
      indexActive: 0,
      file: "",
      productUpdate: {
        name: "",
        description: "",
        price: "",
        category_id: "",
        trade_mark: "",
      },
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

    updateProduct() {
      let nameUpdate = document.querySelector(".name")["value"];
      let trade_markUpdate = document.querySelector("#TradeMark")["value"];
      let priceUpdate = document.querySelector(".price")["value"];
      let descriptionUpdate = document.querySelector(".description")["value"];
      axios
        .patch("http://localhost:3000/products", {
          id: this.product.id,
          name: nameUpdate,
          trade_mark: trade_markUpdate,
          price: Number(priceUpdate),
          description: descriptionUpdate,
        })
        .then((res) => {
          console.log(res);
        });

      alert("updated successfully");
      location.reload();
    },

    handleFileUpload() {
      this.file = this.$refs.file.files[0];
    },

    submitFile() {
      console.log("http://localhost:3000/uploads/" + String(this.product.id));
      let formData = new FormData();
      formData.append('file', "/home/cuong/Videos/1/79190b65025fd4342211d976b7b9cfbb.jpeg");
      axios
        .post(
          "http://localhost:3000/uploads/" + String(this.product.id),
          formData,
                {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
                }
        )
        .then(function () {
          console.log("SUCCESS!!");
        })
        .catch(function () {
          console.log("FAILURE!!");
        });
    },

   
  },
};
</script>
<style scoped>
.main {
  height: 215vh;
  padding-right: 20px;
  margin: 0 !important;
  position: absolute;
  top: 0px;
  left: 0px !important;
  right: 0px !important;
  bottom: 0px;
  width: 100vw;
  background-color: yellow;
  z-index: 10;
  background-color: white;
}

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
  margin-left: 14px;
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