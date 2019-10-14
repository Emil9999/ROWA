<template>
  <div class="container">
    <div class="row justify-content-xl-center">
      <img src="../assets/logo_rowa.png" />
      <hr>
    </div>
    <div class="row">
      <h1>{{ msg }}</h1>
    </div>
    <div class="row">
      <button type="button" class="btn btn-primary">
        <router-link to="/plant">Plant</router-link>
      </button>
      <button type="button" class="btn btn-primary">
        <router-link to="/harvest">Harvest</router-link>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      msg: "Farm Overview"
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
