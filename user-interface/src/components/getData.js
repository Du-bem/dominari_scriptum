async function getDataFromAPI(){
    const data = {
        name: "earth",
        time: 2433,
        position: [1432, 123, 234],
        velocity: [242, 43, 213],
        checksum: "68ac906495480a3404beee4874ed853a037a7a8f"
      };
    return data;
}

async function getChainFromAPI(){
    const data = [
        {link:"hhtps://asd", hash:"AFWmdasoqen"},
        {link:"hhtps://asd", hash:"AFWmdasoqen"},
        {link:"hhtps://asd", hash:"AFWmdasoqen"},
        {link:"hhtps://asd", hash:"AFWmdasoqen"}
    ];
    return data;
}

export {getDataFromAPI, getChainFromAPI}