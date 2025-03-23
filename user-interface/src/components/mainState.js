export const usePlanetsStore = defineStore("planet", () => {
  const data = ref({});
  const checksum = ref({});
  async function loadData() {
    data.value = await getDataFromAPI();
    checksum.value = await getChainFromAPI();
  }

  return { data, checksum, loadData };
});
