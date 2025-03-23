import { defineStore } from "pinia";
import { ref } from "vue";
import { getDataFromAPI, getChainFromAPI } from "./getData.js";

export const usePlanetsStore = defineStore("planet", () => {
  const data = ref({});
  const checksum = ref({});
  async function loadData() {
    data.value = await getDataFromAPI();
    checksum.value = await getChainFromAPI();
  }

  function setDataPosAndVel(pos, vel) {
    data.value.position = pos;
    data.value.velocity = vel;
    data.value.time += 86400000;
  }

  return { data, checksum, loadData, setDataPosAndVel };
});
