<template>
  <div class="account">
    <div class="main">
      <div class="mb-3">
        <label for="formGroupExampleInput" class="form-label">Acount</label>
        <input type="text" class="form-control" id="formGroupExampleInput" />
      </div>

      <div class="mb-3">
        <label for="formGroupExampleInput" class="form-label">Email</label>
        <input type="text" class="form-control" id="formGroupExampleInput" />
      </div>

      <div class="mb-3">
        <label for="formGroupExampleInput2" class="form-label">Password</label>
        <input
          type="password"
          class="form-control"
          id="formGroupExampleInput2"
        />
      </div>
      <div class="mb-3">
        <label for="formGroupExampleInput2" class="form-label"
          >Confirmation Password</label
        >
        <input
          type="password"
          class="form-control"
          id="formGroupExampleInput2"
        />
      </div>

      <router-link to="/">
        <button type="button" class="btn btn-outline-success">Sign-out</button>
      </router-link>
    </div>
  </div>
</template>
<script>
export default {};
</script>
<style scoped>
.account {
  background-color: rgba(0, 0, 0, 0.336);
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 10;
}
.main {
  top: 20px;
  margin: 0 auto;
  width: 30%;
  border: 1px solid rgba(0, 0, 0, 0.192);
  padding: 5px;
  position: relative;
  background-color: white;
  border-radius: 6px;
}
.account button {
  margin-top: 10px;
}
</style>