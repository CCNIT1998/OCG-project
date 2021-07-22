<template>
  <div class="container">
    <div class="">
      <router-link to="/">
        <button type="button" class="btn btn-outline-primary cart-home">
          Home
        </button>
      </router-link>
      <ul class="products">
        <div class="row">
          <div class="col text-left" style="width: 60px; height: 60px">
            Image
          </div>
          <div class="col">Name</div>
          <div class="col">Quantity</div>
          <div class="col">Price</div>
          <div class="col"></div>
        </div>

        <div
          class="row product-detail"
          v-for="(product,index) in products"
          :key="product.id"
        >
          <div class="col text-left" style="width: 60px; height: 60px">
            <img :src="rootImageUrl(product.image?product.image[0].url:'')" />
          </div>
          <div class="col n product-name">{{ product.name }}</div>
          <div class="col n">{{ product.quantity }}</div>
          <div class="col n">{{ formatMoney(product.totalPrice) }}<sup>vnđ</sup> </div>
          <div @click="deleteProduct(index)" class="col delete-product n">
            <i class="bi bi-x text-right"></i>
          </div>
        </div>

        <!-- <li class="row" v-for="product in $store?.state.cart" :key="product.ID">
          <div class="col left">
            <div class="thumbnail">
              <a href="#"
                ><img
                  :src="`http://127.0.0.1:3000/${product.Image}`"
                  style="width: 60px; height: 60px"
              /></a>
            </div>
            <div class="detail">
              <div class="name">
                <a href="">{{ product.Name }}</a>
              </div>
              <div class="price">{{ 1000 }}</div>
            </div>
          </div>
          <div class="col right">
            <div class="quantity">
              <input
                type="number"
                class="quantity"
                v-model="product.quantity"
                placeholder="1"
              />
            </div>
            <div class="remove">
              <i @click="removeItem(index)" class="close far fa-trash-alt"></i>
            </div>
          </div>
        </li> -->
      </ul>
    </div>

    <div class="row">
      <div class="col-10">
        <div class="row">
          <div class="input-group mb-3">
            <input
              v-model="coupon"
              type="text"
              class="form-control"
              placeholder="Recipient's username"
              aria-label="Recipient's username"
              aria-describedby="basic-addon2"
            />
            <span @click="addCoupon" class="input-group-text" id="basic-addon2"
              ><i class="bi bi-arrow-right-circle"></i
            ></span>
          </div>

          <div class="row">
            <button type="button" class="btn btn-outline-primary checkout">
              Checkout
            </button>
          </div>
        </div>
      </div>
      <div class="col-2">
        <div class="row">
          <div class="col-6">subTotal:</div>
          <div class="col-6">{{ formatMoney(subTotal) }}<sup>vnđ</sup></div>
        </div>
        
        <div class="row">
          <div class="col-6">Discount:</div>
          <div class="col-6">{{discount}}</div>
        </div>
        <div class="row">
          <div class="col-6">Total:</div>
          <div class="col-6">{{ formatMoney(total) }}<sup>vnđ</sup></div>
        </div>
      </div>
    </div>

    <ProductSale />
  </div>
</template>
<script>
import ProductSale from "@/pages/product/productSale.vue";
import { mapState, mapGetters } from "vuex";
import rootImageUrl from "@/utils/rootImageUrl";
import formatMoney from "@/utils/formatMoney.js"
export default {
  name: "Cart",
  components: {
    ProductSale,
  },
  data() {
    return {
      coupon: ""
    }
  },
  computed: {
    ...mapState("cart", ["discount", "products", "isLoading"]),
    ...mapGetters("cart", ["total","totalItems", "subTotal"]),
  },
  methods: {
    formatMoney,
    rootImageUrl,
    changeQuanti(value, index) {
      let convert = parseInt(value, 10);
      this.cloneCart[index].quantity = convert;
      if (value === "" || convert === 0) {
        this.cloneCart[index].quantity = 1;
      }
    },
    addCoupon() {
      console.log(this.coupon);
      this.$store.dispatch("cart/addCoupon", this.coupon);
    },


    deleteProduct(item) {
      this.$store.dispatch('cart/deleteProduct',item)
    },
  },
  created() {
    this.cloneCart = [...this.products];
  },
};
</script>

<style scoped>
.product-name{
   display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
  overflow: hidden;
}
img{
  max-width: 60px;
  max-height: 60px;
}

#basic-addon2 {
  cursor: pointer;
}
.products {
  border: 1px solid #e5e7e8;
}

.product-detail {
  border-top: 1px solid #e5e7e8;
  line-height: 60px;
  margin-bottom: 5px;
}

.delete-product {
  cursor: pointer;
  position: relative;
}

.delete-product i {
  position: absolute;
  right: 30px;
  font-size: 30px;
}
.cart-home {
  margin: 10px 0px;
}
.checkout {
  width: auto;
  margin-left: 12px;
}
</style>