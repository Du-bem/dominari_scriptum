// const API_URL = `http://127.0.0.1:8000/ui/search/${id}`;
//
async function getDataFromAPI() {
  const data = {
    name: "Earth (399)",
    time: Date.parse("2024-08-13 00:00:00"),
    position: [
      116804704591.193175833296891141799278557300567626953125 * 6.68459e-12,
      -96592333563.197653223397765032132156193256378173828125 * 6.68459e-12,
      5090599.90494556028153318172346786241178051568567752838134765625 *
        6.68459e-12
    ],
    velocity: [
      18490.06438058809197139440192201062826724309060308668348524305555555555555555555555555555555555555555 *
        5.77548e-7,
      22849.85669343816658696182880166563538498141699367099338107638888888888888888888888888888888888888888 *
        5.77548e-7,
      -0.6983613243274335312647398979415546837612038399567137713785524721498842592592592592592592592592592591 *
        5.77548e-7
    ],
    checksum: "Hi2sm3So2mzs2SGW3ks1L2o9SnB8xs3"
  };
  return data;
}

async function getChainFromAPI() {
  const data = [
    { link: "hhtps://asd", hash: "AFWmdasoqen" },
    { link: "hhtps://asd", hash: "AFWmdasoqen" },
    { link: "hhtps://asd", hash: "AFWmdasoqen" },
    { link: "hhtps://asd", hash: "AFWmdasoqen" }
  ];
  return data;
}

export { getDataFromAPI, getChainFromAPI };
