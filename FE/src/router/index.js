import { createRouter, createWebHistory } from "vue-router";

import ProductDetail from "@/pages/product/productDetail.vue";
import Product from "@/pages/product/product.vue";
import Cart from "@/pages/cart/index.vue";
import SignIn from "@/pages/authentication/login.vue";
import SignOut from "@/pages/authentication/resgister.vue";
import CollectionDetail from "@/pages/product/collectionDetail.vue";

import AdminProduct from "@/admin/product/product.vue";
import ProductAdd from "@/admin/product/product-add.vue";
import productUpdate from "@/admin/product/productUpdate.vue";
import CategoryAdd from "@/admin/category/category-add.vue";
import Oder from "@/admin/order/order.vue";
import Username from "@/admin/user/useredit.vue";
import Listuser from "@/admin/user/list-user.vue";
// import About from "@/views/About.vue";

const routes = [
  {path: "/",name: "Product",component: Product},
  {path: "/:pathMatch(.*)*", component: () => import("@/components/404.vue") },
  {path: "/products/:id",component: ProductDetail},
  {path: "/collection/name",component: CollectionDetail},

  // {path: "/productDetail/:id",name: "ProductDetail",component: ProductDetail},
  {path: "/cart", component: Cart},
  {path: "/login", component: SignIn},
  {path: "/register", component: SignOut},


  {path: "/admin/product", name: "ProductAdmin", component: AdminProduct},
  {path: "/admin/product/add", component: ProductAdd},
  {path: "/admin/product/add/:id", component: productUpdate},
  {path: "/admin/category/add", component: CategoryAdd},
  {path: "/admin/order", component: Oder},
  {path: "/admin/username", component: Username},
  {path: "/admin/list-user", component: Listuser},


//   {
//     path: "/about",
//     name: "About",
//     component: About,
//   },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;