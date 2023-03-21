const express = require('express');
const axios = require('axios');

const app = express();
const port = 3000;

app.get('/api/users/:id', async (req, res) => {
  try {
    const { id } = req.params;
    const response = await axios.get(`https://jsonplaceholder.typicode.com/users/${id}`);
    res.send(response.data);
  } catch (error) {
    console.error(error);
    res.status(500).send('Error');
  }
});

app.listen(port, () => {
  console.log(`Microservice listening at http://localhost:${port}`);
});
