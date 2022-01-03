import _ from 'lodash';
import React from 'react';
import UseConfig from './useConfig.js';

function App() {
  const config = UseConfig();
  return <h1>{_.join(['Hello', 'webpack', config?.uuid], ' ')}</h1>;
}

export default App
