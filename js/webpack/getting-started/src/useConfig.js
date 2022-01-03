import { useEffect, useState } from 'react';
import axios from 'axios';

let loaded = false;
let _config = null;
const url = 'https://httpbin.org/uuid';

function UseConfig() {
  const [config, setConfig] = useState(_config);
  useEffect(() => {
    console.log("triggered", config);
    if (loaded) {
      return;
    }
    axios.get(url).then(function(response){
      setConfig(response.data);
      loaded = true;
    }).catch(function(error){
      console.log(`failed to get config with ${error}`);
    });
  }, [config]);
  return config;
}

export default UseConfig
