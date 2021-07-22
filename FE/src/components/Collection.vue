<template>
  <div class="container">
    <div class="row">
      <div class="col-lg-3 col-md-3 col-6 item" v-for="(item, index) in collections" :key="index">
        <router-link to="/collection/name">
          <img :src="item.thumbnail" alt="" />
          <p>{{item.name}}</p> 
        </router-link>
      </div>
      
    </div>
  </div>
</template>
<script>
// import axios from "axios"
export default {
  data(){
    return{
      collections: []
    }
  },

  async created() {
    const response = await fetch("http://localhost:9000/collections");
    let data = await response.json();
    this.collections = data['data']
  },
};
</script>
<style scoped>
a{
  text-decoration: none;
  color: black;
}

img{
  max-width: 100px;
  max-height: 100px;
  object-fit: cover;
  border: 1px solid rgba(0, 0, 0, 0.226);
  padding: 5px;
}

.item {
    text-align: center;
    padding: 20px 0px;
}
.item img{
    transform: scale(0.9);
    border-radius: 50%;
}
.item img:hover{
    transform: scale(1);
    transition:ease 0.6s;
}

</style>