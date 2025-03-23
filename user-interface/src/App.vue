<script setup>
import Plot3d from "./components/3d.vue";
import RightComp from "./components/right_comp.vue";
import Transactions from "./components/transactions.vue";
import NavBar from "./components/nav_bar.vue";
</script>

<template>
  <NavBar />
  <div class="main">
    <RightComp />
    <Plot3d class="plot" />
    <Transactions class="transactions" />
  </div>
</template>

<style scoped>
.plot {
  border-top-right-radius: 12px;
  border-bottom-right-radius: 12px;
  border: 1px solid rgba(65, 90, 160, 0.3);
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.5), 0 0 20px rgba(65, 120, 255, 0.1);
}
.main {
  display: grid;
  grid-template-columns: 1fr 2.1fr;
}

.transactions {
  grid-column: 1/3;
}
</style>
