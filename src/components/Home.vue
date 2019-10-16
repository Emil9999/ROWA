<template>
  <div class="container">
    <div class="row justify-content-xl-center">
      <img src="../assets/logo_rowa.png" />
      <hr />
    </div>
    <div class="row">
      <h1>{{ msg }}</h1>
    </div>
    <div class="row">
      <router-link to="/plant">
        <button type="button" class="btn btn-primary">Plant</button>
      </router-link>
      <router-link to="/harvest">
        <button type="button" class="btn btn-primary">Harvest</button>
      </router-link>
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
