const express = require('express');
const app = express();

const echo = require('echo');

/*eslint-disable no-undef */
const PORT = process.env.PORT || 3000;

app.get('/', (req, res) => {
  console.log(echo("foobarbaz"));
  res.send('Hello from the simplest Express server!');
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
